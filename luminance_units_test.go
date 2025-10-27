package units

import (
	"testing"
)

func Test_Luminance_Conversions(t *testing.T) {
	var conversionTests = []conversionTest{
		// Base conversions (cd/m² as SI base)
		{"cd/m²", "cd/ft²", "0.092903"},
		{"cd/m²", "ftL", "0.2918635"},
		
		// Reverse conversions
		{"cd/ft²", "cd/m²", "10.76391"},
		{"ftL", "cd/m²", "3.426259"},

		// Identity
		{"cd/m²", "cd/m²", "1"},
		{"cd/ft²", "cd/ft²", "1"},
		{"ftL", "ftL", "1"},
	}

	testConversions(t, conversionTests)
}
