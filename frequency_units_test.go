package units

import (
	"math"
	"testing"

	"github.com/alecthomas/assert/v2"
)

func Test_Frequency_Conversions(t *testing.T) {
	tau := 2 * math.Pi
	var conversionTests = []conversionTest{
		// Core relations (representative and precise values)
		{"hertz", "radian per second", tau},      // 2π
		{"hertz", "degree per second", 360.0},    // 360° per cycle
		{"hertz", "revolution per minute", 60.0}, // 60 rev/min per Hz
		{"hertz", "kilohertz", 0.001},            // metric prefix
		{"gigahertz", "fresnel", 0.001},          // 1 GHz = 0.001 THz (fresnel symbol maps to THz)

		// Cross-unit exact reciprocals / ratios
		{"radian per second", "revolution per second", 1.0 / tau}, // 1/(2π)
		{"degree per second", "revolution per second", 1.0 / 360.0},
		{"revolution per minute", "revolution per second", 1.0 / 60.0},
	}

	testConversions(t, conversionTests)
}

func Test_Frequency_Systems(t *testing.T) {
	si := SiSystem
	assert.Equal(t, si, Hertz.System())
	assert.Equal(t, si, KiloHertz.System())
	assert.Equal(t, si, MegaHertz.System())
	assert.Equal(t, si, GigaHertz.System())
	assert.Equal(t, si, MilliHertz.System())
	assert.Equal(t, si, MicroHertz.System())

	// derived SI frequency units (angles/revolutions) are also in SI
	assert.Equal(t, si, RadianPerSecond.System())
	assert.Equal(t, si, DegreePerSecond.System())
	assert.Equal(t, si, RevolutionPerSecond.System())
}

func Test_Frequency_BaseUnits(t *testing.T) {
	// metric factories should report Hertz as base unit
	assert.Equal(t, Hertz, KiloHertz.Base())
	assert.Equal(t, Hertz, MilliHertz.Base())
	assert.Equal(t, Hertz, CentiHertz.Base())
}

func Test_Lookup_Frequency_Names_and_Symbols(t *testing.T) {
	tests := lookUpTests{
		{Hertz, "hertz"}, {Hertz, "Hz"}, {Hertz, "cps"}, {Hertz, "1/s"}, {Hertz, "cycles per second"},
		{KiloHertz, "kilohertz"}, {KiloHertz, "kHz"},
		{GigaHertz, "gigahertz"}, {GigaHertz, "GHz"},
		{TeraHertz, "terahertz"}, {TeraHertz, "fresnel"}, {TeraHertz, "THz"},
		{RadianPerSecond, "radian per second"}, {RadianPerSecond, "rad/s"},
		{DegreePerHour, "degree per hour"}, {DegreePerHour, "°/h"},
		{RevolutionPerDay, "revolution per day"}, {RevolutionPerDay, "rev/d"},
	}

	testLookupNamesAndSymbols(t, tests)
}
