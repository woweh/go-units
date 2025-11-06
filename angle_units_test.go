package units

import (
	"math"
	"testing"

	"github.com/alecthomas/assert/v2"
)

func Test_Angle_Conversions(t *testing.T) {
	conversionTests := []conversionTest{
		// Revit and internal conversion factors
		{"radian", "degree", 57.2957795130824},
		{"degree", "radian", 0.0174532925199433},

		{"radian", "gon", 63.6619772367582},
		{"gon", "radian", 0.0157079632679489},

		{"radian", "DMS", 57.174481},
		{"DMS", "radian", 0.0174532925199433},

		// Internal implementation conversions
		{"turn", "radian", 2 * math.Pi},
		{"turn", "degree", 360},
		{"turn", "gon", 400},
		{"radian", "degree", 180 / math.Pi}, // already above, but for completeness
		{"degree", "gon", 10.0 / 9.0},
		{"radian", "milliradian", 1000},
		{"radian", "microradian", 1000000},
		{"gon", "milligon", 1000},
		{"gon", "centigon", 100},
		{"turn", "milliradian", 2 * math.Pi * 1000},

		// DMS <-> degree (implementation)
		{"DMS", "degree", 12.5125}, // 12°30'45'' = 12.5125 deg
		{"degree", "DMS", 12.3045}, // 12.5125 deg = 12°30'45''
		{"DMS", "degree", -5.3020}, // -5°30'12'' = -5.503333... deg
		{"degree", "DMS", -5.3012}, // -5.503333... deg = -5°30'12''
	}
	testConversions(t, conversionTests)
}

func Test_Angle_UnitSystems(t *testing.T) {
	// SI units
	assert.Equal(t, SiSystem, Radian.System())
	assert.Equal(t, SiSystem, MilliRadian.System())
	assert.Equal(t, SiSystem, MicroRadian.System())
	assert.Equal(t, SiSystem, Gon.System())
	assert.Equal(t, SiSystem, MilliGon.System())
	assert.Equal(t, SiSystem, CentiGon.System())
	assert.Equal(t, SiSystem, DecaGon.System())
	assert.Equal(t, SiSystem, KiloGon.System())
	// Non-SI units
	assert.NotEqual(t, SiSystem, Degree.System())
	assert.NotEqual(t, SiSystem, Turn.System())
	assert.NotEqual(t, SiSystem, DMS.System())
}

func Test_Angle_BaseUnits(t *testing.T) {
	assert.Equal(t, Radian, Radian.Base())
	assert.Equal(t, Radian, MilliRadian.Base())
	assert.Equal(t, Radian, MicroRadian.Base())
	assert.Equal(t, Gon, MilliGon.Base())
	assert.Equal(t, Gon, CentiGon.Base())
	assert.Equal(t, Gon, DeciGon.Base())
}

func Test_Lookup_Angle_Names_and_Symbols(t *testing.T) {
	tests := lookUpTests{
		// Turn
		{Turn, "turn"}, {Turn, "tr"}, {Turn, "revolution"}, {Turn, "revolutions"}, {Turn, "revs"}, {Turn, "cycle"}, {Turn, "cycles"}, {Turn, "pla"}, {Turn, "rev"}, {Turn, "cyc"},
		// Radian
		{Radian, "radian"}, {Radian, "rad"},
		// Degree
		{Degree, "degree"}, {Degree, "degrees"}, {Degree, "degree of arc"}, {Degree, "degrees of arc"}, {Degree, "arc degree"}, {Degree, "arcdegree"}, {Degree, "°"}, {Degree, "deg"},
		// Gon
		{Gon, "gon"}, {Gon, "grad"}, {Gon, "gradian"}, {Gon, "gradians"}, {Gon, "grads"},
		// Metric-derived Gon units (symbols only for a few)
		{MilliGon, "milligon"}, {CentiGon, "centigon"},
		// DMS
		{DMS, "degree minute second"}, {DMS, "DMS"},
	}
	testLookupNamesAndSymbols(t, tests)
}
