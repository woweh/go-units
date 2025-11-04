package units

import (
	"testing"

	"github.com/alecthomas/assert/v2"
)

func TestValue_AsBaseUnit(t *testing.T) {
	tests := []struct {
		name  string
		v     Value
		baseV Value
	}{
		{"KiloHertz is converted to Hertz", NewValue(1, KiloHertz), NewValue(1000, Hertz)},
		{"Negative MegaHertz is converted to negative Hertz", NewValue(-1, MegaHertz), NewValue(-1000*1000, Hertz)},
		{"MilliMeter is converted to Meter", NewValue(1, MilliMeter), NewValue(0.001, Meter)},
		{"Negative MilliMeter is converted to negative Meter", NewValue(-1, MilliMeter), NewValue(-0.001, Meter)},
		{"Non-metric unit Foot remains unchanged", NewValue(1, Foot), NewValue(1, Foot)},
		{"KiloOhm is converted to Ohm", NewValue(2.5, KiloOhm), NewValue(2500, Ohm)},
		{"Base unit Hertz remains unchanged", NewValue(1000, Hertz), NewValue(1000, Hertz)},
		{"KiloByte is converted to Byte", NewValue(2.1, KiloByte), NewValue(2100, Byte)},
		// MH2O and derived units
		{"MH2O base unit remains unchanged", NewValue(1, MH2O), NewValue(1, MH2O)},
		{"MilliMH2O is converted to MH2O", NewValue(1000, MilliMH2O), NewValue(1, MH2O)},
		{"CentiMH2O is converted to MH2O", NewValue(100, CentiMH2O), NewValue(1, MH2O)},
		{"Fractional MH2O is converted to MH2O", NewValue(0.001, MH2O), NewValue(0.001, MH2O)},
		{"Negative fractional MH2O is converted to MH2O", NewValue(-0.001, MH2O), NewValue(-0.001, MH2O)},
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
		{"Hertz base unit remains unchanged", NewValue(1, Hertz), NewValue(1, Hertz)},
		{"Hertz is humanized to kilohertz", NewValue(1000, Hertz), NewValue(1, KiloHertz)},
		{"Negative hertz is humanized to negative kilohertz", NewValue(-1000, Hertz), NewValue(-1, KiloHertz)},
		{"Meter is humanized to millimeter", NewValue(0.001, Meter), NewValue(1, MilliMeter)},
		{"Negative meter is humanized to negative millimeter", NewValue(-0.001, Meter), NewValue(-1, MilliMeter)},
		{"Non-metric unit foot remains unchanged", NewValue(1000, Foot), NewValue(1000, Foot)},
		{"Ohm is humanized to kiloohm", NewValue(2500, Ohm), NewValue(2.5, KiloOhm)},
		{"Decihertz is humanized to megahertz in multiple steps", NewValue(50000000, DeciHertz), NewValue(5, MegaHertz)},
		{"Hertz does not need humanization", NewValue(200, Hertz), NewValue(200, Hertz)},
		{"Byte is humanized to kilobyte", NewValue(2048, Byte), NewValue(2.048, KiloByte)},
		// MH2O and derived units
		{"MH2O base unit remains unchanged", NewValue(1, MH2O), NewValue(1, MH2O)},
		{"MilliMH2O is humanized to MH2O", NewValue(1000, MilliMH2O), NewValue(1, MH2O)},
		{"CentiMH2O is humanized to MH2O", NewValue(100, CentiMH2O), NewValue(1, MH2O)},
		{"MH2O fraction is humanized to MilliMH2O", NewValue(0.001, MH2O), NewValue(1, MilliMH2O)},
		{"Negative MH2O fraction is humanized to negative MilliMH2O", NewValue(-0.001, MH2O), NewValue(-1, MilliMH2O)},
	}

	for _, tst := range tests {
		t.Run(tst.name, func(t *testing.T) {
			humV := tst.v.Humanize()
			assert.Equal(t, tst.humV, humV)
		})
	}
}

func TestValue_NewValue_Unit_Float(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name string
		val  float64
		unit Unit
	}{
		{"basic", 42, Hertz},
		{"zero", 0, Meter},
		{"negative", -5, Ohm},
		{"nil unit", 1.23, nil},
	}
	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			v := NewValue(tc.val, tc.unit)
			assert.Equal(t, tc.val, v.Float())
			assert.Equal(t, tc.unit, v.Unit())
		})
	}
}

func TestValue_String_Fmt(t *testing.T) {
	t.Parallel()
	v := NewValue(1234.5678, Hertz)
	cases := []struct {
		name string
		val  Value
		opt  FmtOptions
		exp  string
	}{
		{"default", v, DefaultFmtOptions, "1234.5678 hertz"},
		{"short", v, FmtOptions{Label: true, Short: true, Precision: 2, TrimTrailingZeroes: false, Separator: "-"}, "1234.57-Hz"},
		{"no label", v, FmtOptions{Label: false, Short: false, Precision: 1, TrimTrailingZeroes: false, Separator: " "}, "1234.6"},
		{"trim zeroes", NewValue(1.2000, Meter), DefaultFmtOptions, "1.2 meters"},
		{"nil unit", NewValue(5, nil), DefaultFmtOptions, "5"},
		{"singular 1", NewValue(1, Meter), DefaultFmtOptions, "1 meter"},
		{"singular -1", NewValue(-1, Meter), DefaultFmtOptions, "-1 meter"},
		{"zero", NewValue(0, Meter), DefaultFmtOptions, "0 meters"},
		{"fractional < +1", NewValue(0.99, Meter), DefaultFmtOptions, "0.99 meters"},
		{"fractional > +1", NewValue(1.2, Meter), DefaultFmtOptions, "1.2 meters"},
		{"negative fractional < -1", NewValue(-1.2, Meter), DefaultFmtOptions, "-1.2 meters"},
		{"negative fractional > -1", NewValue(-0.99, Meter), DefaultFmtOptions, "-0.99 meters"},
		{"negative zero", NewValue(-0.0, Meter), DefaultFmtOptions, "0 meters"},
		{"DMS", NewValue(12.3045, DMS), DefaultFmtOptions, "12°30'45''"},
	}
	for _, tc := range cases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			got := tc.val.Fmt(tc.opt)
			assert.Equal(t, tc.exp, got)
		})
	}
}

func Test_trimTrailing(t *testing.T) {
	t.Parallel()
	cases := []struct {
		in  string
		out string
	}{
		{"123.45000", "123.45"},
		{"100.000", "100"},
		{"0.000", "0"},
		{"42", "42"},
		{"1.230", "1.23"},
		{"0.0", "0"},
		{"", ""},
	}
	for _, tc := range cases {
		tc := tc
		t.Run(tc.in, func(t *testing.T) {
			got := trimTrailing(tc.in)
			assert.Equal(t, tc.out, got)
		})
	}
}

func TestValue_MustConvert_Convert(t *testing.T) {
	t.Parallel()
	v := NewValue(1, KiloHertz)
	// Success
	v2 := v.MustConvert(Hertz)
	assert.Equal(t, NewValue(1000, Hertz), v2)
	// Same unit
	v3 := v.MustConvert(KiloHertz)
	assert.Equal(t, v, v3)
	// Nil unit
	v4 := NewValue(1, nil)
	assert.Equal(t, v4, v4.MustConvert(nil))
	// Error case (simulate by passing an incompatible unit)
	defer func() {
		if r := recover(); r == nil {
			t.Error("expected panic for invalid conversion")
		}
	}()
	_ = v.MustConvert(Meter)
}
