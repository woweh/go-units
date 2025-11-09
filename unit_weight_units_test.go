package units

import (
	"testing"
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

func Test_UnitWeight_UnitSystems(t *testing.T) {
	tests := []unitSystemTest{
		{KiloNewtonPerCubicMeter, SiSystem},
		{PoundForcePerCubicFoot, BiSystem},
		{KipPerCubicInch, BiSystem},
	}
	testUnitSystems(t, tests)
}
