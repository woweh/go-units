package units

import (
	"testing"

	"github.com/stretchr/testify/assert"
	ns "github.com/woweh/go-units/numericstring"
)

var powerConvTests = []conversionTest{
	{from: "watt", to: "watt", val: "1"},
	{from: "watt", to: "deciwatt", val: ns.Txp1},
	{from: "watt", to: "centiwatt", val: ns.Txp2},
	{from: "watt", to: "milliwatt", val: ns.Txp3},
	{from: "watt", to: "microwatt", val: ns.Txp6},
	{from: "watt", to: "nanowatt", val: ns.Txp9},
	{from: "watt", to: "picowatt", val: ns.Txp12},
	// the following produce rounding errors
	//{from: "watt", to: "femtowatt", val: ns.Txp15},
	//{from: "watt", to: "attowatt", val: ns.Txp18},
	//{from: "watt", to: "zeptowatt", val: ns.Txp21},
	//{from: "watt", to: "yoctowatt", val: ns.Txp24},
	{from: "watt", to: "decaWatt", val: ns.Tmn1},
	{from: "watt", to: "kilowatt", val: ns.Tmn3},
	{from: "watt", to: "megawatt", val: ns.Tmn6},
	{from: "watt", to: "gigawatt", val: ns.Tmn9},
	{from: "watt", to: "terawatt", val: ns.Tmn12},
	{from: "watt", to: "petawatt", val: ns.Tmn15},
	{from: "watt", to: "exawatt", val: ns.Tmn18},
	{from: "watt", to: "zettawatt", val: ns.Tmn21},
	{from: "watt", to: "yottawatt", val: ns.Tmn24},
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
}
