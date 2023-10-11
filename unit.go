package units

import "errors"

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

	// unitMap is a map of all registered units
	unitMap = make(map[string]*Unit)
)

// Unit represents a unit of measurement
type Unit struct {
	Name     string
	Symbol   string
	Quantity string
	plural   string // either "none", "auto", or a specific plural name
	aliases  []string
	system   string
}

// NewUnit registers a new Unit within the package, returning the newly created Unit.
func NewUnit(name, symbol string, opts ...UnitOption) (*Unit, error) {
	if _, ok := unitMap[name]; ok {
		return nil, errors.New("duplicate unit name: " + name)
	}

	u := Unit{
		Name:   name,
		Symbol: symbol,
		plural: "auto",
	}

	for _, opt := range opts {
		u = opt(u)
	}

	unitMap[name] = &u

	return &u, nil
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

// Names returns all names and symbols this unit may be referred to
func (u *Unit) Names() []string {
	names := []string{u.Name}
	if u.Symbol != "" {
		names = append(names, u.Symbol)
	}
	if u.plural != "none" && u.plural != "auto" {
		names = append(names, u.PluralName())
	}
	return append(names, u.aliases...)
}

// String returns the name of this unit
func (u *Unit) String() string {
	return u.Name
}

// AddAliases adds names or symbols that this unit may be referred to
func (u *Unit) AddAliases(a ...string) {
	u.aliases = append(u.aliases, a...)
}

// System returns the system of units this Unit belongs to, if any
func (u *Unit) System() string { return u.system }

// PluralName returns the plural name for this unit
func (u *Unit) PluralName() string {
	switch u.plural {
	case "none":
		return u.Name
	case "auto":
		return u.Name + "s"
	default: // custom plural name
		return u.plural
	}
}

// HasAlias returns true if the provided string matches the provided Unit's name, symbol, or aliases
func (u *Unit) HasAlias(alias string) bool {
	return matchUnit(alias, u, false)
}

// UnitOption defines an option that may be passed to newUnit
type UnitOption func(Unit) Unit

// UnitOptionPlural sets the plural name for this unit
// Either "none", "auto", or a custom plural unit name
// "none" - labels will use the unmodified unit name in a plural context
// "auto" - labels for this unit will be created with a plural suffix when appropriate (default)
func UnitOptionPlural(s string) UnitOption {
	return func(u Unit) Unit {
		u.plural = s
		return u
	}
}

// UnitOptionAliases sets additional names, spellings, or symbols that this unit may be referred to as
func UnitOptionAliases(a ...string) UnitOption {
	return func(u Unit) Unit {
		u.aliases = append(u.aliases, a...)
		return u
	}
}

// UnitOptionSystem sets the system of units for which this Unit belongs
func UnitOptionSystem(s string) UnitOption {
	return func(u Unit) Unit {
		u.system = s
		return u
	}
}

// UnitOptionQuantity sets a quantity label for which this Unit belongs
func UnitOptionQuantity(s string) UnitOption {
	return func(u Unit) Unit {
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
