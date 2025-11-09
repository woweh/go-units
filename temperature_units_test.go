package units

import (
	"testing"
)

func Test_Temperature_Conversions(t *testing.T) {
	conversionTests := []conversionTest{
		// Revit conversions (from/to internal, see RevitUnits.csv)
		{"celsius", "kelvin", 274.15},
		{"kelvin", "celsius", -272.15},
		{"fahrenheit", "kelvin", 255.927777777778},
		{"kelvin", "fahrenheit", -457.87},
		{"rankine", "kelvin", 0.555555555555556},
		{"kelvin", "rankine", 1.8},
		// Cross-system
		{"celsius", "fahrenheit", 33.8},
		{"fahrenheit", "celsius", -17.222222},
		{"fahrenheit", "rankine", 460.67},
		{"rankine", "fahrenheit", -458.67},
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
