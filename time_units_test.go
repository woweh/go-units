package units

import (
	"testing"
)

func Test_Time_Conversions(t *testing.T) {
	conversionTests := []conversionTest{
		// SI/metric
		{from: "second", to: "millisecond", exp: 1000},
		{from: "millisecond", to: "second", exp: 0.001},
		{from: "second", to: "minute", exp: 0.01666667},
		{from: "minute", to: "second", exp: 60},
		{from: "hour", to: "second", exp: 3600},
		{from: "second", to: "hour", exp: 0.0002777778},
		// Metric factory
		{from: "kilosecond", to: "second", exp: 1000},
		{from: "second", to: "kilosecond", exp: 0.001},
		{from: "centisecond", to: "second", exp: 0.01},
		{from: "second", to: "centisecond", exp: 100},
		// Calendar
		{from: "day", to: "hour", exp: 24},
		{from: "hour", to: "day", exp: 0.04166667},
		{from: "year", to: "day", exp: 365.25},
		{from: "day", to: "year", exp: 0.002737851},
		// Esoteric
		{from: "fortnight", to: "day", exp: 14},
		{from: "planck time", to: "second", exp: 5.391247e-44},
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
