package units

// SpecificHeat is a unit quantity for specific heat
var (
	// SpecificHeat is the unit quantity for specific heat.
	SpecificHeat = NewUnitQuantity("specific heat")
	// SI base unit: joule per kilogram degree Celsius
	JoulePerKilogramCelsius = SpecificHeat.MustCreateUnit("joule per kilogram degree Celsius", "J/(kg·°C)", SI)

	// Other SI unit
	JoulePerGramCelsius = SpecificHeat.MustCreateUnit("joule per gram degree Celsius", "J/(g·°C)", SI)

	// Imperial unit
	BritishThermalUnitPerPoundFahrenheit = SpecificHeat.MustCreateUnit("British thermal unit per pound degree Fahrenheit", "BTU/(lb·°F)", BI)
)

func initSpecificHeatUnits() {
	// 1 J/(g·°C) = 1000 J/(kg·°C)
	NewRatioConversion(JoulePerGramCelsius, JoulePerKilogramCelsius, 1000.0)

	// 1 BTU/(lb·°F) = 4186.8 J/(kg·°C)
	NewRatioConversion(BritishThermalUnitPerPoundFahrenheit, JoulePerKilogramCelsius, 4186.8)

	JoulePerKilogramCelsius.AddAliases("joules per kilogram degree Celsius", "J/kg/C", "J/(kg*°C)")
	JoulePerGramCelsius.AddAliases("joules per gram degree Celsius", "J/g/C", "J/(g*°C)")
	BritishThermalUnitPerPoundFahrenheit.AddAliases("British thermal units per pound degree Fahrenheit", "BTU/lb/F", "Btu/(lb·°F)")
}
