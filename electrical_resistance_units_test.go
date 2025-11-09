package units

import (
	"testing"
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
