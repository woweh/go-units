package units

import (
	"testing"
)

func Test_ThermalConductivity_Conversions(t *testing.T) {
	var conversionTests = []conversionTest{
		// Base conversions (adjusted for library precision)
		{"BTU/(h·ft·°F)", "W/(m·K)", "1.730735"},
		{"W/(m·K)", "BTU/(h·ft·°F)", "0.577789"},

		// Identity
		{"W/(m·K)", "W/(m·K)", "1"},
		{"BTU/(h·ft·°F)", "BTU/(h·ft·°F)", "1"},
	}

	testConversions(t, conversionTests)
}
