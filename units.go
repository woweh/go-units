// Package units is a library for manipulating and converting between various units of measurement
package units

import (
	"errors"
	"sort"
	"strings"
)

// All returns all registered Units, sorted by name and quantity
func All() []*Unit {
	units := make(UnitList, 0, len(unitMap))
	for _, u := range unitMap {
		units = append(units, u)
	}
	sort.Sort(units)
	return units
}

// MustConvertFloat converts a provided float from one Unit to another, PANICKING on error
func MustConvertFloat(x float64, from, to *Unit) Value {
	val, err := ConvertFloat(x, from, to)
	if err != nil {
		panic(err)
	}
	return val
}

// ConvertFloat converts a provided float from one Unit to another
func ConvertFloat(x float64, from, to *Unit) (Value, error) {
	// allow converting to same unit
	if from == to {
		return Value{x, *to}, nil
	}

	// find conversion path
	path, err := ResolveConversion(from, to)
	if err != nil {
		return Value{}, err
	}

	//convert value, applying each conversion function in the path
	for _, c := range path {
		x = c.Fn(x)
	}

	return Value{x, *to}, nil
}

// Find a Unit matching the given name, symbol or alias
func Find(s string) (*Unit, error) {

	// first try case-sensitive match on name
	u, ok := unitMap[s]
	if ok {
		return u, nil
	}

	// the try case-sensitive match on symbol
	// symbols are case-sensitive!
	u, ok = symbolMap[s]
	if ok {
		return u, nil
	}

	// then case-insensitive match on name
	for _, u := range unitMap {
		if matchesNameOrAlias(s, u, false) {
			return u, nil
		}
	}

	// finally, try stripping customPlural suffix
	if strings.HasSuffix(s, "s") || strings.HasSuffix(s, "S") {
		s = s[:len(s)-1]
		for _, u := range unitMap {
			if matchesNameOrAlias(s, u, false) {
				return u, nil
			}
		}
	}

	return nil, errors.New("unit \"" + s + "\" not found")
}

// matchesNameOrAlias returns true if the provided string matches the provided Unit's name or aliases
func matchesNameOrAlias(s string, u *Unit, matchCase bool) bool {
	for _, name := range u.Names() {
		if matchCase {
			if name == s {
				return true
			}
		} else {
			if strings.EqualFold(s, name) {
				return true
			}
		}
	}

	return false
}

// matchesSymbol returns true if the provided string matches the provided Unit's symbol
func matchesSymbol(s string, u *Unit) bool {
	// symbols are case-sensitive!
	for _, sym := range u.Symbols() {
		if sym == s {
			return true
		}
	}
	return false
}
