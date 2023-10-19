package units

var (
	Temperature = UnitOptionQuantity("temperature")

	Celsius    = newUnit("celsius", "C", Temperature, UnitOptionPlural(None, ""), SI)
	Fahrenheit = newUnit("fahrenheit", "F", Temperature, UnitOptionPlural(None, ""), US)
	Kelvin     = newUnit("kelvin", "K", Temperature, UnitOptionPlural(None, ""), SI)
	Rankine    = newUnit("rankine", "R", Temperature, UnitOptionPlural(None, ""), US)
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
	Celsius.AddSymbols("°C")

	Fahrenheit.AddAliases("degree Fahrenheit")
	Fahrenheit.AddSymbols("°F")

	Kelvin.AddAliases("degree Kelvin")

	Rankine.AddAliases("degree Rankine")
	Rankine.AddSymbols("°R", "°Ra")
}
