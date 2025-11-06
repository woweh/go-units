package units

import (
	"testing"

	"github.com/alecthomas/assert/v2"
)

func Test_PowerPerLength_Conversions(t *testing.T) {
	conversionTests := []conversionTest{
		// Revit conversions
		{"W/m", "W/ft", 0.3048},
		{"W/ft", "W/m", 3.28083989501312},
	}
	testConversions(t, conversionTests)
}

func Test_PowerPerLength_UnitSystems(t *testing.T) {
	assert.Equal(t, SiSystem, WattPerMeter.System())
	assert.Equal(t, BiSystem, WattPerFoot.System())
}

func Test_Lookup_PowerPerLength_Names_and_Symbols(t *testing.T) {
	tests := lookUpTests{
		{WattPerMeter, "watt per meter"},
		{WattPerMeter, "W/m"},
		{WattPerMeter, "watts per meter"},
		{WattPerMeter, "watt per metre"},
		{WattPerMeter, "watts per metre"},
		{WattPerFoot, "watt per foot"},
		{WattPerFoot, "W/ft"},
		{WattPerFoot, "watts per foot"},
	}
	testLookupNamesAndSymbols(t, tests)
}
