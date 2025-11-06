package units

import (
	"testing"

	"github.com/alecthomas/assert/v2"
)

func Test_SpecificHeat_Conversions(t *testing.T) {
	conversionTests := []conversionTest{
		// Library and Revit conversions (from internal and to internal)
		{"J/(g·°C)", "J/(kg·°C)", 1000},
		{"J/(kg·°C)", "J/(g·°C)", 0.001},
		{"BTU/(lb·°F)", "J/(kg·°C)", 4186.8},
		{"J/(kg·°C)", "BTU/(lb·°F)", 0.0002388459},
	}
	testConversions(t, conversionTests)
}

func Test_SpecificHeat_UnitSystem(t *testing.T) {
	assert.Equal(t, SiSystem, JoulePerKilogramCelsius.System())
	assert.Equal(t, SiSystem, JoulePerGramCelsius.System())
	assert.Equal(t, BiSystem, BritishThermalUnitPerPoundFahrenheit.System())
}

// No metric factories for specific heat, so no base unit tests are needed.

func Test_SpecificHeat_Lookup_Names_and_Symbols(t *testing.T) {
	tests := lookUpTests{
		// SI base
		{JoulePerKilogramCelsius, "J/(kg·°C)"},
		{JoulePerKilogramCelsius, "joule per kilogram degree Celsius"},
		{JoulePerKilogramCelsius, "joules per kilogram degree Celsius"},
		{JoulePerKilogramCelsius, "J/kg/C"},
		{JoulePerKilogramCelsius, "J/(kg*°C)"},
		// SI other
		{JoulePerGramCelsius, "J/(g·°C)"},
		{JoulePerGramCelsius, "joule per gram degree Celsius"},
		{JoulePerGramCelsius, "joules per gram degree Celsius"},
		{JoulePerGramCelsius, "J/g/C"},
		{JoulePerGramCelsius, "J/(g*°C)"},
		// Imperial
		{BritishThermalUnitPerPoundFahrenheit, "BTU/(lb·°F)"},
		{BritishThermalUnitPerPoundFahrenheit, "British thermal unit per pound degree Fahrenheit"},
		{BritishThermalUnitPerPoundFahrenheit, "British thermal units per pound degree Fahrenheit"},
		{BritishThermalUnitPerPoundFahrenheit, "BTU/lb/F"},
		{BritishThermalUnitPerPoundFahrenheit, "Btu/(lb·°F)"},
	}
	testLookupNamesAndSymbols(t, tests)
}
