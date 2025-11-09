package units

import (
	"testing"
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

func Test_TemperatureDifference_UnitSystems(t *testing.T) {
	tests := []unitSystemTest{
		{KelvinInterval, SiSystem},
		{CelsiusInterval, SiSystem},
		{FahrenheitInterval, BiSystem},
		{RankineInterval, BiSystem},
	}
	testUnitSystems(t, tests)
}
