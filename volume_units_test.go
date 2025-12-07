package units

import (
	"testing"
)

func Test_Volume_Conversions(t *testing.T) {
	conversionTests := []conversionTest{
		// SI/metric
		{from: "cubic meter", to: "liter", exp: 1000},
		{from: "liter", to: "cubic meter", exp: 0.001},
		{from: "cubic centimeter", to: "liter", exp: 0.001},
		{from: "liter", to: "cubic centimeter", exp: 1000},
		// Metric factory
		{from: "kiloliter", to: "liter", exp: 1000},
		{from: "liter", to: "kiloliter", exp: 0.001},
		{from: "centiliter", to: "liter", exp: 0.01},
		{from: "liter", to: "centiliter", exp: 100},
		// Imperial/US
		{from: "gallon", to: "liter", exp: 4.54609},
		{from: "liter", to: "gallon", exp: 0.219969248},
		{from: "cubic foot", to: "liter", exp: 28.316846592},
		{from: "liter", to: "cubic foot", exp: 0.0353146667214886},
		// Cross-system
		{from: "cubic inch", to: "liter", exp: 0.016387064},
		{from: "liter", to: "cubic inch", exp: 61.0237440947323},
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
