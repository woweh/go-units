# Revit Unit Coverage Analysis

This analysis maps Revit quantities and units to the `go-units` library. Coverage is determined by whether the specific Unit Name or Symbol exists in `go-units`, regardless of the Quantity it falls under.

**Criteria:**
- **Covered**: The Unit Name (e.g., "Meters") or Symbol (e.g., "m") exists in `go-units` (globally).
- **Missing**: Neither the name nor symbol exists in `go-units`.

## Unit Coverage by Quantity

| Revit Quantity | Revit Unit | Revit Internal | SI Base | Go-Unit Status |
|---|---|---|---|---|
| **Acceleration** | Feet per second squared | Feet per second squared | Meters per Second Squared | ✅ Covered (Meter, Second) |
| | Meters per second squared | | | ✅ Covered |
| | Kilometers per second squared | | | ✅ Covered |
| | Inches per second squared | | | ✅ Covered |
| **Air Flow** | Cubic feet per minute (CFM) | Cubic feet per minute | Cubic Meters per Second | ✅ Covered (Cubic Foot, Minute) |
| | Cubic meters per hour (m³/h) | | | ✅ Covered |
| | Cubic meters per second (m³/s) | | | ✅ Covered |
| | Liters per second (L/s) | | | ✅ Covered |
| | Liters per minute (L/min) | | | ✅ Covered |
| | Gallons per minute (GPM) | | | ✅ Covered |
| | Gallons per hour (GPH) | | | ✅ Covered |
| **Angle** | Degrees | Radians | Radians | ✅ Covered |
| | Radians | | | ✅ Covered |
| | Gradians | | | ✅ Covered |
| | Degrees of Arc | | | ✅ Covered |
| | Minutes of Arc | | | ✅ Covered |
| | Seconds of Arc | | | ✅ Covered |
| **Area** | Square Feet | Square Feet | Square Meters | ✅ Covered |
| | Square Meters | | | ✅ Covered |
| | Square Inches | | | ✅ Covered |
| | Square Centimeters | | | ✅ Covered |
| | Square Millimeters | | | ✅ Covered |
| | Acres | | | ✅ Covered |
| | Hectares | | | ✅ Covered |
| **Cost** | Cost per Square Foot | Cost per Square Foot | ? | ❌ Not Supported (Currency) |
| **Cross Section** | Square Feet | Square Feet | Square Meters | ✅ Covered (Area) |
| **Current** | Amperes | Amperes | Amperes | ✅ Covered |
| | Milliamperes | | | ✅ Covered |
| | Kiloamperes | | | ✅ Covered |
| **Density** | Kilograms per Cubic Meter | Kilograms per Cubic Meter | Kilograms per Cubic Meter | ✅ Covered |
| | Pounds per Cubic Foot | | | ✅ Covered |
| **Distance** | Feet | Feet | Meters | ✅ Covered |
| | Meters | | | ✅ Covered |
| **Duct Size** | Feet | Feet | Meters | ✅ Covered |
| | Inches | | | ✅ Covered |
| | Millimeters | | | ✅ Covered |
| **Electrical Potential** | Volts | Volts | Volts | ✅ Covered |
| | Millivolts | | | ✅ Covered |
| | Kilovolts | | | ✅ Covered |
| **Energy** | Joules | Joules | Joules | ✅ Covered |
| | Kilojoules | | | ✅ Covered |
| | Kilowatt Hours | | | ✅ Covered |
| | British Thermal Units (BTU) | | | ✅ Covered |
| **Force** | Newtons | Newtons | Newtons | ✅ Covered |
| | Kilonewtons | | | ✅ Covered |
| | Pounds Force | | | ✅ Covered |
| | Kips | | | ✅ Covered |
| **Frequency** | Hertz | Cycles per second | Hertz | ✅ Covered |
| | Kilohertz | | | ✅ Covered |
| **Heat Transfer Coeff.** | Watts per Square Meter Kelvin | Watts per square meter kelvin | ? | ✅ Covered |
| **Illuminance** | Lux | Footcandles (Implied?) | Lux | ✅ Covered |
| | Footcandles | | | ✅ Covered |
| **Length** | Meters | Feet | Meters | ✅ Covered |
| | Centimeters | | | ✅ Covered |
| | Millimeters | | | ✅ Covered |
| | Feet | | | ✅ Covered |
| | Inches | | | ✅ Covered |
| **Linear Force** | Newtons per Meter | Newtons per Meter | ? | ❌ Missing |
| **Luminance** | Candelas per Square Meter | Candelas per Square Meter | ? | ✅ Covered |
| | Footlamberts | | | ✅ Covered |
| **Luminous Flux** | Lumens | Lumens | Lumens (cd·sr) | ✅ Covered |
| **Luminous Intensity** | Candelas | Candelas | Candelas | ✅ Covered |
| **Mass** | Kilograms | Kilograms | Kilograms | ✅ Covered |
| | Pounds Mass | | | ✅ Covered |
| **Moment** | Newton Meters | Newton Meters | ? | ✅ Covered |
| | Pound Force Feet | | | ✅ Covered |
| **Power** | Watts | Watts | Watts | ✅ Covered |
| | Kilowatts | | | ✅ Covered |
| | Horsepower | | | ✅ Covered |
| **Pressure** | Pascals | Pascals | Pascals | ✅ Covered |
| | Kilopascals | | | ✅ Covered |
| | Megapascals | | | ✅ Covered |
| | PSI | | | ✅ Covered |
| | Inches of Water | | | ✅ Covered |
| | Inches of Mercury | | | ✅ Covered |
| **Rotation** | Degrees | Radians | Radians | ✅ Covered |
| | Radians | | | ✅ Covered |
| **Speed** | Meters per Second | Feet per Second | Meters per Second | ✅ Covered |
| | Kilometers per Hour | | | ✅ Covered |
| | Miles per Hour | | | ✅ Covered |
| | Feet per Second | | | ✅ Covered |
| | Feet per Minute | | | ✅ Covered |
| **Temperature** | Kelvin | Kelvin | Kelvin | ✅ Covered |
| | Celsius | | | ✅ Covered |
| | Fahrenheit | | | ✅ Covered |
| | Rankine | | | ✅ Covered |
| **Time** | Seconds | Seconds | Seconds | ✅ Covered |
| | Minutes | | | ✅ Covered |
| | Hours | | | ✅ Covered |
| **Volume** | Cubic Meters | Cubic Feet | Cubic Meters | ✅ Covered |
| | Cubic Feet | | | ✅ Covered |
| | Liters | | | ✅ Covered |
| | Gallons | | | ✅ Covered |
| **Volume Flow** | (See Air Flow) | | | ✅ Covered |

