# Units in Revit That Require Formulae (Not Just Multiplicative Factors)

This document lists all 'exotic' and derived units in Revit (from RevitUnits.csv) that require a conversion formula involving more than a simple multiplication/division (i.e., require an offset or other operation). For each, the required formula is provided where possible.

## 1. Temperature Units (already covered)
- Celsius, Fahrenheit, Kelvin, Rankine (absolute temperature, not intervals)
- Require offset and scaling.
- **Status:** All temperature formulae are implemented in `temperature_units.go`.

## 2. Units That Might Need Formulae

### A. Pressure, Friction, and Related Units with Reference Temperatures
Some pressure and friction units are defined at specific reference temperatures (e.g., "Feet of water (39.2 °F)", "Inches of mercury (32 °F)"). However, these are still handled by a conversion factor, as the reference temperature is part of the unit definition, not a variable in the conversion. **No formula needed beyond the factor.**
- **Status:** All conversions are multiplicative and implemented.

### B. Slope Units
- Slope can be expressed as a ratio (rise/run), percentage, per mille, or degrees.
- Converting between ratio and degrees requires a trigonometric function:
  - **From ratio to degrees:**
    - `degrees = atan(ratio) * 180 / π`
  - **From degrees to ratio:**
    - `ratio = tan(degrees * π / 180)`
- Converting between percentage and ratio is multiplicative, but between degrees and percentage/ratio, a formula is needed.
- **Status:** All required formulae are implemented in `slope_units.go`.

### C. Angle Units (Degrees, Radians, Gradians, DMS)
- All can be handled by factors except for Degrees Minutes Seconds (DMS), which requires parsing/formatting:
  - **DMS to degrees:**
    - `degrees = D + M/60 + S/3600`
  - **Degrees to DMS:**
    - `D = int(degrees)`
    - `M = int((degrees - D) * 60)`
    - `S = ((degrees - D) * 60 - M) * 60`
- **Status:** Only ratio-based conversions (degree, radian, gon, turn) are implemented in `angle_units.go`. **DMS conversions are not implemented.**

### D. Compound Units with Offsets (None found)
- All other derived units (e.g., energy per area, mass per time, etc.) use only multiplicative factors.
- **Status:** All conversions are multiplicative and implemented.

## 3. Summary Table
| Unit Type         | Example Unit(s)                | Formula Needed? | Formula/Note | Status |
|-------------------|-------------------------------|-----------------|-------------|--------|
| Slope             | Degrees, Ratio, Percentage     | Yes (trig)      | See above   | Implemented |
| Angle (DMS)       | Degrees Minutes Seconds        | Yes (parse)     | See above   | **Missing** |
| Pressure/Friction | Feet of water (39.2 °F), etc.  | No              | Factor only | Implemented |
| Temperature       | Celsius, Fahrenheit, etc.      | Yes (offset)    | See above   | Implemented |
| All others        |                               | No              | Factor only | Implemented |

## 4. Conclusion
- **Temperature** (already covered) and **Slope/Angle** (when involving degrees/ratios/DMS) are the only cases requiring a formula.
- **DMS angle conversions are not implemented in the current codebase.**
- All other exotic/derived units in RevitUnits.csv can be converted using only a multiplicative factor, and are implemented.

---
Generated: 2025-11-04
