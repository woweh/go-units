package units

import (
	"testing"

	"github.com/alecthomas/assert/v2"
)

func Test_MassPerTime_Conversions(t *testing.T) {
	conversionTests := []conversionTest{
		// Revit conversions from internal (kg/s)
		{"kg/s", "kg/h", 3600},
		{"kg/h", "kg/s", 0.000277777777777778},
		{"kg/s", "kg/min", 60},
		{"kg/min", "kg/s", 0.0166666666666667},
		{"kg/s", "lb/h", 7936.64143865559},
		{"lb/h", "kg/s", 0.000125997880555556},
		{"kg/s", "lb/min", 132.277357310927},
		{"lb/min", "kg/s", 0.00755987283333333},
		{"kg/s", "lb/s", 2.20462262184878},
		{"lb/s", "kg/s", 0.45359237},

		// Cross conversions (SI time units)
		{"kg/min", "kg/h", 60},
		{"kg/h", "kg/min", 0.0166666666666667},

		// Cross conversions (Imperial time units)
		{"lb/min", "lb/h", 60},
		{"lb/h", "lb/min", 0.0166666666666667},
		{"lb/s", "lb/min", 60},
		{"lb/min", "lb/s", 0.0166666666666667},
	}
	testConversions(t, conversionTests)
}

func Test_MassPerTime_UnitSystems(t *testing.T) {
	assert.Equal(t, SiSystem, KilogramPerSecond.System())
	assert.Equal(t, SiSystem, KilogramPerMinute.System())
	assert.Equal(t, SiSystem, KilogramPerHour.System())
	assert.Equal(t, BiSystem, PoundPerSecond.System())
	assert.Equal(t, BiSystem, PoundPerMinute.System())
	assert.Equal(t, BiSystem, PoundPerHour.System())
}

func Test_Lookup_MassPerTime_Names_and_Symbols(t *testing.T) {
	tests := lookUpTests{
		{KilogramPerSecond, "kilogram per second"},
		{KilogramPerSecond, "kg/s"},
		{KilogramPerSecond, "kilograms per second"},
		{KilogramPerSecond, "kg per second"},
		{KilogramPerMinute, "kilogram per minute"},
		{KilogramPerMinute, "kg/min"},
		{KilogramPerMinute, "kilograms per minute"},
		{KilogramPerMinute, "kg per minute"},
		{KilogramPerHour, "kilogram per hour"},
		{KilogramPerHour, "kg/h"},
		{KilogramPerHour, "kilograms per hour"},
		{KilogramPerHour, "kg per hour"},
		{PoundPerSecond, "pound per second"},
		{PoundPerSecond, "lb/s"},
		{PoundPerSecond, "pounds per second"},
		{PoundPerSecond, "pounds mass per second"},
		{PoundPerSecond, "lbm/s"},
		{PoundPerMinute, "pound per minute"},
		{PoundPerMinute, "lb/min"},
		{PoundPerMinute, "pounds per minute"},
		{PoundPerMinute, "pounds mass per minute"},
		{PoundPerMinute, "lbm/min"},
		{PoundPerHour, "pound per hour"},
		{PoundPerHour, "lb/h"},
		{PoundPerHour, "pounds per hour"},
		{PoundPerHour, "pounds mass per hour"},
		{PoundPerHour, "lbm/h"},
	}
	testLookupNamesAndSymbols(t, tests)
}
