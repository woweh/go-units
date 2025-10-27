package units

import (
	"testing"

	"github.com/alecthomas/assert/v2"
)

var magNames = []string{
	"exa",
	"peta",
	"tera",
	"giga",
	"mega",
	"kilo",
	"hecto",
	"deca",
	"deci",
	"centi",
	"milli",
	"micro",
	"nano",
	"pico",
	"femto",
	"atto",
}

type magFn func(Unit, ...UnitOption) Unit

func Test_Magnitudes(t *testing.T) {
	u := newUnit("dong", "₫")
	for i, mfn := range []magFn{
		Exa, Peta, Tera, Giga, Mega, Kilo, Hecto, Deca, Deci, Centi, Milli, Micro, Nano, Pico, Femto, Atto,
	} {
		mu := mfn(u)
		t.Logf("created mag unit: %s (%s)", mu.Name, mu.Symbol)
		assert.Equal(t, mu.Name, magNames[i]+"dong")
	}
}

func Test_MagnitudeForExp(t *testing.T) {
	assert.Equal(t, mags["kilo"], magnitudeForExp(3.0))
	assert.Equal(t, mags["centi"], magnitudeForExp(-2.0))
	assert.Equal(t, mags["atto"], magnitudeForExp(-18.0))
	assert.Equal(t, mags["yotta"], magnitudeForExp(24.0))

	// does not exist
	assert.Zero(t, magnitudeForExp(4.0).Power)
	assert.Zero(t, magnitudeForExp(1.2).Power)
	assert.Zero(t, magnitudeForExp(0).Power)
}

func Test_FindMaxUnitForExp(t *testing.T) {
	tests := []struct {
		name   string
		exp    float64
		expect *Unit
		base   *Unit
	}{
		{"Zero", 0, Hertz, Hertz},
		{"Non metric", 3, Foot, Foot},
		{"Simple", 3, KiloHertz, Hertz},
		{"Simple negative", -3, MilliHertz, Hertz},
		{"Simple, unusual", 2, HectoHertz, Hertz},
		{"Not defined prefix", 2, VoltAmpere, VoltAmpere},
		{"Partially defined, works", 3, KiloVoltAmpere, VoltAmpere},
		{"Simple, not existing prefix", 5, KiloHertz, Hertz},
		{"Partially defined, over", 9, MegaVoltAmpere, VoltAmpere},
		{"Partially defined negative", -10, VoltAmpere, VoltAmpere},
		{"Data", 3, KiloByte, Byte},
	}

	for _, tst := range tests {
		t.Run(tst.name, func(t *testing.T) {
			calcUnit := findMaxUnitForExp(tst.base, tst.exp)
			assert.Equal(t, tst.expect, calcUnit)
		})
	}
}
