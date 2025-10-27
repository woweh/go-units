package units

import (
	"testing"
)

func Test_Illuminance_Conversions(t *testing.T) {
	var conversionTests = []conversionTest{
		// Base conversions (lux as SI base)
		{"lx", "Ftc", "0.092903"},
		{"Ftc", "lx", "10.76391"},

		// Identity
		{"lx", "lx", "1"},
		{"Ftc", "Ftc", "1"},
	}

	testConversions(t, conversionTests)
}
