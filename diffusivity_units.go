package units

var (
	Diffusivity = Quantity("diffusivity")

	// Base unit: square foot per second (Revit base with CF=1.0)
	SquareFootPerSecond = newUnit("square foot per second", "ft²/s", Diffusivity, BI)

	// SI unit
	SquareMeterPerSecond = newUnit("square meter per second", "m²/s", Diffusivity, SI)
)

func init() {
	// From RevitUnits.json:
	// ft²/s: CF = 1.0
	// m²/s: CF = 0.09290304
	// This means: 1 ft²/s = 0.09290304 m²/s
	NewRatioConversion(SquareFootPerSecond, SquareMeterPerSecond, 0.09290304)

	SquareFootPerSecond.AddAliases("square feet per second", "ft2/s")
	SquareMeterPerSecond.AddAliases("square meters per second", "square metre per second", "square metres per second", "m2/s")
}
