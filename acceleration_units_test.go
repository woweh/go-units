package units

import (
	"testing"
)

func Test_Acceleration_Conversions(t *testing.T) {
	conversionTests := []conversionTest{
		// SI base to Imperial/US (Revit and internal conversion factors)
		{from: "m/s²", to: "ft/s²", exp: 3.2808398950131},
		{from: "ft/s²", to: "m/s²", exp: 0.3048},
		{from: "m/s²", to: "in/s²", exp: 39.370078740157},
		{from: "in/s²", to: "m/s²", exp: 0.0254},
		{from: "m/s²", to: "km/s²", exp: 0.001},
		{from: "km/s²", to: "m/s²", exp: 1000},
		{from: "m/s²", to: "mi/s²", exp: 0.00062137119223733},
		{from: "mi/s²", to: "m/s²", exp: 1609.344},

		// Revit CSV: ft/s² <-> in/s²
		{from: "ft/s²", to: "in/s²", exp: 12},
		{from: "in/s²", to: "ft/s²", exp: 0.0833333333333333},

		// Revit CSV: ft/s² <-> km/s²
		{from: "ft/s²", to: "km/s²", exp: 0.0003048},
		{from: "km/s²", to: "ft/s²", exp: 3280.83989501312},

		// Revit CSV: ft/s² <-> mi/s²
		{from: "ft/s²", to: "mi/s²", exp: 0.000189393939393939},
		{from: "mi/s²", to: "ft/s²", exp: 5280},
	}
	testConversions(t, conversionTests)
}

func Test_Acceleration_UnitSystems(t *testing.T) {
	tests := []unitSystemTest{
		{MeterPerSecondSquared, SiSystem},
		{KilometerPerSecondSquared, SiSystem},
		{FootPerSecondSquared, BiSystem},
		{InchPerSecondSquared, BiSystem},
		{MilePerSecondSquared, BiSystem},
	}
	testUnitSystems(t, tests)
}
