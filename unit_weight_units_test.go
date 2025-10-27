package units

import (
	"testing"
)

func Test_UnitWeight_Conversions(t *testing.T) {
	var conversionTests = []conversionTest{
		{"kN/m³", "lbf/ft³", "6.36588"},
		{"kip/in³", "lbf/ft³", "1728000"},
		{"lbf/ft³", "lbf/ft³", "1"},
	}

	testConversions(t, conversionTests)
}
