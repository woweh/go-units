package units

import (
	"testing"
)

func Test_Time_Conversions(t *testing.T) {
	conversionTests := []conversionTest{
		// SI/metric
		{"second", "millisecond", 1000},
		{"millisecond", "second", 0.001},
		{"second", "minute", 0.01666667},
		{"minute", "second", 60},
		{"hour", "second", 3600},
		{"second", "hour", 0.0002777778},
		// Metric factory
		{"kilosecond", "second", 1000},
		{"second", "kilosecond", 0.001},
		{"centisecond", "second", 0.01},
		{"second", "centisecond", 100},
		// Calendar
		{"day", "hour", 24},
		{"hour", "day", 0.04166667},
		{"year", "day", 365.25},
		{"day", "year", 0.002737851},
		// Esoteric
		{"fortnight", "day", 14},
		{"planck time", "second", 5.391247e-44},
	}
	testConversions(t, conversionTests)
}

func Test_Time_UnitSystems(t *testing.T) {
	tests := []unitSystemTest{
		{Second, SiSystem},
		{KiloSecond, SiSystem},
		{CentiSecond, SiSystem},
		{MilliSecond, SiSystem},
		{MicroSecond, SiSystem},
		{NanoSecond, SiSystem},
		{PicoSecond, SiSystem},
		{FemtoSecond, SiSystem},
		{AttoSecond, SiSystem},
		{Minute, NoSystem},
		{Hour, NoSystem},
		{Day, NoSystem},
		{Year, NoSystem},
		{Fortnight, NoSystem},
		{PlanckTime, NoSystem},
	}
	testUnitSystems(t, tests)
}
