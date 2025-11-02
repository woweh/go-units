package units

var (
	Diffusivity = Quantity("diffusivity")

	// SI base unit: square meter per second
	SquareMeterPerSecond = mustCreateNewUnit("square meter per second", "m²/s", Diffusivity, SI)

	// Imperial/US unit
	SquareFootPerSecond = mustCreateNewUnit("square foot per second", "ft²/s", Diffusivity, BI)
)

func init() {
	// SI base unit: m²/s
	// Conversion: 1 m²/s = 10.7639 ft²/s
	NewRatioConversion(SquareMeterPerSecond, SquareFootPerSecond, 10.763910416709722)

	SquareMeterPerSecond.AddAliases("square meters per second", "square metre per second", "square metres per second", "m2/s")
	SquareFootPerSecond.AddAliases("square feet per second", "ft2/s")
}
