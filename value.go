package units

import (
	"math"
	"strconv"
	"strings"
)

// DefaultFmtOptions is the default formatting options for Values
var DefaultFmtOptions = FmtOptions{
	Label:              true,
	Short:              false,
	Precision:          6,
	TrimTrailingZeroes: true,
	Separator:          " ",
}

// FmtOptions are formatting options for Values
type FmtOptions struct {
	Label              bool   // if false, unit label/symbol will be omitted
	Short              bool   // if true, use unit shortname or symbol
	Precision          int    // maximum meaningful precision to truncate value
	TrimTrailingZeroes bool   // trim trailing zeroes
	Separator          string // separating value from unit
}

// Value represents a value with a unit
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

// Fmt returns a string representation of this Value, using either the provided FmtOptions or the FormatFn of the Unit.
// If the Unit has a FormatFn, the FmtOptions are ignored.
func (v Value) Fmt(opts FmtOptions) string {
	if v.unit == nil {
		opts.Label = false
	} else if v.unit.fmtFn != nil {
		return v.unit.fmtFn(v.val)
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
		// Make label plural if needed. In English, only use the singular for "1" and "-1"
		if vstr != "1" && vstr != "-1" {
			label = v.unit.PluralName()
		}
	}
	return vstr + opts.Separator + label
}

// trimTrailing trims trailing zeros from a formatted float string
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
	if v.unit == nil || v.unit == to {
		return v, nil
	}
	return ConvertFloat(v.val, v.unit, to)
}

// AsBaseUnit converts this Value to its base unit (i.e., 2km to 2000m).
func (v Value) AsBaseUnit() Value {
	if v.unit == nil || !v.unit.HasBase() {
		return v
	}
	// we can use "Must" here because we know that this conversion is defined.
	return v.MustConvert(v.unit.Base())
}

// Humanize converts this Value to a human-readable format (i.e., 2000m to 2km).
func (v Value) Humanize() Value {
	if v.val == 0 || v.unit == nil || !(v.unit.IsBase() || v.unit.HasBase()) {
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

	scaledUnit := findBestMatchingUnit(baseV.unit, int(exponent))
	return baseV.MustConvert(scaledUnit)
}
