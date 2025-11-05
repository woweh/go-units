package units

import (
	"testing"

	"github.com/alecthomas/assert/v2"
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
	si := SiSystem

	assert.Equal(t, si, Coulomb.System())
	assert.Equal(t, si, KiloCoulomb.System())
	assert.Equal(t, si, MilliCoulomb.System())
	assert.Equal(t, si, MicroCoulomb.System())
	assert.Equal(t, si, AmpereHour.System())
	assert.Equal(t, si, KiloAmpereHour.System())
	assert.Equal(t, si, MilliAmpereHour.System())
	assert.Equal(t, si, AmpereMinute.System())
	assert.Equal(t, si, KiloAmpereMinute.System())
	assert.Equal(t, si, MilliAmpereMinute.System())
	assert.Equal(t, si, AmpereSecond.System())
	assert.Equal(t, si, KiloAmpereSecond.System())
	assert.Equal(t, si, MilliAmpereSecond.System())
}

func Test_ElectricCharge_BaseUnits(t *testing.T) {
	// Only a few representative metric units
	assert.Equal(t, Coulomb, Coulomb.Base())
	assert.Equal(t, Coulomb, KiloCoulomb.Base())
	assert.Equal(t, Coulomb, MilliCoulomb.Base())
	assert.Equal(t, Coulomb, AmpereHour.Base())
	assert.Equal(t, Coulomb, AmpereMinute.Base())
	assert.Equal(t, Coulomb, AmpereSecond.Base())
}

func Test_Lookup_ElectricCharge_Names_and_Symbols(t *testing.T) {
	tests := lookUpTests{
		{Coulomb, "coulomb"}, {Coulomb, "C"},
		{KiloCoulomb, "kilocoulomb"}, {KiloCoulomb, "kC"},
		{MilliCoulomb, "millicoulomb"}, {MilliCoulomb, "mC"},
		{MicroCoulomb, "microcoulomb"}, {MicroCoulomb, "μC"}, {MicroCoulomb, "uC"},
		{AmpereHour, "ampere-hour"}, {AmpereHour, "A·h"}, {AmpereHour, "A⋅h"}, {AmpereHour, "A*h"}, {AmpereHour, "A.h"}, {AmpereHour, "Ah"}, {AmpereHour, "AHr"},
		{KiloAmpereHour, "kiloampere-hour"}, {KiloAmpereHour, "kA·h"},
		{MilliAmpereHour, "milliampere-hour"}, {MilliAmpereHour, "mA·h"},
		{AmpereMinute, "ampere-minute"}, {AmpereMinute, "A·min"}, {AmpereMinute, "A⋅min"}, {AmpereMinute, "A*min"}, {AmpereMinute, "A.min"}, {AmpereMinute, "Amin"},
		{KiloAmpereMinute, "kiloampere-minute"}, {KiloAmpereMinute, "kA·min"},
		{MilliAmpereMinute, "milliampere-minute"}, {MilliAmpereMinute, "mA·min"},
		{AmpereSecond, "ampere-second"}, {AmpereSecond, "A·s"}, {AmpereSecond, "A⋅s"}, {AmpereSecond, "A*s"}, {AmpereSecond, "A.s"}, {AmpereSecond, "As"},
		{KiloAmpereSecond, "kiloampere-second"}, {KiloAmpereSecond, "kA·s"},
		{MilliAmpereSecond, "milliampere-second"}, {MilliAmpereSecond, "mA·s"},
	}
	testLookupNamesAndSymbols(t, tests)
}
