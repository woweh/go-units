package units

import (
	"testing"
)

func Test_ThermalResistance_Conversions(t *testing.T) {
	var conversionTests = []conversionTest{
		// Base conversions (adjusted for library precision)
		{"(m²·K)/W", "(h·ft²·°F)/BTU", "5.678263"},
		{"(h·ft²·°F)/BTU", "(m²·K)/W", "0.1761102"},

		// Identity
		{"(m²·K)/W", "(m²·K)/W", "1"},
		{"(h·ft²·°F)/BTU", "(h·ft²·°F)/BTU", "1"},
	}

	testConversions(t, conversionTests)
}
