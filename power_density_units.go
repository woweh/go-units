package units

// PowerDensity is a unit quantity for power density
const PowerDensity UnitQuantity = "power density"

var (
	_powerDensity = Quantity(PowerDensity)

	// SI base unit: watt per square meter
	WattPerSquareMeter = mustCreateNewUnit("watt per square meter", "W/m²", _powerDensity, SI)

	// Imperial/US units
	WattPerSquareFoot                   = mustCreateNewUnit("watt per square foot", "W/ft²", _powerDensity, BI)
	BritishThermalUnitPerHourSquareFoot = mustCreateNewUnit("British thermal unit per hour square foot", "Btu/(h·ft²)", _powerDensity, BI)
)

func initPowerDensityUnits() {
	// SI base unit: W/m²
	// Conversion: 1 W/m² = 0.09290304 W/ft²
	NewRatioConversion(WattPerSquareMeter, WattPerSquareFoot, 0.09290304)

	// 1 W/ft² = 3.41214 Btu/(h·ft²)
	NewRatioConversion(WattPerSquareFoot, BritishThermalUnitPerHourSquareFoot, 3.41214)

	WattPerSquareMeter.AddAliases("watts per square meter", "watt per square metre", "watts per square metre", "W/m2")
	WattPerSquareFoot.AddAliases("watts per square foot", "W/ft2")
	BritishThermalUnitPerHourSquareFoot.AddAliases("British thermal units per hour square foot", "BTU/h/ft²", "BTU/(h·ft²)")
}
