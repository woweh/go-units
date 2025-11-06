package units

import (
	"testing"

	"github.com/alecthomas/assert/v2"
)

func Test_Voltage_Conversions(t *testing.T) {
	conversionTests := []conversionTest{
		// Metric factory conversions (SI multiples)
		{"kV", "V", 1000},
		{"V", "kV", 0.001},
		{"MV", "V", 1000000},
		{"V", "MV", 0.000001},
		{"mV", "V", 0.001},
		{"V", "mV", 1000},
	}
	testConversions(t, conversionTests)
}

func Test_Voltage_Systems(t *testing.T) {
	si := SiSystem
	assert.Equal(t, si, Volt.System())
	assert.Equal(t, si, YottaVolt.System())
	assert.Equal(t, si, ZettaVolt.System())
	assert.Equal(t, si, ExaVolt.System())
	assert.Equal(t, si, PetaVolt.System())
	assert.Equal(t, si, TeraVolt.System())
	assert.Equal(t, si, GigaVolt.System())
	assert.Equal(t, si, MegaVolt.System())
	assert.Equal(t, si, KiloVolt.System())
	assert.Equal(t, si, HectoVolt.System())
	assert.Equal(t, si, DecaVolt.System())
	assert.Equal(t, si, DeciVolt.System())
	assert.Equal(t, si, CentiVolt.System())
	assert.Equal(t, si, MilliVolt.System())
	assert.Equal(t, si, MicroVolt.System())
	assert.Equal(t, si, NanoVolt.System())
	assert.Equal(t, si, PicoVolt.System())
}

func Test_Voltage_UnitSystem(t *testing.T) {
	assert.Equal(t, SiSystem, Volt.System())
	assert.Equal(t, SiSystem, YottaVolt.System())
	assert.Equal(t, SiSystem, ZettaVolt.System())
	assert.Equal(t, SiSystem, ExaVolt.System())
	assert.Equal(t, SiSystem, PetaVolt.System())
	assert.Equal(t, SiSystem, TeraVolt.System())
	assert.Equal(t, SiSystem, GigaVolt.System())
	assert.Equal(t, SiSystem, MegaVolt.System())
	assert.Equal(t, SiSystem, KiloVolt.System())
	assert.Equal(t, SiSystem, HectoVolt.System())
	assert.Equal(t, SiSystem, DecaVolt.System())
	assert.Equal(t, SiSystem, DeciVolt.System())
	assert.Equal(t, SiSystem, CentiVolt.System())
	assert.Equal(t, SiSystem, MilliVolt.System())
	assert.Equal(t, SiSystem, MicroVolt.System())
	assert.Equal(t, SiSystem, NanoVolt.System())
	assert.Equal(t, SiSystem, PicoVolt.System())
}

func Test_Voltage_MetricFactory_BaseUnits(t *testing.T) {
	assert.Equal(t, Volt, CentiVolt.Base())
	assert.Equal(t, Volt, KiloVolt.Base())
	assert.Equal(t, Volt, MegaVolt.Base())
}

func Test_Voltage_Lookup_Names_and_Symbols(t *testing.T) {
	tests := lookUpTests{
		{Volt, "V"},
		{Volt, "volt"},
		{YottaVolt, "YV"},
		{ZettaVolt, "ZV"},
		{ExaVolt, "EV"},
		{PetaVolt, "PV"},
		{TeraVolt, "TV"},
		{GigaVolt, "GV"},
		{MegaVolt, "MV"},
		{KiloVolt, "kV"},
		{HectoVolt, "hV"},
		{DecaVolt, "daV"},
		{DeciVolt, "dV"},
		{CentiVolt, "cV"},
		{MilliVolt, "mV"},
		{MicroVolt, "μV"},
		{NanoVolt, "nV"},
		{PicoVolt, "pV"},
	}
	testLookupNamesAndSymbols(t, tests)
}
