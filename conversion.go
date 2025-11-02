package units

import (
	"errors"
	"fmt"
	"regexp"
)

var (
	// _conversionMap holds all registered conversions between units.
	//
	// This is an adjacency list for a directed, labeled graph:
	//
	//- Keys are graph nodes (Unit).
	//- Values are outgoing edges ([]Conversion) from that node.
	//- Each Conversion is a directed, labeled edge from `FromUnit` to `ToUnit` (carries a function and formula).
	//- Used as the graph structure for BFS in `findPath` to find a path from `start` to `end`.
	_conversionMap = make(map[Unit]Conversions)

	// _fmtFormulaRe matches float numbers in formula strings
	_fmtFormulaRe = regexp.MustCompile(`-?\d+(\.\d+)?([eE][-+]?\d+)?`)
	// _sciNotationRe matches scientific notation
	_sciNotationRe = regexp.MustCompile(`[eE]`)
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
func NewRatioConversion(from, to Unit, ratio float64) {
	// use standard formatting for ratio in formula string: %g => %e for large exponents, %f otherwise.
	ratioStr := fmt.Sprintf("%g", ratio)
	NewConversionFromFn(from, to, func(x float64) float64 { return x * ratio }, "x * "+ratioStr)
	NewConversionFromFn(to, from, func(x float64) float64 { return x / ratio }, "x / "+ratioStr)
}

// NewConversionFromFn registers a new conversion formula from one Unit to another.
//
// NOTE:
//   - When using `NewConversionFromFn` directly, you must define _conversions in both directions!
//
// Example:
//
//	NewConversionFromFn(SlopeValue, SlopeDegree, slopeValueToDegree, "math.Atan(x) * 180 / math.Pi")
//	NewConversionFromFn(SlopeDegree, SlopeValue, slopeDegreeToValue, "math.Tan(x * math.Pi / 180)")
func NewConversionFromFn(from, to Unit, f ConversionFn, formula string) {
	c := &conversion{from, to, f, formula}
	_conversionMap[from] = append(_conversionMap[from], c)
}

// ResolveConversion resolves a path of one or more Conversions between two units.
func ResolveConversion(from, to Unit) (Conversions, error) {
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
	visited := make(map[Unit]bool)
	queue := []Conversions{{}} // start with empty path
	for len(queue) > 0 {
		path := queue[0]
		queue = queue[1:]
		current := start
		if len(path) > 0 {
			current = path[len(path)-1].ToUnit
		}
		if visited[current] {
			continue
		}
		visited[current] = true
		if current == end {
			return path
		}
		for _, conv := range _conversionMap[current] {
			if !visited[conv.ToUnit] {
				newPath := make(Conversions, len(path)+1)
				copy(newPath, path)
				newPath[len(path)] = conv
				queue = append(queue, newPath)
			}
		}
	}
	return nil
}

// CheckConversionIntegrity checks if every conversion in internal _conversionMap has a corresponding reverse conversion.
func CheckConversionIntegrity() error {
	// loop through the entire _conversionMap
	for fromUnit, conversions := range _conversionMap {
		// loop through the conversions of fromUnit
		for _, conv := range conversions {
			if conv == nil {
				// can this ever happen??? I don't think so...
				return fmt.Errorf("conversion map integrity error: nil conversion in conversions for unit %s", fromUnit.Name)
			}

			e := conv.isValid()
			if e != nil {
				return fmt.Errorf("conversion map integrity error: invalid conversion for unit %s: %w", fromUnit.Name, e)
			}

			if fromUnit != conv.FromUnit {
				return fmt.Errorf("conversion map integrity error: expected from unit %s, got %s", fromUnit.Name, conv.FromUnit.Name)
			}

			cc := _conversionMap[conv.ToUnit]
			if cc == nil {
				return fmt.Errorf("conversion map integrity error: no reverse conversion found for conversion from %v to %v", conv.FromUnit, conv.ToUnit)
			}
			if !cc.HasTo(fromUnit) {
				return fmt.Errorf("conversion map integrity error: no reverse conversion found for conversion from %v to %v", conv.FromUnit, conv.ToUnit)
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
