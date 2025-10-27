package units

import (
	"errors"
	"fmt"
	"strings"
)

const (
	// PluralNone is a keyword for a unit that has no plural
	PluralNone = "<|_none_|>"
	// PluralAuto is a keyword for a unit that has an auto-generated plural
	PluralAuto = "<|_auto_|>"
)

// UnitSystem is a system of units
type UnitSystem string

// Systems of units
func (s UnitSystem) String() string { return string(s) }

const (
	// SiSystem provides the internal name for International System of Units
	SiSystem UnitSystem = "metric"
	// BiSystem provides the internal name the British Imperial system of units
	BiSystem UnitSystem = "imperial"
	// UsSystem provides the internal name the United States customary system of units
	UsSystem UnitSystem = "us"
	// IecSystem provides the internal name the International Electrotechnical Commission system of units
	IecSystem UnitSystem = "iec"
	// CgsSystem provides the internal name the centimeter-gram-second system of units
	CgsSystem UnitSystem = "cgs"
	// MKpSSystem provides the internal name the MKpS system of units (from French mètre–kilogramme-poids–seconde)
	MKpSSystem UnitSystem = "MKpS"
)

// UnitQuantity is a quantity label for which a unit belongs
type UnitQuantity string

// Quantity labels for which units may belong
func (q UnitQuantity) String() string { return string(q) }

// Shorthands for pre-defined unit systems:
var (
	// SI is the International System of Units
	SI = System(SiSystem)
	// BI is the British Imperial system of units
	BI = System(BiSystem)
	// US is the United States customary system of units
	US = System(UsSystem)
	// IEC is the International Electrotechnical Commission system of units
	IEC = System(IecSystem)
	// CGS is the centimeter-gram-second system of units
	CGS = System(CgsSystem)
	// MKpS is the MKpS system of units (from French mètre–kilogramme-poids–seconde)
	MKpS = System(MKpSSystem)
)

var (
	// unitMap is a map of all registered units.
	// The key is the unit name or alias, the value is the Unit.
	unitMap = make(map[string]Unit)
	// symbolMap is a map of all registered units.
	// The key is the unit symbol, the value is the Unit.
	symbolMap = make(map[string]Unit)
)

// unit represents a unit of measurement -- internal representation
type unit struct {
	// Name is the (English) name of this unit. The name is mandatory and case-insensitive.
	Name string
	// Symbol is the (main) symbol for this unit, e.g. "m" for meters. The symbol is mandatory and case-sensitive.
	Symbol string
	// Quantity is the quantity label for which this unit belongs, e.g. "length" or "area"
	Quantity UnitQuantity
	// plural is the plural name for this unit, either PluralNone, PluralAuto, or a specific plural name
	plural string
	// aliases are additional names, translations, spellings that this unit may be referred to as,
	// e.g., "vierkante meter" for "square meter".
	// aliases are case-insensitive, like the Name.
	aliases []string
	// symbols are additional symbols that this unit may be referred to as.
	// symbols are case-sensitive, like the (main) Symbol.
	symbols []string
	// system is the system of units this Unit belongs to, if any.
	system UnitSystem
	// base is the base unit for metric units
	base *Unit
	// isBaseUnit is true if this unit is the base unit for metric units
	isBaseUnit bool
}

// Unit represents a unit of measurement
type Unit = *unit

