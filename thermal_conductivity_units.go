package units

// ThermalConductivity is a unit quantity for thermal conductivity
const ThermalConductivity UnitQuantity = "thermal conductivity"

var (
	_thermalConductivity = Quantity(ThermalConductivity)

	// SI base unit: watt per meter kelvin (W/(m·K))
	WattPerMeterKelvin = mustCreateNewUnit("watt per meter kelvin", "W/(m·K)", _thermalConductivity, SI)

	// Imperial unit: BTU/(h·ft·°F)
	BritishThermalUnitPerHourFootFahrenheit = mustCreateNewUnit("British thermal unit per hour foot degree Fahrenheit", "BTU/(h·ft·°F)", _thermalConductivity, BI)
)

func init() {
	// From RevitUnits.json:
	// W/(m·K): CF = 0.3048
	// BTU/(h·ft·°F): CF = 0.17611018368230585
	// Ratio: 0.3048 / 0.17611018368230585 = 1.730735
	// This means: 1 W/(m·K) = 0.57779 BTU/(h·ft·°F)
	// Or: 1 BTU/(h·ft·°F) = 1.730735 W/(m·K)
	NewRatioConversion(BritishThermalUnitPerHourFootFahrenheit, WattPerMeterKelvin, 1.730735)

	WattPerMeterKelvin.AddAliases("watts per meter kelvin", "watt per metre kelvin", "watts per metre kelvin", "W/(m*K)", "W/m/K")
	BritishThermalUnitPerHourFootFahrenheit.AddAliases("British thermal units per hour foot degree Fahrenheit", "BTU/h/ft/F", "Btu/(h·ft·°F)")
}
