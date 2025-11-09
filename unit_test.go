package units

import (
	"errors"
	"strings"
	"testing"
)

// quantityForUnitTests is a fake quantity for testing.
const quantityForUnitTests = "<|_quantity_for_unit_tests_|>"

func Test_NewUnit(t *testing.T) {
	tests := []struct {
		name       string
		unitName   string
		unitSymbol string
		wantErr    bool
	}{
		{
			name:       "Add TestUnit_1",
			unitName:   "TestUnit_1",
			unitSymbol: "T1",
			wantErr:    false,
		},
		{
			name:       "Test empty name",
			unitName:   "",
			unitSymbol: "T1",
			wantErr:    true,
		},
		{
			name:       "Test duplicate name (TestUnit_1)",
			unitName:   "TestUnit_1",
			unitSymbol: "",
			wantErr:    true,
		},
		{
			name:       "Testing for duplicate symbol (T1)",
			unitName:   "TestUnit_1.0",
			unitSymbol: "T1",
			wantErr:    true,
		},
		{
			name:       "Testing for case sensitive symbol (t1 != T1)",
			unitName:   "TestUnit_1.1",
			unitSymbol: "t1",
			wantErr:    false,
		},
	}

	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				u, err := NewUnit(tt.unitName, tt.unitSymbol, Quantity(quantityForUnitTests))
				if (err != nil) != tt.wantErr {
					t.Errorf("NewUnit() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
				if u != nil {
					if u.Name != tt.unitName {
						t.Errorf("NewUnit() = %v, want %v", u.Name, tt.unitName)
					}
					if u.Symbol != tt.unitSymbol {
						t.Errorf("NewUnit() = %v, want %v", u.Symbol, tt.unitSymbol)
					}
				}
			},
		)
	}
}

func TestUnit_Names(t *testing.T) {
	testUnit, err := NewUnit("TestUnit_Names", "TU_Names", Quantity(quantityForUnitTests))
	if err != nil || testUnit == nil {
		t.Errorf("NewUnit() failed: %v", err)
	}
	testUnitWithPlural, err := NewUnit("TestUnitPlural_Names", "TUP_Names", Plural("TestUnitPlurals_Names"), Quantity(quantityForUnitTests))
	if err != nil || testUnitWithPlural == nil {
		t.Errorf("NewUnit() failed: %v", err)
	}
	testUnitWithAliases, err := NewUnit("TestUnitAlias_Names", "TUA_Names", Aliases("Alias1_Names", "Alias2_Names"), Quantity(quantityForUnitTests))
	if err != nil || testUnitWithAliases == nil {
		t.Errorf("NewUnit() failed: %v", err)
	}
	tests := []struct {
		unit     Unit
		expected []string
	}{
		{testUnit, []string{"TestUnit_Names"}},
		{testUnitWithPlural, []string{"TestUnitPlural_Names", "TestUnitPlurals_Names"}},
		{testUnitWithAliases, []string{"TestUnitAlias_Names", "Alias1_Names", "Alias2_Names"}},
	}
	for _, tt := range tests {
		got := tt.unit.Names()
		for _, want := range tt.expected {
			found := false
			for _, g := range got {
				if g == want {
					found = true
					break
				}
			}
			if !found {
				t.Errorf("Names() missing %q in %v", want, got)
			}
		}
	}
}

func TestUnit_Aliases(t *testing.T) {
	testUnitWithAliases, err := NewUnit("TestUnitAlias_Aliases", "TUA_Aliases", Aliases("Alias1_Aliases", "Alias2_Aliases"), Quantity(quantityForUnitTests))
	if err != nil || testUnitWithAliases == nil {
		t.Errorf("NewUnit() failed: %v", err)
	}
	got := testUnitWithAliases.Aliases()
	if len(got) != 2 || got[0] != "Alias1_Aliases" || got[1] != "Alias2_Aliases" {
		t.Errorf("Aliases() = %v, want [Alias1_Aliases Alias2_Aliases]", got)
	}
}

func TestUnit_Symbols(t *testing.T) {
	testUnitWithSymbols, err := NewUnit("TestUnitSymbol_Symbols", "TUS_Symbols", Symbols("Sym1_Symbols", "Sym2_Symbols"), Quantity(quantityForUnitTests))
	if err != nil || testUnitWithSymbols == nil {
		t.Errorf("NewUnit() failed: %v", err)
	}
	got := testUnitWithSymbols.Symbols()
	if len(got) != 3 || got[0] != "TUS_Symbols" || got[1] != "Sym1_Symbols" || got[2] != "Sym2_Symbols" {
		t.Errorf("Symbols() = %v, want [TUS_Symbols Sym1_Symbols Sym2_Symbols]", got)
	}
}

