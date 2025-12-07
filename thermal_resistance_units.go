package units

// ThermalResistance is a unit quantity for thermal resistance
var (
	// ThermalResistance is the unit quantity for thermal resistance.
	ThermalResistance = NewUnitQuantity("thermal resistance")
	// SI base unit: square meter kelvin per watt (m²·K)/W
	SquareMeterKelvinPerWatt = ThermalResistance.MustCreateUnit("square meter kelvin per watt", "(m²·K)/W", SI)

	// Imperial unit: hour square foot degree Fahrenheit per British thermal unit
	HourSquareFootFahrenheitPerBritishThermalUnit = ThermalResistance.MustCreateUnit("hour square foot degree Fahrenheit per British thermal unit", "(h·ft²·°F)/BTU", BI)
)

func initThermalResistanceUnits() {
	NewRatioConversion(SquareMeterKelvinPerWatt, HourSquareFootFahrenheitPerBritishThermalUnit, 5.678263341113486)

	SquareMeterKelvinPerWatt.AddAliases("square meter kelvins per watt", "square metre kelvin per watt", "square metre kelvins per watt", "(m2*K)/W", "m2*K/W")
	HourSquareFootFahrenheitPerBritishThermalUnit.AddAliases("hour square foot degrees Fahrenheit per British thermal unit", "h*ft2*F/BTU", "(h·ft²·°F)/Btu")
}
