
import csv
import glob
import re
import os
import collections
import sys

sys.stdout.reconfigure(encoding='utf-8')

GO_UNITS_DIR = r"e:\GitHub\go-units"
REVIT_CSV = r"e:\GitHub\go-units\RevitUnits.csv"

# 1. Parse existing Go units
existing_units = set()
existing_symbols = set()

unit_pattern = re.compile(r'(?:MustCreateUnit|NewUnit)\("([^"]+)",\s*"([^"]+)"')
alias_pattern = re.compile(r'AddAliases\(([^)]+)\)')

for filepath in glob.glob(os.path.join(GO_UNITS_DIR, "*_units.go")):
    with open(filepath, 'r', encoding='utf-8') as f:
        content = f.read()
        # Find defined units
        for match in unit_pattern.finditer(content):
            name, symbol = match.groups()
            existing_units.add(name.lower())
            existing_symbols.add(symbol)
        
        # Find aliases (rough parsing)
        for match in alias_pattern.finditer(content):
            # args might be "foo", "bar"
            raw_args = match.group(1)
            aliases = [a.strip().strip('"') for a in raw_args.split(',')]
            for a in aliases:
                existing_units.add(a.lower())

# 2. Parse Revit CSV
# Format: Quantity,Discipline,Quantity TypeId,Type Catalog String,Unit Name,Unit TypeId,Unit Symbol,Conversion Factor From Internal,Conversion Factor To Internal,Is Valid

revit_data = collections.defaultdict(list)
internal_units = {} # Quantity -> Unit Name
conversion_needed = []
no_conversion_single = []
no_conversion_identical = []
missing_units = []

with open(REVIT_CSV, 'r', encoding='utf-8') as f:
    reader = csv.reader(f)
    header = next(reader)
    for row in reader:
        if len(row) < 9: continue
        qty = row[0]
        unit_name = row[4]
        symbol = row[6]
        try:
             factor_to_internal = float(row[8])
        except ValueError:
             # Try to find the factor column? 
             # Or just skip
             print(f"Skipping row due to float error: {row}", flush=True)
             continue
             
        revit_data[qty].append({
            'name': unit_name,
            'symbol': symbol,
            'factor': factor_to_internal
        })
        
        # Identify internal unit
        # Usually factor 1.0. If multiple have 1.0, any can be base, but usually "Meters" or "Feet".
        if abs(factor_to_internal - 1.0) < 1e-9:
            # Prefer standard names if multiple match
            if qty not in internal_units:
                internal_units[qty] = unit_name
            else:
                # heuristic: prefer SI or simple names if multiple 1.0 exist?
                # Actually, usually there is only one defined as "Internal" in Revit's logic implicitly,
                # but physically multiple might have factor 1.0 (e.g. Kelvin and CelsiusInterval).
                # We'll just keep the first one found or overwrite if "Meters" or "Feet" found.
                cur = internal_units[qty]
                if unit_name in ["Meters", "Feet", "Seconds", "Kilograms", "Kelvin", "Radians", "Watts", "Newtons"]:
                     internal_units[qty] = unit_name

# 3. Analyze
for qty, units in revit_data.items():
    # Check coverage
    qty_missing = []
    for u in units:
        # Check name or symbol match
        if u['name'].lower() in existing_units or u['symbol'] in existing_symbols :
            pass 
        # Check if mapped by alias (e.g. "Centimeters" -> "centimeter")?
        elif u['name'].lower().rstrip('s') in existing_units:
            pass
        else:
            qty_missing.append(u['name'] + " (" + u['symbol'] + ")")
    
    # Classification
    unique_factors = set(u['factor'] for u in units)
    
    if len(units) == 1:
        no_conversion_single.append(qty)
    elif len(unique_factors) == 1:
        no_conversion_identical.append(qty)
    else:
        # Check if cost
        if "Cost" in qty or "Currency" in qty:
             # Skip or note? User said "assume currency stays same" -> simple conversion.
             # So actually they DO need conversion if factors differ.
             # In Revit data, Cost Rate Energy: $/Btu vs $/W-h. Factors differ.
             conversion_needed.append(qty)
             pass
        else:
            conversion_needed.append(qty)

    if qty_missing:
        missing_units.append((qty, qty_missing))

# 4. Generate Output Markdown

print("# Detailed Analysis: Revit Units vs go-units")
print("")
print("## Internal Base Units")
print("| Quantity | Revit Internal Unit | SI Base Unit | Notes |")
print("|---|---|---|---|")

files_si_map = {
    "Length": "Meters",
    "Area": "Square Meters",
    "Volume": "Cubic Meters",
    "Mass": "Kilograms",
    "Time": "Seconds",
    "Temperature": "Kelvin",
    "Force": "Newtons",
    "Pressure": "Pascals",
    "Energy": "Joules",
    "Power": "Watts",
    "Electrical Potential": "Volts",
    "Current": "Amperes",
    "Frequency": "Hertz",
    "Velocity": "Meters per Second",
    "Acceleration": "Meters per Second Squared",
    "Density": "Kilograms per Cubic Meter",
    "Angle": "Radians" 
} # Heuristic map

for qty in sorted(internal_units.keys()):
    internal = internal_units[qty]
    si_base = files_si_map.get(qty, "?")
    
    # Check if internal is SI
    note = "Matches SI" if internal.lower() in [s.lower() for s in files_si_map.values()] or internal.lower() == si_base.lower() else "Non-SI Internal"
    if internal == "Feet": note = "Imperial Internal"
    if internal == "Meters": note = "SI Internal"
    
    print(f"| {qty} | {internal} | {si_base} | {note} |")

print("")
print("## Conversion Requirements")
print("")
print("### 1. No Conversion Required (Single Unit)")
print("The following quantities have only one unit defined in Revit, so no conversion logic is needed (identity only).")
for q in sorted(no_conversion_single):
    # Check if we have it
    status = "[OK] Covered"
    # Find missing for this qty
    miss = next((m[1] for m in missing_units if m[0] == q), [])
    if miss:
        status = f"[MISSING] Missing: {', '.join(miss)}"
    print(f"- **{q}**: {status}")

print("")
print("### 2. No Conversion Required (Identical Factors)")
print("The following quantities have multiple units but all share the same conversion factor (effectively synonyms or same unit).")
for q in sorted(no_conversion_identical):
    status = "[OK] Covered"
    miss = next((m[1] for m in missing_units if m[0] == q), [])
    if miss:
        status = f"[MISSING] Missing: {', '.join(miss)}"
    print(f"- **{q}**: {status}")

print("")
print("### 3. Conversion Required")
print("The following quantities require conversion implementation.")
for q in sorted(conversion_needed):
    # Determine covered status
    miss = next((m[1] for m in missing_units if m[0] == q), [])
    if not miss:
        print(f"- [OK] **{q}**: All units covered.")
    else:
        # Check if ALL missing?
        revit_u = revit_data[q]
        if len(miss) == len(revit_u):
             print(f"- [MISSING] **{q}**: Not implemented. Missing: {', '.join(miss)}")
        else:
             print(f"- [PARTIAL] **{q}**: Partially implemented. Missing: {', '.join(miss)}")

