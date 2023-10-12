package units

var (
	Temp = UnitOptionQuantity("temperature")

	Celsius    = newUnit("celsius", "C", Temp, UnitOptionPlural(None, ""), SI)
	Fahrenheit = newUnit("fahrenheit", "F", Temp, UnitOptionPlural(None, ""), US)
	Kelvin     = newUnit("kelvin", "K", Temp, UnitOptionPlural(None, ""), SI)
)

func init() {
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
}
