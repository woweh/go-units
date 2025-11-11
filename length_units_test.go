package units

import (
	"testing"
)

func Test_Length_Conversions(t *testing.T) {
	var conversionTests = []conversionTest{
		// Metric prefix conversions (representative sample)
		{from: "meter", to: "kilometer", exp: 0.001},
		{from: "meter", to: "centimeter", exp: 100.0},
		{from: "meter", to: "millimeter", exp: 1000.0},

		// Imperial conversions
		{from: "foot", to: "meter", exp: 0.3048},
		{from: "inch", to: "meter", exp: 0.0254},
		{from: "yard", to: "meter", exp: 0.9144},
		{from: "mile", to: "meter", exp: 1609.344},

		// Cross-system conversions
		{from: "foot", to: "inch", exp: 12.0},
		{from: "yard", to: "foot", exp: 3.0},
		{from: "mile", to: "yard", exp: 1760.0},

		// Scientific units
		{from: "angstrom", to: "meter", exp: 1e-10},
		{from: "furlong", to: "meter", exp: 201.168},
		{from: "league", to: "meter", exp: 4828.032},
	}

	testConversions(t, conversionTests)
}

func Test_Length_UnitSystems(t *testing.T) {
	tests := []unitSystemTest{
		{Meter, SiSystem},
		{KiloMeter, SiSystem},
		{CentiMeter, SiSystem},
		{MilliMeter, SiSystem},
		{MicroMeter, SiSystem},
		{Angstrom, SiSystem},
		{Inch, BiSystem},
		{Foot, BiSystem},
		{Yard, BiSystem},
		{Mile, BiSystem},
		{League, BiSystem},
		{Furlong, BiSystem},
		{USSurveyFoot, UsSystem},
	}
	testUnitSystems(t, tests)
}