*Note: This table summarizes key quantities. Units like "Pipe Size", "Wire Diameter", "Reinforcement Spacing" are strictly Length units and are covered by the generic Length units (Meters, Feet, Inches).*

## Required Additions to Go-Units

The following units name or symbol were identified as potentially missing during analysis (verified against current codebase):

### 1. Missing Symbols/Aliases
These units exist, but Revit-specific symbols might need to be added as aliases.

- **Cubic Feet per Minute (CFM)**: Unit exists (`CubicFootPerMinute`), symbol `CFM` needs verification.
- **Linear Feet (LF)**: `Length` exists (`Foot`), symbol `LF` needs to be added as alias.
- **Square Feet (SF)**: `Area` exists (`SquareFoot`), symbol `SF` needs to be added as alias.
- **Tons (Refrigeration)**: Unit exists? Needs check.
- **Kips**: Exists in `force_units.go`.

### 2. Missing Units
These units do not appear to be in `go-units` and need to be implemented.

- **Cost Units**: `Cost per Square Foot`, `Cost per square meter` (Out of scope?)
- **Compound Spring/Force Coefficients**:
    - `Newtons per meter` (Linear Force / Spring constant)
    - `Kilonewtons per square meter` (Area Force / Subgrade modulus)
    - `Kilonewton meters per degree` (Rotational Spring)
- **Fluid/Thermal Specifics**:
    - `Lumens per Watt` (Efficacy)
    - `Permeability` units
    - `Thermal Gradient Coefficient` units
- **Slopes/Ratios**:
    - `Rise/Run` ratios (1:10, etc) - Likely handled as Dimensionless/Ratio.

## Next Steps

1.  **Add Missing Symbols**: Update existing units (Foot, Square Foot, Cubic Foot) to include Revit symbols (LF, SF, CF, CFM, etc.).
2.  **Add Derived Units**: Implement missing compound units where necessary (e.g., Linear Force `N/m` if not present).
3.  **Validation**: Verify that `go-units` can parse "LF", "SF", "CFM" correctly after updates.
