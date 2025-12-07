package units

import (
	"testing"
)

func Test_PowerDensity_Conversions(t *testing.T) {
	conversionTests := []conversionTest{
		// Revit conversions from internal (W/m²)
		{from: "W/m²", to: "W/ft²", exp: 0.09290304},
		{from: "W/ft²", to: "W/m²", exp: 10.7639104167097},
		{from: "W/m²", to: "Btu/(h·ft²)", exp: 0.316998330628151, tol: fPtr(1e-6)},
		{from: "Btu/(h·ft²)", to: "W/m²", exp: 3.15459074506305, tol: fPtr(1e-5)},

		// Cross conversions
		{from: "W/ft²", to: "Btu/(h·ft²)", exp: 3.41214},
		{from: "Btu/(h·ft²)", to: "W/ft²", exp: 0.29307107017222, tol: fPtr(1e-6)},
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
