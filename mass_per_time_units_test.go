package units

import (
	"testing"
)

func Test_MassPerTime_Conversions(t *testing.T) {
	var conversionTests = []conversionTest{
		{"kg/s", "kg/min", "60"},
		{"kg/s", "kg/h", "3600"},
		{"kg/s", "lb/s", "2.204623"},
		{"kg/s", "kg/s", "1"},
	}

	testConversions(t, conversionTests)
}
