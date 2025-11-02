package units

var (
	Temperature = Quantity("temperature")

	Celsius    = mustCreateNewUnit("celsius", "°C", Temperature, Plural(PluralNone), SI)
	Fahrenheit = mustCreateNewUnit("fahrenheit", "°F", Temperature, Plural(PluralNone), US)
	Kelvin     = mustCreateNewUnit("kelvin", "°K", Temperature, Plural(PluralNone), SI)
	Rankine    = mustCreateNewUnit("rankine", "°R", Temperature, Plural(PluralNone), US)
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
