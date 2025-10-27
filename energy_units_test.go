package units

import (
	"testing"

	"github.com/alecthomas/assert/v2"
)

func Test_Energy_Conversions(t *testing.T) {
	conversionTests := []conversionTest{
		// SI base and metric
		{"joule", "calorie", "0.2390057"},
		{"joule", "electronvolt", "6241507648655547392"},
		{"joule", "megawatt-hour", "0.0000000002777778"},
		{"joule", "kilowatt-hour", "0.0000002777778"},
		{"joule", "megajoule", "0.000001"},
		{"joule", "gigajoule", "0.000000001"},
		// Imperial/other
		{"joule", "british thermal unit", "0.000947817"},
		{"calorie", "joule", "4.184"},
		{"watt-hour", "joule", "3600"},
		{"british thermal unit", "joule", "1055.05585262"},
	}
	testConversions(t, conversionTests)
}

func Test_Energy_UnitSystems(t *testing.T) {
	si := SiSystem
	assert.Equal(t, si, Joule.System())
	assert.Equal(t, si, KiloJoule.System())
	assert.Equal(t, si, MegaJoule.System())
	assert.Equal(t, si, GigaJoule.System())
	assert.Equal(t, si, TeraJoule.System())
	assert.Equal(t, si, PetaJoule.System())
	assert.Equal(t, si, ExaJoule.System())
	assert.Equal(t, si, ZettaJoule.System())
	assert.Equal(t, si, YottaJoule.System())
	assert.Equal(t, si, MilliJoule.System())
	assert.Equal(t, si, MicroJoule.System())
	assert.Equal(t, si, NanoJoule.System())
	assert.Equal(t, si, PicoJoule.System())
	assert.Equal(t, si, FemtoJoule.System())
	assert.Equal(t, si, AttoJoule.System())
}

func Test_Energy_BaseUnits(t *testing.T) {
	assert.Equal(t, Joule, Joule.Base())
	assert.Equal(t, Joule, KiloJoule.Base())
	assert.Equal(t, Joule, MegaJoule.Base())
	assert.Equal(t, Joule, GigaJoule.Base())
	assert.Equal(t, Joule, TeraJoule.Base())
	assert.Equal(t, Joule, PetaJoule.Base())
	assert.Equal(t, Joule, ExaJoule.Base())
	assert.Equal(t, Joule, ZettaJoule.Base())
	assert.Equal(t, Joule, YottaJoule.Base())
}
