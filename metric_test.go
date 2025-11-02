package units

import (
	"testing"
)

var magNames = []string{
	"exa",
	"peta",
	"tera",
	"giga",
	"mega",
	"kilo",
	"hecto",
	"deca",
	"deci",
	"centi",
	"milli",
	"micro",
	"nano",
	"pico",
	"femto",
	"atto",
}

type magFn func(Unit, ...UnitOption) Unit

func Test_getUnitForExponent(t *testing.T) {
	t.Parallel()
	u := mustCreateNewUnit("testunit_gufe", "TUGUFE")
	Kilo(u)
	tests := []struct {
		name     string
		baseName string
		exp      int
		wantNil  bool
		wantName string
	}{
		{"Exact match", "testunit_gufe", 3, false, "kilotestunit_gufe"},
		{"No match", "testunit_gufe", 5, true, ""},
		{"Wrong base", "notexist", 3, true, ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := getUnitForExponent(tt.baseName, tt.exp)
			if tt.wantNil {
				if got != nil {
					t.Fatalf("getUnitForExponent(%q, %d) = %v, want nil", tt.baseName, tt.exp, got)
				}
			} else {
				if got == nil {
					t.Fatalf("getUnitForExponent(%q, %d) = nil, want %q", tt.baseName, tt.exp, tt.wantName)
				}
				if got.Name != tt.wantName {
					t.Fatalf("getUnitForExponent(%q, %d) = %q, want %q", tt.baseName, tt.exp, got.Name, tt.wantName)
				}
			}
		})
	}
}

func Test_findNextLowerUnit_and_findNextHigherUnit(t *testing.T) {
	t.Parallel()
	u := mustCreateNewUnit("testunit_fnlufnhu", "TUFNLU")
	Kilo(u)
	Mega(u)
	Deci(u)
	Centi(u)
	tests := []struct {
		name       string
		exp        int
		wantLower  string
		wantHigher string
	}{
		{"Between kilo and mega", 5, "kilotestunit_fnlufnhu", "megatestunit_fnlufnhu"},
		{"Below deci", -2, "", "decitestunit_fnlufnhu"},                                // No lower unit for exp -2, only higher
		{"No lower, only higher", 2, "decitestunit_fnlufnhu", "kilotestunit_fnlufnhu"}, // Lower is deci, higher is kilo
		{"No higher, only lower", 6, "kilotestunit_fnlufnhu", ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lower, _ := findNextLowerUnit(u, tt.exp)
			higher, _ := findNextHigherUnit(u, tt.exp)
			if tt.wantLower == "" {
				if lower != nil {
					t.Fatalf("findNextLowerUnit(%d) = %v, want nil", tt.exp, lower)
				}
			} else {
				if lower == nil || lower.Name != tt.wantLower {
					t.Fatalf("findNextLowerUnit(%d) = %v, want %q", tt.exp, lower, tt.wantLower)
				}
			}
			if tt.wantHigher == "" {
				if higher != nil {
					t.Fatalf("findNextHigherUnit(%d) = %v, want nil", tt.exp, higher)
				}
			} else {
				if higher == nil || higher.Name != tt.wantHigher {
					t.Fatalf("findNextHigherUnit(%d) = %v, want %q", tt.exp, higher, tt.wantHigher)
				}
			}
		})
	}
}

func Test_Magnitudes(t *testing.T) {
	u := mustCreateNewUnit("dong", "₫")
	for i, mfn := range []magFn{
		Exa, Peta, Tera, Giga, Mega, Kilo, Hecto, Deca, Deci, Centi, Milli, Micro, Nano, Pico, Femto, Atto,
	} {
		mu := mfn(u)
		if mu.Name != magNames[i]+"dong" {
			t.Fatalf("created mag unit: %s, want %s", mu.Name, magNames[i]+"dong")
		}
	}
}

func Test_findBestMatchingUnit(t *testing.T) {
	tests := []struct {
		name   string
		exp    int
		expect Unit
		base   Unit
	}{
		{"Zero", 0, Hertz, Hertz},
		{"Non metric", 3, Foot, Foot},
		{"Simple", 3, KiloHertz, Hertz},
		{"Simple negative", -3, MilliHertz, Hertz},
		{"Simple, unusual", 2, HectoHertz, Hertz},
		{"Not defined prefix", 2, VoltAmpere, VoltAmpere},
		{"Partially defined, works", 3, KiloVoltAmpere, VoltAmpere},
		{"Simple, not existing prefix", 5, KiloHertz, Hertz},
		{"Partially defined, over", 9, MegaVoltAmpere, VoltAmpere},
		{"Partially defined negative", -10, VoltAmpere, VoltAmpere},
		{"Data", 3, KiloByte, Byte},
	}

	for _, tst := range tests {
		t.Run(tst.name, func(t *testing.T) {
			calcUnit := findBestMatchingUnit(tst.base, tst.exp)
			if calcUnit != tst.expect {
				t.Fatalf("findBestMatchingUnit(%v, %d) = %v, want %v", tst.base, tst.exp, calcUnit, tst.expect)
			}
		})
	}
}
