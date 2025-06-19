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
func aggrByQuantity() map[UnitQuantity][]*Unit {
	m := make(map[UnitQuantity][]*Unit)

	for _, u := range All() {
		if _, ok := m[u.Quantity]; !ok {
			m[u.Quantity] = []*Unit{}
		}
		m[u.Quantity] = append(m[u.Quantity], u)
	}

	return m
}

func Test_UnitLookup(t *testing.T) {
	for _, name := range aggrNames() {
		u, err := Find(name)
		if err != nil {
			t.Error(err.Error())
			continue
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
				t.Errorf("overlap in unit names: %s, %s (%s)", u.Name, existing.Name, name)
				failed++
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
			v1 := NewValue(1.0, *u1)
			for _, u2 := range qunits {
				if u1.Name == u2.Name {
					continue
				}
				_, err := v1.Convert(*u2)
				if err != nil {
					t.Errorf("failed to resolve path: %s -> %s", u1.Name, u2.Name)
				}
			}
		}
	}
}

func Test_ListUnitsAsCsv(t *testing.T) {
	f, err := ListUnitsAsCsv()
	if err != nil {
		t.Error(err)
		return
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
