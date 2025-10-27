package units

import (
	"testing"
)

func Test_Velocity_Conversions(t *testing.T) {
	var conversionTests = []conversionTest{
		// Base conversions (adjusted for library precision)
		{"ft/s", "ft/min", "60"},
		{"ft/s", "m/s", "0.3048"},
		{"ft/s", "cm/min", "1828.8"},
		{"ft/s", "km/h", "1.09728"},
		{"ft/s", "mph", "0.681818"},

		// Common conversions
		{"m/s", "km/h", "3.6"},
		{"mph", "km/h", "1.609344"},

		// Identity
		{"ft/s", "ft/s", "1"},
		{"m/s", "m/s", "1"},
	}

	testConversions(t, conversionTests)
}
