package units

var (
	// Temperature is the unit quantity for temperature.
	Temperature = NewUnitQuantity("temperature")

	Celsius    = Temperature.MustCreateUnit("celsius", "°C", SI, Plural(PluralNone))
	Fahrenheit = Temperature.MustCreateUnit("fahrenheit", "°F", US, Plural(PluralNone))
	Kelvin     = Temperature.MustCreateUnit("kelvin", "K", SI, Plural(PluralNone))
	Rankine    = Temperature.MustCreateUnit("rankine", "°R", US, Plural(PluralNone))
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
