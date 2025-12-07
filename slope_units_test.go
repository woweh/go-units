package units

import (
	"testing"
)

func Test_Slope_Conversions(t *testing.T) {
	conversionTests := []conversionTest{
		// Revit conversions (from internal and to internal, see RevitUnits.csv)
		{from: "slope value", to: "slope percent", exp: 100},
		{from: "slope percent", to: "slope value", exp: 0.01},
		{from: "slope value", to: "slope permille", exp: 1000},
		{from: "slope permille", to: "slope value", exp: 0.001},
		{from: "slope value", to: "slope ratio", exp: 1},
		{from: "slope ratio", to: "slope value", exp: 1},
		{from: "slope value", to: "slope degree", exp: 45},
		{from: "slope degree", to: "slope value", exp: 0.017455065, tol: fPtr(1e-8)}, // tan(1 degree)
		// Core conversions (no identity, no redundant metric)
		{from: "slope value", to: "inverse slope ratio", exp: 1},
		{from: "slope percent", to: "slope degree", exp: 0.572938698, tol: fPtr(1e-8)}, // atan(0.01) * 180/pi
		{from: "slope percent", to: "slope permille", exp: 10},
		{from: "slope permille", to: "slope degree", exp: 0.05729573, tol: fPtr(1e-7)},   // atan(0.001) * 180/pi
		{from: "slope degree", to: "slope percent", exp: 1.745506493, tol: fPtr(1e-8)},   // tan(1) * 100
		{from: "slope degree", to: "slope permille", exp: 17.455064928, tol: fPtr(1e-7)}, // tan(1) * 1000
		{from: "slope degree", to: "slope ratio", exp: 57.289961631, tol: fPtr(1e-7)},    // 1 / tan(1)
	}
	testConversions(t, conversionTests)
}

func Test_Slope_UnitSystem(t *testing.T) {
	tests := []unitSystemTest{
		{SlopeValue, SiSystem},
		{SlopeRatio, SiSystem},
		{SlopeInverseRatio, SiSystem},
		{SlopeDegree, SiSystem},
		{SlopePercent, SiSystem},
		{SlopePermille, SiSystem},
	}
	testUnitSystems(t, tests)
}
