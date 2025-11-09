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

// Quetta creates a new Unit with the quetta prefix applied to the base Unit (10^30)
// This will also create a conversion and inverse conversion to the base Unit.
func Quetta(b Unit, o ...UnitOption) Unit { return _mags[30].makeUnit(b, o...) }

// Ronna creates a new Unit with the ronna prefix applied to the base Unit (10^27)
// This will also create a conversion and inverse conversion to the base Unit.
func Ronna(b Unit, o ...UnitOption) Unit { return _mags[27].makeUnit(b, o...) }

// Yotta creates a new Unit with the yotta prefix applied to the base Unit (10^24)
// This will also create a conversion and inverse conversion to the base Unit.
func Yotta(b Unit, o ...UnitOption) Unit { return _mags[24].makeUnit(b, o...) }

// Zetta creates a new Unit with the zetta prefix applied to the base Unit (10^21)
// This will also create a conversion and inverse conversion to the base Unit.
func Zetta(b Unit, o ...UnitOption) Unit { return _mags[21].makeUnit(b, o...) }

// Exa creates a new Unit with the exa prefix applied to the base Unit (10^18)
// This will also create a conversion and inverse conversion to the base Unit.
func Exa(b Unit, o ...UnitOption) Unit { return _mags[18].makeUnit(b, o...) }

// Peta creates a new Unit with the peta prefix applied to the base Unit (10^15)
// This will also create a conversion and inverse conversion to the base Unit.
func Peta(b Unit, o ...UnitOption) Unit { return _mags[15].makeUnit(b, o...) }

// Tera creates a new Unit with the tera prefix applied to the base Unit (10^12)
// This will also create a conversion and inverse conversion to the base Unit.
func Tera(b Unit, o ...UnitOption) Unit { return _mags[12].makeUnit(b, o...) }

// Giga creates a new Unit with the giga prefix applied to the base Unit (10^9)
// This will also create a conversion and inverse conversion to the base Unit.
func Giga(b Unit, o ...UnitOption) Unit { return _mags[9].makeUnit(b, o...) }

// Mega creates a new Unit with the mega prefix applied to the base Unit (10^6)
// This will also create a conversion and inverse conversion to the base Unit.
func Mega(b Unit, o ...UnitOption) Unit { return _mags[6].makeUnit(b, o...) }

// Kilo creates a new Unit with the kilo prefix applied to the base Unit (10^3)
// This will also create a conversion and inverse conversion to the base Unit.
func Kilo(b Unit, o ...UnitOption) Unit { return _mags[3].makeUnit(b, o...) }

// Hecto creates a new Unit with the hecto prefix applied to the base Unit (10^2)
// This will also create a conversion and inverse conversion to the base Unit.
func Hecto(b Unit, o ...UnitOption) Unit { return _mags[2].makeUnit(b, o...) }

// Deca creates a new Unit with the deca prefix applied to the base Unit (10^1)
// This will also create a conversion and inverse conversion to the base Unit.
func Deca(b Unit, o ...UnitOption) Unit { return _mags[1].makeUnit(b, o...) }

// Deci creates a new Unit with the deci prefix applied to the base Unit (10^-1)
// This will also create a conversion and inverse conversion to the base Unit.
func Deci(b Unit, o ...UnitOption) Unit { return _mags[-1].makeUnit(b, o...) }

// Centi creates a new Unit with the centi prefix applied to the base Unit (10^-2)
// This will also create a conversion and inverse conversion to the base Unit.
func Centi(b Unit, o ...UnitOption) Unit { return _mags[-2].makeUnit(b, o...) }

// Milli creates a new Unit with the milli prefix applied to the base Unit (10^-3)
// This will also create a conversion and inverse conversion to the base Unit.
func Milli(b Unit, o ...UnitOption) Unit { return _mags[-3].makeUnit(b, o...) }

// Micro creates a new Unit with the micro prefix applied to the base Unit (10^-6)
// This will also create a conversion and inverse conversion to the base Unit.
func Micro(b Unit, o ...UnitOption) Unit { return _mags[-6].makeUnit(b, o...) }

// Nano creates a new Unit with the nano prefix applied to the base Unit (10^-9)
// This will also create a conversion and inverse conversion to the base Unit.
func Nano(b Unit, o ...UnitOption) Unit { return _mags[-9].makeUnit(b, o...) }

// Pico creates a new Unit with the pico prefix applied to the base Unit (10^-12)
// This will also create a conversion and inverse conversion to the base Unit.
func Pico(b Unit, o ...UnitOption) Unit { return _mags[-12].makeUnit(b, o...) }

// Femto creates a new Unit with the femto prefix applied to the base Unit (10^-15)
// This will also create a conversion and inverse conversion to the base Unit.
func Femto(b Unit, o ...UnitOption) Unit { return _mags[-15].makeUnit(b, o...) }

// Atto creates a new Unit with the atto prefix applied to the base Unit (10^-18)
// This will also create a conversion and inverse conversion to the base Unit.
func Atto(b Unit, o ...UnitOption) Unit { return _mags[-18].makeUnit(b, o...) }

// Zepto creates a new Unit with the zepto prefix applied to the base Unit (10^-21)
// This will also create a conversion and inverse conversion to the base Unit.
func Zepto(b Unit, o ...UnitOption) Unit { return _mags[-21].makeUnit(b, o...) }

// Yocto creates a new Unit with the yocto prefix applied to the base Unit (10^-24)
// This will also create a conversion and inverse conversion to the base Unit.
func Yocto(b Unit, o ...UnitOption) Unit { return _mags[-24].makeUnit(b, o...) }

// Ronto creates a new Unit with the ronto prefix applied to the base Unit (10^-27)
// This will also create a conversion and inverse conversion to the base Unit.
func Ronto(b Unit, o ...UnitOption) Unit { return _mags[-27].makeUnit(b, o...) }

// Quecto creates a new Unit with the quecto prefix applied to the base Unit (10^-30)
// This will also create a conversion and inverse conversion to the base Unit.
func Quecto(b Unit, o ...UnitOption) Unit { return _mags[-30].makeUnit(b, o...) }

// makeUnit creates a magnitude unit and conversion given a base unit.
//
// NOTE:
// This should be the only place where base units are marked as such (`base.isBaseUnit = true`)!
func (mag magnitude) makeUnit(base Unit, addOpts ...UnitOption) Unit {
	// mark the base unit as such
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
	u.base = base

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
//
// NOTE:
// Also, non-SI units ca have derived units, like MilliMH2O or MilliMHg.
// => Do not check for `!baseUnit.IsMetric()`!
func findBestMatchingUnit(baseUnit Unit, exp int) Unit {
	if !baseUnit.IsBase() || exp == 0 {
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
