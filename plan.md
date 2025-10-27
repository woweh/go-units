# Plan to Enhance go-units Library for Revit Compatibility

## Scope Update

- Cost-related quantities are deferred: Cost Rate Energy, Cost Rate Power, Cost per Area, and Currency — skip for now.
- No-conversion required categories (single-unit in Revit): Color Temperature (K), Efficacy (lm/W), Electrical Resistivity (ohm·m), Luminous Flux (lm), Luminous Intensity (cd), Pulsation (rad/s), Wattage (W). These require quantity/unit presence and symbol alignment but no ratio mappings.
- All other multi-unit quantities require conversions based on RevitUnits.json.

## Executive Summary

**Goal**: Add support for all 150 Revit quantities (782 units) to go-units library

**Current State**: 
- go-units has 22 quantities with ~200 units
- Revit requires 150 quantities with 782 units
- Gap: 39 new quantities needed + updates to 6 existing quantities

**Approach**: 
1. Compare by SYMBOL (most accurate)
2. Fix existing quantities first (add missing units/symbols)
3. Add new quantities systematically by priority
4. Validate conversions against RevitUnits.json

**Timeline**: 6 weeks total (5 implementation phases + 1 testing)

## Current Coverage Analysis

### Existing go-units Quantities (22)
- Angle, Area, Data, Density, Dimensionless Ratio
- Electric Charge, Electric Current, Electrical Resistance
- Energy, Force, Frequency
- Length, Mass
- Power, Pressure
- Slope, Temperature, Thermal Insulance, Time
- Voltage, Volume Flow Rate, Volume

### Units Added During Analysis Phase
- power_units.go: Btu/h, Btu/s, cal/s, hp, kcal/s, MBH
- energy_units.go: Btu, kcal, therm  
- length_units.go: USft (US survey foot)
- mass_units.go: ShortTon
- force_units.go: Kip, ShortTonForce
- pressure_units.go: FT, in-wg, psig (partial)

## Implementation Phases

### Phase 1: Fix Existing Quantities - CRITICAL

**Objective**: Ensure all existing quantities have complete Revit symbol coverage

#### Tasks:
1. **pressure_units.go**
   - Review water column units implementation
   - Ensure proper symbols: FT, in-wg, mH2O, mmH2O
   - Fix any duplicate symbol issues
   - Test: No conflicts, conversions accurate

2. **electrical_resistance_units.go**
   - Add resistivity unit with symbol "ohm·m"
   - Add proper conversions

3. **frequency_units.go**
   - Add alias "cps" → "Hz" (cycles per second)

4. **length_units.go**
   - Verify DeciMeter exists (via Deci prefix)
   - Verify USft is working

5. **Validation**
   - Run: `go test ./...` - All must pass
   - Check: No duplicate symbols
   - Verify: Conversions match RevitUnits.json ±0.1%

### Phase 2: High-Priority New Quantities

**Objective**: Add the 8 most commonly used Revit quantities

#### Part A: Lighting & Motion
1. **lighting_units.go** - Illuminance (lux, ftc), Luminous Flux (lm), Luminous Intensity (cd), Efficacy (lm/W), Luminance (cd/m², cd/ft², ftL)
2. **velocity_units.go** - m/s, ft/s, ft/min, mph, km/h, cm/min
3. **acceleration_units.go** - m/s², ft/s², in/s², km/s², mi/s²
4. **angular_speed_units.go** - RPM, RPS, rev/min, rev/s

#### Part B: Loads & Structural
5. **cooling_load_units.go** - kW, W, Btu/h, Btu/s, ton of refrigeration
6. **heating_load_units.go** - kW, W, Btu/h, Btu/s
7. **moment_of_inertia_units.go** - m⁴, ft⁴, in⁴, cm⁴, mm⁴
8. **stress_units.go** - Pa, psi, ksi, MPa, kPa (can reuse/alias pressure)

**Validation**: Test after each file, run full suite regularly

### Phase 3: Thermal Quantities

**Objective**: Add thermal-related quantities for energy analysis

1. **thermal_conductivity_units.go** - W/(m·K), Btu/(h·ft·°F)
2. **specific_heat_units.go** - J/(kg·°C), Btu/(lb·°F)
3. **thermal_mass_units.go** - J/K, Btu/°F, kJ/K
4. **coefficient_of_heat_transfer_units.go** - W/(m²·K), Btu/(h·ft²·°F)
5. **temperature_difference_units.go** - K, °C, °F, °R (interval scales)

