package units

import (
	"testing"
)

func Test_Moment_Conversions(t *testing.T) {
	conversionTests := []conversionTest{
		// SI metric unit conversions
		{from: "N-m", to: "daN-m", exp: 0.1},
		{from: "daN-m", to: "N-m", exp: 10},
		{from: "N-m", to: "kN-m", exp: 0.001},
		{from: "kN-m", to: "N-m", exp: 1000},
		{from: "N-m", to: "MN-m", exp: 0.000001},
		{from: "MN-m", to: "N-m", exp: 1000000},

		// SI to Imperial conversions
		{from: "N-m", to: "lb-ft", exp: 0.737562085, tol: fPtr(1e-8)},
		{from: "lb-ft", to: "N-m", exp: 1.355817948, tol: fPtr(1e-6)},
		{from: "kN-m", to: "lb-ft", exp: 737.562085483, tol: fPtr(1e-8)},
		{from: "lb-ft", to: "kN-m", exp: 0.00135581794833},

		// Imperial conversions
		{from: "kip-ft", to: "lb-ft", exp: 1000},
		{from: "lb-ft", to: "kip-ft", exp: 0.001},

		// Force-based conversions
		{from: "kgf-m", to: "N-m", exp: 9.80665},
		{from: "N-m", to: "kgf-m", exp: 0.10197162129779},
		{from: "Tf-m", to: "N-m", exp: 9806.65},
		{from: "N-m", to: "Tf-m", exp: 0.00010197162129779},

		// Cross conversions
		{from: "kip-ft", to: "kN-m", exp: 1.355817948, tol: fPtr(1e-6)},
		{from: "kN-m", to: "kip-ft", exp: 0.737562085, tol: fPtr(1e-8)},
		{from: "kgf-m", to: "lb-ft", exp: 7.23301385120968, tol: fPtr(1e-5)},
		{from: "Tf-m", to: "kip-ft", exp: 7.23301385120968, tol: fPtr(1e-5)},
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
