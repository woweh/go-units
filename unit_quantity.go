package units

var (
	// _quantityMap is a map of all registered quantities.
	// The key is the quantity name, the value is the UnitQuantity.
	_quantityMap = make(map[string]UnitQuantity)
)

// conversionMap is a map of conversions between units.
// The keys are the "from" unit name, the values are maps of "to" unit name to Conversion.
type conversionMap map[string]map[string]Conversion

// unitQuantity is the internal representation of a unit quantity
type unitQuantity struct {
	name string
	// baseQuantity points to the root quantity if this is an alias.
	// If nil, this is a base quantity.
	baseQuantity UnitQuantity
	// aliases tracks quantities that alias this one (optional, for traversal/cleanup)
	aliases []UnitQuantity
	// conversions stores the conversion map. Only initialized for Base quantities.
	conversions conversionMap
}

// UnitQuantity represents a unit quantity
type UnitQuantity = *unitQuantity

// NewUnitQuantity creates a new UnitQuantity and registers it globally.
// PANICS if a quantity with the same name already exists.
// Use AliasOf to create an alias quantity.
func NewUnitQuantity(name string) UnitQuantity {
	if name == "" {
		panic("quantity name cannot be empty")
	}
	if _, exists := _quantityMap[name]; exists {
		panic("duplicate quantity name: " + name)
	}

	uq := &unitQuantity{
		name:        name,
		conversions: make(conversionMap),
	}
	_quantityMap[name] = uq
	return uq
}

// AliasOf creates a new UnitQuantity as an alias of another quantity.
// Both quantities will share the same conversion map.
// The alias is NOT registered separately in the quantity map (only the base is).
func AliasOf(base UnitQuantity) UnitQuantity {
	baseQuantity := base.Base()
	uq := &unitQuantity{
		name:         base.name,
		baseQuantity: baseQuantity,
		conversions:  nil, // will use base's conversions
	}
	baseQuantity.aliases = append(baseQuantity.aliases, uq)
	return uq
}

// FindQuantity finds a quantity by name (searches both base quantities and aliases).
// Returns nil if not found.
func FindQuantity(name string) UnitQuantity {
	// First try direct lookup in registered base quantities
	if q, ok := _quantityMap[name]; ok {
		return q
	}

	// Search through aliases of all base quantities
	for _, q := range _quantityMap {
		for _, alias := range q.aliases {
			if alias.name == name {
				return alias
			}
		}
	}

	return nil
}

// AllQuantities returns all registered base quantities.
// If includeAliases is true, also includes all alias quantities.
func AllQuantities(includeAliases bool) []UnitQuantity {
	if !includeAliases {
		quantities := make([]UnitQuantity, 0, len(_quantityMap))
		for _, q := range _quantityMap {
			quantities = append(quantities, q)
		}
		return quantities
	}

	// Include aliases
	quantities := make([]UnitQuantity, 0, len(_quantityMap)*2) // rough estimate
	for _, q := range _quantityMap {
		quantities = append(quantities, q)
		quantities = append(quantities, q.aliases...)
	}
	return quantities
}

// Base returns the canonical base quantity.
// If the quantity is already a base, it returns itself.
func (q UnitQuantity) Base() UnitQuantity {
	if q == nil {
		return nil
	}
	if q.baseQuantity != nil {
		return q.baseQuantity.Base()
	}
	return q
}

// NewUnit creates a new Unit within the package, associated with this quantity.
func (q UnitQuantity) NewUnit(name, symbol string, opts ...UnitOption) (Unit, error) {
	// Inject the quantity option
	opts = append(opts, quantity(q))
	return newUnit(name, symbol, opts...)
}

// MustCreateUnit registers a new Unit within the package, associated with this quantity.
// Returns the newly created Unit.
// PANICS if the unit already exists (calls mustCreateNewUnit())!
func (q UnitQuantity) MustCreateUnit(name, symbol string, opts ...UnitOption) Unit {
	// Inject the quantity option
	opts = append(opts, quantity(q))
	return mustCreateUnit(name, symbol, opts...)
}

// Name returns the name of the quantity
func (q UnitQuantity) Name() string {
	if q == nil {
		return ""
	}
	return q.name
}

// String returns the name of the quantity
func (q UnitQuantity) String() string {
	if q == nil {
		return ""
	}
	return q.name
}

// IsCompatible returns true if the two quantities are compatible (same or aliases)
func (q UnitQuantity) IsCompatible(other UnitQuantity) bool {
	if q == nil || other == nil {
		return q == other
	}
	return q.Base() == other.Base()
}
