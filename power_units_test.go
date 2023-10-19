package units

import (
	"testing"

	"github.com/stretchr/testify/assert"
	ns "github.com/woweh/go-units/numericstring"
)

var powerConvTests = []conversionTest{
	{"watt", "watt", "1"},
	{"watt", "deciwatt", ns.Txp1},
	{"watt", "centiwatt", ns.Txp2},
	{"watt", "milliwatt", ns.Txp3},
	{"watt", "microwatt", ns.Txp6},
	{"watt", "nanowatt", ns.Txp9},
	{"watt", "picowatt", ns.Txp12},
	// the following produce rounding errors
	//{ "watt",  "femtowatt",  ns.Txp15},
	//{ "watt",  "attowatt",  ns.Txp18},
	//{ "watt",  "zeptowatt",  ns.Txp21},
	//{ "watt",  "yoctowatt",  ns.Txp24},
	{"watt", "decaWatt", ns.Tmn1},
	{"watt", "kilowatt", ns.Tmn3},
	{"watt", "megawatt", ns.Tmn6},
	{"watt", "gigawatt", ns.Tmn9},
	{"watt", "terawatt", ns.Tmn12},
	{"watt", "petawatt", ns.Tmn15},
	{"watt", "exawatt", ns.Tmn18},
	{"watt", "zettawatt", ns.Tmn21},
	{"watt", "yottawatt", ns.Tmn24},
	{"volt-ampere", "volt-ampere", "1"},
	{"volt-ampere", "kilovolt-ampere", ns.Tmn3},
	{"volt-ampere", "megavolt-ampere", ns.Tmn6},
}

func Test_PowerConversions(t *testing.T) {
	testConversions(t, powerConvTests)
}

func Test_PowerSystems(t *testing.T) {
	si := SiSystem
	assert.Equal(t, si, Watt.System())
	assert.Equal(t, si, DeciWatt.System())
	assert.Equal(t, si, CentiWatt.System())
	assert.Equal(t, si, MilliWatt.System())
	assert.Equal(t, si, MicroWatt.System())
	assert.Equal(t, si, NanoWatt.System())
	assert.Equal(t, si, PicoWatt.System())
	assert.Equal(t, si, FemtoWatt.System())
	assert.Equal(t, si, AttoWatt.System())
	assert.Equal(t, si, ZeptoWatt.System())
	assert.Equal(t, si, YoctoWatt.System())
	assert.Equal(t, si, DecaWatt.System())
	assert.Equal(t, si, HectoWatt.System())
	assert.Equal(t, si, KiloWatt.System())
	assert.Equal(t, si, MegaWatt.System())
	assert.Equal(t, si, GigaWatt.System())
	assert.Equal(t, si, TeraWatt.System())
	assert.Equal(t, si, PetaWatt.System())
	assert.Equal(t, si, ExaWatt.System())
	assert.Equal(t, si, ZettaWatt.System())
	assert.Equal(t, si, YottaWatt.System())

	assert.Equal(t, si, VoltAmpere.System())
	assert.Equal(t, si, KiloVoltAmpere.System())
	assert.Equal(t, si, MegaVoltAmpere.System())
}
