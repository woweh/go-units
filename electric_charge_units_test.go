package units

import (
	"testing"

	"github.com/alecthomas/assert/v2"
)

func Test_ElectricCharge_Conversions(t *testing.T) {
	var conversionTests = []conversionTest{
		{"coulomb", "C", "1"},
		{"ampere-hour", "coulomb", "3600"},
		{"ampere-minute", "coulomb", "60"},
		{"coulomb", "ampere-second", "1"},
		{"kiloampere-hour", "kilocoulomb", "3600"},
		{"kiloampere-minute", "kilocoulomb", "60"},
		{"kiloampere-second", "kilocoulomb", "1"},
		{"milliampere-hour", "millicoulomb", "3600"},
		{"milliampere-minute", "millicoulomb", "60"},
		{"milliampere-second", "millicoulomb", "1"},
	}

	testConversions(t, conversionTests)
}

func Test_ElectricCharge_Symbols(t *testing.T) {
	si := SiSystem

	assert.Equal(t, si, Coulomb.System())
	assert.Equal(t, si, ExaCoulomb.System())
	assert.Equal(t, si, PetaCoulomb.System())
	assert.Equal(t, si, TeraCoulomb.System())
	assert.Equal(t, si, GigaCoulomb.System())
	assert.Equal(t, si, MegaCoulomb.System())
	assert.Equal(t, si, KiloCoulomb.System())
	assert.Equal(t, si, HectoCoulomb.System())
	assert.Equal(t, si, DecaCoulomb.System())
	assert.Equal(t, si, DeciCoulomb.System())
	assert.Equal(t, si, CentiCoulomb.System())
	assert.Equal(t, si, MilliCoulomb.System())
	assert.Equal(t, si, MicroCoulomb.System())
	assert.Equal(t, si, NanoCoulomb.System())
	assert.Equal(t, si, PicoCoulomb.System())
	assert.Equal(t, si, FemtoCoulomb.System())
	assert.Equal(t, si, AttoCoulomb.System())

	assert.Equal(t, si, AmpereHour.System())
	assert.Equal(t, si, KiloAmpereHour.System())
	assert.Equal(t, si, MilliAmpereHour.System())
}
