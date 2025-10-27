package units

import (
	"testing"
)

func Test_DynamicViscosity_Conversions(t *testing.T) {
	var conversionTests = []conversionTest{
		{"Pa-s", "N·s/m²", "1"},
		{"Pa-s", "cP", "1000"},
		{"kg/(m·h)", "kg/(m·s)", "3600"},
		{"Pa-s", "Pa-s", "1"},
	}

	testConversions(t, conversionTests)
}
