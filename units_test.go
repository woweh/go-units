package units

import (
	"fmt"
	"strings"
	"testing"
)

// collectUnitNames collects All Unit names and aliases
func collectUnitNames() (a []string) {
	for _, u := range All() {
		a = append(a, u.Names()...)
	}
	return a
}

func Test_UnitLookup_ByName(t *testing.T) {
	for _, name := range collectUnitNames() {
		u, err := Find(name)
		if err != nil {
			t.Errorf("Find(%q) error: %v", name, err)
		}
		t.Logf("found unit by name: %s => %s", name, u.Name)
	}
}

func Test_UnitLookup_BySymbol(t *testing.T) {
	for _, u := range All() {
		for _, symbol := range u.Symbols() {
			if symbol == "" {
				continue
			}
			u2, err := Find(symbol)
			if err != nil {
				t.Errorf("Find(%q) error: %v", symbol, err)
			}
			if u.Name != u2.Name {
				t.Errorf("Find(%q) got unit %q, want %q", symbol, u2.Name, u.Name)
			}
		}
	}
}

func Test_UnitNameOverlap(t *testing.T) {
	nameMap := make(map[string]Unit)

	var total, failed int
	for _, u := range All() {
		for _, name := range u.Names() {
			if existing, ok := nameMap[name]; ok {
				t.Errorf("overlap in unit names: %s, %s (%s)", u.Name, existing.Name, name)
			} else {
				nameMap[name] = u
			}
			total++
		}
	}
	t.Logf("tested %d unit names, %d overlaps", total, failed)
}

// groupUnitsByQuantity groups all units by their UnitQuantity
func groupUnitsByQuantity() map[UnitQuantity][]Unit {
	m := make(map[UnitQuantity][]Unit)

	for _, u := range All() {
		if _, ok := m[u.Quantity]; !ok {
			m[u.Quantity] = []Unit{}
		}
		m[u.Quantity] = append(m[u.Quantity], u)
	}

	return m
}

// ensure all units within the same quantity resolve a conversion path
func Test_PathResolve(t *testing.T) {
	for qname, qunits := range groupUnitsByQuantity() {
		if qname == quantityForUnitTests {
			// skip temporary units created by unit tests
			continue
		}
		t.Logf("testing conversion paths for quantity: %s", qname)
		for _, u1 := range qunits {
			v1 := NewValue(1.0, u1)
			for _, u2 := range qunits {
				if u1.Name == u2.Name {
					continue
				}
				_, err := v1.Convert(u2)
				if err != nil {
					t.Errorf("failed to resolve path: %s -> %s: %v", u1.Name, u2.Name, err)
				}
			}
		}
	}
}

func Test_BaseUnits(t *testing.T) {
	for qname, qunits := range groupUnitsByQuantity() {
		if qname == quantityForUnitTests {
			continue
		}
		testName := fmt.Sprintf("Test BaseUnits of Quantity '%s'", qname)
		t.Run(testName, func(t *testing.T) {
			for _, u := range qunits {
				// Test HasBase() consistency
				hasBase := u.HasBase()
				base := u.Base()

				if hasBase && base == nil {
					t.Errorf("unit %s HasBase() is true but Base() returns nil", u.Name)
				}
				if !hasBase && base != nil {
					t.Errorf("unit %s HasBase() is false but Base() returns non-nil", u.Name)
				}

				if base == nil {
					continue
				}

				// Verify base unit is in same quantity
				if u.Quantity != base.Quantity {
					t.Errorf("unit %s has base %s with different quantity", u.Name, base.Name)
				}

				// Verify base unit can be looked up
				_, err := Find(base.Name)
				if err != nil {
					t.Errorf("base unit %s not found", base.Name)
				}

				// Verify base unit is marked as base
				if !base.IsBase() {
					t.Errorf("base unit %s for %s is not marked as base unit", base.Name, u.Name)
				}

				// Verify derived unit is not marked as base
				if u.IsBase() {
					t.Errorf("derived unit %s should not be marked as base unit", u.Name)
				}

				// Verify base unit's base is nil (base units don't have bases)
				if base.Base() != nil {
					t.Errorf("base unit %s should not have a base unit", base.Name)
				}
			}
		})
	}
}

func Test_GetCsv(t *testing.T) {
	csv := GetCsv()
	if len(csv) == 0 {
		t.Errorf("GetCsv() returned empty slice")
	}
	if !strings.HasPrefix(csv[0], "Name,Symbol,PluralName,Quantity,System,Aliases & Symbols") {
		t.Errorf("GetCsv() header incorrect: %v", csv[0])
	}
	found := false
	for _, line := range csv[1:] {
		if strings.Contains(line, ",") {
			found = true
			break
		}
	}
	if !found {
		t.Errorf("GetCsv() did not return any data lines with commas")
	}
}

func Test_UnitList_Sorting(t *testing.T) {
	units := UnitList{
		&unit{Name: "B", Quantity: NewUnitQuantity("Q2"), system: "sys2"},
		&unit{Name: "A", Quantity: NewUnitQuantity("Q1"), system: "sys1"},
		&unit{Name: "C", Quantity: NewUnitQuantity("Q3"), system: "sys2"},
	}

	if units.Len() != 3 {
		t.Errorf("Len() = %d, want 3", units.Len())
	}

	units.Swap(0, 1)
	if units[0].Name != "A" || units[1].Name != "B" {
		t.Errorf("Swap() failed, got %v, want [A B C]", []string{units[0].Name, units[1].Name, units[2].Name})
	}

	tests := []struct {
		i, j int
		want bool
	}{
		{0, 1, true},  // Q1 < Q2
		{1, 2, true},  // Q2 < Q3, so Less(1,2) should be true
		{2, 0, false}, // Q3 > Q1
	}
	for _, tt := range tests {
		got := units.Less(tt.i, tt.j)
		if got != tt.want {
			t.Errorf("Less(%d, %d) = %v, want %v", tt.i, tt.j, got, tt.want)
		}
	}
}

func Test_All(t *testing.T) {
	units := All()
	if len(units) == 0 {
		t.Errorf("All() returned 0 units, want >0")
	}
	for i := 1; i < len(units); i++ {
		if units[i-1].Quantity.Name() > units[i].Quantity.Name() {
			t.Errorf("All() not sorted by Quantity: %s > %s", units[i-1].Quantity, units[i].Quantity)
		}
	}
}

// unitSystemTest provides a test case for verifying the UnitSystem of a unit
type unitSystemTest struct {
	u    Unit
	want UnitSystem
}

// testUnitSystems tests that units return the correct UnitSystem
func testUnitSystems(t *testing.T, tests []unitSystemTest) {
	for _, tt := range tests {
		got := tt.u.System()
		if got != tt.want {
			t.Errorf("Unit %s: want %s, got %s", tt.u.Name, tt.want, got)
		}
	}
}
