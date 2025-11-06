package units

import (
	"testing"

	"github.com/alecthomas/assert/v2"
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
	assert.Equal(t, SiSystem, NewtonMeter.System())
	assert.Equal(t, SiSystem, DekaNewtonMeter.System())
	assert.Equal(t, SiSystem, KiloNewtonMeter.System())
	assert.Equal(t, SiSystem, MegaNewtonMeter.System())
	assert.Equal(t, BiSystem, PoundForceFoot.System())
	assert.Equal(t, BiSystem, KipFoot.System())
	assert.Equal(t, MKpSSystem, KilogramForceMeter.System())
	assert.Equal(t, MKpSSystem, TonneForceMeter.System())
}

func Test_Lookup_Moment_Names_and_Symbols(t *testing.T) {
	tests := lookUpTests{
		{NewtonMeter, "newton meter"},
		{NewtonMeter, "N-m"},
		{NewtonMeter, "newton meters"},
		{NewtonMeter, "N*m"},
		{NewtonMeter, "Nm"},
		{DekaNewtonMeter, "dekanewton meter"},
		{DekaNewtonMeter, "daN-m"},
		{DekaNewtonMeter, "dekanewton meters"},
		{KiloNewtonMeter, "kilonewton meter"},
		{KiloNewtonMeter, "kN-m"},
		{KiloNewtonMeter, "kilonewton meters"},
		{MegaNewtonMeter, "meganewton meter"},
		{MegaNewtonMeter, "MN-m"},
		{MegaNewtonMeter, "meganewton meters"},
		{PoundForceFoot, "pound force foot"},
		{PoundForceFoot, "lb-ft"},
		{PoundForceFoot, "pound force feet"},
		{PoundForceFoot, "pound-foot"},
		{PoundForceFoot, "pound-feet"},
		{PoundForceFoot, "lbf-ft"},
		{KipFoot, "kip foot"},
		{KipFoot, "kip-ft"},
		{KipFoot, "kip feet"},
		{KipFoot, "kip-feet"},
		{KilogramForceMeter, "kilogram force meter"},
		{KilogramForceMeter, "kgf-m"},
		{KilogramForceMeter, "kilogram force meters"},
		{TonneForceMeter, "tonne force meter"},
		{TonneForceMeter, "Tf-m"},
		{TonneForceMeter, "tonne force meters"},
	}
	testLookupNamesAndSymbols(t, tests)
}
