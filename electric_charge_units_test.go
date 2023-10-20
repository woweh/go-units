package units

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var electricChargeUnitConvTests = []conversionTest{
	{"coulomb", "C", "1"},
	{"ampere-hour", "coulomb", "3600"},
	{"ampere-minute", "coulomb", "60"},
	{"coulomb", "ampere-second", "1"},
	{"kiloampere-hour", "kilocoulomb", "3600"},
	{"kiloampere-minute", "kilocoulomb", "60"},
	{"kiloampere-second", "kilocoulomb", "1"},
}

func Test_ElectricChargeUnitConversions(t *testing.T) {
	testConversions(t, electricChargeUnitConvTests)
}

func Test_ElectricChargeSymbols(t *testing.T) {
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
