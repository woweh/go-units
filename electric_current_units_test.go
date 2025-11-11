package units

import (
	"testing"
)

func Test_ElectricCurrent_Conversions(t *testing.T) {
	conversionTests := []conversionTest{
		{from: "kiloampere", to: "ampere", exp: 1000},
		{from: "megaampere", to: "kiloampere", exp: 1000},
		{from: "milliampere", to: "ampere", exp: 0.001},
		{from: "microampere", to: "ampere", exp: 0.000001},
		{from: "nanoampere", to: "ampere", exp: 0.000000001},
		{from: "gigaampere", to: "megaampere", exp: 1000},
		{from: "teraampere", to: "gigaampere", exp: 1000},
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
