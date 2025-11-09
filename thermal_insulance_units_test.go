package units

import (
	"testing"
)

func Test_ThermalInsulance_Conversions(t *testing.T) {
	conversionTests := []conversionTest{
		// Library/NIST conversions (see thermal_insulance_units.go)
		{"°F⋅hr⋅ft²/Btu", "K·m²/W", 0.1761102},
		{"K·m²/W", "°F⋅hr⋅ft²/Btu", 5.678263},
	}
	testConversions(t, conversionTests)
}

func Test_ThermalInsulance_UnitSystems(t *testing.T) {
	tests := []unitSystemTest{
		{KelvinSquareMeterPerWatt, SiSystem},
		{DegreeFahrenheitHourSquareFootPerBtu, BiSystem},
	}
	testUnitSystems(t, tests)
}
