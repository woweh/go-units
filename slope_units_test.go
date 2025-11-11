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
		{from: "slope degree", to: "slope value", exp: 1},

		// Core conversions (no identity, no redundant metric)
		{from: "slope value", to: "inverse slope ratio", exp: 1},
		{from: "slope percent", to: "slope degree", exp: 0.572939},
		{from: "slope percent", to: "slope permille", exp: 10},
		{from: "slope permille", to: "slope degree", exp: 0.0572958},
		{from: "slope degree", to: "slope percent", exp: 1.745506},
		{from: "slope degree", to: "slope permille", exp: 17.455065},
		{from: "slope degree", to: "slope ratio", exp: 57.289962},
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
