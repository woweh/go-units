package units

var (
	PowerDensity = Quantity("power density")

	// Base unit: watt per square meter (Revit base with CF=1.0)
	WattPerSquareMeter = mustCreateNewUnit("watt per square meter", "W/m²", PowerDensity, SI)

	// Other units
	WattPerSquareFoot                   = mustCreateNewUnit("watt per square foot", "W/ft²", PowerDensity)
	BritishThermalUnitPerHourSquareFoot = mustCreateNewUnit("British thermal unit per hour square foot", "Btu/(h·ft²)", PowerDensity)
)

func init() {
	// From RevitUnits.json:
	// W/m²: CF = 1.0
	// W/ft²: CF = 0.09290304
	// Btu/(h·ft²): CF = 0.3169983306281505

	// 1 W/m² = 0.09290304 W/ft²
	NewRatioConversion(WattPerSquareMeter, WattPerSquareFoot, 0.09290304)

	// From ratios: 0.3169983306281505 / 0.09290304 = 3.41214
	// 1 W/ft² = 3.41214 Btu/(h·ft²)
	NewRatioConversion(WattPerSquareFoot, BritishThermalUnitPerHourSquareFoot, 3.41214)

	WattPerSquareMeter.AddAliases("watts per square meter", "watt per square metre", "watts per square metre", "W/m2")
	WattPerSquareFoot.AddAliases("watts per square foot", "W/ft2")
	BritishThermalUnitPerHourSquareFoot.AddAliases("British thermal units per hour square foot", "BTU/h/ft²", "BTU/(h·ft²)")
}
