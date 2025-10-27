package units

import (
	"testing"
)

func Test_Luminance_Conversions(t *testing.T) {
	var conversionTests = []conversionTest{
		// Base conversions (adjusted for library precision)
		{"cd/ft²", "cd/m²", "10.76391"},
		{"cd/ft²", "ftL", "3.141593"},
		
		// Reverse conversions
		{"cd/m²", "cd/ft²", "0.092903"},
		{"ftL", "cd/ft²", "0.31831"},

		// Identity
		{"cd/ft²", "cd/ft²", "1"},
		{"cd/m²", "cd/m²", "1"},
		{"ftL", "ftL", "1"},
	}

	testConversions(t, conversionTests)
}
