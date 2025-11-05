package units

// ThermalResistance is a unit quantity for thermal resistance
const ThermalResistance UnitQuantity = "thermal resistance"

var (
	_thermalResistance = Quantity(ThermalResistance)

	// SI base unit: square meter kelvin per watt (m²·K)/W
	SquareMeterKelvinPerWatt = mustCreateNewUnit("square meter kelvin per watt", "(m²·K)/W", _thermalResistance, SI)

	// Imperial unit: hour square foot degree Fahrenheit per British thermal unit
	HourSquareFootFahrenheitPerBritishThermalUnit = mustCreateNewUnit("hour square foot degree Fahrenheit per British thermal unit", "(h·ft²·°F)/BTU", _thermalResistance, BI)
)

func init() {
	// From RevitUnits.json:
	// (m²·K)/W: CF = 1.0
	// (h·ft²·°F)/BTU: CF = 5.678263341113486
	// This means: 1 (m²·K)/W = 5.678263 (h·ft²·°F)/BTU
	NewRatioConversion(SquareMeterKelvinPerWatt, HourSquareFootFahrenheitPerBritishThermalUnit, 5.678263341113486)

	SquareMeterKelvinPerWatt.AddAliases("square meter kelvins per watt", "square metre kelvin per watt", "square metre kelvins per watt", "(m2*K)/W", "m2*K/W")
	HourSquareFootFahrenheitPerBritishThermalUnit.AddAliases("hour square foot degrees Fahrenheit per British thermal unit", "h*ft2*F/BTU", "(h·ft²·°F)/Btu")
}
