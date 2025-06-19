package units

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValue_AsBaseUnit(t *testing.T) {
	tests := []struct {
		name  string
		v     Value
		baseV Value
	}{
		{"simple", NewValue(1, KiloHertz), NewValue(1000, Hertz)},
		{"negative", NewValue(-1, MegaHertz), NewValue(-1000*1000, Hertz)},
		{"fraction", NewValue(1, MilliMeter), NewValue(0.001, Meter)},
		{"neg. fraction", NewValue(-1, MilliMeter), NewValue(-0.001, Meter)},
		{"non metric", NewValue(1, Foot), NewValue(1, Foot)},
		{"unequal to one", NewValue(2.5, KiloOhm), NewValue(2500, Ohm)},
		{"base", NewValue(1000, Hertz), NewValue(1000, Hertz)},
		{"data", NewValue(2.1, KiloByte), NewValue(2100, Byte)},
		{"nil", NewValue(2, nil), NewValue(2, nil)},
		{"zero", Value{}, Value{}},
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
		{"base", NewValue(1, Hertz), NewValue(1, Hertz)},
		{"simple", NewValue(1000, Hertz), NewValue(1, KiloHertz)},
		{"negative", NewValue(-1000, Hertz), NewValue(-1, KiloHertz)},
		{"fraction", NewValue(0.001, Meter), NewValue(1, MilliMeter)},
		{"neg. fraction", NewValue(-0.001, Meter), NewValue(-1, MilliMeter)},
		{"non metric", NewValue(1000, Foot), NewValue(1000, Foot)},
		{"unequal to one", NewValue(2500, Ohm), NewValue(2.5, KiloOhm)},
		{"multiple steps", NewValue(50000000, DeciHertz), NewValue(5, MegaHertz)},
		{"only in steps 3", NewValue(200, Hertz), NewValue(200, Hertz)},
		{"data", NewValue(2048, Byte), NewValue(2.048, KiloByte)},
		{"nil", NewValue(2, nil), NewValue(2, nil)},
		{"zero", Value{}, Value{}},
	}

	for _, tst := range tests {
		t.Run(tst.name, func(t *testing.T) {
			humV := tst.v.Humanize()
			assert.Equal(t, tst.humV, humV)
		})
	}
}
