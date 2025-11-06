package units

import (
	"testing"

	"github.com/alecthomas/assert/v2"
)

func Test_Friction_Conversions(t *testing.T) {
	var conversionTests = []conversionTest{
		// Metric / SI relations
		{"mH2O/m", "mmH2O/m", 1000.0}, // 1 mH2O/m = 1000 mmH2O/m
		{"mmH2O/m", "Pa/m", 9.80665},  // 1 mmH2O/m = 9.80665 Pa/m
		{"mH2O/m", "Pa/m", 9806.65},   // 1 mH2O/m = 9806.65 Pa/m

		// Imperial relations (representative)
		{"mmH2O/m", "FT/100ft", 0.1},   // 1 mmH2O/m = 0.1 FT/100ft (per definitions)
		{"FT/100ft", "Pa/m", 0.980665}, // 1 FT/100ft = 0.980665 Pa/m
		{"in-wg/100ft", "Pa/m", 11.76798},
	}

	testConversions(t, conversionTests)
}

func Test_Friction_Systems(t *testing.T) {
	// Pascal per meter is SI
	si := SiSystem
	assert.Equal(t, si, PascalPerMeter.System())

	// Metric hydraulic units should be metric (no explicit system set => SI considered)
	assert.Equal(t, si, MillimeterOfWaterPerMeter.System())
	assert.Equal(t, si, MeterOfWaterPerMeter.System())

	// Imperial units are in BI
	bi := BiSystem
	assert.Equal(t, bi, FootOfWaterPer100Feet.System())
	assert.Equal(t, bi, InchOfWaterPer100Feet.System())
}

func Test_Lookup_Friction_Names_and_Symbols(t *testing.T) {
	tests := lookUpTests{
		{MillimeterOfWaterPerMeter, "millimeter of water column per meter"}, {MillimeterOfWaterPerMeter, "mmH2O/m"},
		{MillimeterOfWaterPerMeter, "millimeters of water column per meter"}, {MillimeterOfWaterPerMeter, "mmH2O per m"},
		{MeterOfWaterPerMeter, "meter of water column per meter"}, {MeterOfWaterPerMeter, "mH2O/m"},
		{PascalPerMeter, "pascal per meter"}, {PascalPerMeter, "Pa/m"},
		{FootOfWaterPer100Feet, "foot of water per 100 feet"}, {FootOfWaterPer100Feet, "FT/100ft"},
		{InchOfWaterPer100Feet, "inch of water per 100 feet"}, {InchOfWaterPer100Feet, "in-wg/100ft"},
	}

	testLookupNamesAndSymbols(t, tests)
}
