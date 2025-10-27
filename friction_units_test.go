package units

import (
	"testing"
)

func Test_Friction_Conversions(t *testing.T) {
	var conversionTests = []conversionTest{
		{"mH2O/m", "mmH2O/m", "1000"},
		{"mmH2O/m", "Pa/m", "9.80665"},
		{"mmH2O/m", "FT/100ft", "0.1"},
		{"mmH2O/m", "mmH2O/m", "1"},
	}

	testConversions(t, conversionTests)
}
