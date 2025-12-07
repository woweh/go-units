package units

import (
	"testing"
)

func Test_ElectricalResistance_Conversions(t *testing.T) {
	conversionTests := []conversionTest{
		// Metric progressions (sampled)
		{from: "kiloohm", to: "ohm", exp: 1000},
		{from: "megaohm", to: "kiloohm", exp: 1000},
		{from: "milliohm", to: "ohm", exp: 0.001},
		{from: "microohm", to: "ohm", exp: 0.000001},
		{from: "gigaohm", to: "megaohm", exp: 1000},
		{from: "teraohm", to: "gigaohm", exp: 1000},
	}
	testConversions(t, conversionTests)
}

func Test_ElectricalResistance_UnitSystems(t *testing.T) {
	tests := []unitSystemTest{
		{Ohm, SiSystem},
		{KiloOhm, SiSystem},
		{MegaOhm, SiSystem},
		{MilliOhm, SiSystem},
		{MicroOhm, SiSystem},
		{GigaOhm, SiSystem},
		{TeraOhm, SiSystem},
		{DecaOhm, SiSystem},
		{HectoOhm, SiSystem},
		{DeciOhm, SiSystem},
		{CentiOhm, SiSystem},
		{PicoOhm, SiSystem},
		{FemtoOhm, SiSystem},
		{AttoOhm, SiSystem},
		{ZeptoOhm, SiSystem},
		{YoctoOhm, SiSystem},
		{PetaOhm, SiSystem},
		{ExaOhm, SiSystem},
		{ZettaOhm, SiSystem},
		{YottaOhm, SiSystem},
		{OhmMeter, SiSystem},
	}
	testUnitSystems(t, tests)
}
