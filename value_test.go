package units

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func nv(val float64, u *Unit) Value {
	return NewValue(val, *u)
}

func TestValue_AsBaseUnit(t *testing.T) {
	tests := []struct {
		name  string
		v     Value
		baseV Value
	}{
		{"simple", nv(1, KiloHertz), nv(1000, Hertz)},
		{"negative", nv(-1, MegaHertz), nv(-1000*1000, Hertz)},
		{"fraction", nv(1, MilliMeter), nv(0.001, Meter)},
		{"neg. fraction", nv(-1, MilliMeter), nv(-0.001, Meter)},
		{"non metric", nv(1, Foot), nv(1, Foot)},
		{"unequal to one", nv(2.5, KiloOhm), nv(2500, Ohm)},
		{"base", nv(1000, Hertz), nv(1000, Hertz)},
		{"data", nv(2.1, KiloByte), nv(2100, Byte)},
	}

	for _, tst := range tests {
		t.Run(tst.name, func(t *testing.T) {
			baseV := tst.v.AsBaseUnit()
			assert.Equal(t, tst.baseV, baseV)
		})
	}
}

func TestValue_Humanize(t *testing.T) {
	tests := []struct {
		name string
		v    Value
		humV Value
	}{
		{"base", nv(1, Hertz), nv(1, Hertz)},
		{"simple", nv(1000, Hertz), nv(1, KiloHertz)},
		{"negative", nv(-1000, Hertz), nv(-1, KiloHertz)},
		{"fraction", nv(0.001, Meter), nv(1, MilliMeter)},
		{"neg. fraction", nv(-0.001, Meter), nv(-1, MilliMeter)},
		{"non metric", nv(1000, Foot), nv(1000, Foot)},
		{"unequal to one", nv(2500, Ohm), nv(2.5, KiloOhm)},
		{"multiple steps", nv(50000000, DeciHertz), nv(5, MegaHertz)},
		{"only in steps 3", nv(200, Hertz), nv(200, Hertz)},
		{"data", nv(2048, Byte), nv(2.048, KiloByte)},
	}

	for _, tst := range tests {
		t.Run(tst.name, func(t *testing.T) {
			humV := tst.v.Humanize()
			assert.Equal(t, tst.humV, humV)
		})
	}
}
