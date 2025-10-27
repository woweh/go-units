package units

import (
	"testing"

	"github.com/alecthomas/assert/v2"
)

func Test_ThermalInsulance_Symbols(t *testing.T) {
	si := SiSystem
	assert.Equal(t, si, KelvinSquareMeterPerWatt.System())
	bi := BiSystem
	assert.Equal(t, bi, DegreeFahrenheitHourSquareFootPerBtu.System())
}

func Test_ThermalInsulance_Conversion(t *testing.T) {
	conversionTests := []conversionTest{
		{"K·m²/W", "°C⋅m²/W", "1"},
		{"K*m2/W", "m²·K/W", "1"},
		{"°F⋅hr⋅ft²/Btu", "K·m²/W", "0.1761102"},
		{"K·m²/W", "°F⋅hr⋅ft²/Btu", "5.678263"},
		{"°F⋅ft²⋅h/BTU", "°F*ft2*h/BTU", "1"},
	}

	testConversions(t, conversionTests)
}
