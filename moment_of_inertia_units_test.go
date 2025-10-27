package units

import (
	"testing"
)

func Test_MomentOfInertia_Conversions(t *testing.T) {
	var conversionTests = []conversionTest{
		// Base conversions (adjusted for library precision)
		{"ft⁴", "in⁴", "20736"},
		{"ft⁴", "m⁴", "0.00863097"},
		{"ft⁴", "cm⁴", "863097.484124"},
		{"ft⁴", "mm⁴", "8630974841.2416"},

		// Common conversions (1 m = 100 cm, so 1 m⁴ = 100,000,000 cm⁴)
		{"m⁴", "cm⁴", "100000000"},

		// Identity
		{"ft⁴", "ft⁴", "1"},
		{"m⁴", "m⁴", "1"},
	}

	testConversions(t, conversionTests)
}
