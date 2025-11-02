package units

import (
	"math"
)

// magnitude defines a metric prefix with its symbol and power of ten
type magnitude struct {
	// Symbol is the prefix symbol
	Symbol string
	// Prefix is the full name of the prefix
	Prefix string
	// Power is the power of ten associated with the prefix
	Power int
}

const (
	// MaxExponent defines the maximum exponent for metric prefixes
	MaxExponent = 30
	// MinExponent defines the minimum exponent for metric prefixes
	MinExponent = -30
)

// _mags holds the standard SI metric prefixes with their symbols and powers, keyed by exponent
var _mags = map[int]magnitude{
	30:  {"Q", "quetta", 30},
	27:  {"R", "ronna", 27},
	24:  {"Y", "yotta", 24},
	21:  {"Z", "zetta", 21},
	18:  {"E", "exa", 18},
	15:  {"P", "peta", 15},
	12:  {"T", "tera", 12},
	9:   {"G", "giga", 9},
	6:   {"M", "mega", 6},
	3:   {"k", "kilo", 3},
	2:   {"h", "hecto", 2},
	1:   {"da", "deca", 1},
	-1:  {"d", "deci", -1},
	-2:  {"c", "centi", -2},
	-3:  {"m", "milli", -3},
	-6:  {"μ", "micro", -6},
	-9:  {"n", "nano", -9},
	-12: {"p", "pico", -12},
	-15: {"f", "femto", -15},
	-18: {"a", "atto", -18},
	-21: {"z", "zepto", -21},
	-24: {"y", "yocto", -24},
	-27: {"r", "ronto", -27},
	-30: {"q", "quecto", -30},
}

// Magnitude prefix methods create and return a new Unit, while automatically registering
// conversions to and from the provided base Unit
func Quetta(b Unit, o ...UnitOption) Unit { return _mags[30].makeUnit(b, o...) }
func Ronna(b Unit, o ...UnitOption) Unit  { return _mags[27].makeUnit(b, o...) }
func Yotta(b Unit, o ...UnitOption) Unit  { return _mags[24].makeUnit(b, o...) }
func Zetta(b Unit, o ...UnitOption) Unit  { return _mags[21].makeUnit(b, o...) }
func Exa(b Unit, o ...UnitOption) Unit    { return _mags[18].makeUnit(b, o...) }
func Peta(b Unit, o ...UnitOption) Unit   { return _mags[15].makeUnit(b, o...) }
func Tera(b Unit, o ...UnitOption) Unit   { return _mags[12].makeUnit(b, o...) }
func Giga(b Unit, o ...UnitOption) Unit   { return _mags[9].makeUnit(b, o...) }
func Mega(b Unit, o ...UnitOption) Unit   { return _mags[6].makeUnit(b, o...) }
func Kilo(b Unit, o ...UnitOption) Unit   { return _mags[3].makeUnit(b, o...) }
func Hecto(b Unit, o ...UnitOption) Unit  { return _mags[2].makeUnit(b, o...) }
func Deca(b Unit, o ...UnitOption) Unit   { return _mags[1].makeUnit(b, o...) }
func Deci(b Unit, o ...UnitOption) Unit   { return _mags[-1].makeUnit(b, o...) }
func Centi(b Unit, o ...UnitOption) Unit  { return _mags[-2].makeUnit(b, o...) }
func Milli(b Unit, o ...UnitOption) Unit  { return _mags[-3].makeUnit(b, o...) }
func Micro(b Unit, o ...UnitOption) Unit  { return _mags[-6].makeUnit(b, o...) }
func Nano(b Unit, o ...UnitOption) Unit   { return _mags[-9].makeUnit(b, o...) }
func Pico(b Unit, o ...UnitOption) Unit   { return _mags[-12].makeUnit(b, o...) }
func Femto(b Unit, o ...UnitOption) Unit  { return _mags[-15].makeUnit(b, o...) }
func Atto(b Unit, o ...UnitOption) Unit   { return _mags[-18].makeUnit(b, o...) }
func Zepto(b Unit, o ...UnitOption) Unit  { return _mags[-21].makeUnit(b, o...) }
func Yocto(b Unit, o ...UnitOption) Unit  { return _mags[-24].makeUnit(b, o...) }
func Ronto(b Unit, o ...UnitOption) Unit  { return _mags[-27].makeUnit(b, o...) }
func Quecto(b Unit, o ...UnitOption) Unit { return _mags[-30].makeUnit(b, o...) }

