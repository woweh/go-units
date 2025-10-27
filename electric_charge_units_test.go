package units

import (
	"testing"

	"github.com/alecthomas/assert/v2"
)

func Test_ElectricCharge_Conversions(t *testing.T) {
	conversionTests := []conversionTest{
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

func Test_ElectricCharge_UnitSystems(t *testing.T) {
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

func Test_ElectricCharge_BaseUnits(t *testing.T) {
	assert.Equal(t, Coulomb, Coulomb.Base())
	assert.Equal(t, Coulomb, ExaCoulomb.Base())
	assert.Equal(t, Coulomb, PetaCoulomb.Base())
	assert.Equal(t, Coulomb, TeraCoulomb.Base())
	assert.Equal(t, Coulomb, GigaCoulomb.Base())
	assert.Equal(t, Coulomb, MegaCoulomb.Base())
	assert.Equal(t, Coulomb, KiloCoulomb.Base())
	assert.Equal(t, Coulomb, HectoCoulomb.Base())
	assert.Equal(t, Coulomb, DecaCoulomb.Base())
	assert.Equal(t, Coulomb, DeciCoulomb.Base())
	assert.Equal(t, Coulomb, CentiCoulomb.Base())
	assert.Equal(t, Coulomb, MilliCoulomb.Base())
	assert.Equal(t, Coulomb, MicroCoulomb.Base())
	assert.Equal(t, Coulomb, NanoCoulomb.Base())
	assert.Equal(t, Coulomb, PicoCoulomb.Base())
	assert.Equal(t, Coulomb, FemtoCoulomb.Base())
	assert.Equal(t, Coulomb, AttoCoulomb.Base())

	assert.Equal(t, Coulomb, AmpereHour.Base())
	assert.Equal(t, Coulomb, KiloAmpereHour.Base())
	assert.Equal(t, Coulomb, MilliAmpereHour.Base())
}
