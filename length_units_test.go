package units

import (
	"testing"
)

func Test_Length_Conversions(t *testing.T) {
	var conversionTests = []conversionTest{
		// Metric prefix conversions (representative sample)
		{"meter", "kilometer", 0.001},
		{"meter", "centimeter", 100.0},
		{"meter", "millimeter", 1000.0},

		// Imperial conversions
		{"foot", "meter", 0.3048},
		{"inch", "meter", 0.0254},
		{"yard", "meter", 0.9144},
		{"mile", "meter", 1609.344},

		// Cross-system conversions
		{"foot", "inch", 12.0},
		{"yard", "foot", 3.0},
		{"mile", "yard", 1760.0},

		// Scientific units
		{"angstrom", "meter", 1e-10},
		{"furlong", "meter", 201.168},
		{"league", "meter", 4828.032},
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