// makeUnit creates a magnitude unit and conversion given a base unit.
// This also marks the base unit as such (> set unit.isBaseUnit = true).
func (mag magnitude) makeUnit(base Unit, addOpts ...UnitOption) Unit {
	// make sure the base unit is marked as such
	base.isBaseUnit = true

	name := mag.Prefix + base.Name
	symbol := mag.Symbol + base.Symbol

	// set system to metric by default
	opts := []UnitOption{SI}

	// create prefixed aliases if needed
	for _, a := range base.aliases {
		magAlias := mag.Prefix + a
		opts = append(opts, Aliases(magAlias))
	}

	// create prefixed symbols if needed
	for _, s := range base.symbols {
		magSymbol := mag.Symbol + s
		opts = append(opts, Symbols(magSymbol))
	}

	// append any supplemental options
	opts = append(opts, addOpts...)

	// append quantity name opt
	opts = append(opts, Quantity(base.Quantity))

	u := mustCreateNewUnit(name, symbol, opts...)
	u.base = &base

	// create conversions to and from the base unit
	ratio := 1.0 * math.Pow(10.0, float64(mag.Power))

	NewRatioConversion(u, base, ratio)

	return u
}

// findBestMatchingUnit finds the best matching unit for the given base unit and exponent.
// If no better match is found, the base unit is returned.
func findBestMatchingUnit(baseUnit Unit, exp int) Unit {
	if !baseUnit.IsMetric() || exp == 0 {
		return baseUnit
	}

	// first, see if we have an exact match
	bestUnit := getUnitForExponent(baseUnit.Name, exp)
	if bestUnit != nil {
		return bestUnit
	}

	// no exact match; find the next lower and higher units
	lowerUnit, lowerExp := findNextLowerUnit(baseUnit, exp)
	higherUnit, higherExp := findNextHigherUnit(baseUnit, exp)

	// if we have neither, return the base unit
	if lowerUnit == nil && higherUnit == nil {
		return baseUnit
	}
	// if we have only one, return it only if it's within 3 exponents, else return baseUnit
	if lowerUnit == nil {
		if math.Abs(float64(higherExp-exp)) < 3 {
			return higherUnit
		}
		return baseUnit
	}
	if higherUnit == nil {
		if math.Abs(float64(exp-lowerExp)) < 3 {
			return lowerUnit
		}
		return baseUnit
	}

	// calculate the difference in exponents
	lowerDiff := math.Abs(float64(exp - lowerExp))
	if lowerDiff < 3 {
		// example use case: prefer 200m over 0.2km (assuming no 'deka' and 'hecto' are defined)
		// or perhaps better: prefer 200km over 0.2Mm
		// => if the difference is less than 3, prefer the lower unit
		return lowerUnit
	}

	// pick the closest one; in case of a tie, pick the lower one;
	higherDiff := math.Abs(float64(higherExp - exp))
	if lowerDiff <= higherDiff {
		return lowerUnit
	}
	return higherUnit
}

// getUnitForExponent returns the Unit for the given base name and exponent, or nil if not found
func getUnitForExponent(baseName string, exp int) Unit {
	if mag, ok := _mags[exp]; ok {
		name := mag.Prefix + baseName
		if candidate, ok := _unitMap[name]; ok {
			return candidate
		}
	}
	return nil
}

// findNextLowerUnit finds the next lower unit for the given base unit and exponent.
func findNextLowerUnit(baseUnit Unit, exp int) (Unit, int) {
	for e := exp - 1; e >= MinExponent; e-- {
		if u := getUnitForExponent(baseUnit.Name, e); u != nil {
			return u, e
		}
	}
	return nil, -1
}

// findNextHigherUnit finds the next higher unit for the given base unit and exponent.
func findNextHigherUnit(baseUnit Unit, exp int) (Unit, int) {
	for e := exp + 1; e <= MaxExponent; e++ {
		if u := getUnitForExponent(baseUnit.Name, e); u != nil {
			return u, e
		}
	}
	return nil, -1
}
