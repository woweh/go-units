package units

import (
	"testing"
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

func Test_Volume_UnitSystems(t *testing.T) {
	tests := []unitSystemTest{
		{Liter, SiSystem},
		{KiloLiter, SiSystem},
		{CentiLiter, SiSystem},
		{CubicMeter, SiSystem},
		{CubicCentiMeter, SiSystem},
		{Gallon, BiSystem},
		{CubicFoot, BiSystem},
		{CubicInch, BiSystem},
	}
	testUnitSystems(t, tests)
}
