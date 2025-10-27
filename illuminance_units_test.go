package units

import (
	"testing"
)

func Test_Illuminance_Conversions(t *testing.T) {
	var conversionTests = []conversionTest{
		// Base conversions (adjusted for library precision)
		{"Ftc", "lx", "10.76391"},
		{"lx", "Ftc", "0.092903"},

		// Identity
		{"lx", "lx", "1"},
		{"Ftc", "Ftc", "1"},
	}

	testConversions(t, conversionTests)
}