### Phase 4: Remaining Quantities

**Objective**: Add remaining 24 specialized quantities

#### High-Usage Remaining:
- mass_per_time_units.go (kg/s, lb/h, etc.)
- friction_units.go (Pa/m, in-wg/100ft, etc.)
- moment_units.go (torque: N·m, lb·ft, kip·ft)
- power_density_units.go (W/m², W/ft²)
- power_per_length_units.go (W/m, W/ft)
- apparent_power_density_units.go (VA/m², VA/ft²)
- viscosity_units.go (Pa·s, cP)
- diffusivity_units.go (ft²/s, m²/s)
- unit_weight_units.go (kN/m³, lbf/ft³)

#### Lower-Usage Remaining:
- cost_per_area_units.go ($/m², $/ft²)
- currency_units.go (¥, $, €, £)
- stationing_units.go (for infrastructure)
- section_modulus_units.go (m³, in³, ft³, cm³)
- warping_constant_units.go (m⁶, in⁶, ft⁶)
- rotational_spring_units.go
- thermal_expansion_coefficient_units.go
- heat_capacity_per_area_units.go
- isothermal_moisture_capacity_units.go
- permeability_units.go
- specific_heat_of_vaporization_units.go
- pulsation_units.go
- cost_rate_energy_units.go
- cost_rate_power_units.go
- weight_per_length_units.go

**Note**: Some quantities may be combined into single files (e.g., all cost-related in one file)

### Phase 5: Testing & Validation

**Objective**: Comprehensive validation and documentation

#### Validation Steps:
1. Run full test suite: `go test ./...`
2. Verify all 150 Revit quantities have equivalents
3. Check all 782 unit symbols are recognized (directly or via aliases)
4. Validate conversions against RevitUnits.json (±0.1% tolerance)
5. Ensure zero duplicate symbols
6. Performance testing for lookup speed

#### Documentation:
1. Update README.md with Revit compatibility
2. Document all new quantities
3. Create conversion reference tables
4. List any known limitations or assumptions

## Success Criteria

- [ ] All 150 Revit quantities supported
- [ ] All 782 unit symbols recognized
- [ ] Conversion accuracy ±0.1% vs RevitUnits.json
- [ ] Zero duplicate symbols
- [ ] All tests passing
- [ ] Complete documentation

## Implementation Guidelines

### For Each New Quantity File:

1. **Create `{quantity}_units.go`**
   - Define Quantity variable
   - Add all base units with exact Revit symbols
   - Use SI prefixes (Kilo, Milli, etc.) where applicable
   - Define conversions using RevitUnits.json ConversionFactor values

2. **Create `{quantity}_units_test.go`**
   - Test basic conversions
   - Validate against known values
   - Test edge cases

3. **Validation**
   - Run: `go test ./...`
   - Check: No duplicate symbols
   - Verify: Conversions accurate

### Symbol Comparison Strategy:

- Primary: Match by UnitSymbol from RevitUnits.json
- Secondary: Use AddAliases() for variant names
- Tertiary: Use AddSymbols() for alternate symbols
- Priority: Symbol uniqueness over name matching

### Conversion Factor Usage:

From RevitUnits.json:
```
"ConversionFactor": X means 1 base_unit = X * this_unit
```

In go-units:
```go
NewRatioConversion(ThisUnit, BaseUnit, X)
```

## Risk Mitigation

### Potential Issues:
1. **Duplicate Symbols**: Some Revit symbols may conflict
   - Solution: Use most common meaning, add aliases for alternates

2. **Interval vs Absolute Scales**: Temperature difference vs absolute temperature
   - Solution: Create separate quantities for intervals where needed

3. **Complex Conversions**: Some units may need formula-based conversions
   - Solution: Use NewConversionFromFn() for non-linear conversions

4. **Performance**: 782 units may impact lookup speed
   - Solution: Test and optimize if needed, current implementation should be fine

## Implementation Progress

Work through phases sequentially, testing after each file.

**Target Completion**: TBD

## Next Steps

1. Begin Phase 1 with pressure_units.go fixes
2. Follow weekly schedule in phases above
3. Test incrementally after each file
4. Update this plan as issues arise
5. Document any deviations or assumptions

## References

- **RevitUnits.json**: Source of truth for all unit definitions and conversions
- **Existing _units.go files**: Code patterns and structure to follow
- **REVIT_ANALYSIS.md**: Detailed breakdown of all 150 quantities

