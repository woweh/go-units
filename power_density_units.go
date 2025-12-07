package units

var (
	// PowerDensity is the unit quantity for power density.
	PowerDensity = NewUnitQuantity("power density")

	// SI base unit: watt per square meter
	WattPerSquareMeter = PowerDensity.MustCreateUnit("watt per square meter", "W/m²", SI)

	// Imperial/US units
	WattPerSquareFoot                   = PowerDensity.MustCreateUnit("watt per square foot", "W/ft²", BI)
	BritishThermalUnitPerHourSquareFoot = PowerDensity.MustCreateUnit("British thermal unit per hour square foot", "Btu/(h·ft²)", BI)
)

func initPowerDensityUnits() {
	// Derive from area units (power/area)
	NewRatioConversion(WattPerSquareMeter, WattPerSquareFoot, powerDensityFactor(Foot))

	// 1 W/ft² = 3.41214 Btu/(h·ft²)
	NewRatioConversion(WattPerSquareFoot, BritishThermalUnitPerHourSquareFoot, 3.41214)

	WattPerSquareMeter.AddAliases("watts per square meter", "watt per square metre", "watts per square metre", "W/m2")
	WattPerSquareFoot.AddAliases("watts per square foot", "W/ft2")
	BritishThermalUnitPerHourSquareFoot.AddAliases("British thermal units per hour square foot", "BTU/h/ft²", "BTU/(h·ft²)")
}

// powerDensityFactor computes the power density conversion factor from area units.
// power density = power / area, for same power unit: 1 / area
func powerDensityFactor(areaLength Unit) float64 {
	return areaFactor(areaLength)
}
