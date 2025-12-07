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

var (
	// _unitMap is a map of all registered units.
	// The key is the unit name or alias, the value is the Unit.
	_unitMap = make(map[string]Unit)
	// _symbolMap is a map of all registered units.
	// The key is the unit symbol, the value is the Unit.
	_symbolMap = make(map[string]Unit)
)

// FormatFn is a function that formats a float64 value into a string representation.
type FormatFn func(float64) string

// unit is the internal representation of a unit of measurement.
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
	base *unit
	// isBaseUnit is true if this unit is the base unit for metric units
	// NOTE:
	// isBaseUnit must only be set from method `makeUnit` in 'metric.go'.
	isBaseUnit bool
	// fmtFn is a function that formats a float64 value into a string representation.
	fmtFn FormatFn
}

// Unit represents a unit of measurement.
type Unit = *unit

// newUnit registers a new Unit within the package, returning the newly created Unit.
// Returns an error if the unit already exists.
// The name is mandatory and must be unique.
// The symbol is optional, but if provided, must be unique.
func newUnit(name, symbol string, opts ...UnitOption) (Unit, error) {
	if name == "" {
		return nil, errors.New("unit name cannot be empty")
	}
	if _, ok := _unitMap[name]; ok {
		return nil, errors.New("duplicate unit name: " + name)
	}
	if symbol != "" {
		if _, ok := _symbolMap[symbol]; ok {
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
	_unitMap[name] = u
	if symbol != "" {
		_symbolMap[symbol] = u
	}

	return u, nil
}

// mustCreateUnit registers a new Unit within the package, returning the newly created Unit.
// PANICS if the unit already exists!
func mustCreateUnit(name, symbol string, opts ...UnitOption) Unit {
	u, err := newUnit(name, symbol, opts...)
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

// FactorTo returns the multiplicative conversion factor from this unit to the target unit.
// For example, Meter.FactorTo(Foot) returns 3.28084 (1 meter = 3.28084 feet).
// Returns an error if the units are incompatible or no conversion path exists.
func (u Unit) FactorTo(to Unit) (float64, error) {
	result, err := ConvertFloat(1.0, u, to)
	if err != nil {
		return 0, err
	}
	return result.Float(), nil
}

// to is an internal helper that returns the conversion factor, panicking on error.
// This is useful during initialization when conversion paths are known to exist.
// For example: Meter.to(Foot) returns 3.28084.
func (u Unit) to(target Unit) float64 {
	factor, err := u.FactorTo(target)
	if err != nil {
		panic(fmt.Sprintf("failed to get conversion factor from %s to %s: %v", u.Name, target.Name, err))
	}
	return factor
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

// HasBase returns true if this unit has a base unit.
func (u Unit) HasBase() bool {
	return u.base != nil
}

// Base returns the base unit for metric units, or nil for non-metric units.
func (u Unit) Base() Unit {
	return u.base
}

// IsBase returns true if this unit is the base unit for derived (usually metric) units.
func (u Unit) IsBase() bool {
	return u.isBaseUnit
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
		ar.Err = fmt.Errorf("failed to add any %s for unit '%s'", strings.ToLower(ar.What), ar.Unit.Name)
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
		if _, ok := _unitMap[a]; ok {
			result.Failures[a] = errors.New("alias " + a + " is already registered as a unit name or alias")
			// skip aliases that are already registered as a unit name
			continue
		}
		if _, ok := _symbolMap[a]; ok {
			result.Failures[a] = errors.New("alias " + a + " is already registered as a unit symbol")
			// skip aliases that are already registered as a unit symbol
			continue
		}
		// else add alias
		u.aliases = append(u.aliases, a)
		result.Added = append(result.Added, a)
		_unitMap[a] = u
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
		if _, ok := _unitMap[s]; ok {
			result.Failures[s] = errors.New("symbol " + s + " is already registered as a unit name or alias")
			// skip aliases that are already registered as a unit name
			continue
		}
		if _, ok := _symbolMap[s]; ok {
			result.Failures[s] = errors.New("symbol " + s + " is already registered as a unit symbol")
			// skip aliases that are already registered as a unit symbol
			continue
		}
		// else add symbol
		u.symbols = append(u.symbols, s)
		result.Added = append(result.Added, s)
		_symbolMap[s] = u
	}

	result.validate()

	return result
}

// AddFormatFn adds a formatting function to the unit, which can be used to customize how values are displayed.
func (u Unit) AddFormatFn(fn FormatFn) Unit {
	u.fmtFn = fn
	return u
}

// IsCompatible returns true if the other unit is compatible with this unit.
// Two units are compatible if their quantities are compatible.
func (u Unit) IsCompatible(other unit) bool {
	return UnitsAreCompatible(u, &other)
}

// UnitsAreCompatible returns true if the two units are compatible.
// Two units are compatible if their quantities are compatible.
func UnitsAreCompatible(u1, u2 Unit) bool {
	// nil units are not compatible
	if u1 == nil || u2 == nil {
		return false
	}
	return u1.Quantity.IsCompatible(u2.Quantity)
}

// UnitOption defines an option that may be passed to mustCreateNewUnit
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

// quantity sets a quantity label for which this Unit belongs
func quantity(uq UnitQuantity) UnitOption {
	return func(u Unit) Unit {
		u.Quantity = uq
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
