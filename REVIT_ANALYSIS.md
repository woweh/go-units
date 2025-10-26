# Detailed Analysis: Revit Units vs go-units

## Revit Quantities Found (150 total)

### Common/Basic Quantities (COMPLETE)
- ✅ Acceleration (m/s², ft/s²) - MISSING: create acceleration_units.go
- ✅ Angle (°, rad, grad) - PRESENT
- ✅ Area (m², ft², acres, hectares) - PRESENT
- ✅ Currency (¥) - MISSING: create currency_units.go
- ✅ Data/Bytes (B, KB, MB, GB, TB, PB) - PRESENT
- ✅ Density (kg/m³, lb/ft³) - PRESENT
- ✅ Distance/Length (m, ft, in, cm, mm, km) - PRESENT (add USft, dm)
- ✅ Time (s, min, h, ms, d, yr) - PRESENT
- ✅ Volume (m³, ft³, L, gal) - PRESENT
- ✅ Mass (kg, lb, t, Tons) - PRESENT
- ✅ Temperature (°C, °F, K, °R) - PRESENT

### Electrical Quantities (MOSTLY PRESENT)
- ✅ Amperes (A, kA, mA) - PRESENT
- ✅ Apparent Power (VA, kVA, MVA) - PRESENT (in power_units)
- ✅ Apparent Power Density (VA/ft², VA/m²) - MISSING
- ✅ Cable Tray Size (length units) - PRESENT
- ✅ Color Temperature (K) - PRESENT (temperature)
- ✅ Conduit Size (length units) - PRESENT
- ✅ Demand Factor (%, fixed) - PRESENT (dimensionless)
- ✅ Efficacy (lm/W) - MISSING: create lighting_units.go
- ✅ Electric Charge (A·h, A·s, A·min, C) - PRESENT
- ✅ Electrical Potential (V, kV, mV) - PRESENT
- ✅ Electrical Resistivity (ohm·m) - MISSING
- ✅ Frequency (Hz, cycles/s, rev/min, deg/s) - PRESENT
- ✅ Illuminance (lux, ftc) - MISSING
- ✅ Luminance (cd/m², cd/ft², ftL) - MISSING
- ✅ Luminous Flux (lm) - MISSING
- ✅ Luminous Intensity (cd) - MISSING
- ✅ Power (W, kW, Btu/h, hp, kVA, cal/s, kcal/s, MBH) - PRESENT (need to verify all)
- ✅ Power Density (W/m², W/ft²) - MISSING
- ✅ Power per Length (W/m, W/ft) - MISSING
- ✅ Voltage (V, kV, mV) - PRESENT
- ✅ Wattage (W) - PRESENT

### Energy & Thermal Quantities (PARTIAL)
- ✅ Coefficient of Heat Transfer (W/(m²·K), Btu/(h·ft²·°F)) - MISSING
- ✅ Demand Factor (%) - PRESENT
- ✅ Energy (J, kJ, Btu, cal, kWh, therm) - PRESENT
- ✅ Heat Capacity per Area - MISSING
- ✅ Heat Transfer Coefficient - MISSING
- ✅ Isothermal Moisture Capacity - MISSING
- ✅ Permeability - MISSING
- ✅ Specific Heat (J/(kg·°C), Btu/(lb·°F)) - MISSING
- ✅ Specific Heat of Vaporization - MISSING
- ✅ Thermal Conductivity (W/(m·K), Btu/(h·ft·°F)) - MISSING
- ✅ Thermal Expansion Coefficient - MISSING
- ✅ Thermal Insulance (°F·h·ft²/Btu, K·m²/W) - PRESENT (check both symbols)
- ✅ Thermal Mass (J/K, Btu/°F) - MISSING
- ✅ Thermal Resistance - MISSING

### HVAC/Flow Quantities (MOSTLY MISSING)
- ✅ Air Flow (CFM, m³/h, L/min) - PRESENT (volume flow rate)
- ✅ Air Flow Density (CFM/ft², LPS/m²) - MISSING
- ✅ Air Flow divided by Cooling Load (CFM/ton, L/(s·kW)) - MISSING
- ✅ Air Flow divided by Volume (CFM/ft³) - MISSING
- ✅ Angular Speed (RPM, RPS) - MISSING
- ✅ Area divided by Cooling Load (m²/kW, SF/MBh, SF/ton) - MISSING
- ✅ Area divided by Heating Load - MISSING
- ✅ Cooling Load (kW, W, Btu/h, ton) - MISSING
- ✅ Cooling Load divided by Area (W/m², Btu/(h·ft²)) - MISSING
- ✅ Cooling Load divided by Volume (W/m³, Btu/(h·ft³)) - MISSING
- ✅ Cross Section (m², ft², cm²) - PRESENT (as area)
- ✅ Density (kg/m³, lb/ft³) - PRESENT
- ✅ Diffusivity (ft²/s, m²/s) - MISSING
- ✅ Duct Insulation Thickness - PRESENT (as length)
- ✅ Duct Size - PRESENT (as length)
- ✅ Factor (%, fixed) - PRESENT
- ✅ Flow per Power - MISSING
- ✅ Friction (Pa/m, in-wg/100ft) - MISSING
- ✅ Heat Gain (W, Btu/h) - MISSING
- ✅ Heating Load (W, kW, Btu/h) - MISSING
- ✅ Mass per Time (kg/s, lb/h) - MISSING
- ✅ Power (W, Btu/h, kW) - PRESENT
- ✅ Power Density - MISSING
- ✅ Power per Flow - MISSING
- ✅ Pressure (Pa, psi, Torr, in-wg, bar, kPa) - PRESENT
- ✅ Roughness - PRESENT (as length)
- ✅ Slope - PRESENT
- ✅ Temperature - PRESENT
- ✅ Temperature Difference - MISSING (interval scales)
- ✅ Velocity/Speed (ft/s, m/s, mph, km/h) - MISSING
- ✅ Viscosity (Pa·s, cP) - MISSING

