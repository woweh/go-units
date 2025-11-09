package units

import (
	"testing"
)

func Test_ElectricCharge_Conversions(t *testing.T) {
	conversionTests := []conversionTest{
		// SI/metric progressions (sampled)
		{"kilocoulomb", "coulomb", 1000},
		{"millicoulomb", "coulomb", 0.001},
		{"microcoulomb", "coulomb", 0.000001},
		// Ampere-hour/minute/second to coulomb (from implementation)
		{"ampere-hour", "coulomb", 3600},
		{"ampere-minute", "coulomb", 60},
		{"ampere-second", "coulomb", 1},
		// Reverse conversions
		{"coulomb", "ampere-hour", 1.0 / 3600},
		{"coulomb", "ampere-minute", 1.0 / 60},
		{"coulomb", "ampere-second", 1},
		// Kilo/milli ampere-hour/minute/second
		{"kiloampere-hour", "kilocoulomb", 3600},
		{"kiloampere-minute", "kilocoulomb", 60},
		{"kiloampere-second", "kilocoulomb", 1},
		{"milliampere-hour", "millicoulomb", 3600},
		{"milliampere-minute", "millicoulomb", 60},
		{"milliampere-second", "millicoulomb", 1},
	}

	testConversions(t, conversionTests)
}

func Test_ElectricCharge_UnitSystems(t *testing.T) {
	tests := []unitSystemTest{
		{Coulomb, SiSystem},
		{KiloCoulomb, SiSystem},
		{MilliCoulomb, SiSystem},
		{MicroCoulomb, SiSystem},
		{AmpereHour, SiSystem},
		{KiloAmpereHour, SiSystem},
		{MilliAmpereHour, SiSystem},
		{AmpereMinute, SiSystem},
		{KiloAmpereMinute, SiSystem},
		{MilliAmpereMinute, SiSystem},
		{AmpereSecond, SiSystem},
		{KiloAmpereSecond, SiSystem},
		{MilliAmpereSecond, SiSystem},
	}
	testUnitSystems(t, tests)
}
