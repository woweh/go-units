package units

import (
	"testing"

	ns "github.com/woweh/go-units/numericstring"
	"github.com/alecthomas/assert/v2"
)

func Test_Force_Conversions(t *testing.T) {
	var conversionTests = []conversionTest{
		{"newton", "kilonewton", ns.Thousandth},
		{"newton", "newton", ns.One},
		{"newton", "millinewton", ns.Thousand},
		{"newton", "micronewton", ns.Million},
		{"newton", "nanonewton", ns.Billion},
		{"newton", "dyne", ns.HundredThousand},
		{"newton", "pound force", "0.2248089"},
		{"newton", "poundal", "7.233014"},
		{"newton", "kilogram-force", "0.1019716"},
		{"newton", "tonne-force", "0.0001019716"},
	}

	testConversions(t, conversionTests)
}

func Test_Force_Systems(t *testing.T) {
	si := SiSystem
	assert.Equal(t, si, Newton.System())
	assert.Equal(t, si, CentiNewton.System())
	assert.Equal(t, si, DeciNewton.System())
	assert.Equal(t, si, MilliNewton.System())
	assert.Equal(t, si, MicroNewton.System())
	assert.Equal(t, si, NanoNewton.System())
	assert.Equal(t, si, PicoNewton.System())
	assert.Equal(t, si, FemtoNewton.System())
	assert.Equal(t, si, AttoNewton.System())
	assert.Equal(t, si, ZeptoNewton.System())
	assert.Equal(t, si, YoctoNewton.System())
	assert.Equal(t, si, DecaNewton.System())
	assert.Equal(t, si, HectoNewton.System())
	assert.Equal(t, si, KiloNewton.System())
	assert.Equal(t, si, MegaNewton.System())
	assert.Equal(t, si, GigaNewton.System())
	assert.Equal(t, si, TeraNewton.System())
	assert.Equal(t, si, PetaNewton.System())
	assert.Equal(t, si, ExaNewton.System())
	assert.Equal(t, si, ZettaNewton.System())
	assert.Equal(t, si, YottaNewton.System())

	bi := BiSystem
	assert.Equal(t, bi, PoundForce.System())

	cgs := CgsSystem
	assert.Equal(t, cgs, Dyne.System())

	mkps := MKpSSystem
	assert.Equal(t, mkps, KilogramForce.System())
	assert.Equal(t, mkps, TonneForce.System())
}
