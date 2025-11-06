package units

import (
	"testing"

	"github.com/alecthomas/assert/v2"
)

func Test_Volume_Conversions(t *testing.T) {
	conversionTests := []conversionTest{
		// SI/metric
		{"cubic meter", "liter", 1000},
		{"liter", "cubic meter", 0.001},
		{"cubic centimeter", "liter", 0.001},
		{"liter", "cubic centimeter", 1000},
		// Metric factory
		{"kiloliter", "liter", 1000},
		{"liter", "kiloliter", 0.001},
		{"centiliter", "liter", 0.01},
		{"liter", "centiliter", 100},
		// Imperial/US
		{"gallon", "liter", 4.54609},
		{"liter", "gallon", 0.2199692},
		{"cubic foot", "liter", 28.316846592},
		{"liter", "cubic foot", 0.0353146667214886},
		// Cross-system
		{"cubic inch", "liter", 0.016387064},
		{"liter", "cubic inch", 61.0237440947323},
	}
	testConversions(t, conversionTests)
}

func Test_Volume_UnitSystem(t *testing.T) {
	assert.Equal(t, SiSystem, Liter.System())
	assert.Equal(t, SiSystem, KiloLiter.System())
	assert.Equal(t, SiSystem, CentiLiter.System())
	assert.Equal(t, SiSystem, CubicMeter.System())
	assert.Equal(t, SiSystem, CubicCentiMeter.System())
	assert.Equal(t, BiSystem, Gallon.System())
	assert.Equal(t, BiSystem, CubicFoot.System())
	assert.Equal(t, BiSystem, CubicInch.System())
}

func Test_Volume_MetricFactory_BaseUnits(t *testing.T) {
	assert.Equal(t, Liter, CentiLiter.Base())
	assert.Equal(t, Liter, KiloLiter.Base())
}

func Test_Volume_Lookup_Names_and_Symbols(t *testing.T) {
	tests := lookUpTests{
		// SI/metric
		{Liter, "l"},
		{Liter, "liter"},
		{Liter, "litre"},
		{KiloLiter, "kl"},
		{CentiLiter, "cl"},
		{CubicMeter, "m³"},
		{CubicMeter, "cubic meter"},
		{CubicMeter, "cubic metre"},
		{CubicMeter, "m3"},
		{CubicCentiMeter, "cm³"},
		{CubicCentiMeter, "cubic centimeter"},
		{CubicCentiMeter, "cubic centimetre"},
		{CubicCentiMeter, "cc"},
		// Imperial/US
		{Gallon, "gal"},
		{Gallon, "gallon"},
		{CubicFoot, "ft³"},
		{CubicFoot, "cubic foot"},
		{CubicInch, "in³"},
		{CubicInch, "cubic inch"},
	}
	testLookupNamesAndSymbols(t, tests)
}
