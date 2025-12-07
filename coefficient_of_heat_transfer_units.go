package units

var (
	// CoefficientOfHeatTransfer is a unit quantity for coefficient of heat transfer
	CoefficientOfHeatTransfer = NewUnitQuantity("coefficient of heat transfer")

	// Base unit: watt per square meter kelvin (Revit base with CF=1.0)
	WattPerSquareMeterKelvin = CoefficientOfHeatTransfer.MustCreateUnit("watt per square meter kelvin", "W/(m²·K)", SI)

	// Imperial unit
	BritishThermalUnitPerHourSquareFootFahrenheit = CoefficientOfHeatTransfer.MustCreateUnit("British thermal unit per hour square foot degree Fahrenheit", "BTU/(h·ft²·°F)", BI)
)

func initCoefficientOfHeatTransferUnits() {
	NewRatioConversion(BritishThermalUnitPerHourSquareFootFahrenheit, WattPerSquareMeterKelvin, 5.6782633411134871)

	WattPerSquareMeterKelvin.AddAliases("watts per square meter kelvin", "watt per square metre kelvin", "watts per square metre kelvin", "W/(m2*K)", "W/m2/K")
	BritishThermalUnitPerHourSquareFootFahrenheit.AddAliases("British thermal units per hour square foot degree Fahrenheit", "BTU/h/ft2/F", "Btu/(h·ft²·°F)")
}
