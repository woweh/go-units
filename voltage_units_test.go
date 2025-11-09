package units

import (
	"testing"
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

func Test_Voltage_UnitSystems(t *testing.T) {
	tests := []unitSystemTest{
		{Volt, SiSystem},
		{YottaVolt, SiSystem},
		{ZettaVolt, SiSystem},
		{ExaVolt, SiSystem},
		{PetaVolt, SiSystem},
		{TeraVolt, SiSystem},
		{GigaVolt, SiSystem},
		{MegaVolt, SiSystem},
		{KiloVolt, SiSystem},
		{HectoVolt, SiSystem},
		{DecaVolt, SiSystem},
		{DeciVolt, SiSystem},
		{CentiVolt, SiSystem},
		{MilliVolt, SiSystem},
		{MicroVolt, SiSystem},
		{NanoVolt, SiSystem},
		{PicoVolt, SiSystem},
	}
	testUnitSystems(t, tests)
}
