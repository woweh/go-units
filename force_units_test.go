package units

import (
	"testing"
)

func Test_Force_Conversions(t *testing.T) {
	var conversionTests = []conversionTest{
		// metric prefixes (representative subset)
		{"newton", "kilonewton", 0.001},   // 1 N = 0.001 kN
		{"newton", "millinewton", 1000.0}, // 1 N = 1000 mN
		{"newton", "micronewton", 1e6},    // 1 N = 1e6 µN

		// CGS
		{"newton", "dyne", 100000.0},

		// Imperial / other
		{"newton", "pound force", 0.22480894387096}, // 1 N = 0.22480894387096 lbf
		{"newton", "poundal", 7.2330139987},         // 1 N = 7.2330139987 pdl
		{"newton", "kilogram-force", 0.10197162129779},
		{"newton", "tonne-force", 0.00010197162129779},
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
