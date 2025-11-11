package units

const CoefficientOfHeatTransfer UnitQuantity = "coefficient of heat transfer"

var (
	_coefficientOfHeatTransfer = Quantity(CoefficientOfHeatTransfer)

	// Base unit: watt per square meter kelvin (Revit base with CF=1.0)
	WattPerSquareMeterKelvin = mustCreateNewUnit("watt per square meter kelvin", "W/(m²·K)", _coefficientOfHeatTransfer, SI)

	// Imperial unit
	BritishThermalUnitPerHourSquareFootFahrenheit = mustCreateNewUnit("British thermal unit per hour square foot degree Fahrenheit", "BTU/(h·ft²·°F)", _coefficientOfHeatTransfer, BI)
)

func initCoefficientOfHeatTransferUnits() {
	NewRatioConversion(BritishThermalUnitPerHourSquareFootFahrenheit, WattPerSquareMeterKelvin, 5.6782633411134871)

	WattPerSquareMeterKelvin.AddAliases("watts per square meter kelvin", "watt per square metre kelvin", "watts per square metre kelvin", "W/(m2*K)", "W/m2/K")
	BritishThermalUnitPerHourSquareFootFahrenheit.AddAliases("British thermal units per hour square foot degree Fahrenheit", "BTU/h/ft2/F", "Btu/(h·ft²·°F)")
}