func TestUnit_String(t *testing.T) {
	testUnit, err := NewUnit("TestUnit_String", "TU_String", Quantity(quantityForUnitTests))
	if err != nil || testUnit == nil {
		t.Errorf("NewUnit() failed: %v", err)
	}
	if got := testUnit.String(); got != "TestUnit_String" {
		t.Errorf("String() = %v, want TestUnit_String", got)
	}
}

func TestUnit_System(t *testing.T) {
	unit, err := NewUnit("TestUnitSys_System", "TUSY_System", System(SiSystem), Quantity(quantityForUnitTests))
	if err != nil || unit == nil {
		t.Errorf("NewUnit() failed: %v", err)
	}
	if got := unit.System(); got != SiSystem {
		t.Errorf("System() = %v, want %v", got, SiSystem)
	}
}

func TestUnit_PluralName(t *testing.T) {
	testUnit, err := NewUnit("TestUnit_PluralName", "TU_PluralName", Quantity(quantityForUnitTests))
	if err != nil || testUnit == nil {
		t.Errorf("NewUnit() failed: %v", err)
	}
	testUnitWithPlural, err := NewUnit("TestUnitPlural_PluralName", "TUP_PluralName", Plural("TestUnitPlurals_PluralName"), Quantity(quantityForUnitTests))
	if err != nil || testUnitWithPlural == nil {
		t.Errorf("NewUnit() failed: %v", err)
	}
	if got := testUnitWithPlural.PluralName(); got != "TestUnitPlurals_PluralName" {
		t.Errorf("PluralName() = %v, want TestUnitPlurals_PluralName", got)
	}
	if got := testUnit.PluralName(); got != "TestUnit_PluralName"+"s" {
		t.Errorf("PluralName() = %v, want TestUnit_PluralNames", got)
	}
}

func TestUnit_HasName(t *testing.T) {
	testUnitWithAliases, err := NewUnit("TestUnitAlias_HasName", "TUA_HasName", Aliases("Alias1_HasName", "Alias2_HasName"), Quantity(quantityForUnitTests))
	if err != nil || testUnitWithAliases == nil {
		t.Errorf("NewUnit() failed: %v", err)
	}
	if !testUnitWithAliases.HasName("Alias1_HasName") {
		t.Errorf("HasName(Alias1_HasName) = false, want true")
	}
	if !testUnitWithAliases.HasName("TestUnitAlias_HasName") {
		t.Errorf("HasName(TestUnitAlias_HasName) = false, want true")
	}
	if testUnitWithAliases.HasName("NotAnAlias_HasName") {
		t.Errorf("HasName(NotAnAlias_HasName) = true, want false")
	}
}

func TestUnit_HasSymbol(t *testing.T) {
	testUnitWithSymbols, err := NewUnit("TestUnitSymbol_HasSymbol", "TUS_HasSymbol", Symbols("Sym1_HasSymbol", "Sym2_HasSymbol"), Quantity(quantityForUnitTests))
	if err != nil || testUnitWithSymbols == nil {
		t.Errorf("NewUnit() failed: %v", err)
	}
	if !testUnitWithSymbols.HasSymbol("Sym1_HasSymbol") {
		t.Errorf("HasSymbol(Sym1_HasSymbol) = false, want true")
	}
	if !testUnitWithSymbols.HasSymbol("TUS_HasSymbol") {
		t.Errorf("HasSymbol(TUS_HasSymbol) = false, want true")
	}
	if testUnitWithSymbols.HasSymbol("NotASymbol_HasSymbol") {
		t.Errorf("HasSymbol(NotASymbol_HasSymbol) = true, want false")
	}
}

func TestUnit_CsvLine(t *testing.T) {
	unit, err := NewUnit("TestUnitCSV_CsvLine", "TUC_CsvLine", Plural("TestUnitCSVs_CsvLine"), Aliases("A1_CsvLine", "A2_CsvLine"), Symbols("S1_CsvLine", "S2_CsvLine"), Quantity(quantityForUnitTests))
	if err != nil || unit == nil {
		t.Errorf("NewUnit() failed: %v", err)
	}
	got := unit.CsvLine()
	wantPrefix := "TestUnitCSV_CsvLine,TUC_CsvLine,TestUnitCSVs_CsvLine,<|_quantity_for_unit_tests_|>,,A1_CsvLine,A2_CsvLine,S1_CsvLine,S2_CsvLine"
	if !strings.HasPrefix(got, wantPrefix[:20]) {
		t.Errorf("CsvLine() = %v, want prefix %v", got, wantPrefix)
	}
}

