package units

import (
	"testing"

	"github.com/alecthomas/assert/v2"
)

func Test_Illuminance_Conversions(t *testing.T) {
	var conversionTests = []conversionTest{
		// Revit conversion factors (Footcandle is Revit internal unit)
		// 1 Ftc -> 10.7639104167097 lx (from Revit: CF from internal)
		{"Ftc", "lx", 10.7639104167097},
		// 1 lx -> 0.09290304 Ftc (from Revit: CF to internal)
		{"lx", "Ftc", 0.09290304},
	}

	testConversions(t, conversionTests)
}

func Test_Illuminance_Systems(t *testing.T) {
	assert.Equal(t, SiSystem, Lux.System())
	assert.Equal(t, BiSystem, Footcandle.System())
}

func Test_Lookup_Illuminance_Names_and_Symbols(t *testing.T) {
	tests := lookUpTests{
		{Lux, "lux"}, {Lux, "lx"},
		{Footcandle, "footcandle"}, {Footcandle, "Ftc"},
		{Footcandle, "footcandles"}, {Footcandle, "foot-candle"},
		{Footcandle, "foot-candles"}, {Footcandle, "fc"},
	}

	testLookupNamesAndSymbols(t, tests)
}
