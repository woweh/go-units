package units

import (
	"testing"
)

func Test_Stress_Conversions(t *testing.T) {
	var conversionTests = []conversionTest{
		// Stress-specific units to Pascal (from pressure) - adjusted for library precision
		{"ksf", "Pa", "47880.26"},
		{"ksi", "Pa", "6894757.293168"},

		// Common conversions
		{"ksi", "kPa", "6894.757293"},
		{"MPa", "ksi", "0.1450377"},

		// Identity
		{"ksf", "ksf", "1"},
		{"ksi", "ksi", "1"},
	}

	testConversions(t, conversionTests)
}
