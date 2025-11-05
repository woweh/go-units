package units

import (
	"testing"

	"github.com/alecthomas/assert/v2"
)

func Test_Acceleration_Conversions(t *testing.T) {
	conversionTests := []conversionTest{
		// SI base to Imperial/US (Revit and internal conversion factors)
		{"m/s²", "ft/s²", 3.2808398950131},
		{"ft/s²", "m/s²", 0.3048},
		{"m/s²", "in/s²", 39.370078740157},
		{"in/s²", "m/s²", 0.0254},
		{"m/s²", "km/s²", 0.001},
		{"km/s²", "m/s²", 1000},
		{"m/s²", "mi/s²", 0.00062137119223733},
		{"mi/s²", "m/s²", 1609.344},

		// Revit CSV: ft/s² <-> in/s²
		{"ft/s²", "in/s²", 12},
		{"in/s²", "ft/s²", 0.0833333333333333},

		// Revit CSV: ft/s² <-> km/s²
		{"ft/s²", "km/s²", 0.0003048},
		{"km/s²", "ft/s²", 3280.83989501312},

		// Revit CSV: ft/s² <-> mi/s²
		{"ft/s²", "mi/s²", 0.000189393939393939},
		{"mi/s²", "ft/s²", 5280},
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
