package units

import (
	"testing"

	"github.com/alecthomas/assert/v2"
)

func Test_ElectricCurrent_Conversions(t *testing.T) {
	conversionTests := []conversionTest{
		{"kiloampere", "ampere", 1000},
		{"megaampere", "kiloampere", 1000},
		{"milliampere", "ampere", 0.001},
		{"microampere", "ampere", 0.000001},
		{"nanoampere", "ampere", 0.000000001},
		{"gigaampere", "megaampere", 1000},
		{"teraampere", "gigaampere", 1000},
	}
	testConversions(t, conversionTests)
}

func Test_ElectricCurrent_UnitSystems(t *testing.T) {
	si := SiSystem
	assert.Equal(t, si, Ampere.System())
	assert.Equal(t, si, MilliAmpere.System())
	assert.Equal(t, si, MicroAmpere.System())
	assert.Equal(t, si, NanoAmpere.System())
	assert.Equal(t, si, PicoAmpere.System())
	assert.Equal(t, si, FemtoAmpere.System())
	assert.Equal(t, si, AttoAmpere.System())
	assert.Equal(t, si, ZeptoAmpere.System())
	assert.Equal(t, si, YoctoAmpere.System())
	assert.Equal(t, si, KiloAmpere.System())
	assert.Equal(t, si, MegaAmpere.System())
	assert.Equal(t, si, GigaAmpere.System())
	assert.Equal(t, si, TeraAmpere.System())
	assert.Equal(t, si, PetaAmpere.System())
	assert.Equal(t, si, ExaAmpere.System())
	assert.Equal(t, si, ZettaAmpere.System())
	assert.Equal(t, si, YottaAmpere.System())
}

func Test_ElectricCurrent_BaseUnits(t *testing.T) {
	// Only a few representative metric units
	assert.Equal(t, Ampere, Ampere.Base())
	assert.Equal(t, Ampere, MilliAmpere.Base())
	assert.Equal(t, Ampere, KiloAmpere.Base())
}

func Test_Lookup_ElectricCurrent_Names_and_Symbols(t *testing.T) {
	tests := lookUpTests{
		{Ampere, "ampere"}, {Ampere, "A"},
		{MilliAmpere, "milliampere"}, {MilliAmpere, "mA"},
		{MicroAmpere, "microampere"}, {MicroAmpere, "μA"}, {MicroAmpere, "uA"},
		{NanoAmpere, "nanoampere"}, {NanoAmpere, "nA"},
		{PicoAmpere, "picoampere"}, {PicoAmpere, "pA"},
		{FemtoAmpere, "femtoampere"}, {FemtoAmpere, "fA"},
		{AttoAmpere, "attoampere"}, {AttoAmpere, "aA"},
		{ZeptoAmpere, "zeptoampere"}, {ZeptoAmpere, "zA"},
		{YoctoAmpere, "yoctoampere"}, {YoctoAmpere, "yA"},
		{KiloAmpere, "kiloampere"}, {KiloAmpere, "kA"},
		{MegaAmpere, "megaampere"}, {MegaAmpere, "MA"},
		{GigaAmpere, "gigaampere"}, {GigaAmpere, "GA"},
		{TeraAmpere, "teraampere"}, {TeraAmpere, "TA"},
		{PetaAmpere, "petaampere"}, {PetaAmpere, "PA"},
		{ExaAmpere, "exaampere"}, {ExaAmpere, "EA"},
		{ZettaAmpere, "zettaampere"}, {ZettaAmpere, "ZA"},
		{YottaAmpere, "yottaampere"}, {YottaAmpere, "YA"},
	}
	testLookupNamesAndSymbols(t, tests)
}
