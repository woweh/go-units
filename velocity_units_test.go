package units

import (
	"testing"
)

func Test_Velocity_Conversions(t *testing.T) {
	conversionTests := []conversionTest{
		// Revit conversions (from/to internal, see RevitUnits.csv)
		{from: "m/s", to: "ft/s", exp: 3.28083989501312},
		{from: "ft/s", to: "m/s", exp: 0.3048},
		{from: "m/s", to: "ft/min", exp: 196.850394},
		{from: "ft/min", to: "m/s", exp: 0.00508},
		{from: "m/s", to: "cm/min", exp: 6000},
		{from: "cm/min", to: "m/s", exp: 0.000166666666666667},
		{from: "m/s", to: "km/h", exp: 3.6},
		{from: "km/h", to: "m/s", exp: 0.277777777777778},
		{from: "m/s", to: "mph", exp: 2.2369362920544},
		{from: "mph", to: "m/s", exp: 0.44704},
		// Cross-system
		{from: "mph", to: "km/h", exp: 1.609344},
		{from: "km/h", to: "mph", exp: 0.621371192237334},
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
