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

func initThermalConductivityUnits() {
	NewRatioConversion(BritishThermalUnitPerHourFootFahrenheit, WattPerMeterKelvin, 1.730735)

	WattPerMeterKelvin.AddAliases("watts per meter kelvin", "watt per metre kelvin", "watts per metre kelvin", "W/(m*K)", "W/m/K")
	BritishThermalUnitPerHourFootFahrenheit.AddAliases("British thermal units per hour foot degree Fahrenheit", "BTU/h/ft/F", "Btu/(h·ft·°F)")
}
