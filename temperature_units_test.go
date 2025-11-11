package units

import (
	"testing"
)

func Test_Temperature_Conversions(t *testing.T) {
	conversionTests := []conversionTest{
		// Revit conversions (from/to internal, see RevitUnits.csv)
		{from: "celsius", to: "kelvin", exp: 274.15},
		{from: "kelvin", to: "celsius", exp: -272.15},
		{from: "fahrenheit", to: "kelvin", exp: 255.927777777778},
		{from: "kelvin", to: "fahrenheit", exp: -457.87},
		{from: "rankine", to: "kelvin", exp: 0.555555555555556},
		{from: "kelvin", to: "rankine", exp: 1.8},
		// Cross-system
		{from: "celsius", to: "fahrenheit", exp: 33.8},
		{from: "fahrenheit", to: "celsius", exp: -17.222222},
		{from: "fahrenheit", to: "rankine", exp: 460.67},
		{from: "rankine", to: "fahrenheit", exp: -458.67},
	}
	testConversions(t, conversionTests)
}

func Test_Temperature_UnitSystems(t *testing.T) {
	tests := []unitSystemTest{
		{Celsius, SiSystem},
		{Kelvin, SiSystem},
		{Fahrenheit, UsSystem},
		{Rankine, UsSystem},
	}
	testUnitSystems(t, tests)
}
