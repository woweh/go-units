package units

import (
	"testing"
)

func Test_Diffusivity_Conversions(t *testing.T) {
	conversionTests := []conversionTest{
		{from: "m²/s", to: "ft²/s", exp: 10.7639104167097},
		{from: "ft²/s", to: "m²/s", exp: 0.09290304},
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
