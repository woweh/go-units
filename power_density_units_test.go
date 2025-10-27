package units

import (
	"testing"
)

func Test_PowerDensity_Conversions(t *testing.T) {
	var conversionTests = []conversionTest{
		{"W/m²", "W/ft²", "0.092903"},
		{"W/ft²", "Btu/(h·ft²)", "3.41214"},
		{"W/m²", "W/m²", "1"},
	}

	testConversions(t, conversionTests)
}
