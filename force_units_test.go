package units

import (
	"testing"
)

func Test_Force_Conversions(t *testing.T) {
	var conversionTests = []conversionTest{
		// metric prefixes (representative subset)
		{from: "newton", to: "kilonewton", exp: 0.001},   // 1 N = 0.001 kN
		{from: "newton", to: "millinewton", exp: 1000.0}, // 1 N = 1000 mN
		{from: "newton", to: "micronewton", exp: 1e6},    // 1 N = 1e6 µN

		// CGS
		{from: "newton", to: "dyne", exp: 100000.0},

		// Imperial / other
		{from: "newton", to: "pound force", exp: 0.22480894387096, tol: fPtr(1e-7)}, // 1 N = 0.22480894387096 lbf
		{from: "newton", to: "poundal", exp: 7.2330139987, tol: fPtr(1e-6)},         // 1 N = 7.2330139987 pdl
		{from: "newton", to: "kilogram-force", exp: 0.10197162129779},
		{from: "newton", to: "tonne-force", exp: 0.00010197162129779},
	}

	testConversions(t, conversionTests)
}

func Test_Force_UnitSystems(t *testing.T) {
	tests := []unitSystemTest{
		{Newton, SiSystem},
		{MilliNewton, SiSystem},
		{MicroNewton, SiSystem},
		{KiloNewton, SiSystem},
		{Dyne, CgsSystem},
		{PoundForce, BiSystem},
		{Poundal, NoSystem},
		{KilogramForce, MKpSSystem},
		{TonneForce, MKpSSystem},
	}
	testUnitSystems(t, tests)
}
