package units

import (
	"testing"
)

func Test_ThermalInsulance_Conversions(t *testing.T) {
	conversionTests := []conversionTest{
		// Library/NIST conversions (see thermal_insulance_units.go)
		{from: "°F⋅hr⋅ft²/Btu", to: "K·m²/W", exp: 0.1761102},
		{from: "K·m²/W", to: "°F⋅hr⋅ft²/Btu", exp: 5.67826281},
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
