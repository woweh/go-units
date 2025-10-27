package units

import (
	"testing"
)

func Test_CoefficientOfHeatTransfer_Conversions(t *testing.T) {
	var conversionTests = []conversionTest{
		// Base conversions (adjusted for library precision)
		{"BTU/(h·ft²·°F)", "W/(m²·K)", "5.67826"},
		{"W/(m²·K)", "BTU/(h·ft²·°F)", "0.1761103"},

		// Identity
		{"W/(m²·K)", "W/(m²·K)", "1"},
		{"BTU/(h·ft²·°F)", "BTU/(h·ft²·°F)", "1"},
	}

	testConversions(t, conversionTests)
}
