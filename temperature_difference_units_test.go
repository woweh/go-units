package units

import (
	"testing"
)

func Test_TemperatureDifference_Conversions(t *testing.T) {
	conversionTests := []conversionTest{
		// Revit conversions (from/to internal, see RevitUnits.csv)
		{from: "delta K", to: "delta°C", exp: 1},
		{from: "delta°C", to: "delta K", exp: 1},
		{from: "delta K", to: "delta°F", exp: 1.8},
		{from: "delta°F", to: "delta K", exp: 0.555555555555556},
		{from: "delta K", to: "delta°R", exp: 1.8},
		{from: "delta°R", to: "delta K", exp: 0.555555555555556},
		{from: "delta°C", to: "delta°F", exp: 1.8},
		{from: "delta°F", to: "delta°C", exp: 0.555555555555556},
		{from: "delta°C", to: "delta°R", exp: 1.8},
		{from: "delta°R", to: "delta°C", exp: 0.555555555555556},
		{from: "delta°F", to: "delta°R", exp: 1},
		{from: "delta°R", to: "delta°F", exp: 1},
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
