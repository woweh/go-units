package units

import (
	"testing"
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
	tests := []unitSystemTest{
		{Ampere, SiSystem},
		{MilliAmpere, SiSystem},
		{MicroAmpere, SiSystem},
		{NanoAmpere, SiSystem},
		{PicoAmpere, SiSystem},
		{FemtoAmpere, SiSystem},
		{AttoAmpere, SiSystem},
		{ZeptoAmpere, SiSystem},
		{YoctoAmpere, SiSystem},
		{KiloAmpere, SiSystem},
		{MegaAmpere, SiSystem},
		{GigaAmpere, SiSystem},
		{TeraAmpere, SiSystem},
		{PetaAmpere, SiSystem},
		{ExaAmpere, SiSystem},
		{ZettaAmpere, SiSystem},
		{YottaAmpere, SiSystem},
	}
	testUnitSystems(t, tests)
}
