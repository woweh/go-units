package units

import (
	"testing"

	"github.com/alecthomas/assert/v2"
)

func Test_ElectricalResistance_Conversions(t *testing.T) {
	conversionTests := []conversionTest{
		// Metric progressions (sampled)
		{"kiloohm", "ohm", 1000},
		{"megaohm", "kiloohm", 1000},
		{"miliohm", "ohm", 0.001},
		{"microohm", "ohm", 0.000001},
		{"gigaohm", "megaohm", 1000},
		{"teraohm", "gigaohm", 1000},
	}
	testConversions(t, conversionTests)
}

func Test_ElectricalResistance_UnitSystems(t *testing.T) {
	si := SiSystem
	assert.Equal(t, si, Ohm.System())
	assert.Equal(t, si, KiloOhm.System())
	assert.Equal(t, si, MegaOhm.System())
	assert.Equal(t, si, MilliOhm.System())
	assert.Equal(t, si, MicroOhm.System())
	assert.Equal(t, si, GigaOhm.System())
	assert.Equal(t, si, TeraOhm.System())
	assert.Equal(t, si, DecaOhm.System())
	assert.Equal(t, si, HectoOhm.System())
	assert.Equal(t, si, DeciOhm.System())
	assert.Equal(t, si, CentiOhm.System())
	assert.Equal(t, si, PicoOhm.System())
	assert.Equal(t, si, FemtoOhm.System())
	assert.Equal(t, si, AttoOhm.System())
	assert.Equal(t, si, ZeptoOhm.System())
	assert.Equal(t, si, YoctoOhm.System())
	assert.Equal(t, si, PetaOhm.System())
	assert.Equal(t, si, ExaOhm.System())
	assert.Equal(t, si, ZettaOhm.System())
	assert.Equal(t, si, YottaOhm.System())
	assert.Equal(t, si, OhmMeter.System())
}

func Test_ElectricalResistance_BaseUnits(t *testing.T) {
	// Only a few representative metric units
	assert.Equal(t, Ohm, Ohm.Base())
	assert.Equal(t, Ohm, KiloOhm.Base())
	assert.Equal(t, Ohm, MilliOhm.Base())
	assert.Equal(t, Ohm, OhmMeter.Base())
}

func Test_Lookup_ElectricalResistance_Names_and_Symbols(t *testing.T) {
	tests := lookUpTests{
		{Ohm, "ohm"}, {Ohm, "Ω"}, {Ohm, "Ω"},
		{KiloOhm, "kiloohm"}, {KiloOhm, "kΩ"},
		{MegaOhm, "megaohm"}, {MegaOhm, "MΩ"},
		{MilliOhm, "miliohm"}, {MilliOhm, "mΩ"},
		{MicroOhm, "microohm"}, {MicroOhm, "μΩ"}, {MicroOhm, "uΩ"},
		{GigaOhm, "gigaohm"}, {GigaOhm, "GΩ"},
		{TeraOhm, "teraohm"}, {TeraOhm, "TΩ"},
		{DecaOhm, "decaohm"}, {DecaOhm, "daΩ"},
		{HectoOhm, "hectoohm"}, {HectoOhm, "hΩ"},
		{DeciOhm, "deciohm"}, {DeciOhm, "dΩ"},
		{CentiOhm, "centiohm"}, {CentiOhm, "cΩ"},
		{PicoOhm, "picoohm"}, {PicoOhm, "pΩ"},
		{FemtoOhm, "femtoohm"}, {FemtoOhm, "fΩ"},
		{AttoOhm, "attoohm"}, {AttoOhm, "aΩ"},
		{ZeptoOhm, "zeptoohm"}, {ZeptoOhm, "zΩ"},
		{YoctoOhm, "yoctoohm"}, {YoctoOhm, "yΩ"},
		{PetaOhm, "petaohm"}, {PetaOhm, "PΩ"},
		{ExaOhm, "exaohm"}, {ExaOhm, "EΩ"},
		{ZettaOhm, "zettaohm"}, {ZettaOhm, "ZΩ"},
		{YottaOhm, "yottaohm"}, {YottaOhm, "YΩ"},
		{OhmMeter, "ohm meter"}, {OhmMeter, "ohm-meter"}, {OhmMeter, "ohm·m"},
	}
	testLookupNamesAndSymbols(t, tests)
}
