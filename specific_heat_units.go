package units

var (
	SpecificHeat = Quantity("specific heat")

	// Base unit: joule per kilogram degree Celsius (Revit base with CF=0.09290304)
	JoulePerKilogramCelsius = mustCreateNewUnit("joule per kilogram degree Celsius", "J/(kg·°C)", SpecificHeat, SI)

	// Other units
	JoulePerGramCelsius                  = mustCreateNewUnit("joule per gram degree Celsius", "J/(g·°C)", SpecificHeat, SI)
	BritishThermalUnitPerPoundFahrenheit = mustCreateNewUnit("British thermal unit per pound degree Fahrenheit", "BTU/(lb·°F)", SpecificHeat)
)

func init() {
	// From RevitUnits.json:
	// J/(kg·°C): CF = 0.09290304
	// J/(g·°C): CF = 9.290304e-05
	// BTU/(lb·°F): CF = 2.2189509888220124e-05

	// 1 J/(g·°C) = 1000 J/(kg·°C)
	NewRatioConversion(JoulePerGramCelsius, JoulePerKilogramCelsius, 1000.0)

	// From the ratios: 0.09290304 / 2.2189509888220124e-05 = 4186.8
	// This means: 1 J/(kg·°C) = 0.0002388 BTU/(lb·°F)
	// Or: 1 BTU/(lb·°F) = 4186.8 J/(kg·°C)
	NewRatioConversion(BritishThermalUnitPerPoundFahrenheit, JoulePerKilogramCelsius, 4186.8)

	JoulePerKilogramCelsius.AddAliases("joules per kilogram degree Celsius", "J/kg/C", "J/(kg*°C)")
	JoulePerGramCelsius.AddAliases("joules per gram degree Celsius", "J/g/C", "J/(g*°C)")
	BritishThermalUnitPerPoundFahrenheit.AddAliases("British thermal units per pound degree Fahrenheit", "BTU/lb/F", "Btu/(lb·°F)")
}
