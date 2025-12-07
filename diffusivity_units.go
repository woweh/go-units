package units

var (
	// Diffusivity is the unit quantity for diffusivity.
	Diffusivity = NewUnitQuantity("diffusivity")

	// SI base unit: square meter per second
	SquareMeterPerSecond = Diffusivity.MustCreateUnit("square meter per second", "m²/s", SI)

	// Imperial/US unit
	SquareFootPerSecond = Diffusivity.MustCreateUnit("square foot per second", "ft²/s", BI)
)

func initDiffusivityUnits() {
	// SI base unit: m²/s
	// Calculated: 1 m²/s = X ft²/s
	NewRatioConversion(SquareMeterPerSecond, SquareFootPerSecond, 1.0/diffusivityFactor(Foot, Second))

	SquareMeterPerSecond.AddAliases("square meters per second", "square metre per second", "square metres per second", "m2/s")
	SquareFootPerSecond.AddAliases("square feet per second", "ft2/s")
}

// diffusivityFactor calculates the conversion factor for diffusivity units.
// Diffusivity = area / time = length² / time
// Base unit: SquareMeterPerSecond = m²/s
//
// Example: ft²/s
// - areaRatio: how many m² in 1 ft²
// - timeRatio: how many s in 1 s
// - diffusivityFactor = areaRatio / timeRatio
func diffusivityFactor(length, time Unit) float64 {
	// How many m² in 1 unit of target area (length²)
	areaRatio := areaFactor(length)
	// How many seconds in 1 unit of target time
	timeRatio := time.to(Second)
	// Diffusivity factor: area per time
	return areaRatio / timeRatio
}
