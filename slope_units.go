package units

import "math"

// https://en.wikipedia.org/wiki/Slope
// https://en.wikipedia.org/wiki/Grade_(slope)

var (
	// Slope is a unit option for slope units, also known as gradient.
	Slope = UnitOptionQuantity("slope")

	// SlopeValue m = rise (delta h) / run (distance) = y/x = tan(alpha)
	// In the case of a vertical line, the slope is infinite.
	// In the case of a horizontal line, the slope is zero.
	SlopeValue = newUnit("slope value", "", Slope, UnitOptionPlural(None, ""))
	// SlopeRatio is a ratio of one part rise to so many parts run (e.g. 1:10). == SlopeInverseValue
	SlopeRatio = newUnit("slope ratio", "", Slope, UnitOptionPlural(None, ""))
	// SlopeInverseRatio is a ratio of many parts run to one part rise (e.g. 10:1).
	SlopeInverseRatio = newUnit("inverse slope ratio", "", Slope, UnitOptionPlural(None, ""))
	// SlopeDegree is the angle of inclination in degrees (e.g. 45°).
	SlopeDegree = newUnit("slope degree", "", Slope, UnitOptionPlural(None, ""))
	// SlopePercent 100 * m = 100 * (rise/run) = 100 * tan(α)
	SlopePercent = newUnit("slope percent", "", Slope, UnitOptionPlural(None, ""))
	// SlopePermille 1000 * m = 1000 * (rise/run) = 1000 * tan(α)
	SlopePermille = newUnit("slope permille", "", Slope, UnitOptionPlural(None, ""))
)

func init() {
	// Use SlopeValue as base unit for all conversions.
	// When using `NewConversionFromFn`, you must define conversions in both directions.

	NewConversionFromFn(SlopeValue, SlopeRatio, slopeValueToRatio, "1 / x")
	NewConversionFromFn(SlopeRatio, SlopeValue, slopeRatioToValue, "1 / x")

	NewConversionFromFn(SlopeValue, SlopeInverseRatio, returnSameValue, "x")
	NewConversionFromFn(SlopeInverseRatio, SlopeValue, returnSameValue, "x")

	NewConversionFromFn(SlopeValue, SlopeDegree, slopeValueToDegree, "math.Atan(x) * 180 / math.Pi")
	NewConversionFromFn(SlopeDegree, SlopeValue, slopeDegreeToValue, "math.Tan(x * math.Pi / 180)")

	NewConversionFromFn(SlopeValue, SlopePercent, valueToPercent, "x * 100")
	NewConversionFromFn(SlopePercent, SlopeValue, percentToValue, "x / 100")

	NewConversionFromFn(SlopeValue, SlopePermille, valueToPermille, "x * 1000")
	NewConversionFromFn(SlopePermille, SlopeValue, permilleToValue, "x / 1000")
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

// valueToPercent converts a decimal value to percent.
func valueToPercent(x float64) float64 {
	if math.IsInf(x, 0) {
		return x
	}
	return x * 100
}

// percentToValue converts percent to decimal value.
func percentToValue(x float64) float64 {
	if math.IsInf(x, 0) {
		return x
	}
	return x / 100
}

// valueToPermille converts a decimal value to permille.
func valueToPermille(x float64) float64 {
	if math.IsInf(x, 0) {
		return x
	}
	return x * 1000
}

// permilleToValue converts permille to a decimal value.
func permilleToValue(x float64) float64 {
	if math.IsInf(x, 0) {
		return x
	}
	return x / 1000
}
