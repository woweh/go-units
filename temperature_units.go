package units

var (
	Temperature = Quantity("temperature")

	Celsius    = newUnit("celsius", "°C", Temperature, Plural(PluralNone), SI)
	Fahrenheit = newUnit("fahrenheit", "°F", Temperature, Plural(PluralNone), US)
	Kelvin     = newUnit("kelvin", "°K", Temperature, Plural(PluralNone), SI)
	Rankine    = newUnit("rankine", "°R", Temperature, Plural(PluralNone), US)
)

func init() {
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
