package units

import (
	"math"
	"testing"
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

func Test_Frequency_UnitSystems(t *testing.T) {
	tests := []unitSystemTest{
		{Hertz, SiSystem},
		{KiloHertz, SiSystem},
		{MegaHertz, SiSystem},
		{GigaHertz, SiSystem},
		{MilliHertz, SiSystem},
		{MicroHertz, SiSystem},
		{RadianPerSecond, SiSystem},
		{DegreePerSecond, SiSystem},
		{RevolutionPerSecond, SiSystem},
	}
	testUnitSystems(t, tests)
}
