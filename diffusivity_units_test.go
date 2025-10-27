package units

import (
	"testing"
)

func Test_Diffusivity_Conversions(t *testing.T) {
	var conversionTests = []conversionTest{
		{"ft²/s", "m²/s", "0.092903"},
		{"m²/s", "ft²/s", "10.76391"},
		{"ft²/s", "ft²/s", "1"},
	}

	testConversions(t, conversionTests)
}
