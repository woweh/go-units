package units

import (
	"testing"
)

func Test_Acceleration_Conversions(t *testing.T) {
	var conversionTests = []conversionTest{
		// Base conversions (adjusted for library precision)
		{"ft/s²", "in/s²", "12"},
		{"ft/s²", "m/s²", "0.3048"},
		{"ft/s²", "km/s²", "0.0003048"},
		{"ft/s²", "mi/s²", "0.0001893939"},

		// Common conversions
		{"m/s²", "ft/s²", "3.28084"},

		// Identity
		{"ft/s²", "ft/s²", "1"},
		{"m/s²", "m/s²", "1"},
	}

	testConversions(t, conversionTests)
}
