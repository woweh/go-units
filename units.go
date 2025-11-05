// Package units is a library for looking up units,
// and for manipulating and converting between various units of measurement.
package units

import (
	"fmt"
	"sort"
	"strings"
)

// CsvHeader is the header row for the CSV output (> see func GetCsv).
// It matches the format of the Unit.CsvLine method.
const CsvHeader = "Name,Symbol,PluralName,Quantity,System,Aliases & Symbols"

// UnitList is a slice of Units. UnitList implements sort.Interface
type UnitList []Unit

// Len returns the length of the UnitList
func (ul UnitList) Len() int {
	return len(ul)
}

// Swap swaps the Units at the given indices
func (ul UnitList) Swap(i, j int) {
	ul[i], ul[j] = ul[j], ul[i]
}

// Less returns whether the Unit at index i is less than the Unit at index j
func (ul UnitList) Less(i, j int) bool {
	// sorting rules:
	// 1. sort by quantity
	// 2. sort by system
	// 3. sort by name
	if ul[i].Quantity != ul[j].Quantity {
		return ul[i].Quantity < ul[j].Quantity
	}
	if ul[i].system != ul[j].system {
		return ul[i].system < ul[j].system
	}
	return ul[i].Name < ul[j].Name
}

// All returns all registered Units, sorted by name and quantity
func All() UnitList {
	units := make(UnitList, 0, len(_unitMap))
	for _, u := range _unitMap {
		units = append(units, u)
	}
	sort.Sort(units)
	return units
}

// MustConvertFloat converts a provided float from one Unit to another, PANICKING on error
func MustConvertFloat(x float64, from, to Unit) Value {
	val, err := ConvertFloat(x, from, to)
	if err != nil {
		panic(err)
	}
	return val
}

// ConvertFloat converts a provided float from one Unit to another
func ConvertFloat(x float64, from, to Unit) (Value, error) {
	if from == nil {
		return Value{}, fmt.Errorf("no unit specified to convert from")
	}
	if to == nil {
		return Value{}, fmt.Errorf("no unit specified to convert to")
	}
	if from.Quantity != to.Quantity {
		return Value{}, fmt.Errorf("cannot convert from unit of quantity %s to %s", from.Quantity, to.Quantity)
	}
	// allow converting to same unit
	if from == to {
		return Value{x, to}, nil
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

	return Value{x, to}, nil
}

// Find a Unit matching the given name, symbol or alias
func Find(s string) (Unit, error) {

	// first try case-sensitive match on name
	u, ok := _unitMap[s]
	if ok {
		return u, nil
	}

	// then try case-sensitive match on symbol
	// symbols are case-sensitive!
	u, ok = _symbolMap[s]
	if ok {
		return u, nil
	}

	// then case-insensitive match on name
	for _, u := range _unitMap {
		if matchesNameOrAlias(s, u, false) {
			return u, nil
		}
	}

	// finally, try stripping plural suffix
	if strings.HasSuffix(s, "s") || strings.HasSuffix(s, "S") {
		s = s[:len(s)-1]
		for _, u := range _unitMap {
			if matchesNameOrAlias(s, u, false) {
				return u, nil
			}
		}
	}

	return nil, fmt.Errorf("unit '%s' not found", s)
}

// matchesNameOrAlias returns true if the provided string matches the provided Unit's name or aliases
func matchesNameOrAlias(s string, u Unit, matchCase bool) bool {
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
func matchesSymbol(s string, u Unit) bool {
	// symbols are case-sensitive!
	for _, sym := range u.Symbols() {
		if sym == s {
			return true
		}
	}
	return false
}

// GetCsv returns a CSV representation of all registered Units.
func GetCsv() []string {

	// _unitMap contains 'duplicate' units, because they are registered multiple times with different names/aliases
	uniqueUnits := make(map[string]Unit)
	for _, u := range _unitMap {
		uniqueUnits[u.Name] = u
	}

	csvLines := make([]string, 0, len(uniqueUnits)+1)
	csvLines = append(csvLines, CsvHeader)
	for _, u := range uniqueUnits {
		csvLines = append(csvLines, u.CsvLine())
	}

	return csvLines
}
