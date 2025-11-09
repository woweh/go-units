package units

import (
	"fmt"
	"math"
	"testing"
)

// conversionTest struct defines test data for a conversion, where a value of 1.0 in
// the `from` unit name is expected to equal `val` when converted to the `to` unit name
type conversionTest struct {
	from string
	to   string
	val  float64
}

// floatEquals returns true if a and b are within tolerance 'epsilon' of each other.
func floatEquals(a, b, epsilon float64) bool {
	return math.Abs(a-b) <= epsilon
}

func testConversions(t *testing.T, convTests []conversionTest) {
	const (
		// epsilon is the tolerance used for comparing floating-point numbers
		epsilon   = 1e-9
		baseValue = 1.0
	)
	for _, cTest := range convTests {
		label := fmt.Sprintf("%s <-> %s conversion", cTest.from, cTest.to)
		t.Run(
			label, func(t *testing.T) {
				u1, err := Find(cTest.from)
				if err != nil {
					t.Errorf("failed to find unit %s: %v", cTest.from, err)
				}
				u2, err := Find(cTest.to)
				if err != nil {
					t.Errorf("failed to find unit %s: %v", cTest.to, err)
				}

				converted := MustConvertFloat(baseValue, u1, u2)
				got := converted.Float()
				if !floatEquals(got, cTest.val, epsilon) {
					t.Errorf("%s -> %s conversion failed: want %f, got %f", cTest.from, cTest.to, cTest.val, got)
				}

				// test inverse conversion
				inverse := MustConvertFloat(got, u2, u1)
				gotInv := inverse.Float()
				if !floatEquals(gotInv, baseValue, epsilon) {
					t.Errorf("%s <- %s inverse conversion failed: want 1.0, got %f", cTest.from, cTest.to, gotInv)
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
