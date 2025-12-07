package units

import (
	"errors"
	"fmt"
)

// ConversionFn defines a function that converts a float64 value.
type ConversionFn func(float64) float64

// conversion is the internal representation of a conversion between two Units.
type conversion struct {
	FromUnit Unit
	ToUnit   Unit
	Fn       ConversionFn
	Formula  string
}

// Conversion represents a conversion between 2 Units (alias for *conversion).
type Conversion = *conversion

// Name returns a short label like "'meter -> foot'".
func (c Conversion) Name() string {
	if c == nil {
		return "<nil>"
	}
	from := "<nil>"
	to := "<nil>"
	if c.FromUnit != nil && c.FromUnit.Name != "" {
		from = c.FromUnit.Name
	}
	if c.ToUnit != nil && c.ToUnit.Name != "" {
		to = c.ToUnit.Name
	}
	return fmt.Sprintf(`'%s -> %s'`, from, to)
}

// String returns a string representation of the conversion.
func (c Conversion) String() string {
	if c == nil {
		return "<nil conversion>"
	}
	return fmt.Sprintf("%s (%s)", c.Name(), safeFormula(c.Formula))
}

// safeFormula returns a safe string representation of the formula.
func safeFormula(s string) string {
	if s == "" {
		return "<no formula>"
	}
	return s
}

// isValid checks if the conversion is valid.
func (c Conversion) isValid() error {
	if c == nil {
		return errors.New("conversion is nil")
	}
	if c.FromUnit == nil {
		return fmt.Errorf("conversion FromUnit is nil for %s", c.String())
	}
	if c.ToUnit == nil {
		return fmt.Errorf("conversion ToUnit is nil for %s", c.String())
	}
	if c.Fn == nil {
		return fmt.Errorf("conversion function is nil for %s", c.String())
	}
	if c.Formula == "" {
		return fmt.Errorf("conversion formula is empty for %s", c.String())
	}
	if c.FromUnit.Name == "" || c.ToUnit.Name == "" {
		return fmt.Errorf("conversion has unit with empty name: %s", c.String())
	}
	return nil
}

// Conversions is a slice of Conversion pointers.
type Conversions []Conversion

// NewRatioConversion registers a conversion formula and the **inverse**, given the ratio between `from` and `to` Unit.
// PANICS if the units have incompatible quantities.
func NewRatioConversion(from, to Unit, ratio float64) {
	if !from.Quantity.IsCompatible(to.Quantity) {
		panic(fmt.Sprintf("cannot create conversion between incompatible quantities: %s and %s",
			from.Quantity.Name(), to.Quantity.Name()))
	}
	// use standard formatting for ratio in formula string: %g => %e for large exponents, %f otherwise.
	ratioStr := fmt.Sprintf("%g", ratio)
	NewConversionFromFn(from, to, func(x float64) float64 { return x * ratio }, "x * "+ratioStr)
	NewConversionFromFn(to, from, func(x float64) float64 { return x / ratio }, "x / "+ratioStr)
}

// NewConversionFromFn registers a new conversion formula from one Unit to another.
// PANICS if the units have incompatible quantities.
func NewConversionFromFn(from, to Unit, f ConversionFn, formula string) {
	if !from.Quantity.IsCompatible(to.Quantity) {
		panic(fmt.Sprintf("cannot create conversion between incompatible quantities: %s and %s",
			from.Quantity.Name(), to.Quantity.Name()))
	}

	c := &conversion{from, to, f, formula}
	base := from.Quantity.Base()

	// ensure initialization (just in case, though NewUnitQuantity handles it)
	if base.conversions == nil {
		base.conversions = make(conversionMap)
	}

	if base.conversions[from.Name] == nil {
		base.conversions[from.Name] = make(map[string]Conversion)
	}
	base.conversions[from.Name][to.Name] = c
}

// ResolveConversion resolves a path of one or more Conversions between two units.
func ResolveConversion(from, to Unit) (Conversions, error) {
	if !from.Quantity.IsCompatible(to.Quantity) {
		return nil, fmt.Errorf("incompatible quantities: %s and %s", from.Quantity.Name(), to.Quantity.Name())
	}

	path := findPath(from, to)
	if path == nil {
		return nil, errors.New("no conversion path found from " + from.Name + " to " + to.Name)
	}
	return path, nil
}

// findPath finds a path of Conversions from start to end, using a Breadth-First Search (BFS) algorithm.
func findPath(start, end Unit) Conversions {
	if start == end {
		return nil
	}

	// Use the conversion map from the starting unit's quantity base
	conversions := start.Quantity.Base().conversions

	visited := make(map[string]bool)
	queue := []Conversions{{}} // start with empty path
	for len(queue) > 0 {
		path := queue[0]
		queue = queue[1:]
		current := start
		if len(path) > 0 {
			current = path[len(path)-1].ToUnit
		}
		if visited[current.Name] {
			continue
		}
		visited[current.Name] = true
		if current == end {
			return path
		}

		// Get outgoing edges using the nested map
		outgoing := conversions[current.Name]
		for _, conv := range outgoing {
			if !visited[conv.ToUnit.Name] {
				newPath := make(Conversions, len(path)+1)
				copy(newPath, path)
				newPath[len(path)] = conv
				queue = append(queue, newPath)
			}
		}
	}
	return nil
}

// CheckConversionIntegrity checks if every conversion in internal conversion maps has a corresponding reverse conversion.
func CheckConversionIntegrity() error {
	// We need to iterate over all registered units to find all Quantities,
	// then check each Quantity's conversion map.

	checkedQuantities := make(map[UnitQuantity]bool)

	for _, u := range All() {
		q := u.Quantity.Base()
		if checkedQuantities[q] {
			continue
		}
		checkedQuantities[q] = true

		for fromUnitName, outgoing := range q.conversions {
			for _, conv := range outgoing {
				if conv == nil {
					return fmt.Errorf("conversion map integrity error: nil conversion in conversions for unit %s", fromUnitName)
				}
				e := conv.isValid()
				if e != nil {
					return fmt.Errorf("conversion map integrity error: invalid conversion for unit %s: %w", fromUnitName, e)
				}
				if fromUnitName != conv.FromUnit.Name {
					return fmt.Errorf("conversion map integrity error: expected from unit %s, got %s", fromUnitName, conv.FromUnit.Name)
				}

				// Check reverse
				reverseMap := q.conversions[conv.ToUnit.Name]
				if reverseMap == nil {
					return fmt.Errorf("conversion map integrity error: no outgoing conversions from %v (reverse of %v -> %v)", conv.ToUnit, conv.FromUnit, conv.ToUnit)
				}
				if _, ok := reverseMap[fromUnitName]; !ok {
					return fmt.Errorf("conversion map integrity error: no reverse conversion found for conversion from %v to %v", conv.FromUnit, conv.ToUnit)
				}
			}
		}
	}
	return nil
}

// HasTo returns true if there is a conversion to the specified unit in the Conversions slice.
func (cc Conversions) HasTo(to Unit) bool {
	for _, c := range cc {
		if c.ToUnit == to {
			return true
		}
	}
	return false
}
