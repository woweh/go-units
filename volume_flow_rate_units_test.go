package units

import (
	"testing"
)

func Test_VolumeFlowRate_Conversions(t *testing.T) {
	conversionTests := []conversionTest{
		// SI conversions
		{from: "cubic meter per second", to: "cubic meter per minute", exp: 60},
		{from: "cubic meter per second", to: "cubic meter per hour", exp: 3600},
		{from: "cubic meter per second", to: "cubic meter per day", exp: 86400},
		{from: "cubic meter per second", to: "cubic decimeter per second", exp: 1000},
		{from: "cubic meter per second", to: "cubic centimeter per second", exp: 1000000},
		// Imperial conversions
		{from: "cubic meter per second", to: "cubic inch per second", exp: 61023.744094732},
		{from: "cubic meter per second", to: "cubic foot per second", exp: 35.314666721},
		{from: "cubic meter per second", to: "cubic yard per second", exp: 1.307950619},
		// Cross-system
		{from: "cubic foot per second", to: "cubic meter per second", exp: 0.028316847},
		{from: "cubic inch per second", to: "cubic meter per second", exp: 0.000016387064},
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
