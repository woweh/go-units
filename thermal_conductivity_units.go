package units

// ThermalConductivity is a unit quantity for thermal conductivity
var (
	// ThermalConductivity is the unit quantity for thermal conductivity.
	ThermalConductivity = NewUnitQuantity("thermal conductivity")
	// SI base unit: watt per meter kelvin (W/(m·K))
	WattPerMeterKelvin = ThermalConductivity.MustCreateUnit("watt per meter kelvin", "W/(m·K)", SI)

	// Imperial unit: BTU/(h·ft·°F)
	BritishThermalUnitPerHourFootFahrenheit = ThermalConductivity.MustCreateUnit("British thermal unit per hour foot degree Fahrenheit", "BTU/(h·ft·°F)", BI)
)

func initThermalConductivityUnits() {
	NewRatioConversion(BritishThermalUnitPerHourFootFahrenheit, WattPerMeterKelvin, 1.730734666)

	WattPerMeterKelvin.AddAliases("watts per meter kelvin", "watt per metre kelvin", "watts per metre kelvin", "W/(m*K)", "W/m/K")
	BritishThermalUnitPerHourFootFahrenheit.AddAliases("British thermal units per hour foot degree Fahrenheit", "BTU/h/ft/F", "Btu/(h·ft·°F)")
}
