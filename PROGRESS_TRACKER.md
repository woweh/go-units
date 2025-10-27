[ARCHIVED] This checklist has been consolidated. For current tracking, see STATUS.md and plan.md.

# Implementation Progress Tracker

## Phase 1: Fix Existing Quantities - CRITICAL

### pressure_units.go
- [x] Review water column units (FT, in-wg, mH2O, mmH2O)
- [x] Ensure FootH2O has symbol "FT"
- [x] Ensure InchesOfWater has symbol "in-wg"
- [x] Add proper MeterOfWater (reuse existing MH2O if present)
- [x] Add MillimeterOfWater (reuse MilliMH2O if present)
- [x] Test: No duplicate symbols
- [x] Test: Conversions match RevitUnits.json

### electrical_resistance_units.go
- [x] Add Resistivity unit with symbol "ohm·m"
- [x] Define conversion factors
- [x] Add aliases if needed
- [x] Test: Symbol unique

### frequency_units.go
- [x] Add alias "cps" for Hertz
- [x] Add alias "cycles per second"
- [x] Test: Lookup works for all aliases

### length_units.go
- [x] Verify DeciMeter exists (via Deci() prefix)
- [x] Verify USft exists with correct symbol
- [x] Test: All expected symbols present

### power_units.go
- [x] Verify all Revit symbols present
- [x] Check: Btu/h, Btu/s, cal/s, hp, kcal/s, MBH
- [x] Test: No missing symbols

### energy_units.go
- [x] Verify Btu, kcal, therm present
- [x] Test: Conversions accurate

### Validation
- [x] Run: `go test ./...` - All pass
- [x] Check: No duplicate symbols
- [x] Verify: All conversions within ±0.1% of RevitUnits.json

---

## Phase 2: High-Priority New Quantities

### Part A: Lighting (lighting_units.go)
- [ ] Create lighting_units.go with 5 quantities:
  - [ ] Illuminance (lux, ftc)
  - [ ] Luminous Flux (lumen)
  - [ ] Luminous Intensity (candela)
  - [ ] Efficacy (lm/W)
  - [ ] Luminance (cd/m², cd/ft², ftL)
- [ ] Define all units with correct symbols
- [ ] Add conversion factors from RevitUnits.json
- [ ] Create lighting_units_test.go
- [ ] Test: All conversions pass

### Part B: Motion Units
- [ ] Create velocity_units.go
  - [ ] m/s, ft/s, ft/min, mph, km/h, cm/min
  - [ ] Test all conversions
- [ ] Create acceleration_units.go
  - [ ] m/s², ft/s², in/s², km/s², mi/s²
  - [ ] Test all conversions
- [ ] Create angular_speed_units.go
  - [ ] RPM, RPS, rev/min, rev/s, rad/s
  - [ ] Test all conversions

### Part C: Load Units
- [ ] Create cooling_load_units.go
  - [ ] kW, W, Btu/h, Btu/s, ton
  - [ ] Define conversions
  - [ ] Test: All accurate
- [ ] Create heating_load_units.go
  - [ ] kW, W, Btu/h, Btu/s
  - [ ] Define conversions
  - [ ] Test: All accurate

### Part D: Structural Units
- [ ] Create moment_of_inertia_units.go
  - [ ] m⁴, ft⁴, in⁴, cm⁴, mm⁴
  - [ ] Define conversions
  - [ ] Test: All accurate
- [ ] Create stress_units.go (or update pressure_units.go)
  - [ ] Pa, psi, ksi, MPa, kPa
  - [ ] Define conversions
  - [ ] Test: All accurate

### Validation After Phase 2
- [ ] Run: `go test ./...` - All pass
- [ ] Check: No duplicate symbols
- [ ] Verify: All conversions within ±0.1%
- [ ] Count: 8 new quantity files created

---

## Phase 3: Thermal Quantities

- [ ] Create thermal_conductivity_units.go
  - [ ] W/(m·K), Btu/(h·ft·°F)
  - [ ] Define conversions
  - [ ] Test

- [ ] Create specific_heat_units.go
  - [ ] J/(kg·°C), Btu/(lb·°F)
  - [ ] Define conversions
  - [ ] Test

