package units

import (
	"os"
	"strings"
	"testing"
)

// aggregate all unit names, aliases, etc
func aggrNames() (a []string) {
	for _, u := range All() {
		a = append(a, u.Names()...)
	}
	return a
}

// aggregate units by quantity
func aggrByQuantity() map[UnitQuantity][]Unit {
	m := make(map[UnitQuantity][]Unit)

	for _, u := range All() {
		if _, ok := m[u.Quantity]; !ok {
			m[u.Quantity] = []Unit{}
		}
		m[u.Quantity] = append(m[u.Quantity], u)
	}

	return m
}

func Test_UnitLookup(t *testing.T) {
	for _, name := range aggrNames() {
		u, err := Find(name)
		if err != nil {
			t.Fatalf("Find(%q) error: %v", name, err)
		}
		t.Logf("found unit by name: %s (%s)", name, u.Name)
	}
}

func Test_UnitNameOverlap(t *testing.T) {
	nameMap := make(map[string]Unit)

	var total, failed int
	for _, u := range nameMap {
		for _, name := range u.Names() {
			if existing, ok := nameMap[name]; ok {
				t.Fatalf("overlap in unit names: %s, %s (%s)", u.Name, existing.Name, name)
			} else {
				nameMap[name] = u
			}
			total++
		}
	}
	t.Logf("tested %d unit names, %d overlaps", total, failed)
}

// ensure all units within the same quantity resolve a conversion path
func Test_PathResolve(t *testing.T) {
	for qname, qunits := range aggrByQuantity() {
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
					t.Fatalf("failed to resolve path: %s -> %s: %v", u1.Name, u2.Name, err)
				}
			}
		}
	}
}

func Test_ListUnitsAsCsv(t *testing.T) {
	f, err := ListUnitsAsCsv()
	if err != nil {
		t.Fatalf("ListUnitsAsCsv() error: %v", err)
	}
	t.Log("created file: " + f)
}

func ListUnitsAsCsv() (fileName string, err error) {

	fileName = "units.csv"

	csvLines := GetCsv()

	// Add UTF-8 BOM at the beginning
	bom := []byte{0xEF, 0xBB, 0xBF}
	windowsNewLine := "\r\n"
	content := append(bom, []byte(strings.Join(csvLines, windowsNewLine))...)

	err = os.WriteFile(fileName, content, 0644)

	return fileName, nil
}

func Test_UnitList_Len_Swap_Less(t *testing.T) {
	units := UnitList{
		&unit{Name: "B", Quantity: "Q2", system: "sys2"},
		&unit{Name: "A", Quantity: "Q1", system: "sys1"},
		&unit{Name: "C", Quantity: "Q1", system: "sys2"},
	}

	if units.Len() != 3 {
		t.Fatalf("Len() = %d, want 3", units.Len())
	}

	units.Swap(0, 1)
	if units[0].Name != "A" || units[1].Name != "B" {
		t.Fatalf("Swap() failed, got %v, want [A B C]", []string{units[0].Name, units[1].Name, units[2].Name})
	}

	tests := []struct {
		i, j int
		want bool
	}{
		{0, 1, true},  // Q1 < Q2
		{1, 2, false}, // Q2 > Q1, so Less(1,2) should be false
		{2, 0, false}, // Q1 == Q1, sys2 > sys1
	}
	for _, tt := range tests {
		got := units.Less(tt.i, tt.j)
		if got != tt.want {
			t.Fatalf("Less(%d, %d) = %v, want %v", tt.i, tt.j, got, tt.want)
		}
	}
}

func Test_All(t *testing.T) {
	units := All()
	if len(units) == 0 {
		t.Fatalf("All() returned 0 units, want >0")
	}
	for i := 1; i < len(units); i++ {
		if units[i-1].Quantity > units[i].Quantity {
			t.Fatalf("All() not sorted by Quantity: %s > %s", units[i-1].Quantity, units[i].Quantity)
		}
	}
}

func Test_MustConvertFloat(t *testing.T) {
	unitA, err := NewUnit("TestUnit_MustConvertFloat_A", "TUMCFA", Quantity("mustconvfloat"))
	if err != nil || unitA == nil {
		t.Fatalf("NewUnit() failed: %v", err)
	}
	unitB, err := NewUnit("TestUnit_MustConvertFloat_B", "TUMCFB", Quantity("mustconvfloat"))
	if err != nil || unitB == nil {
		t.Fatalf("NewUnit() failed: %v", err)
	}
	NewRatioConversion(unitA, unitB, 2.0)
	got := MustConvertFloat(3, unitA, unitB)
	if got.val != 6 {
		t.Fatalf("MustConvertFloat(3, unitA, unitB) = %v, want 6", got.val)
	}
	defer func() {
		if r := recover(); r == nil {
			t.Fatalf("MustConvertFloat should panic on error, but did not")
		}
	}()
	_ = MustConvertFloat(1, nil, unitB)
}

func Test_ConvertFloat(t *testing.T) {
	unitA, err := NewUnit("TestUnit_ConvertFloat_A", "TUCFA", Quantity("convfloat"))
	if err != nil || unitA == nil {
		t.Fatalf("NewUnit() failed: %v", err)
	}
	unitB, err := NewUnit("TestUnit_ConvertFloat_B", "TUCFB", Quantity("convfloat"))
	if err != nil || unitB == nil {
		t.Fatalf("NewUnit() failed: %v", err)
	}
	NewRatioConversion(unitA, unitB, 2.0)

	tests := []struct {
		x       float64
		from    Unit
		to      Unit
		want    float64
		wantErr bool
	}{
		{3, unitA, unitB, 6, false},
		{6, unitB, unitA, 3, false},
		{1, unitA, unitA, 1, false},
		{1, nil, unitA, 0, true},
		{1, unitA, nil, 0, true},
	}
	for _, tt := range tests {
		got, err := ConvertFloat(tt.x, tt.from, tt.to)
		if tt.wantErr {
			if err == nil {
				t.Fatalf("ConvertFloat(%v, %v, %v) did not error, want error", tt.x, tt.from, tt.to)
			}
			continue
		}
		if err != nil {
			t.Fatalf("ConvertFloat(%v, %v, %v) error: %v", tt.x, tt.from, tt.to, err)
		}
		if got.val != tt.want {
			t.Fatalf("ConvertFloat(%v, %v, %v) = %v, want %v", tt.x, tt.from, tt.to, got.val, tt.want)
		}
	}
}

func Test_GetCsv(t *testing.T) {
	csv := GetCsv()
	if len(csv) == 0 {
		t.Fatalf("GetCsv() returned empty slice")
	}
	if !strings.HasPrefix(csv[0], "Name,Symbol,PluralName,Quantity,System,Aliases & Symbols") {
		t.Fatalf("GetCsv() header incorrect: %v", csv[0])
	}
	found := false
	for _, line := range csv[1:] {
		if strings.Contains(line, ",") {
			found = true
			break
		}
	}
	if !found {
		t.Fatalf("GetCsv() did not return any data lines with commas")
	}
}
