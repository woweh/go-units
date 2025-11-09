package units

import (
	"testing"
)

func Test_VolumeFlowRate_Conversions(t *testing.T) {
	conversionTests := []conversionTest{
		// SI conversions
		{"cubic meter per second", "cubic meter per minute", 60},
		{"cubic meter per second", "cubic meter per hour", 3600},
		{"cubic meter per second", "cubic meter per day", 86400},
		{"cubic meter per second", "cubic decimeter per second", 1000},
		{"cubic meter per second", "cubic centimeter per second", 1000000},
		// Imperial conversions
		{"cubic meter per second", "cubic inch per second", 61023.744095},
		{"cubic meter per second", "cubic foot per second", 35.314667},
		{"cubic meter per second", "cubic yard per second", 1.307951},
		// Cross-system
		{"cubic foot per second", "cubic meter per second", 0.0283168},
		{"cubic inch per second", "cubic meter per second", 0.0000163871},
	}
	testConversions(t, conversionTests)
}

func Test_VolumeFlowRate_UnitSystems(t *testing.T) {
	tests := []unitSystemTest{
		{CubicMeterPerSecond, SiSystem},
		{CubicMeterPerMinute, SiSystem},
		{CubicMeterPerHour, SiSystem},
		{CubicMeterPerDay, SiSystem},
		{CubicDecimeterPerSecond, SiSystem},
		{CubicDecimeterPerMinute, SiSystem},
		{CubicDecimeterPerHour, SiSystem},
		{CubicDecimeterPerDay, SiSystem},
		{CubicCentimeterPerSecond, SiSystem},
		{CubicCentimeterPerMinute, SiSystem},
		{CubicCentimeterPerHour, SiSystem},
		{CubicCentimeterPerDay, SiSystem},
		{CubicInchPerSecond, BiSystem},
		{CubicInchPerMinute, BiSystem},
		{CubicInchPerHour, BiSystem},
		{CubicInchPerDay, BiSystem},
		{CubicFootPerSecond, BiSystem},
		{CubicFootPerMinute, BiSystem},
		{CubicFootPerHour, BiSystem},
		{CubicFootPerDay, BiSystem},
		{CubicYardPerSecond, BiSystem},
		{CubicYardPerMinute, BiSystem},
		{CubicYardPerHour, BiSystem},
		{CubicYardPerDay, BiSystem},
	}
	testUnitSystems(t, tests)
}
