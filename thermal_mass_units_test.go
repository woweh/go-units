package units

import (
	"testing"
)

func Test_ThermalMass_Conversions(t *testing.T) {
	conversionTests := []conversionTest{
		// Revit conversions (from/to internal, see RevitUnits.csv)
		{from: "kJ/K", to: "J/K", exp: 1000},
		{from: "J/K", to: "kJ/K", exp: 0.001},
		{from: "BTU/°F", to: "J/K", exp: 1899.1},
		{from: "J/K", to: "BTU/°F", exp: 0.000526565},
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
