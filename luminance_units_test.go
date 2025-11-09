package units

import (
	"testing"
)

func Test_Luminance_Conversions(t *testing.T) {
	conversionTests := []conversionTest{
		// Revit conversions from internal (cd/ft²)
		{"cd/ft²", "cd/m²", 10.7639104167097},
		{"cd/m²", "cd/ft²", 0.09290304},
		{"cd/ft²", "ftL", 3.14159265358979},
		{"ftL", "cd/ft²", 0.318309886183791},

		// Cross conversions
		{"cd/m²", "ftL", 0.29186351},
		{"ftL", "cd/m²", 3.426259},
	}
	testConversions(t, conversionTests)
}

func Test_Luminance_UnitSystems(t *testing.T) {
	tests := []unitSystemTest{
		{CandelaPerSquareMeter, SiSystem},
		{CandelaPerSquareFoot, BiSystem},
		{Footlambert, BiSystem},
	}
	testUnitSystems(t, tests)
}
