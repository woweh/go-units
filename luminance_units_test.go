package units

import (
	"testing"
)

func Test_Luminance_Conversions(t *testing.T) {
	conversionTests := []conversionTest{
		// Revit conversions from internal (cd/ft²)
		{from: "cd/ft²", to: "cd/m²", exp: 10.7639104167097},
		{from: "cd/m²", to: "cd/ft²", exp: 0.09290304},
		{from: "cd/ft²", to: "ftL", exp: 3.14159265358979, tol: fPtr(1e-7)},
		{from: "ftL", to: "cd/ft²", exp: 0.318309886183791},

		// Cross conversions
		{from: "cd/m²", to: "ftL", exp: 0.29186351},
		{from: "ftL", to: "cd/m²", exp: 3.426259, tol: fPtr(1e-6)},
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
