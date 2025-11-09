package units

import (
	"testing"
)

func Test_Moment_Conversions(t *testing.T) {
	conversionTests := []conversionTest{
		// SI metric unit conversions
		{"N-m", "daN-m", 0.1},
		{"daN-m", "N-m", 10},
		{"N-m", "kN-m", 0.001},
		{"kN-m", "N-m", 1000},
		{"N-m", "MN-m", 0.000001},
		{"MN-m", "N-m", 1000000},

		// SI to Imperial conversions
		{"N-m", "lb-ft", 0.737562},
		{"lb-ft", "N-m", 1.35581794833},
		{"kN-m", "lb-ft", 737.562},
		{"lb-ft", "kN-m", 0.00135581794833},

		// Imperial conversions
		{"kip-ft", "lb-ft", 1000},
		{"lb-ft", "kip-ft", 0.001},

		// Force-based conversions
		{"kgf-m", "N-m", 9.80665},
		{"N-m", "kgf-m", 0.10197162129779},
		{"Tf-m", "N-m", 9806.65},
		{"N-m", "Tf-m", 0.00010197162129779},

		// Cross conversions
		{"kip-ft", "kN-m", 1.35581794833},
		{"kN-m", "kip-ft", 0.737562},
		{"kgf-m", "lb-ft", 7.23301385120968},
		{"Tf-m", "kip-ft", 7.23301385120968},
	}
	testConversions(t, conversionTests)
}

func Test_Moment_UnitSystems(t *testing.T) {
	tests := []unitSystemTest{
		{NewtonMeter, SiSystem},
		{DekaNewtonMeter, SiSystem},
		{KiloNewtonMeter, SiSystem},
		{MegaNewtonMeter, SiSystem},
		{PoundForceFoot, BiSystem},
		{KipFoot, BiSystem},
		{KilogramForceMeter, MKpSSystem},
		{TonneForceMeter, MKpSSystem},
	}
	testUnitSystems(t, tests)
}
