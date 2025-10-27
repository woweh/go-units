package units

import (
	"testing"
)

func Test_MomentOfInertia_Conversions(t *testing.T) {
	var conversionTests = []conversionTest{
		// Base conversions (m⁴ as SI base)
		{"m⁴", "ft⁴", "115.861767"},
		{"m⁴", "in⁴", "2402509.610029"},
		{"m⁴", "cm⁴", "100000000"},
		{"m⁴", "mm⁴", "1000000000000"},

		// Reverse conversions
		{"ft⁴", "m⁴", "0.00863097"},
		{"in⁴", "m⁴", "0.000000416231"},

		// Identity
		{"m⁴", "m⁴", "1"},
		{"ft⁴", "ft⁴", "1"},
	}

	testConversions(t, conversionTests)
}
