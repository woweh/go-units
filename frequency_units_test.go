package units

import (
	"math"
	"testing"
)

func Test_Frequency_Conversions(t *testing.T) {
	tau := 2 * math.Pi
	var conversionTests = []conversionTest{
		// Core relations (representative and precise values)
		{from: "hertz", to: "radian per second", exp: tau},      // 2π
		{from: "hertz", to: "degree per second", exp: 360.0},    // 360° per cycle
		{from: "hertz", to: "revolution per minute", exp: 60.0}, // 60 rev/min per Hz
		{from: "hertz", to: "kilohertz", exp: 0.001},            // metric prefix
		{from: "gigahertz", to: "fresnel", exp: 0.001},          // 1 GHz = 0.001 THz (fresnel symbol maps to THz)

		// Cross-unit exact reciprocals / ratios
		{from: "radian per second", to: "revolution per second", exp: 1.0 / tau}, // 1/(2π)
		{from: "degree per second", to: "revolution per second", exp: 1.0 / 360.0},
		{from: "revolution per minute", to: "revolution per second", exp: 1.0 / 60.0},
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
