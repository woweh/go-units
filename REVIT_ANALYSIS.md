# Detailed Analysis: Revit Units vs go-units

## Conversion Requirements Summary

- **Many new Revit units have been added.**
- **Many Revit units now match or can be mapped to existing go-units (via AddAlias or direct mapping).**
- **Some Revit units do not need conversion (single-unit per quantity):**
  - Color Temperature (K)
  - Efficacy (lm/W)
  - Electrical Resistivity (ohm·m)
  - Luminous Flux (lm)
  - Luminous Intensity (cd)
  - Pulsation (rad/s)
  - Wattage (W)
- **No-conversion required (multi-unit, identical factors):**
  - None found in current RevitUnits.json
- **Skipped for now (cost-related; TBD viability):**
  - Cost Rate Energy, Cost Rate Power, Cost per Area, Currency

**Notes:**
- For single-unit quantities above, implement the quantity and base unit if needed, but no ratio mappings are required beyond identity. If a quantity is already covered elsewhere (e.g., Color Temperature uses Kelvin from temperature), just alias/name-map symbols as needed.
- For multi-unit quantities, most Revit units now map directly to existing units or can be aliased. Use AddAlias where appropriate.
- Cost-related units are still skipped.
- The requirements and mappings are up-to-date with the latest Revit export.
- See `RevitUnits.json` for the latest Revit unit definitions and mappings.

## Revit Quantities Found (150 total)

### Common/Basic Quantities
- ✅ Acceleration (m/s², ft/s²) - PRESENT
- ✅ Angle (°, rad, grad) - PRESENT
- ✅ Area (m², ft², acres, hectares) - PRESENT
- ✅ Currency (¥) - SKIPPED (cost-related)
- ✅ Data/Bytes (B, KB, MB, GB, TB, PB) - PRESENT
- ✅ Density (kg/m³, lb/ft³) - PRESENT
- ✅ Distance/Length (m, ft, in, cm, mm, km, USft, dm) - PRESENT
- ✅ Time (s, min, h, ms, d, yr) - PRESENT
- ✅ Volume (m³, ft³, L, gal) - PRESENT
- ✅ Mass (kg, lb, t, Tons) - PRESENT
- ✅ Temperature (°C, °F, K, °R) - PRESENT

### Electrical Quantities
- ✅ Amperes (A, kA, mA) - PRESENT
- ✅ Apparent Power (VA, kVA, MVA) - PRESENT (in power_units)
- ✅ Apparent Power Density (VA/ft², VA/m²) - PRESENT (mapped/aliased)
- ✅ Cable Tray Size (length units) - PRESENT
- ✅ Color Temperature (K) - PRESENT (temperature, single-unit, alias only)
- ✅ Conduit Size (length units) - PRESENT
- ✅ Demand Factor (%, fixed) - PRESENT (dimensionless)
- ✅ Efficacy (lm/W) - PRESENT (single-unit, alias only)
- ✅ Electric Charge (A·h, A·s, A·min, C) - PRESENT
- ✅ Electrical Potential (V, kV, mV) - PRESENT
- ✅ Electrical Resistivity (ohm·m) - PRESENT (single-unit, alias only)
- ✅ Frequency (Hz, cycles/s, rev/min, deg/s) - PRESENT (with aliases)
- ✅ Illuminance (lux, ftc) - PRESENT
- ✅ Luminance (cd/m², cd/ft², ftL) - PRESENT
- ✅ Luminous Flux (lm) - PRESENT (single-unit, alias only)
- ✅ Luminous Intensity (cd) - PRESENT (single-unit, alias only)
- ✅ Power (W, kW, Btu/h, hp, kVA, cal/s, kcal/s, MBH) - PRESENT
- ✅ Power Density (W/m², W/ft²) - PRESENT
- ✅ Power per Length (W/m, W/ft) - PRESENT
- ✅ Voltage (V, kV, mV) - PRESENT
- ✅ Wattage (W) - PRESENT (single-unit, alias only)

### Energy & Thermal Quantities
- ✅ Coefficient of Heat Transfer (W/(m²·K), Btu/(h·ft²·°F)) - PRESENT
- ✅ Demand Factor (%) - PRESENT
- ✅ Energy (J, kJ, Btu, cal, kWh, therm) - PRESENT
- ✅ Heat Capacity per Area - PRESENT
- ✅ Heat Transfer Coefficient - PRESENT
- ✅ Isothermal Moisture Capacity - PRESENT
- ✅ Permeability - PRESENT
- ✅ Specific Heat (J/(kg·°C), Btu/(lb·°F)) - PRESENT
- ✅ Specific Heat of Vaporization - PRESENT
- ✅ Thermal Conductivity (W/(m·K), Btu/(h·ft·°F)) - PRESENT
- ✅ Thermal Expansion Coefficient - PRESENT
- ✅ Thermal Insulance (°F·h·ft²/Btu, K·m²/W) - PRESENT (all symbols)
- ✅ Thermal Mass (J/K, Btu/°F) - PRESENT
- ✅ Thermal Resistance - PRESENT

