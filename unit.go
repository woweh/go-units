package units

import (
	"errors"
	"fmt"
	"strings"
)

// PluralType is the type of customPlural name for a unit
type PluralType int

const (
	None PluralType = iota
	Auto
	Custom
)

// UnitSystem is a system of units
type UnitSystem string

// UnitQuantity is a quantity label for which a unit belongs
type UnitQuantity string

var (
	// Shorthands for pre-defined unit systems:
	//----------------------------------------------------------------------------------------------

	// BI is the British Imperial system of units
	BI = UnitOptionSystem("imperial")
	// SI is the International System of Units
	SI = UnitOptionSystem("metric")
	// US is the United States customary system of units
	US = UnitOptionSystem("us")
	// IEC is the International Electrotechnical Commission system of units
	IEC = UnitOptionSystem("iec")
	//----------------------------------------------------------------------------------------------

	// unitMap is a map of all registered units.
	// The key is the unit name or alias, the value is the Unit.
	unitMap = make(map[string]*Unit)
	// symbolMap is a map of all registered units.
	// The key is the unit symbol, the value is the Unit.
	symbolMap = make(map[string]*Unit)
)

// Unit represents a unit of measurement
type Unit struct {
	// Name is the (english) name of this unit. The name is mandatory and case-insensitive.
	Name string
	// Symbol is the (main) symbol for this unit, e.g. "m" for meters. The symbol is mandatory and case-sensitive.
	Symbol string
	// Quantity is the quantity label for which this unit belongs, e.g. "length" or "area"
	Quantity UnitQuantity
	// pluralType is the type of customPlural name for this unit, either "none", "auto", or "custom"
	pluralType PluralType
	// customPlural is the customPlural name for this unit, either "none", "auto", or a specific customPlural name
	customPlural string
	// aliases are additional names, translations, spellings that this unit may be referred to as,
	// e.g., "vierkante meter" for "square meter".
	// aliases are case-insensitive, like the Name.
	aliases []string
	// symbols are additional symbols that this unit may be referred to as.
	// symbols are case-sensitive, like the (main) Symbol.
	symbols []string
	// system is the system of units this Unit belongs to, if any.
	system UnitSystem
}

// NewUnit registers a new Unit within the package, returning the newly created Unit.
// Returns an error if the unit already exists.
// The name is mandatory and must be unique.
// The symbol is optional, but if provided, must be unique.
func NewUnit(name, symbol string, opts ...UnitOption) (*Unit, error) {
	if name == "" {
		return nil, errors.New("unit name cannot be empty")
	}
	if _, ok := unitMap[name]; ok {
		return nil, errors.New("duplicate unit name: " + name)
	}
	if symbol != "" {
		if _, ok := symbolMap[symbol]; ok {
			return nil, errors.New("duplicate unit symbol: " + symbol)
		}
	}

	u := &Unit{
		Name:         name,
		Symbol:       symbol,
		pluralType:   Auto,
		customPlural: "",
	}

	for _, opt := range opts {
		u = opt(u)
	}

	// register unit:
	unitMap[name] = u
	if symbol != "" {
		symbolMap[symbol] = u
	}

	return u, nil
}

// newUnit registers a new Unit within the package, returning the newly created Unit.
// PANICS if the unit already exists!
func newUnit(name, symbol string, opts ...UnitOption) *Unit {
	u, err := NewUnit(name, symbol, opts...)
	if err != nil {
		panic(err)
	}
	return u
}

// Names returns all names and aliases this unit may be referred to.
// The main name is always the first in the list.
// Names and aliases are NOT case-sensitive!
func (u *Unit) Names() []string {
	names := []string{u.Name}
	if u.pluralType == Custom {
		names = append(names, u.PluralName())
	}
	return append(names, u.aliases...)
}

// Aliases returns all aliases this unit may be referred to.
func (u *Unit) Aliases() []string {
	return u.aliases
}

// Symbols returns all symbols this unit may be referred to.
// The main symbol is always the first in the list.
// Symbols are case-sensitive!
func (u *Unit) Symbols() []string {
	return append([]string{u.Symbol}, u.symbols...)
}

// String returns the name of this unit
func (u *Unit) String() string {
	return u.Name
}

// AddResult is the result of adding aliases or symbols to a unit
type AddResult struct {
	What     string           // "Aliases" or "Symbols", depending on what was added
	Unit     *Unit            // the unit to which the aliases or symbols were added
	Added    []string         // the aliases or symbols that were added
	Failures map[string]error // the aliases or symbols that failed to be added, and the reason why
	Err      error            // the overall error that occurred, if any
	Summary  string           // a summary of the result
}