// NewUnit registers a new Unit within the package, returning the newly created Unit.
// Returns an error if the unit already exists.
// The name is mandatory and must be unique.
// The symbol is optional, but if provided, must be unique.
func NewUnit(name, symbol string, opts ...UnitOption) (Unit, error) {
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

	u := &unit{
		Name:   name,
		Symbol: symbol,
		plural: PluralAuto,
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
func newUnit(name, symbol string, opts ...UnitOption) Unit {
	u, err := NewUnit(name, symbol, opts...)
	if err != nil {
		panic(err)
	}
	return u
}

// Names returns all names and aliases this unit may be referred to.
// The main name is always the first in the list.
// Names and aliases are NOT case-sensitive!
func (u Unit) Names() []string {
	names := []string{u.Name}
	if u.plural != PluralNone && u.plural != PluralAuto {
		names = append(names, u.PluralName())
	}
	return append(names, u.aliases...)
}

// Aliases returns all aliases this unit may be referred to.
func (u Unit) Aliases() []string {
	return u.aliases
}

// Symbols returns all symbols this unit may be referred to.
// The main symbol is always the first in the list.
// Symbols are case-sensitive!
func (u Unit) Symbols() []string {
	return append([]string{u.Symbol}, u.symbols...)
}

// String returns the name of this unit
func (u Unit) String() string {
	return u.Name
}

// AddResult is the result of adding aliases or symbols to a unit
type AddResult struct {
	What     string           // "Aliases" or "Symbols", depending on what was added
	Unit     Unit             // the unit to which the aliases or symbols were added
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
func (u Unit) AddAliases(aliases ...string) *AddResult {

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
func (u Unit) AddSymbols(symbols ...string) *AddResult {

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
func (u Unit) System() UnitSystem { return u.system }

// PluralName returns the plural name for this unit
func (u Unit) PluralName() string {
	switch u.plural {
	case PluralNone:
		return u.Name
	case PluralAuto:
		return u.Name + "s"
	default: // custom plural name
		return u.plural
	}
}

// HasName returns true if the provided string matches the provided Unit's Name or Aliases
func (u Unit) HasName(alias string) bool {
	return matchesNameOrAlias(alias, u, false) || matchesNameOrAlias(alias, u, true)
}

// HasSymbol returns true if the provided string matches the provided Unit's Symbol or Symbols
func (u Unit) HasSymbol(symbol string) bool {
	return matchesSymbol(symbol, u)
}

// CsvLine returns a CSV line for this unit
// The line contains the following fields:
// Name, Symbol, PluralName, Quantity, System, Aliases, Symbols
func (u Unit) CsvLine() string {
	line := fmt.Sprintf("%s,%s,%s,%s,%s,", u.Name, u.Symbol, u.PluralName(), u.Quantity, u.System())
	if len(u.aliases) > 0 {
		line += strings.Join(u.aliases, ",")
	}
	if len(u.symbols) > 0 {
		if len(u.aliases) > 0 {
			line += ","
		}
		line += strings.Join(u.symbols, ",")
	}
	return line
}

// ConvertTo converts the provided value from this unit to the provided unit
func (u Unit) ConvertTo(value float64, to Unit) (Value, error) {
	return ConvertFloat(value, u, to)
}

// IsMetric returns true if the UnitSystem is metric (SiSystem)
func (u Unit) IsMetric() bool {
	return u.system == SiSystem
}

// IsImperial returns true if the UnitSystem is British Imperial (BiSystem)
func (u Unit) IsImperial() bool {
	return u.system == BiSystem
}

// IsUS returns true if the UnitSystem is United States customary (UsSystem)
func (u Unit) IsUS() bool {
	return u.system == UsSystem
}

// IsIEC returns true if the UnitSystem is IEC (IecSystem)
func (u Unit) IsIEC() bool {
	return u.system == IecSystem
}

// IsCGS returns true if the UnitSystem is CGS (CgsSystem)
func (u Unit) IsCGS() bool {
	return u.system == CgsSystem
}

// IsMKpS returns true if the UnitSystem is MKpS (MKpSSystem)
func (u Unit) IsMKpS() bool {
	return u.system == MKpSSystem
}

// Base returns the base unit for metric units, or nil for non-metric units.
func (u Unit) Base() Unit {
	return *u.base
}

// IsBase returns true if this unit is the base unit for metric units.
func (u Unit) IsBase() bool {
	return u.isBaseUnit
}

// UnitOption defines an option that may be passed to newUnit
type UnitOption func(Unit) Unit

// Plural sets the plural name for this unit, either PluralNone, PluralAuto, or a custom plural unit name.
//   - PluralNone - labels will use the unmodified unit name in a plural context
//   - PluralAuto - labels for this unit will be created with a plural suffix when appropriate (default)
func Plural(s string) UnitOption {
	return func(u Unit) Unit {
		u.plural = s
		return u
	}
}

// System sets the system of units for which this Unit belongs
func System(s UnitSystem) UnitOption {
	return func(u Unit) Unit {
		u.system = s
		return u
	}
}

// Quantity sets a quantity label for which this Unit belongs
func Quantity(s UnitQuantity) UnitOption {
	return func(u Unit) Unit {
		u.Quantity = s
		return u
	}
}

// Aliases sets the aliases for this Unit
func Aliases(aliases ...string) UnitOption {
	return func(u Unit) Unit {
		u.System()
		u.AddAliases(aliases...)
		return u
	}
}

// Symbols sets the symbols for this Unit
func Symbols(symbols ...string) UnitOption {
	return func(u Unit) Unit {
		u.AddSymbols(symbols...)
		return u
	}
}

// BaseSiUnit marks this unit as the base unit for metric units (SiSystem).
//
// This is defined as a var to make it as easy to use as the shorthands UnitSystem's, like SI.
var BaseSiUnit = func(u Unit) Unit {
	u.isBaseUnit = true
	u.system = SiSystem
	return u
}

// UnitList is a slice of Units. UnitList implements sort.Interface
type UnitList []Unit

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
	// sorting rules:
	// 1. sort by quantity
	// 2. sort by system
	// 3. sort by name
	if a[i].Quantity != a[j].Quantity {
		return a[i].Quantity < a[j].Quantity
	}
	if a[i].system != a[j].system {
		return a[i].system < a[j].system
	}
	return a[i].Name < a[j].Name
}
