import csv
from collections import defaultdict

units_by_quantity = defaultdict(set)

with open('RevitUnits.csv', 'r', encoding='utf-8') as f:
    reader = csv.reader(f)
    headers = next(reader)
    
    for row in reader:
        if len(row) < 7:
            continue
        quantity = row[0].strip('"')
        unit_symbol = row[6].strip('"')
        
        if unit_symbol:
            units_by_quantity[quantity].add(unit_symbol)

sorted_quantities = sorted(units_by_quantity.items())

print('=' * 90)
print('REVIT UNITS SUMMARY - ORGANIZED BY QUANTITY TYPE')
print('=' * 90)
print()

for quantity, symbols in sorted_quantities:
    sorted_symbols = sorted(symbols)
    symbols_str = ', '.join(sorted_symbols)
    print(f'{quantity}:')
    print(f'  Unit Symbols ({len(sorted_symbols)}): {symbols_str}')
    print()

print('=' * 90)
print(f'TOTAL QUANTITY TYPES: {len(sorted_quantities)}')
total_unique_symbols = sum(len(symbols) for symbols in units_by_quantity.values())
print(f'TOTAL UNIQUE UNIT SYMBOLS: {total_unique_symbols}')
print('=' * 90)
