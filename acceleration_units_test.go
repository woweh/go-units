package units

import (
	"testing"
)

func Test_Acceleration_Conversions(t *testing.T) {
	var conversionTests = []conversionTest{
		// Base conversions (m/s² as SI base)
		{"m/s²", "ft/s²", "3.28084"},
		{"m/s²", "in/s²", "39.370079"},
		{"m/s²", "km/s²", "0.001"},
		{"m/s²", "mi/s²", "0.000621371"},

		// Reverse conversions
		{"ft/s²", "m/s²", "0.3048"},
		{"in/s²", "m/s²", "0.0254"},
		{"km/s²", "m/s²", "1000"},

		// Identity
		{"m/s²", "m/s²", "1"},
		{"ft/s²", "ft/s²", "1"},
	}

	testConversions(t, conversionTests)
}
