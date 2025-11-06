package units

import (
	"testing"

	"github.com/alecthomas/assert/v2"
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

func Test_Velocity_UnitSystem(t *testing.T) {
	assert.Equal(t, SiSystem, MeterPerSecond.System())
	assert.Equal(t, SiSystem, CentimeterPerMinute.System())
	assert.Equal(t, SiSystem, KilometerPerHour.System())
	assert.Equal(t, BiSystem, FootPerSecond.System())
	assert.Equal(t, BiSystem, FootPerMinute.System())
	assert.Equal(t, BiSystem, MilePerHour.System())
}

func Test_Velocity_MetricFactory_BaseUnits(t *testing.T) {
	// CentimeterPerMinute is a metric factory unit
	assert.Equal(t, MeterPerSecond, CentimeterPerMinute.Base())
}

func Test_Velocity_Lookup_Names_and_Symbols(t *testing.T) {
	tests := lookUpTests{
		// SI
		{MeterPerSecond, "m/s"},
		{MeterPerSecond, "meters per second"},
		{MeterPerSecond, "metre per second"},
		{MeterPerSecond, "metres per second"},
		{MeterPerSecond, "mps"},
		{CentimeterPerMinute, "cm/min"},
		{CentimeterPerMinute, "centimeter per minute"},
		{CentimeterPerMinute, "centimeters per minute"},
		{CentimeterPerMinute, "centimetre per minute"},
		{CentimeterPerMinute, "centimetres per minute"},
		{KilometerPerHour, "km/h"},
		{KilometerPerHour, "kilometer per hour"},
		{KilometerPerHour, "kilometers per hour"},
		{KilometerPerHour, "kilometre per hour"},
		{KilometerPerHour, "kilometres per hour"},
		{KilometerPerHour, "kph"},
		// Imperial/US
		{FootPerSecond, "ft/s"},
		{FootPerSecond, "foot per second"},
		{FootPerSecond, "feet per second"},
		{FootPerSecond, "fps"},
		{FootPerMinute, "ft/min"},
		{FootPerMinute, "foot per minute"},
		{FootPerMinute, "feet per minute"},
		{FootPerMinute, "fpm"},
		{MilePerHour, "mph"},
		{MilePerHour, "mile per hour"},
		{MilePerHour, "miles per hour"},
	}
	testLookupNamesAndSymbols(t, tests)
}
