package units

import (
	"testing"
)

func Test_MassPerTime_Conversions(t *testing.T) {
	conversionTests := []conversionTest{
		// Revit conversions from internal (kg/s)
		{from: "kg/s", to: "kg/h", exp: 3600},
		{from: "kg/h", to: "kg/s", exp: 0.000277777777777778},
		{from: "kg/s", to: "kg/min", exp: 60},
		{from: "kg/min", to: "kg/s", exp: 0.0166666666666667},
		{from: "kg/s", to: "lb/h", exp: 7936.64143865559},
		{from: "lb/h", to: "kg/s", exp: 0.000125997880555556},
		{from: "kg/s", to: "lb/min", exp: 132.277357310927},
		{from: "lb/min", to: "kg/s", exp: 0.00755987283333333},
		{from: "kg/s", to: "lb/s", exp: 2.20462262184878},
		{from: "lb/s", to: "kg/s", exp: 0.45359237},

		// Cross conversions (SI time units)
		{from: "kg/min", to: "kg/h", exp: 60},
		{from: "kg/h", to: "kg/min", exp: 0.0166666666666667},

		// Cross conversions (Imperial time units)
		{from: "lb/min", to: "lb/h", exp: 60},
		{from: "lb/h", to: "lb/min", exp: 0.0166666666666667},
		{from: "lb/s", to: "lb/min", exp: 60},
		{from: "lb/min", to: "lb/s", exp: 0.0166666666666667},
	}
	testConversions(t, conversionTests)
}

func Test_MassPerTime_UnitSystems(t *testing.T) {
	tests := []unitSystemTest{
		{KilogramPerSecond, SiSystem},
		{KilogramPerMinute, SiSystem},
		{KilogramPerHour, SiSystem},
		{PoundPerSecond, BiSystem},
		{PoundPerMinute, BiSystem},
		{PoundPerHour, BiSystem},
	}
	testUnitSystems(t, tests)
}
