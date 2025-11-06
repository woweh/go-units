package units

import "math"

// https://en.wikipedia.org/wiki/Slope
// https://en.wikipedia.org/wiki/Grade_(slope)

// Slope is a unit quantity for slope (gradient)
const Slope UnitQuantity = "slope"

var (
	_slope = Quantity(Slope)

	// SlopeValue is the base unit: rise/run (dimensionless, SI system)
	SlopeValue = mustCreateNewUnit("slope value", "", _slope, SI, Plural(PluralNone))
	// SlopeRatio is a ratio of one part rise to so many parts run (e.g. 1:10), dimensionless, SI system
	SlopeRatio = mustCreateNewUnit("slope ratio", "", _slope, SI, Plural(PluralNone))
	// SlopeInverseRatio is a ratio of many parts run to one part rise (e.g. 10:1), dimensionless, SI system
	SlopeInverseRatio = mustCreateNewUnit("inverse slope ratio", "", _slope, SI, Plural(PluralNone))
	// SlopeDegree is the angle of inclination in degrees (e.g. 45°), SI system
	SlopeDegree = mustCreateNewUnit("slope degree", "deg", _slope, SI, Plural(PluralNone))
	// SlopePercent 100 * m = 100 * (rise/run) = 100 * tan(α), SI system - note that the % symbol is already used by Percent
	SlopePercent = mustCreateNewUnit("slope percent", "", _slope, SI, Plural(PluralNone))
	// SlopePermille 1000 * m = 1000 * (rise/run) = 1000 * tan(α), SI system - note that the ‰ symbol is already used by Permille
	SlopePermille = mustCreateNewUnit("slope permille", "", _slope, SI, Plural(PluralNone))
)

func init() {
	// Use SlopeValue as base unit for all conversions.
	// When using `NewConversionFromFn`, you must define conversions in both directions.

	NewConversionFromFn(SlopeValue, SlopeRatio, slopeValueToRatio, "n = 1 / m")
	NewConversionFromFn(SlopeRatio, SlopeValue, slopeRatioToValue, "m = 1 / n")

	NewConversionFromFn(SlopeValue, SlopeInverseRatio, returnSameValue, "x")
	NewConversionFromFn(SlopeInverseRatio, SlopeValue, returnSameValue, "x")

	NewConversionFromFn(SlopeValue, SlopeDegree, slopeValueToDegree, "math.Atan(x) * 180 / math.Pi")
	NewConversionFromFn(SlopeDegree, SlopeValue, slopeDegreeToValue, "math.Tan(x * math.Pi / 180)")

	NewRatioConversion(SlopeValue, SlopePercent, 100)

	NewRatioConversion(SlopeValue, SlopePermille, 1000)
}

// slopeValueToRatio converts a slope value (m) to a slope ratio (1:n).
func slopeValueToRatio(m float64) float64 {
	if math.IsInf(m, 0) {
		return 0
	}
	if m == 0 {
		return math.Inf(1)
	}
	// m=1:n => n=1/m
	return 1 / m
}

// slopeRatioToValue converts a slope ratio (1:n) to a slope value (m).
func slopeRatioToValue(n float64) float64 {
	if math.IsInf(n, 0) {
		return 0
	}
	if n == 0 {
		return math.Inf(1)
	}
	// m=1:n
	return 1 / n
}

// returnSameValue returns the same value.
func returnSameValue(x float64) float64 {
	return x
}

// slopeValueToDegree converts a slope value to a slope degree.
// - Positive infinity (= vertical upward) is 90°
// - Negative infinity (= vertical downward) is -90° (or 270°)
func slopeValueToDegree(x float64) float64 {
	// check for infinity
	if math.IsInf(x, 1) {
		return 90
	}
	if math.IsInf(x, -1) {
		return -90
	}
	return math.Atan(x) * 180 / math.Pi
}

// slopeDegreeToValue converts a slope degree to a slope value.
func slopeDegreeToValue(x float64) float64 {
	return math.Tan(x * math.Pi / 180)
}
