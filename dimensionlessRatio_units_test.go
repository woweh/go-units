package units

import (
	"testing"

	"github.com/alecthomas/assert/v2"
)

func Test_DimensionlessRatio_Conversions(t *testing.T) {
	conversionTests := []conversionTest{
		// Base conversions from fraction
		{"fraction", "percent", 100},
		{"fraction", "permille", 1000},
		{"fraction", "partsPerMillion", 1000000},
		{"fraction", "partsPerBillion", 1000000000},
		{"fraction", "partsPerTrillion", 1000000000000},
		// Base conversions to fraction
		{"percent", "fraction", 0.01},
		{"permille", "fraction", 0.001},
		{"partsPerMillion", "fraction", 0.000001},
		{"partsPerBillion", "fraction", 0.000000001},
		{"partsPerTrillion", "fraction", 0.000000000001},
		// Additional important conversions
		{"percent", "permille", 10},
		{"permille", "partsPerMillion", 1000},
		{"partsPerMillion", "partsPerBillion", 1000},
		{"partsPerBillion", "partsPerTrillion", 1000},
		// Cross-system (sampled, not all)
		{"percent", "partsPerMillion", 10000},
		{"partsPerMillion", "percent", 0.0001},
	}
	testConversions(t, conversionTests)
}

func Test_Lookup_DimensionlessRatio_Names_and_Symbols(t *testing.T) {
	tests := lookUpTests{
		{Fraction, "fraction"},
		{Percent, "percent"},
		{Percent, "%"},
		{Permille, "permille"},
		{Permille, "‰"},
		{PartsPerMillion, "partsPerMillion"},
		{PartsPerMillion, "ppm"},
		{PartsPerBillion, "partsPerBillion"},
		{PartsPerTrillion, "partsPerTrillion"},
	}
	testLookupNamesAndSymbols(t, tests)
}

func Test_DimensionlessRatio_UnitSystems(t *testing.T) {
	assert.Equal(t, SiSystem, Fraction.System())
	assert.Equal(t, SiSystem, Percent.System())
	assert.Equal(t, SiSystem, Permille.System())
	assert.Equal(t, SiSystem, PartsPerMillion.System())
	assert.Equal(t, SiSystem, PartsPerBillion.System())
	assert.Equal(t, SiSystem, PartsPerTrillion.System())
}
