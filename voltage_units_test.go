package units

import (
	"testing"

	"github.com/stretchr/testify/assert"
	ns "github.com/Necoro/go-units/numericstring"
)

func Test_Voltage_Conversions(t *testing.T) {
	var conversionTests = []conversionTest{
		{from: "volt", to: "volt", val: ns.One},
		{from: "volt", to: "yottavolt", val: ns.Septillionth},
		{from: "volt", to: "zettavolt", val: ns.Sextillionth},
		{from: "volt", to: "exavolt", val: ns.Quintillionth},
		{from: "volt", to: "petavolt", val: ns.Quadrillionth},
		{from: "volt", to: "teravolt", val: ns.Trillionth},
		{from: "volt", to: "gigavolt", val: ns.Billionth},
		{from: "volt", to: "megavolt", val: ns.Millionth},
		{from: "volt", to: "kilovolt", val: ns.Thousandth},
		{from: "volt", to: "hectovolt", val: ns.Hundredth},
		{from: "volt", to: "decavolt", val: ns.Tenth},
		{from: "volt", to: "decivolt", val: ns.Ten},
		{from: "volt", to: "centivolt", val: ns.Hundred},
		{from: "volt", to: "millivolt", val: ns.Thousand},
		{from: "volt", to: "microvolt", val: ns.Million},
		{from: "volt", to: "nanovolt", val: ns.Billion},
		{from: "volt", to: "picovolt", val: ns.Trillion},
	}

	testConversions(t, conversionTests)
}

func Test_Voltage_Systems(t *testing.T) {
	si := SiSystem
	assert.Equal(t, si, Volt.System())
	assert.Equal(t, si, YottaVolt.System())
	assert.Equal(t, si, ZettaVolt.System())
	assert.Equal(t, si, ExaVolt.System())
	assert.Equal(t, si, PetaVolt.System())
	assert.Equal(t, si, TeraVolt.System())
	assert.Equal(t, si, GigaVolt.System())
	assert.Equal(t, si, MegaVolt.System())
	assert.Equal(t, si, KiloVolt.System())
	assert.Equal(t, si, HectoVolt.System())
	assert.Equal(t, si, DecaVolt.System())
	assert.Equal(t, si, DeciVolt.System())
	assert.Equal(t, si, CentiVolt.System())
	assert.Equal(t, si, MilliVolt.System())
	assert.Equal(t, si, MicroVolt.System())
	assert.Equal(t, si, NanoVolt.System())
	assert.Equal(t, si, PicoVolt.System())
}
