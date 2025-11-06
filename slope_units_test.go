package units

import (
	"testing"

	"github.com/alecthomas/assert/v2"
)

func Test_Slope_Conversions(t *testing.T) {
	conversionTests := []conversionTest{
		// Revit conversions (from internal and to internal, see RevitUnits.csv)
		{"slope value", "slope percent", 100},
		{"slope percent", "slope value", 0.01},
		{"slope value", "slope permille", 1000},
		{"slope permille", "slope value", 0.001},
		{"slope value", "slope ratio", 1},
		{"slope ratio", "slope value", 1},
		{"slope value", "slope degree", 45},
		{"slope degree", "slope value", 1},

		// Core conversions (no identity, no redundant metric)
		{"slope value", "inverse slope ratio", 1},
		{"slope percent", "slope degree", 0.572939},
		{"slope percent", "slope permille", 10},
		{"slope permille", "slope degree", 0.0572958},
		{"slope degree", "slope percent", 1.745506},
		{"slope degree", "slope permille", 17.455065},
		{"slope degree", "slope ratio", 57.289962},
	}
	testConversions(t, conversionTests)
}

func Test_Slope_UnitSystem(t *testing.T) {
	assert.Equal(t, SiSystem, SlopeValue.System())
	assert.Equal(t, SiSystem, SlopeRatio.System())
	assert.Equal(t, SiSystem, SlopeInverseRatio.System())
	assert.Equal(t, SiSystem, SlopeDegree.System())
	assert.Equal(t, SiSystem, SlopePercent.System())
	assert.Equal(t, SiSystem, SlopePermille.System())
}

// Slope does not have derived metric units or metric factories, so no base unit tests are needed.

func Test_Slope_Lookup_Names_and_Symbols(t *testing.T) {
	tests := lookUpTests{
		{SlopeDegree, "deg"},
		{SlopePercent, "%"},
		{SlopePermille, "‰"},
		{SlopeDegree, "slope degree"},
		{SlopePercent, "slope percent"},
		{SlopePermille, "slope permille"},
		{SlopeValue, "slope value"},
		{SlopeRatio, "slope ratio"},
		{SlopeInverseRatio, "inverse slope ratio"},
	}
	testLookupNamesAndSymbols(t, tests)
}
