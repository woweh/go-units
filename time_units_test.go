package units

import (
	"testing"

	"github.com/alecthomas/assert/v2"
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

func Test_Time_UnitSystem(t *testing.T) {
	assert.Equal(t, SiSystem, Second.System())
	assert.Equal(t, SiSystem, KiloSecond.System())
	assert.Equal(t, SiSystem, CentiSecond.System())
	assert.Equal(t, SiSystem, MilliSecond.System())
	assert.Equal(t, SiSystem, MicroSecond.System())
	assert.Equal(t, SiSystem, NanoSecond.System())
	assert.Equal(t, SiSystem, PicoSecond.System())
	assert.Equal(t, SiSystem, FemtoSecond.System())
	assert.Equal(t, SiSystem, AttoSecond.System())
}

func Test_Time_MetricFactory_BaseUnits(t *testing.T) {
	assert.Equal(t, Second, CentiSecond.Base())
	assert.Equal(t, Second, KiloSecond.Base())
}

func Test_Time_Lookup_Names_and_Symbols(t *testing.T) {
	tests := lookUpTests{
		// SI/metric
		{Second, "s"},
		{Second, "second"},
		{KiloSecond, "ks"},
		{CentiSecond, "cs"},
		{MilliSecond, "ms"},
		{MicroSecond, "μs"},
		{NanoSecond, "ns"},
		{PicoSecond, "ps"},
		{FemtoSecond, "fs"},
		{AttoSecond, "as"},
		// Calendar
		{Minute, "min"},
		{Hour, "hr"},
		{Day, "d"},
		{Year, "yr"},
		// Esoteric
		{PlanckTime, "𝑡ₚ"},
		{Fortnight, "fortnight"},
		{Score, "score"},
	}
	testLookupNamesAndSymbols(t, tests)
}
