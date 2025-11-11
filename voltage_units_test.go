package units

import (
	"testing"
)

func Test_Voltage_Conversions(t *testing.T) {
	conversionTests := []conversionTest{
		// Metric factory conversions (SI multiples)
		{from: "kV", to: "V", exp: 1000},
		{from: "V", to: "kV", exp: 0.001},
		{from: "MV", to: "V", exp: 1000000},
		{from: "V", to: "MV", exp: 0.000001},
		{from: "mV", to: "V", exp: 0.001},
		{from: "V", to: "mV", exp: 1000},
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
