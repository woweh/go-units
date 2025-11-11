package units

import (
	"testing"
)

func Test_DimensionlessRatio_Conversions(t *testing.T) {
	conversionTests := []conversionTest{
		// Base conversions from fraction
		{from: "fraction", to: "percent", exp: 100},
		{from: "fraction", to: "permille", exp: 1000},
		{from: "fraction", to: "partsPerMillion", exp: 1000000},
		{from: "fraction", to: "partsPerBillion", exp: 1000000000},
		{from: "fraction", to: "partsPerTrillion", exp: 1000000000000},
		// Base conversions to fraction
		{from: "percent", to: "fraction", exp: 0.01},
		{from: "permille", to: "fraction", exp: 0.001},
		{from: "partsPerMillion", to: "fraction", exp: 0.000001},
		{from: "partsPerBillion", to: "fraction", exp: 0.000000001},
		{from: "partsPerTrillion", to: "fraction", exp: 0.000000000001},
		// Additional important conversions
		{from: "percent", to: "permille", exp: 10},
		{from: "permille", to: "partsPerMillion", exp: 1000},
		{from: "partsPerMillion", to: "partsPerBillion", exp: 1000},
		{from: "partsPerBillion", to: "partsPerTrillion", exp: 1000},
		// Cross-system (sampled, not all)
		{from: "percent", to: "partsPerMillion", exp: 10000},
		{from: "partsPerMillion", to: "percent", exp: 0.0001},
	}
	testConversions(t, conversionTests)
}

func Test_DimensionlessRatio_UnitSystems(t *testing.T) {
	tests := []unitSystemTest{
		{Fraction, SiSystem},
		{Percent, SiSystem},
		{Permille, SiSystem},
		{PartsPerMillion, SiSystem},
		{PartsPerBillion, SiSystem},
		{PartsPerTrillion, SiSystem},
	}
	testUnitSystems(t, tests)
}