// validate the AddResult, setting the Error and Summary fields
func (ar *AddResult) validate() {
	addCount := len(ar.Added)
	failCount := len(ar.Failures)
	inCount := addCount + failCount
	ar.Summary = `Add %s for unit '%s':
- processed: %d
- added: %d
- failed: %d`
	ar.Summary = fmt.Sprintf(ar.Summary, ar.What, ar.Unit.Name, inCount, addCount, failCount)
	if failCount == inCount {
		ar.Err = fmt.Errorf("error, failed to add any %s for unit '%s'", strings.ToLower(ar.What), ar.Unit.Name)
	}
}

// String returns a summary of the AddResult
func (ar *AddResult) String() string {
	return ar.Summary
}

// AddAliases adds aliases that this unit may be referred to
func (u *Unit) AddAliases(aliases ...string) *AddResult {

	result := &AddResult{
		What:     "Aliases",
		Unit:     u,
		Added:    []string{},
		Failures: map[string]error{},
		Err:      nil,
		Summary:  "",
	}

	for _, a := range aliases {
		if _, ok := unitMap[a]; ok {
			result.Failures[a] = errors.New("alias " + a + " is already registered as a unit name or alias")
			// skip aliases that are already registered as a unit name
			continue
		}
		if _, ok := symbolMap[a]; ok {
			result.Failures[a] = errors.New("alias " + a + " is already registered as a unit symbol")
			// skip aliases that are already registered as a unit symbol
			continue
		}
		// else add alias
		u.aliases = append(u.aliases, a)
		result.Added = append(result.Added, a)
		unitMap[a] = u
	}

	result.validate()

	return result
}

// AddSymbols adds symbols that this unit may be referred to
func (u *Unit) AddSymbols(symbols ...string) *AddResult {

	result := &AddResult{
		What:     "Symbols",
		Unit:     u,
		Added:    []string{},
		Failures: map[string]error{},
		Err:      nil,
		Summary:  "",
	}

	for _, s := range symbols {
		if _, ok := unitMap[s]; ok {
			result.Failures[s] = errors.New("symbol " + s + " is already registered as a unit name or alias")
			// skip aliases that are already registered as a unit name
			continue
		}
		if _, ok := symbolMap[s]; ok {
			result.Failures[s] = errors.New("symbol " + s + " is already registered as a unit symbol")
			// skip aliases that are already registered as a unit symbol
			continue
		}
		// else add symbol
		u.symbols = append(u.symbols, s)
		result.Added = append(result.Added, s)
		symbolMap[s] = u
	}

	result.validate()

	return result
}

// System returns the system of units this Unit belongs to, if any
func (u *Unit) System() UnitSystem { return u.system }

// PluralName returns the customPlural name for this unit
func (u *Unit) PluralName() string {
	switch u.pluralType {
	case None:
		return u.Name
	case Auto:
		return u.Name + "s"
	default: // custom customPlural name
		return u.customPlural
	}
}

// HasName returns true if the provided string matches the provided Unit's Name or Aliases
func (u *Unit) HasName(alias string) bool {
	return matchesNameOrAlias(alias, u, false) || matchesNameOrAlias(alias, u, true)
}

// HasSymbol returns true if the provided string matches the provided Unit's Symbol or Symbols
func (u *Unit) HasSymbol(symbol string) bool {
	return matchesSymbol(symbol, u)
}

// UnitOption defines an option that may be passed to newUnit
type UnitOption func(*Unit) *Unit

// UnitOptionPlural sets the PluralType and optional customPlural name for this unit
func UnitOptionPlural(pt PluralType, s string) UnitOption {
	return func(u *Unit) *Unit {
		u.pluralType = pt
		if pt == Custom {
			u.customPlural = s
		} else {
			u.customPlural = ""
		}
		return u
	}
}

// UnitOptionSystem sets the system of units for which this Unit belongs
func UnitOptionSystem(s UnitSystem) UnitOption {
	return func(u *Unit) *Unit {
		u.system = s
		return u
	}
}

// UnitOptionQuantity sets a quantity label for which this Unit belongs
func UnitOptionQuantity(s UnitQuantity) UnitOption {
	return func(u *Unit) *Unit {
		u.Quantity = s
		return u
	}
}

// UnitList is a slice of Units. UnitList implements sort.Interface
type UnitList []*Unit

// Len returns the length of the UnitList
func (a UnitList) Len() int {
	return len(a)
}

// Swap swaps the Units at the given indices
func (a UnitList) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

// Less returns whether the Unit at index i is less than the Unit at index j
func (a UnitList) Less(i, j int) bool {
	if a[i].Quantity != a[j].Quantity {
		return a[i].Quantity < a[j].Quantity
	}
	return a[i].Name < a[j].Name
}
