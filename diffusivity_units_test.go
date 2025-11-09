package units

import (
	"testing"
)

func Test_Diffusivity_Conversions(t *testing.T) {
	conversionTests := []conversionTest{
		{"m²/s", "ft²/s", 10.7639104167097},
		{"ft²/s", "m²/s", 0.09290304},
	}
	testConversions(t, conversionTests)
}

func Test_Diffusivity_UnitSystems(t *testing.T) {
	tests := []unitSystemTest{
		{SquareMeterPerSecond, SiSystem},
		{SquareFootPerSecond, BiSystem},
	}
	testUnitSystems(t, tests)
}
