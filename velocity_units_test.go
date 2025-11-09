package units

import (
	"testing"
)

func Test_Velocity_Conversions(t *testing.T) {
	conversionTests := []conversionTest{
		// Revit conversions (from/to internal, see RevitUnits.csv)
		{"m/s", "ft/s", 3.28083989501312},
		{"ft/s", "m/s", 0.3048},
		{"m/s", "ft/min", 196.850394},
		{"ft/min", "m/s", 0.00508},
		{"m/s", "cm/min", 6000},
		{"cm/min", "m/s", 0.000166666666666667},
		{"m/s", "km/h", 3.6},
		{"km/h", "m/s", 0.277777777777778},
		{"m/s", "mph", 2.2369362920544},
		{"mph", "m/s", 0.44704},
		// Cross-system
		{"mph", "km/h", 1.609344},
		{"km/h", "mph", 0.621371192237334},
	}
	testConversions(t, conversionTests)
}

func Test_Velocity_UnitSystems(t *testing.T) {
	tests := []unitSystemTest{
		{MeterPerSecond, SiSystem},
		{CentimeterPerMinute, SiSystem},
		{KilometerPerHour, SiSystem},
		{FootPerSecond, BiSystem},
		{FootPerMinute, BiSystem},
		{MilePerHour, BiSystem},
	}
	testUnitSystems(t, tests)
}