func TestUnit_IsMetric(t *testing.T) {
	unitMetric, err := NewUnit("TestMetric_IsMetric", "TM_IsMetric", System(SiSystem), Quantity(quantityForUnitTests))
	if err != nil || unitMetric == nil {
		t.Errorf("NewUnit() failed: %v", err)
	}
	if !unitMetric.IsMetric() {
		t.Errorf("IsMetric() = false, want true")
	}
}

func TestUnit_IsImperial(t *testing.T) {
	unitImperial, err := NewUnit("TestImperial_IsImperial", "TI_IsImperial", System(BiSystem), Quantity(quantityForUnitTests))
	if err != nil || unitImperial == nil {
		t.Errorf("NewUnit() failed: %v", err)
	}
	if !unitImperial.IsImperial() {
		t.Errorf("IsImperial() = false, want true")
	}
}

func TestUnit_IsUS(t *testing.T) {
	unitUS, err := NewUnit("TestUS_IsUS", "TUUS_IsUS", System(UsSystem), Quantity(quantityForUnitTests))
	if err != nil || unitUS == nil {
		t.Errorf("NewUnit() failed: %v", err)
	}
	if !unitUS.IsUS() {
		t.Errorf("IsUS() = false, want true")
	}
}

func TestUnit_IsIEC(t *testing.T) {
	unitIEC, err := NewUnit("TestIEC_IsIEC", "TIEC_IsIEC", System(IecSystem), Quantity(quantityForUnitTests))
	if err != nil || unitIEC == nil {
		t.Errorf("NewUnit() failed: %v", err)
	}
	if !unitIEC.IsIEC() {
		t.Errorf("IsIEC() = false, want true")
	}
}

func TestUnit_IsCGS(t *testing.T) {
	unitCGS, err := NewUnit("TestCGS_IsCGS", "TCGS_IsCGS", System(CgsSystem), Quantity(quantityForUnitTests))
	if err != nil || unitCGS == nil {
		t.Errorf("NewUnit() failed: %v", err)
	}
	if !unitCGS.IsCGS() {
		t.Errorf("IsCGS() = false, want true")
	}
}

func TestUnit_IsMKpS(t *testing.T) {
	unitMKpS, err := NewUnit("TestMKpS_IsMKpS", "TMKpS_IsMKpS", System(MKpSSystem), Quantity(quantityForUnitTests))
	if err != nil || unitMKpS == nil {
		t.Errorf("NewUnit() failed: %v", err)
	}
	if !unitMKpS.IsMKpS() {
		t.Errorf("IsMKpS() = false, want true")
	}
}

func Test_AddAliases(t *testing.T) {
	testUnit, err := NewUnit("TestAdd_Aliases", "TA_Aliases", Quantity(quantityForUnitTests))
	if err != nil || testUnit == nil {
		t.Errorf("NewUnit() failed: %v", err)
	}
	result := testUnit.AddAliases("AliasA_Aliases", "AliasB_Aliases")
	if len(result.Added) != 2 {
		t.Errorf("AddAliases() Added = %v, want 2", result.Added)
	}
	if result.Err != nil {
		t.Errorf("AddAliases() Err = %v, want nil", result.Err)
	}
}

func Test_AddSymbols(t *testing.T) {
	testUnit, err := NewUnit("TestAdd_Symbols", "TA_Symbols", Quantity(quantityForUnitTests))
	if err != nil || testUnit == nil {
		t.Errorf("NewUnit() failed: %v", err)
	}
	result := testUnit.AddSymbols("SymA_Symbols", "SymB_Symbols")
	if len(result.Added) != 2 {
		t.Errorf("AddSymbols() Added = %v, want 2", result.Added)
	}
	if result.Err != nil {
		t.Errorf("AddSymbols() Err = %v, want nil", result.Err)
	}
}

func Test_AddResult_String(t *testing.T) {
	ar := &AddResult{
		What:     "Aliases",
		Unit:     &unit{Name: "TestUnit_AddResultString"},
		Added:    []string{"A1_AddResultString", "A2_AddResultString"},
		Failures: map[string]error{"F1_AddResultString": errors.New("fail")},
	}
	ar.validate()
	got := ar.String()
	if got == "" {
		t.Errorf("AddResult.String() = empty, want non-empty summary")
	}
}
