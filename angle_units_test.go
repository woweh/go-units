package units

import (
	"math"
	"testing"
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
	tests := []unitSystemTest{
		{Radian, SiSystem},
		{MilliRadian, SiSystem},
		{MicroRadian, SiSystem},
		{Gon, SiSystem},
		{MilliGon, SiSystem},
		{CentiGon, SiSystem},
		{DecaGon, SiSystem},
		{KiloGon, SiSystem},
		{Degree, SiSystem},
		{Turn, NoSystem},
		{DMS, NoSystem},
	}
	testUnitSystems(t, tests)
}
