package units

import (
	"cmp"
	"math"
	"slices"
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

// unitChoice represents a combo of unit and exponent, used by func findBestMatchingUnit
// to find the best matching unit for a given base unit and exponent.
type unitChoice struct {
	unit Unit
	exp  int
}

// findBestMatchingUnit finds the best matching unit for the given base unit and exponent.
// If no better match is found, the base unit is returned.
func findBestMatchingUnit(baseUnit Unit, exp int) Unit {
	if !baseUnit.IsMetric() || exp == 0 {
		return baseUnit
	}
	if exact := getUnitForExponent(baseUnit.Name, exp); exact != nil {
		return exact
	}
	choices := buildUnitChoices(baseUnit.Name, baseUnit)
	if len(choices) == 0 {
		return baseUnit
	}
	if extreme := handleUnitExtremes(choices, exp); extreme != nil {
		return extreme
	}
	lower, higher := findUnitNeighbors(choices, exp)
	return chooseUnit(exp, lower, higher)
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

// buildUnitChoices returns a sorted slice of unitChoice for the given base name.
func buildUnitChoices(baseName string, baseUnit Unit) []unitChoice {
	choices := make([]unitChoice, 0, len(_mags)+1)
	baseUnitAdded := false
	for exp := range _mags {
		u := getUnitForExponent(baseName, exp)
		if u != nil {
			choices = append(choices, unitChoice{u, exp})
			if exp == 0 && u == baseUnit {
				baseUnitAdded = true
			}
		}
	}
	if !baseUnitAdded {
		choices = append(choices, unitChoice{baseUnit, 0})
	}
	slices.SortStableFunc(choices, func(x, y unitChoice) int {
		return cmp.Compare(x.exp, y.exp)
	})
	return choices
}

// handleUnitExtremes returns the unit if exp is outside the choice range, else nil.
func handleUnitExtremes(choices []unitChoice, exp int) Unit {
	if exp < choices[0].exp {
		return choices[0].unit
	}
	if exp > choices[len(choices)-1].exp {
		return choices[len(choices)-1].unit
	}
	return nil
}

// findUnitNeighbors returns the closest lower and higher unitChoices for exp.
func findUnitNeighbors(choices []unitChoice, exp int) (lower, higher *unitChoice) {
	for i := range choices {
		if choices[i].exp < exp {
			lower = &choices[i]
		}
		if choices[i].exp > exp && higher == nil {
			higher = &choices[i]
		}
	}
	return
}

// chooseUnit applies the decision rule to pick the best unitChoice.
func chooseUnit(exp int, lower, higher *unitChoice) Unit {
	switch {
	case lower == nil:
		return higher.unit
	case higher == nil:
		return lower.unit
	default:
		distLower := exp - lower.exp
		distHigher := higher.exp - exp
		if distLower <= distHigher || distLower <= 3 {
			return lower.unit
		}
		return higher.unit
	}
}
