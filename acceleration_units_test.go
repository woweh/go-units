package units

import (
	"testing"

	"github.com/alecthomas/assert/v2"
)

func Test_Acceleration_Conversions(t *testing.T) {
	conversionTests := []conversionTest{
		// SI base to Imperial/US (Revit conversion factors)
		{"m/s²", "ft/s²", "3.2808398950131"},
		{"m/s²", "in/s²", "39.370078740157"},
		{"m/s²", "km/s²", "0.001"},
		{"m/s²", "mi/s²", "0.00062137119223733"},
		// Reverse
		{"ft/s²", "m/s²", "0.3048"},
		{"in/s²", "m/s²", "0.0254"},
		{"km/s²", "m/s²", "1000"},
		// Identity
		{"m/s²", "m/s²", "1"},
		{"ft/s²", "ft/s²", "1"},
	}
	testConversions(t, conversionTests)
}

func Test_Acceleration_UnitSystems(t *testing.T) {
	assert.Equal(t, SiSystem, MeterPerSecondSquared.System())
	assert.Equal(t, SiSystem, KilometerPerSecondSquared.System())
	assert.Equal(t, BiSystem, FootPerSecondSquared.System())
	assert.Equal(t, BiSystem, InchPerSecondSquared.System())
	assert.Equal(t, BiSystem, MilePerSecondSquared.System())
}

func Test_Acceleration_BaseUnits(t *testing.T) {
	assert.Equal(t, MeterPerSecondSquared, MeterPerSecondSquared.Base())
	assert.Equal(t, MeterPerSecondSquared, KilometerPerSecondSquared.Base())
	// Imperial units should not be their own base
	assert.NotEqual(t, FootPerSecondSquared, FootPerSecondSquared.Base())
	assert.NotEqual(t, InchPerSecondSquared, InchPerSecondSquared.Base())
	assert.NotEqual(t, MilePerSecondSquared, MilePerSecondSquared.Base())
}

func Test_Lookup_Acceleration_Names_and_Symbols(t *testing.T) {
	tests := lookUpTests{
		{MeterPerSecondSquared, "meter per second squared"},
		{MeterPerSecondSquared, "m/s²"},
		{MeterPerSecondSquared, "meters per second squared"},
		{MeterPerSecondSquared, "metre per second squared"},
		{MeterPerSecondSquared, "m/sec²"},
		{KilometerPerSecondSquared, "kilometer per second squared"},
		{KilometerPerSecondSquared, "km/s²"},
		{FootPerSecondSquared, "foot per second squared"},
		{FootPerSecondSquared, "ft/s²"},
		{InchPerSecondSquared, "inch per second squared"},
		{InchPerSecondSquared, "in/s²"},
		{MilePerSecondSquared, "mile per second squared"},
		{MilePerSecondSquared, "mi/s²"},
	}
	testLookupNamesAndSymbols(t, tests)
}
