package units

import (
	"testing"

	"github.com/alecthomas/assert/v2"
)

func Test_ThermalInsulance_Conversions(t *testing.T) {
	conversionTests := []conversionTest{
		// Library/NIST conversions (see thermal_insulance_units.go)
		{"°F⋅hr⋅ft²/Btu", "K·m²/W", 0.1761102},
		{"K·m²/W", "°F⋅hr⋅ft²/Btu", 5.678263},
	}
	testConversions(t, conversionTests)
}

func Test_ThermalInsulance_UnitSystem(t *testing.T) {
	assert.Equal(t, SiSystem, KelvinSquareMeterPerWatt.System())
	assert.Equal(t, BiSystem, DegreeFahrenheitHourSquareFootPerBtu.System())
}

// No metric factories for thermal insulance, so no base unit tests are needed.

func Test_ThermalInsulance_Lookup_Names_and_Symbols(t *testing.T) {
	tests := lookUpTests{
		// SI
		{KelvinSquareMeterPerWatt, "K·m²/W"},
		{KelvinSquareMeterPerWatt, "K*m2/W"},
		{KelvinSquareMeterPerWatt, "°C⋅m²/W"},
		{KelvinSquareMeterPerWatt, "°C*m2/W"},
		{KelvinSquareMeterPerWatt, "m2.K.W-1"},
		{KelvinSquareMeterPerWatt, "m²·K/W"},
		{KelvinSquareMeterPerWatt, "m2*K/W"},
		{KelvinSquareMeterPerWatt, "R-value"},
		{KelvinSquareMeterPerWatt, "RSI"},
		{KelvinSquareMeterPerWatt, "RSI-value"},
		// Imperial/US
		{DegreeFahrenheitHourSquareFootPerBtu, "°F⋅hr⋅ft²/Btu"},
		{DegreeFahrenheitHourSquareFootPerBtu, "degree Fahrenheit hour square foot per British thermal unit"},
		{DegreeFahrenheitHourSquareFootPerBtu, "degree Fahrenheit hour square foot per British thermal unitIT"},
		{DegreeFahrenheitHourSquareFootPerBtu, "degree Fahrenheit hour square foot per British thermal unitth"},
		{DegreeFahrenheitHourSquareFootPerBtu, "°F · h · ft2/BtuIT"},
		{DegreeFahrenheitHourSquareFootPerBtu, "°F*hr*ft2/Btu"},
		{DegreeFahrenheitHourSquareFootPerBtu, "°F⋅ft²⋅h/BTU"},
		{DegreeFahrenheitHourSquareFootPerBtu, "°F⋅ft2⋅h/BTU"},
		{DegreeFahrenheitHourSquareFootPerBtu, "°F*ft2*h/BTU"},
	}
	testLookupNamesAndSymbols(t, tests)
}