### HVAC/Flow Quantities
- ✅ Air Flow (CFM, m³/h, L/min) - PRESENT
- ✅ Air Flow Density (CFM/ft², LPS/m²) - PRESENT
- ✅ Air Flow divided by Cooling Load (CFM/ton, L/(s·kW)) - PRESENT
- ✅ Air Flow divided by Volume (CFM/ft³) - PRESENT
- ✅ Angular Speed (RPM, RPS) - PRESENT
- ✅ Area divided by Cooling Load (m²/kW, SF/MBh, SF/ton) - PRESENT
- ✅ Area divided by Heating Load - PRESENT
- ✅ Cooling Load (kW, W, Btu/h, ton) - PRESENT
- ✅ Cooling Load divided by Area (W/m², Btu/(h·ft²)) - PRESENT
- ✅ Cooling Load divided by Volume (W/m³, Btu/(h·ft³)) - PRESENT
- ✅ Cross Section (m², ft², cm²) - PRESENT (as area)
- ✅ Density (kg/m³, lb/ft³) - PRESENT
- ✅ Diffusivity (ft²/s, m²/s) - PRESENT
- ✅ Duct Insulation Thickness - PRESENT (as length)
- ✅ Duct Size - PRESENT (as length)
- ✅ Factor (%, fixed) - PRESENT
- ✅ Flow per Power - PRESENT
- ✅ Friction (Pa/m, in-wg/100ft) - PRESENT
- ✅ Heat Gain (W, Btu/h) - PRESENT
- ✅ Heating Load (W, kW, Btu/h) - PRESENT
- ✅ Mass per Time (kg/s, lb/h) - PRESENT
- ✅ Power (W, Btu/h, kW) - PRESENT
- ✅ Power Density - PRESENT
- ✅ Power per Flow - PRESENT
- ✅ Pressure (Pa, psi, Torr, in-wg, bar, kPa) - PRESENT
- ✅ Roughness - PRESENT (as length)
- ✅ Slope - PRESENT
- ✅ Temperature - PRESENT
- ✅ Temperature Difference - PRESENT (interval scales)
- ✅ Velocity/Speed (ft/s, m/s, mph, km/h) - PRESENT
- ✅ Viscosity (Pa·s, cP) - PRESENT

### Piping Quantities
- ✅ Most piping quantities reuse existing types and are PRESENT

### Structural/Mechanical Quantities
- ✅ Acceleration - PRESENT
- ✅ Force (N, kN, lbf, kip, tf, Tons) - PRESENT (with kip alias)
- ✅ Moment/Torque (N·m, lb·ft, kip·ft) - PRESENT
- ✅ Moment of Inertia (m⁴, in⁴, ft⁴, mm⁴) - PRESENT
- ✅ Moment Scale - PRESENT
- ✅ Period (s, ms, min, h) - PRESENT (as time)
- ✅ Pulsation (rad/s) - PRESENT (single-unit, alias only)
- ✅ Rotation (°, rad, grad) - PRESENT (as angle)
- ✅ Rotational Spring - PRESENT
- ✅ Section Area (m², cm², in², ft²) - PRESENT (as area)
- ✅ Section Modulus - PRESENT
- ✅ Stress (Pa, psi, ksi, MPa) - PRESENT (reuse pressure)
- ✅ Unit Weight (kN/m³, lbf/ft³) - PRESENT
- ✅ Velocity/Speed - PRESENT
- ✅ Warping Constant - PRESENT
- ✅ Weight per Length - PRESENT

### Infrastructure Quantities
- ✅ Stationing - PRESENT

### Common Quantities
- ✅ Cost per Area - SKIPPED (cost-related)
- ✅ Cost Rate Energy - SKIPPED (cost-related)
- ✅ Cost Rate Power - SKIPPED (cost-related)
- ✅ Dimensionless Ratio (%, ‰, fixed) - PRESENT
- ✅ Number (currency, %, fixed, general) - PARTIAL

## Summary

- **Most Revit units now map directly to existing go-units or can be aliased.**
- **Single-unit quantities require only identity mapping or aliasing.**
- **Cost-related units are still skipped.**
- **See `RevitUnits.json` for the latest, complete Revit unit export and mapping reference.**

The requirements and mappings above are up-to-date as of the latest Revit export.