- [ ] Create thermal_mass_units.go
  - [ ] J/K, Btu/°F, kJ/K
  - [ ] Define conversions
  - [ ] Test

- [ ] Create coefficient_of_heat_transfer_units.go
  - [ ] W/(m²·K), Btu/(h·ft²·°F)
  - [ ] Define conversions
  - [ ] Test

- [ ] Create temperature_difference_units.go
  - [ ] K, °C, °F, °R (interval scales)
  - [ ] Handle special conversion (no offset)
  - [ ] Test

### Thermal Validation
- [ ] Run: `go test ./...` - All pass
- [ ] Check: Interval scales work correctly
- [ ] Count: 5 new quantity files created

---

## Phase 4: Remaining Quantities

### Create (26 more files):
- [ ] mass_per_time_units.go (kg/s, kg/min, kg/h, lb/s, lb/min, lb/h)
- [ ] friction_units.go (Pa/m, FT/100ft, in-wg/100ft, mH2O/m, mmH2O/m)
- [ ] moment_units.go (N·m, lb·ft, kip·ft)
- [ ] power_density_units.go (W/m², W/ft²)
- [ ] power_per_length_units.go (W/m, W/ft)
- [ ] apparent_power_density_units.go (VA/m², VA/ft²)
- [ ] cost_per_area_units.go ($/m², $/ft²)
- [ ] currency_units.go (¥, $, €, £, etc.)
- [ ] stationing_units.go (feet, meters, US survey feet)
- [ ] velocity_difference_units.go (if needed)
- [ ] isothermal_moisture_capacity_units.go
- [ ] permeability_units.go
- [ ] specific_heat_of_vaporization_units.go
- [ ] diffusivity_units.go (ft²/s, m²/s)
- [ ] unit_weight_units.go (kN/m³, lbf/ft³)
- [ ] viscosity_units.go (Pa·s, cP)
- [ ] section_modulus_units.go (m³, in³, ft³, cm³)
- [ ] warping_constant_units.go (m⁶, in⁶, ft⁶, cm⁶, mm⁶)
- [ ] weight_per_length_units.go
- [ ] rotational_spring_units.go
- [ ] heat_capacity_per_area_units.go
- [ ] pulsation_units.go (rad/s)
- [ ] cost_rate_energy_units.go
- [ ] cost_rate_power_units.go
- [ ] thermal_expansion_coefficient_units.go
- [ ] (remaining as identified in REVIT_ANALYSIS.md)

### Batch Validation (every 5 files)
- [ ] Run: `go test ./...` - All pass
- [ ] Check: No duplicate symbols
- [ ] Count: Running total of files created

---

## Phase 5: Testing & Validation

### Comprehensive Testing
- [ ] Run: `go test ./...` - ALL pass
- [ ] Count total: 39 new files + 6 updated = 45 changes
- [ ] Verify: All 150 Revit quantities have equivalents

### Symbol Validation
- [ ] Extract all symbols from new files
- [ ] Compare with RevitUnits.json
- [ ] Verify: All 782 units recognized
- [ ] Check: Zero duplicate symbols

### Conversion Accuracy
- [ ] For each quantity, spot-check 3-5 conversions
- [ ] Verify: All within ±0.1% of RevitUnits.json
- [ ] Document: Any deviations and reasons

### Documentation
- [ ] Update plan.md with completion
- [ ] Update STATUS.md with final status
- [ ] Create COMPLETION_REPORT.md
- [ ] Document any deviations from plan
- [ ] List all new quantities and their units

### Final Checklist
- [ ] 150 quantities supported
- [ ] 782 units recognized
- [ ] 0 duplicate symbols
- [ ] 0 test failures
- [ ] ±0.1% conversion accuracy
- [ ] Complete documentation
- [ ] Ready for release

---


## Notes & Issues

### Current Issues
- [ ] (None identified - all tests passing)

### Blocked By
- [ ] (None identified)

### Considerations
- Keep conversion factors from RevitUnits.json
- Validate symbol uniqueness after each phase
- Run full test suite daily during implementation
- Document any assumptions about unit definitions

### Communication
- Update this file regularly
- Document blockers immediately

