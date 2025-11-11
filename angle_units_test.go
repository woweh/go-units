package units

import (
	"math"
	"testing"
)

func Test_Angle_Conversions(t *testing.T) {
	conversionTests := []conversionTest{
		// Revit and internal conversion factors
		{from: "radian", to: "degree", exp: 57.2957795130824},
		{from: "degree", to: "radian", exp: 0.0174532925199433},

		{from: "radian", to: "gon", exp: 63.6619772367582},
		{from: "gon", to: "radian", exp: 0.0157079632679489},

		{from: "radian", to: "DMS", exp: 57.174481, tol: fPtr(1e-6)},
		{from: "DMS", to: "radian", exp: 0.0174532925199433},

		// Internal implementation conversions
		{from: "turn", to: "radian", exp: 2 * math.Pi},
		{from: "turn", to: "degree", exp: 360},
		{from: "turn", to: "gon", exp: 400},
		{from: "degree", to: "gon", exp: 10.0 / 9.0},
		{from: "radian", to: "milliradian", exp: 1000},
		{from: "radian", to: "microradian", exp: 1000000},
		{from: "gon", to: "milligon", exp: 1000},
		{from: "gon", to: "centigon", exp: 100},
		{from: "turn", to: "milliradian", exp: 2 * math.Pi * 1000},

		// 1.0 DMS = 1°00'00'' = 1.0 degree and vice versa
		{from: "DMS", to: "degree", exp: 1.0},
		{from: "degree", to: "DMS", exp: 1.0},
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
