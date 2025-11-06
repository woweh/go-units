package units

import (
	"testing"

	"github.com/alecthomas/assert/v2"
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

func Test_Temperature_UnitSystem(t *testing.T) {
	assert.Equal(t, SiSystem, Celsius.System())
	assert.Equal(t, SiSystem, Kelvin.System())
	assert.Equal(t, UsSystem, Fahrenheit.System())
	assert.Equal(t, UsSystem, Rankine.System())
}

// No metric factories for temperature, so no base unit tests are needed.

func Test_Temperature_Lookup_Names_and_Symbols(t *testing.T) {
	tests := lookUpTests{
		{Celsius, "celsius"},
		{Celsius, "centigrade"},
		{Celsius, "centigrades"},
		{Celsius, "°C"},
		{Fahrenheit, "fahrenheit"},
		{Fahrenheit, "degree Fahrenheit"},
		{Fahrenheit, "°F"},
		{Kelvin, "kelvin"},
		{Kelvin, "degree Kelvin"},
		{Kelvin, "K"},
		{Rankine, "rankine"},
		{Rankine, "degree Rankine"},
		{Rankine, "°R"},
		{Rankine, "°Ra"},
	}
	testLookupNamesAndSymbols(t, tests)
}
