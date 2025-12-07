package units

import (
	"testing"
)

func Test_UnitWeight_Conversions(t *testing.T) {
	conversionTests := []conversionTest{
		// Revit conversions (from/to internal, see RevitUnits.csv)
		{from: "kN/m³", to: "lbf/ft³", exp: 6.365879804, tol: fPtr(1e-8)},
		{from: "lbf/ft³", to: "kN/m³", exp: 0.157087473},
		{from: "kip/in³", to: "lbf/ft³", exp: 1728000},
		{from: "lbf/ft³", to: "kip/in³", exp: 0.0000005787037},
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
