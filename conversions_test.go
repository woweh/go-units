package units

import (
	"fmt"
	"math"
	"testing"
)

// conversionTest defines a test case for unit conversion (> testConversions).
//
// It specifies the source unit, target unit, expected result, and an optional
// tolerance for floating-point comparison. If tolerance is nil, a default
// value (_defaultTolerance) is used.
//
// A value of 1.0 (_baseValue) in the `from` unit should convert to `exp` in the `to` unit.
type conversionTest struct {
	// from is the name, alias, or symbol of the unit to convert from.
	from string
	// to is the name, alias, or symbol of the unit to convert to.
	to string
	// exp is the value expected after conversion.
	exp float64
	// tol is the optional tolerance for floating-point comparison.
	// If nil, the default tolerance is used.
	tol *float64
}

// fPtr returns a pointer to the given float64 value.
func fPtr(f float64) *float64 {
	return &f
}

const (
	// _defaultTolerance is the default tolerance used for comparing floating-point numbers.
	_defaultTolerance = 1e-8
	// _baseValue is the initial/base value used for conversion tests.
	_baseValue = 1.0
)

// getTolerance returns the tolerance for floating-point comparison, defaulting to _defaultTolerance if not specified.
func (ct conversionTest) getTolerance() float64 {
	if ct.tol == nil {
		return _defaultTolerance
	}
	return *ct.tol
}

// floatEquals returns true if a and b are within tolerance t of each other.
func floatEquals(a, b, t float64) bool {
	return math.Abs(a-b) <= t
}

// testConversions executes the conversion tests and verifies the results.
func testConversions(t *testing.T, convTests []conversionTest) {
	for _, cTest := range convTests {
		label := fmt.Sprintf("%s <-> %s conversion", cTest.from, cTest.to)
		t.Run(label, func(t *testing.T) {
			tolerance := cTest.getTolerance()
			u1, err := Find(cTest.from)
			if err != nil {
				t.Fatalf("failed to find unit %s: %v", cTest.from, err)
			}
			u2, err := Find(cTest.to)
			if err != nil {
				t.Fatalf("failed to find unit %s: %v", cTest.to, err)
			}

			converted, err := ConvertFloat(_baseValue, u1, u2)
			if err != nil {
				t.Fatalf("conversion failed: %v", err)
			}
			got := converted.Float()
			if !floatEquals(got, cTest.exp, tolerance) {
				t.Errorf("%s -> %s conversion failed: want %.9f, got %.9f", cTest.from, cTest.to, cTest.exp, got)
			}

			// test inverse conversion
			inverse, err := ConvertFloat(got, u2, u1)
			if err != nil {
				t.Fatalf("inverse conversion failed: %v", err)
			}
			gotInv := inverse.Float()
			if !floatEquals(gotInv, _baseValue, tolerance) {
				t.Errorf("%s <- %s inverse conversion failed: want 1.0, got %.2f", cTest.from, cTest.to, gotInv)
			}
		},
		)
	}
}

// roundFloat is a helper for rounding floats to a given precision.
func roundFloat(f float64, precision uint) float64 {
	r := math.Pow(10, float64(precision))
	return math.Round(f*r) / r
}

func Test_ResolveConversion(t *testing.T) {
	from, err := Find("meter")
	if err != nil {
		t.Errorf("failed to find unit meter: %v", err)
	}
	to, err := Find("foot")
	if err != nil {
		t.Errorf("failed to find unit foot: %v", err)
	}
	path, err := ResolveConversion(from, to)
	if err != nil {
		t.Errorf("ResolveConversion failed: %v", err)
	}
	if len(path) == 0 {
		t.Error("expected non-empty conversion path")
	}
	// Verify the path by applying conversions
	val := 1.0
	for _, c := range path {
		val = c.Fn(val)
	}
	expected := 3.280839895013123 // approximate 1 meter in feet
	if math.Abs(val-expected) > 1e-6 {
		t.Errorf("conversion path incorrect: got %f, expected %f", val, expected)
	}
}

func Test_checkConversionIntegrity(t *testing.T) {
	// Call the checkConversionIntegrity function and check for errors
	err := CheckConversionIntegrity()
	if err != nil {
		t.Errorf("checkConversionIntegrity failed: %v", err)
	} else {
		t.Log("checkConversionIntegrity passed")
	}
}

func Test_roundFloat(t *testing.T) {
	tests := []struct {
		name      string
		input     float64
		precision uint
		expected  float64
	}{
		{"zero", 0, 2, 0},
		{"positive", 1.23456, 2, 1.23},
		{"negative", -1.23456, 2, -1.23},
		{"round up", 1.235, 2, 1.24},
		{"large precision", 1.23456789, 7, 1.2345679},
	}
	for _, ttc := range tests {
		tc := ttc
		t.Run(tc.name, func(t *testing.T) {
			got := roundFloat(tc.input, tc.precision)
			if got != tc.expected {
				t.Errorf("roundFloat(%v, %d) = %v, want %v", tc.input, tc.precision, got, tc.expected)
			}
		})
	}
}
