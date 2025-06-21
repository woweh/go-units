package units

import (
	"math"
	"strconv"
	"strings"
)

var DefaultFmtOptions = FmtOptions{
	Label:              true,
	Short:              false,
	Precision:          6,
	TrimTrailingZeroes: true,
}

type FmtOptions struct {
	Label              bool // if false, unit label/symbol will be omitted
	Short              bool // if true, use unit shortname or symbol
	Precision          int  // maximum meaningful precision to truncate value
	TrimTrailingZeroes bool // trim trailing zeroes
}

type Value struct {
	val  float64
	unit Unit
}

// NewValue creates a new Value instance
func NewValue(v float64, u Unit) Value { return Value{v, u} }

// Unit returns the Unit of this Value
func (v Value) Unit() Unit { return v.unit }

// Float returns the float value of this Value
func (v Value) Float() float64 { return v.val }

// String returns a string representation of this Value
func (v Value) String() string { return v.Fmt(DefaultFmtOptions) }

// Fmt returns a string representation of this Value, using the provided FmtOptions
func (v Value) Fmt(opts FmtOptions) string {
	if v.unit == nil {
		opts.Label = false
	}

	prec := opts.Precision
	// expand precision if needed to present meaningful value
	if v.val < 1 && v.val > 0 {
		prec = int((math.Log10(v.val)-0.5)*-1) + prec
	}

	vstr := strconv.FormatFloat(v.val, 'f', prec, 64)
	if opts.TrimTrailingZeroes {
		vstr = trimTrailing(vstr)
	}

	if !opts.Label {
		return vstr
	}

	var label string

	if opts.Short {
		label = v.unit.Symbol
	} else {
		label = v.unit.Name
		// make label plural if needed
		if v.val > 1.0 {
			label = v.unit.PluralName()
		}
	}
	return vstr + " " + label
}

// MustConvert converts this Value to another Unit, PANICKING on error
func (v Value) MustConvert(to Unit) Value {
	newV, err := v.Convert(to)
	if err != nil {
		panic(err)
	}
	return newV
}

// Convert converts this Value to another Unit
func (v Value) Convert(to Unit) (Value, error) {
	// allow converting to same unit
	if v.unit == nil || v.unit.Name == to.Name {
		return v, nil
	}

	return ConvertFloat(v.val, v.unit, to)
}

// AsBaseUnit converts this Value to its base unit (i.e., 2km to 2000m).
func (v Value) AsBaseUnit() Value {
	if v.unit == nil || !v.unit.IsMetric() {
		return v
	}

	// we can use "Must" here because we know that the unit is metric
	// and then this conversion is defined.
	return v.MustConvert(v.unit.BaseUnit())
}

// Humanize converts this Value to a human-readable format (i.e., 2000m to 2km).
func (v Value) Humanize() Value {
	if v.val == 0 || v.unit == nil || !v.unit.IsMetric() {
		return v
	}

	// Algorithm taken from go-humanize
	baseV := v.AsBaseUnit()
	mag := math.Abs(baseV.val)
	exponent := math.Floor(math.Log10(mag))
	exponent = math.Floor(exponent/3) * 3

	value := mag / math.Pow(10, exponent)

	// Handle special case where value is exactly 1000.0
	// Should return 1 M instead of 1000 k
	if value == 1000.0 {
		exponent += 3
	}

	scaledUnit := findMaxUnitForExp(baseV.unit, exponent)
	return baseV.MustConvert(scaledUnit)
}

// Trim trailing zeros from formatted float string
func trimTrailing(s string) string {
	if !strings.ContainsRune(s, '.') {
		return s
	}
	s = strings.TrimRight(s, "0")
	if s == "" {
		return "0"
	}
	if s[len(s)-1] == '.' {
		s = strings.TrimSuffix(s, ".")
	}
	return s
}
