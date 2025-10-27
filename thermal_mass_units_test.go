package units

import (
	"testing"
)

func Test_ThermalMass_Conversions(t *testing.T) {
	var conversionTests = []conversionTest{
		// Base conversions (adjusted for library precision)
		{"kJ/K", "J/K", "1000"},
		{"BTU/°F", "J/K", "1899.1"},
		{"J/K", "BTU/°F", "0.000526565"},

		// Identity
		{"J/K", "J/K", "1"},
		{"BTU/°F", "BTU/°F", "1"},
	}

	testConversions(t, conversionTests)
}
