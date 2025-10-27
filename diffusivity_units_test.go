package units

import (
	"testing"
)

func Test_Diffusivity_Conversions(t *testing.T) {
	var conversionTests = []conversionTest{
		// Base conversions (m²/s as SI base)
		{"m²/s", "ft²/s", "10.76391"},
		{"ft²/s", "m²/s", "0.092903"},
		
		// Identity
		{"m²/s", "m²/s", "1"},
		{"ft²/s", "ft²/s", "1"},
	}

	testConversions(t, conversionTests)
}
