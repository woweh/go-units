package units

var (
	CoefficientOfHeatTransfer = Quantity("coefficient of heat transfer")

	// Base unit: watt per square meter kelvin (Revit base with CF=1.0)
	WattPerSquareMeterKelvin = mustCreateNewUnit("watt per square meter kelvin", "W/(m²·K)", CoefficientOfHeatTransfer, SI)

	// Imperial unit
	BritishThermalUnitPerHourSquareFootFahrenheit = mustCreateNewUnit("British thermal unit per hour square foot degree Fahrenheit", "BTU/(h·ft²·°F)", CoefficientOfHeatTransfer)
)

func init() {
	// From RevitUnits.json:
	// W/(m²·K): CF = 1.0
	// BTU/(h·ft²·°F): CF = 0.17611018368230585
	// This means: 1 W/(m²·K) = 0.176110 BTU/(h·ft²·°F)
	// Or: 1 BTU/(h·ft²·°F) = 5.67826 W/(m²·K)
	NewRatioConversion(BritishThermalUnitPerHourSquareFootFahrenheit, WattPerSquareMeterKelvin, 5.67826)

	WattPerSquareMeterKelvin.AddAliases("watts per square meter kelvin", "watt per square metre kelvin", "watts per square metre kelvin", "W/(m2*K)", "W/m2/K")
	BritishThermalUnitPerHourSquareFootFahrenheit.AddAliases("British thermal units per hour square foot degree Fahrenheit", "BTU/h/ft2/F", "Btu/(h·ft²·°F)")
}
