package units

import (
	"testing"
)

func Test_PowerDensity_Conversions(t *testing.T) {
	conversionTests := []conversionTest{
		// Revit conversions from internal (W/m²)
		{"W/m²", "W/ft²", 0.09290304},
		{"W/ft²", "W/m²", 10.7639104167097},
		{"W/m²", "Btu/(h·ft²)", 0.316998330628151},
		{"Btu/(h·ft²)", "W/m²", 3.15459074506305},

		// Cross conversions
		{"W/ft²", "Btu/(h·ft²)", 3.41214},
		{"Btu/(h·ft²)", "W/ft²", 0.29307107017222},
	}
	testConversions(t, conversionTests)
}

func Test_PowerDensity_UnitSystems(t *testing.T) {
	tests := []unitSystemTest{
		{WattPerSquareMeter, SiSystem},
		{WattPerSquareFoot, BiSystem},
		{BritishThermalUnitPerHourSquareFoot, BiSystem},
	}
	testUnitSystems(t, tests)
}
