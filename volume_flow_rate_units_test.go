package units

import (
	"testing"

	"github.com/alecthomas/assert/v2"
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

func Test_VolumeFlowRate_UnitSystem(t *testing.T) {
	assert.Equal(t, SiSystem, CubicMeterPerSecond.System())
	assert.Equal(t, SiSystem, CubicMeterPerMinute.System())
	assert.Equal(t, SiSystem, CubicMeterPerHour.System())
	assert.Equal(t, SiSystem, CubicMeterPerDay.System())
	assert.Equal(t, SiSystem, CubicDecimeterPerSecond.System())
	assert.Equal(t, SiSystem, CubicDecimeterPerMinute.System())
	assert.Equal(t, SiSystem, CubicDecimeterPerHour.System())
	assert.Equal(t, SiSystem, CubicDecimeterPerDay.System())
	assert.Equal(t, SiSystem, CubicCentimeterPerSecond.System())
	assert.Equal(t, SiSystem, CubicCentimeterPerMinute.System())
	assert.Equal(t, SiSystem, CubicCentimeterPerHour.System())
	assert.Equal(t, SiSystem, CubicCentimeterPerDay.System())
	assert.Equal(t, BiSystem, CubicInchPerSecond.System())
	assert.Equal(t, BiSystem, CubicInchPerMinute.System())
	assert.Equal(t, BiSystem, CubicInchPerHour.System())
	assert.Equal(t, BiSystem, CubicInchPerDay.System())
	assert.Equal(t, BiSystem, CubicFootPerSecond.System())
	assert.Equal(t, BiSystem, CubicFootPerMinute.System())
	assert.Equal(t, BiSystem, CubicFootPerHour.System())
	assert.Equal(t, BiSystem, CubicFootPerDay.System())
	assert.Equal(t, BiSystem, CubicYardPerSecond.System())
	assert.Equal(t, BiSystem, CubicYardPerMinute.System())
	assert.Equal(t, BiSystem, CubicYardPerHour.System())
	assert.Equal(t, BiSystem, CubicYardPerDay.System())
}

func Test_VolumeFlowRate_MetricFactory_BaseUnits(t *testing.T) {
	assert.Equal(t, CubicMeterPerSecond, CubicCentimeterPerSecond.Base())
	assert.Equal(t, CubicMeterPerSecond, CubicCentimeterPerMinute.Base())
	assert.Equal(t, CubicMeterPerSecond, CubicCentimeterPerHour.Base())
	assert.Equal(t, CubicMeterPerSecond, CubicCentimeterPerDay.Base())
}

func Test_VolumeFlowRate_Lookup_Names_and_Symbols(t *testing.T) {
	tests := lookUpTests{
		// SI
		{CubicMeterPerSecond, "m³/s"},
		{CubicMeterPerSecond, "cubic meter per second"},
		{CubicMeterPerSecond, "cubic metre per second"},
		{CubicMeterPerSecond, "m3/s"},
		{CubicMeterPerSecond, "m3s-1"},
		{CubicMeterPerMinute, "m³/min"},
		{CubicMeterPerMinute, "cubic meter per minute"},
		{CubicMeterPerMinute, "cubic metre per minute"},
		{CubicMeterPerMinute, "m3/min"},
		{CubicMeterPerMinute, "m3m-1"},
		{CubicMeterPerHour, "m³/h"},
		{CubicMeterPerHour, "cubic meter per hour"},
		{CubicMeterPerHour, "cubic metre per hour"},
		{CubicMeterPerHour, "m3/h"},
		{CubicMeterPerHour, "m3h-1"},
		{CubicMeterPerDay, "m³/d"},
		{CubicMeterPerDay, "cubic meter per day"},
		{CubicMeterPerDay, "cubic metre per day"},
		{CubicMeterPerDay, "m3/d"},
		{CubicMeterPerDay, "m3d-1"},
		{CubicDecimeterPerSecond, "dm³/s"},
		{CubicDecimeterPerSecond, "cubic decimeter per second"},
		{CubicDecimeterPerSecond, "cubic decimetre per second"},
		{CubicDecimeterPerSecond, "dm3/s"},
		{CubicDecimeterPerSecond, "dm3s-1"},
		{CubicCentimeterPerSecond, "cm³/s"},
		{CubicCentimeterPerSecond, "cubic centimeter per second"},
		{CubicCentimeterPerSecond, "cubic centimetre per second"},
		{CubicCentimeterPerSecond, "cm3/s"},
		{CubicCentimeterPerSecond, "cm3s-1"},
		// Imperial/US
		{CubicInchPerSecond, "in³/s"},
		{CubicInchPerSecond, "cubic inch per second"},
		{CubicFootPerSecond, "ft³/s"},
		{CubicFootPerSecond, "cubic foot per second"},
		{CubicYardPerSecond, "yd³/s"},
		{CubicYardPerSecond, "cubic yard per second"},
	}
	testLookupNamesAndSymbols(t, tests)
}
