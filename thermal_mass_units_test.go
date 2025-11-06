package units

import (
	"testing"

	"github.com/alecthomas/assert/v2"
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

func Test_ThermalMass_UnitSystem(t *testing.T) {
	assert.Equal(t, SiSystem, JoulePerKelvin.System())
	assert.Equal(t, SiSystem, KiloJoulePerKelvin.System())
	assert.Equal(t, BiSystem, BritishThermalUnitPerFahrenheit.System())
}

// No metric factories for thermal mass, so no base unit tests are needed.

func Test_ThermalMass_Lookup_Names_and_Symbols(t *testing.T) {
	tests := lookUpTests{
		// SI
		{JoulePerKelvin, "J/K"},
		{JoulePerKelvin, "joule per kelvin"},
		{JoulePerKelvin, "joules per kelvin"},
		{JoulePerKelvin, "J per K"},
		{KiloJoulePerKelvin, "kJ/K"},
		{KiloJoulePerKelvin, "kilojoule per kelvin"},
		{KiloJoulePerKelvin, "kilojoules per kelvin"},
		{KiloJoulePerKelvin, "kJ per K"},
		// Imperial/US
		{BritishThermalUnitPerFahrenheit, "BTU/°F"},
		{BritishThermalUnitPerFahrenheit, "British thermal unit per degree Fahrenheit"},
		{BritishThermalUnitPerFahrenheit, "British thermal units per degree Fahrenheit"},
		{BritishThermalUnitPerFahrenheit, "BTU per F"},
		{BritishThermalUnitPerFahrenheit, "Btu/°F"},
	}
	testLookupNamesAndSymbols(t, tests)
}
