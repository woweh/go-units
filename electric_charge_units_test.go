package units

import (
	"testing"
)

func Test_ElectricCharge_Conversions(t *testing.T) {
	conversionTests := []conversionTest{
		// SI/metric progressions (sampled)
		{from: "kilocoulomb", to: "coulomb", exp: 1000},
		{from: "millicoulomb", to: "coulomb", exp: 0.001},
		{from: "microcoulomb", to: "coulomb", exp: 0.000001},
		// Ampere-hour/minute/second to coulomb (from implementation)
		{from: "ampere-hour", to: "coulomb", exp: 3600},
		{from: "ampere-minute", to: "coulomb", exp: 60},
		{from: "ampere-second", to: "coulomb", exp: 1},
		// Reverse conversions
		{from: "coulomb", to: "ampere-hour", exp: 1.0 / 3600},
		{from: "coulomb", to: "ampere-minute", exp: 1.0 / 60},
		{from: "coulomb", to: "ampere-second", exp: 1},
		// Kilo/milli ampere-hour/minute/second
		{from: "kiloampere-hour", to: "kilocoulomb", exp: 3600},
		{from: "kiloampere-minute", to: "kilocoulomb", exp: 60},
		{from: "kiloampere-second", to: "kilocoulomb", exp: 1},
		{from: "milliampere-hour", to: "millicoulomb", exp: 3600},
		{from: "milliampere-minute", to: "millicoulomb", exp: 60},
		{from: "milliampere-second", to: "millicoulomb", exp: 1},
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
