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
					t.Errorf("getUnitForExponent(%q, %d) = %v, want nil", tt.baseName, tt.exp, got)
				}
			} else {
				if got == nil {
					t.Errorf("getUnitForExponent(%q, %d) = nil, want %q", tt.baseName, tt.exp, tt.wantName)
				}
				if got.Name != tt.wantName {
					t.Errorf("getUnitForExponent(%q, %d) = %q, want %q", tt.baseName, tt.exp, got.Name, tt.wantName)
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
			t.Errorf("created mag unit: %s, want %s", mu.Name, magNames[i]+"dong")
		}
	}
}

func Test_findBestMatchingUnit(t *testing.T) {
	tests := []struct {
		name string
		base Unit
		exp  int
		want Unit
	}{
		// Full metric prefix set (Hertz)
		{"Zero exponent", Hertz, 0, Hertz},
		{"Non-metric unit", Foot, 3, Foot},
		{"Exact match - kilo", Hertz, 3, KiloHertz},
		{"Exact match - mega", Hertz, 6, MegaHertz},
		{"Exact match - milli", Hertz, -3, MilliHertz},
		{"Exact match - micro", Hertz, -6, MicroHertz},
		{"Exact match - hecto", Hertz, 2, HectoHertz},
		// Prefer lower if within 3
		{"exp=5: prefer kilo (10^3) over mega (10^6)", Hertz, 5, KiloHertz},
		{"exp=8: prefer mega (10^6) over giga (10^9)", Hertz, 8, MegaHertz},
		{"exp=-4: prefer micro (10^-6) over milli (10^-3)", Hertz, -4, MicroHertz},
		{"exp=-5: prefer micro (10^-6) over milli (10^-3)", Hertz, -5, MicroHertz},
		{"exp=7: exact distance, prefer lower mega", Hertz, 7, MegaHertz},
		// Extreme exponents
		{"exp=-100: return lowest defined (yocto at -24)", Hertz, -100, YoctoHertz},
		{"exp=100: return highest defined (yotta at 24)", Hertz, 100, YottaHertz},
		// Partial prefix set (VoltAmpere: only base, kilo, mega)
		{"Partial: exp=-100", VoltAmpere, -100, VoltAmpere},
		{"Partial: exp=-4", VoltAmpere, -4, VoltAmpere},
		{"Partial: exp=-3", VoltAmpere, -3, VoltAmpere},
		{"Partial: exp=-2", VoltAmpere, -2, VoltAmpere},
		{"Partial: exp=-1", VoltAmpere, -1, VoltAmpere},
		{"Partial: exp=0", VoltAmpere, 0, VoltAmpere},
		{"Partial: exp=1", VoltAmpere, 1, VoltAmpere},
		{"Partial: exp=2", VoltAmpere, 2, VoltAmpere},
		{"Partial: exp=3 (exact)", VoltAmpere, 3, KiloVoltAmpere},
		{"Partial: exp=4", VoltAmpere, 4, KiloVoltAmpere},
		{"Partial: exp=5", VoltAmpere, 5, KiloVoltAmpere},
		{"Partial: exp=6 (exact)", VoltAmpere, 6, MegaVoltAmpere},
		{"Partial: exp=7", VoltAmpere, 7, MegaVoltAmpere},
		{"Partial: exp=8", VoltAmpere, 8, MegaVoltAmpere},
		{"Partial: exp=15", VoltAmpere, 15, MegaVoltAmpere},
		{"Partial: exp=100", VoltAmpere, 100, MegaVoltAmpere},
		// Data units (uses kilo which is 10^3, not 1024)
		{"Data: exp=3", Byte, 3, KiloByte},
		// Edge cases around the "within 3" boundary
		{"Just within 3 - lower", Hertz, 11, GigaHertz},    // 11 is 2 away from giga(9)
		{"At boundary - exactly 3", Hertz, 12, TeraHertz},  // 12 is exact match
		{"Beyond 3 - choose closer", Hertz, 13, TeraHertz}, // 13 is closer to tera(12) than giga(9)
		// Edge: exp exactly at lowest and highest defined
		{"At lowest defined, exp=-24", Hertz, -24, YoctoHertz},
		{"At highest defined, exp=24", Hertz, 24, YottaHertz},
		// Edge: exp just above highest defined
		{"Just above highest, exp=25", Hertz, 25, YottaHertz},
		// Edge: exp just below lowest defined
		{"Just below lowest, exp=-25", Hertz, -25, YoctoHertz},
		// Edge: custom unit with gaps in prefixes
		{"Gap: exp=9 (no giga)", VoltAmpere, 9, MegaVoltAmpere},
		{"Gap: exp=-6 (no micro)", VoltAmpere, -6, VoltAmpere},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := findBestMatchingUnit(tt.base, tt.exp)
			if got != tt.want {
				t.Errorf("findBestMatchingUnit(%v, %d) = %v, want %v", tt.base, tt.exp, got, tt.want)
			}
		})
	}
}
