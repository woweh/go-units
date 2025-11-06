package units

import (
	"testing"

	"github.com/alecthomas/assert/v2"
)

func Test_TemperatureDifference_Conversions(t *testing.T) {
	conversionTests := []conversionTest{
		// Revit conversions (from/to internal, see RevitUnits.csv)
		{"delta K", "delta°C", 1},
		{"delta°C", "delta K", 1},
		{"delta K", "delta°F", 1.8},
		{"delta°F", "delta K", 0.555555555555556},
		{"delta K", "delta°R", 1.8},
		{"delta°R", "delta K", 0.555555555555556},
		{"delta°C", "delta°F", 1.8},
		{"delta°F", "delta°C", 0.555555555555556},
		{"delta°C", "delta°R", 1.8},
		{"delta°R", "delta°C", 0.555555555555556},
		{"delta°F", "delta°R", 1},
		{"delta°R", "delta°F", 1},
	}
	testConversions(t, conversionTests)
}

func Test_TemperatureDifference_UnitSystem(t *testing.T) {
	assert.Equal(t, SiSystem, KelvinInterval.System())
	assert.Equal(t, SiSystem, CelsiusInterval.System())
	assert.Equal(t, BiSystem, FahrenheitInterval.System())
	assert.Equal(t, BiSystem, RankineInterval.System())
}

// No metric factories for temperature difference, so no base unit tests are needed.

func Test_TemperatureDifference_Lookup_Names_and_Symbols(t *testing.T) {
	tests := lookUpTests{
		// SI base
		{KelvinInterval, "delta K"},
		{KelvinInterval, "kelvin interval"},
		{KelvinInterval, "kelvin"},
		{KelvinInterval, "K interval"},
		{KelvinInterval, "delta K"},
		// SI other
		{CelsiusInterval, "delta°C"},
		{CelsiusInterval, "celsius interval"},
		{CelsiusInterval, "Celsius interval"},
		{CelsiusInterval, "°C interval"},
		{CelsiusInterval, "C interval"},
		{CelsiusInterval, "delta C"},
		{CelsiusInterval, "deltaC"},
		// Imperial
		{FahrenheitInterval, "delta°F"},
		{FahrenheitInterval, "fahrenheit interval"},
		{FahrenheitInterval, "Fahrenheit interval"},
		{FahrenheitInterval, "°F interval"},
		{FahrenheitInterval, "F interval"},
		{FahrenheitInterval, "delta F"},
		{FahrenheitInterval, "deltaF"},
		{RankineInterval, "delta°R"},
		{RankineInterval, "rankine interval"},
		{RankineInterval, "Rankine interval"},
		{RankineInterval, "°R interval"},
		{RankineInterval, "R interval"},
		{RankineInterval, "delta R"},
		{RankineInterval, "deltaR"},
	}
	testLookupNamesAndSymbols(t, tests)
}
