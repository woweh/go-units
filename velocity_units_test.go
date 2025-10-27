package units

import (
	"testing"
)

func Test_Velocity_Conversions(t *testing.T) {
	var conversionTests = []conversionTest{
		// Base conversions (m/s as SI base)
		{"m/s", "ft/s", "3.28084"},
		{"m/s", "ft/min", "196.850394"},
		{"m/s", "cm/min", "6000"},
		{"m/s", "km/h", "3.6"},
		{"m/s", "mph", "2.236936"},

		// Reverse conversions
		{"ft/s", "m/s", "0.3048"},
		{"km/h", "m/s", "0.2777778"},
		{"mph", "m/s", "0.44704"},

		// Common conversions
		{"mph", "km/h", "1.609344"},

		// Identity
		{"m/s", "m/s", "1"},
		{"ft/s", "ft/s", "1"},
	}

	testConversions(t, conversionTests)
}
