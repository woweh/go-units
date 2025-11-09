package units

import (
	"testing"
)

func Test_ThermalMass_Conversions(t *testing.T) {
	conversionTests := []conversionTest{
		// Revit conversions (from/to internal, see RevitUnits.csv)
		{"kJ/K", "J/K", 1000},
		{"J/K", "kJ/K", 0.001},
		{"BTU/°F", "J/K", 1899.1},
		{"J/K", "BTU/°F", 0.000526565},
	}
	testConversions(t, conversionTests)
}

func Test_ThermalMass_UnitSystems(t *testing.T) {
	tests := []unitSystemTest{
		{JoulePerKelvin, SiSystem},
		{KiloJoulePerKelvin, SiSystem},
		{BritishThermalUnitPerFahrenheit, BiSystem},
	}
	testUnitSystems(t, tests)
}
