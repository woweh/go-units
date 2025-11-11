package units

// Temperature is a unit quantity for temperature
const Temperature UnitQuantity = "temperature"

var (
	_temperature = Quantity(Temperature)

	Celsius    = mustCreateNewUnit("celsius", "°C", _temperature, SI, Plural(PluralNone))
	Fahrenheit = mustCreateNewUnit("fahrenheit", "°F", _temperature, US, Plural(PluralNone))
	Kelvin     = mustCreateNewUnit("kelvin", "K", _temperature, SI, Plural(PluralNone))
	Rankine    = mustCreateNewUnit("rankine", "°R", _temperature, US, Plural(PluralNone))
)

func initTemperatureUnits() {
	// https://en.wikipedia.org/wiki/Conversion_of_scales_of_temperature

	NewConversionFromFn(
		Celsius, Kelvin, func(x float64) float64 {
			return x + 273.15
		}, "x + 273.15",
	)
	NewConversionFromFn(
		Kelvin, Celsius, func(x float64) float64 {
			return x - 273.15
		}, "x - 273.15",
	)

	NewConversionFromFn(
		Celsius, Fahrenheit, func(x float64) float64 {
			return x*1.8 + 32
		}, "x * 1.8 + 32",
	)
	NewConversionFromFn(
		Fahrenheit, Celsius, func(x float64) float64 {
			return (x - 32) / 1.8
		}, "(x - 32) / 1.8",
	)

	NewRatioConversion(Kelvin, Rankine, 1.8)

	Celsius.AddAliases("centigrade", "centigrades")
	Fahrenheit.AddAliases("degree Fahrenheit")
	Kelvin.AddAliases("degree Kelvin")
	Rankine.AddAliases("degree Rankine")
	Rankine.AddSymbols("°Ra")
}
