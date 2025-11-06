package units

import (
	"testing"

	"github.com/alecthomas/assert/v2"
)

func Test_UnitWeight_Conversions(t *testing.T) {
	conversionTests := []conversionTest{
		// Revit conversions (from/to internal, see RevitUnits.csv)
		{"kN/m³", "lbf/ft³", 6.36588},
		{"lbf/ft³", "kN/m³", 0.157196},
		{"kip/in³", "lbf/ft³", 1728000},
		{"lbf/ft³", "kip/in³", 0.0000005787037},
	}
	testConversions(t, conversionTests)
}

func Test_UnitWeight_UnitSystem(t *testing.T) {
	assert.Equal(t, SiSystem, KiloNewtonPerCubicMeter.System())
	assert.Equal(t, BiSystem, PoundForcePerCubicFoot.System())
	assert.Equal(t, BiSystem, KipPerCubicInch.System())
}

// No metric factories for unit weight, so no base unit tests are needed.

func Test_UnitWeight_Lookup_Names_and_Symbols(t *testing.T) {
	tests := lookUpTests{
		// SI
		{KiloNewtonPerCubicMeter, "kN/m³"},
		{KiloNewtonPerCubicMeter, "kilonewton per cubic meter"},
		{KiloNewtonPerCubicMeter, "kilonewtons per cubic meter"},
		{KiloNewtonPerCubicMeter, "kilonewton per cubic metre"},
		{KiloNewtonPerCubicMeter, "kilonewtons per cubic metre"},
		{KiloNewtonPerCubicMeter, "kN/m3"},
		// Imperial/US
		{PoundForcePerCubicFoot, "lbf/ft³"},
		{PoundForcePerCubicFoot, "pound force per cubic foot"},
		{PoundForcePerCubicFoot, "pounds force per cubic foot"},
		{PoundForcePerCubicFoot, "lb/ft³"},
		{PoundForcePerCubicFoot, "lbf/ft3"},
		{PoundForcePerCubicFoot, "pcf"},
		{KipPerCubicInch, "kip/in³"},
		{KipPerCubicInch, "kip per cubic inch"},
		{KipPerCubicInch, "kips per cubic inch"},
		{KipPerCubicInch, "kip/in3"},
	}
	testLookupNamesAndSymbols(t, tests)
}
