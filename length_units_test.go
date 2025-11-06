package units

import (
	"testing"

	"github.com/alecthomas/assert/v2"
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

func Test_Length_Systems(t *testing.T) {
	si := SiSystem
	bi := BiSystem
	us := UsSystem

	// SI system units
	assert.Equal(t, si, Meter.System())
	assert.Equal(t, si, KiloMeter.System())
	assert.Equal(t, si, CentiMeter.System())
	assert.Equal(t, si, MilliMeter.System())
	assert.Equal(t, si, MicroMeter.System())
	assert.Equal(t, si, Angstrom.System())

	// BI system units
	assert.Equal(t, bi, Inch.System())
	assert.Equal(t, bi, Foot.System())
	assert.Equal(t, bi, Yard.System())
	assert.Equal(t, bi, Mile.System())
	assert.Equal(t, bi, League.System())
	assert.Equal(t, bi, Furlong.System())

	// US system units
	assert.Equal(t, us, USSurveyFoot.System())
}

func Test_Length_BaseUnits(t *testing.T) {
	// Metric factory units should have Meter as base
	assert.Equal(t, Meter, KiloMeter.Base())
	assert.Equal(t, Meter, CentiMeter.Base())
	assert.Equal(t, Meter, MilliMeter.Base())
	assert.Equal(t, Meter, MicroMeter.Base())
	assert.Equal(t, Meter, NanoMeter.Base())
	assert.Equal(t, Meter, DeciMeter.Base())
	assert.Equal(t, Meter, DecaMeter.Base())
	assert.Equal(t, Meter, HectoMeter.Base())
	assert.Equal(t, Meter, MegaMeter.Base())
	assert.Equal(t, Meter, GigaMeter.Base())
	assert.Equal(t, Meter, TeraMeter.Base())
	assert.Equal(t, Meter, PetaMeter.Base())
	assert.Equal(t, Meter, ExaMeter.Base())
	assert.Equal(t, Meter, PicoMeter.Base())
	assert.Equal(t, Meter, FemtoMeter.Base())
	assert.Equal(t, Meter, AttoMeter.Base())
}

func Test_Lookup_Length_Names_and_Symbols(t *testing.T) {
	tests := lookUpTests{
		// Meter and aliases
		{Meter, "meter"}, {Meter, "m"}, {Meter, "metre"},

		// Metric prefixes (representative sample)
		{KiloMeter, "kilometer"}, {KiloMeter, "km"},
		{CentiMeter, "centimeter"}, {CentiMeter, "cm"},
		{MilliMeter, "millimeter"}, {MilliMeter, "mm"},
		{MicroMeter, "micrometer"}, {MicroMeter, "µm"},
		{NanoMeter, "nanometer"}, {NanoMeter, "nm"},

		// Angstrom and aliases
		{Angstrom, "angstrom"}, {Angstrom, "Å"},
		{Angstrom, "ångström"}, {Angstrom, "angstroms"}, {Angstrom, "ångströms"},

		// Imperial units
		{Inch, "inch"}, {Inch, "in"}, {Inch, "inches"}, {Inch, "in."}, {Inch, "″"},
		{Foot, "foot"}, {Foot, "ft"}, {Foot, "feet"}, {Foot, "ft."}, {Foot, "′"},
		{Yard, "yard"}, {Yard, "yd"}, {Yard, "yd."},
		{Mile, "mile"}, {Mile, "mi"}, {Mile, "mi."},
		{League, "league"}, {League, "lea"},
		{Furlong, "furlong"}, {Furlong, "fur"},

		// US Survey
		{USSurveyFoot, "US survey foot"}, {USSurveyFoot, "USft"},
	}

	testLookupNamesAndSymbols(t, tests)
}