### Piping Quantities (MOSTLY MISSING)
- Similar to HVAC, mostly reuse existing quantity types

### Structural/Mechanical Quantities (MOSTLY MISSING)
- ✅ Acceleration - MISSING
- ✅ Force (N, kN, lbf, kip, tf, Tons) - PRESENT (add kip alias)
- ✅ Moment/Torque (N·m, lb·ft, kip·ft) - MISSING
- ✅ Moment of Inertia (m⁴, in⁴, ft⁴, mm⁴) - MISSING
- ✅ Moment Scale - MISSING
- ✅ Period (s, ms, min, h) - PRESENT (as time)
- ✅ Pulsation (rad/s) - MISSING
- ✅ Rotation (°, rad, grad) - PRESENT (as angle)
- ✅ Rotational Spring - MISSING
- ✅ Section Area (m², cm², in², ft²) - PRESENT (as area)
- ✅ Section Modulus - MISSING
- ✅ Stress (Pa, psi, ksi, MPa) - MISSING (reuse pressure)
- ✅ Unit Weight (kN/m³, lbf/ft³) - MISSING
- ✅ Velocity/Speed - MISSING
- ✅ Warping Constant - MISSING
- ✅ Weight per Length - MISSING

### Infrastructure Quantities (MISSING)
- ✅ Stationing - MISSING

### Common Quantities (PARTIAL)
- ✅ Cost per Area - MISSING
- ✅ Cost Rate Energy - MISSING
- ✅ Cost Rate Power - MISSING
- ✅ Dimensionless Ratio (%, ‰, fixed) - PRESENT
- ✅ Number (currency, %, fixed, general) - PARTIAL

## Summary of Missing/Incomplete

### NEW QUANTITIES NEEDED (create new files):
1. acceleration_units.go
2. angular_speed_units.go  
3. apparent_power_density_units.go
4. coefficient_of_heat_transfer_units.go
5. cooling_load_units.go
6. cost_per_area_units.go
7. currency_units.go
8. diffusivity_units.go
9. efficacy_units.go (lumens per watt)
10. friction_units.go
11. heat_capacity_per_area_units.go
12. heat_transfer_coefficient_units.go
13. heating_load_units.go
14. illuminance_units.go
15. isothermal_moisture_capacity_units.go
16. lighting_units.go (candela, lumens, etc.)
17. mass_per_time_units.go
18. moment_of_inertia_units.go
19. permeability_units.go
20. power_density_units.go
21. power_per_flow_units.go
22. power_per_length_units.go
23. resistivity_units.go (electrical)
24. rotational_spring_units.go
25. section_modulus_units.go
26. specific_heat_units.go
27. specific_heat_of_vaporization_units.go
28. stationing_units.go
29. stress_units.go
30. temperature_difference_units.go
31. thermal_conductivity_units.go
32. thermal_expansion_coefficient_units.go
33. thermal_mass_units.go
34. thermal_resistance_units.go
35. unit_weight_units.go
36. velocity_units.go
37. viscosity_units.go
38. warping_constant_units.go
39. weight_per_length_units.go

### EXISTING QUANTITY UPDATES NEEDED:
1. **length_units.go**: Add decimeter (dm) if missing
2. **pressure_units.go**: Add feet of water and water column units properly
3. **thermal_insulance_units.go**: Check both symbol formats
4. **frequency_units.go**: Add cps alias for cycles per second
5. **force_units.go**: Add kip symbol alias
6. **voltage_units.go**: Add any missing symbols
7. **electrical_resistance_units.go**: Add resistivity (ohm·m)
8. **power_units.go**: Verify all symbols present
9. **energy_units.go**: Verify all symbols present

## Reference

For implementation strategy, timeline, and phases, see **plan.md**.

This document provides the detailed technical requirements only.
