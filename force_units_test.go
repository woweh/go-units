package units

import (
	"testing"

	"github.com/alecthomas/assert/v2"
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

func Test_Force_Systems(t *testing.T) {
	// representative SI checks (don't repeat every SI prefix)
	si := SiSystem
	assert.Equal(t, si, Newton.System())
	assert.Equal(t, si, MilliNewton.System())
	assert.Equal(t, si, MicroNewton.System())
	assert.Equal(t, si, KiloNewton.System())

	// other systems
	bi := BiSystem
	assert.Equal(t, bi, PoundForce.System())

	cgs := CgsSystem
	assert.Equal(t, cgs, Dyne.System())

	mkps := MKpSSystem
	assert.Equal(t, mkps, KilogramForce.System())
	assert.Equal(t, mkps, TonneForce.System())
}

func Test_Force_BaseUnits(t *testing.T) {
	// representative checks that metric factories return Newton as the base
	assert.Equal(t, Newton, KiloNewton.Base())
	assert.Equal(t, Newton, MilliNewton.Base())
	assert.Equal(t, Newton, CentiNewton.Base())
}

func Test_Lookup_Force_Names_and_Symbols(t *testing.T) {
	tests := lookUpTests{
		{Newton, "newton"}, {Newton, "N"},
		{KiloNewton, "kilonewton"}, {MilliNewton, "millinewton"},
		{Dyne, "dyne"}, {Dyne, "dyn"},
		{PoundForce, "pound force"}, {PoundForce, "lbf"},
		{Poundal, "poundal"}, {Poundal, "pdl"},
		{KilogramForce, "kilogram-force"}, {KilogramForce, "kgf"}, {KilogramForce, "kilopond"}, {KilogramForce, "kp"},
		{TonneForce, "tonne-force"}, {TonneForce, "tf"}, {TonneForce, "megapond"}, {TonneForce, "Mp"},
		{ShortTonForce, "short ton force"}, {ShortTonForce, "Tons"},
	}

	testLookupNamesAndSymbols(t, tests)
}
