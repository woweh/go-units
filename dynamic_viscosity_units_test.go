package units

import (
	"testing"
)

func Test_DynamicViscosity_Conversions(t *testing.T) {
	conversionTests := []conversionTest{
		// SI base and equivalents (Revit factors)
		{"Pa-s", "N·s/m²", 1.0},
		{"Pa-s", "kg/(m·s)", 1.0},
		{"Pa-s", "cP", 1000.0},
		{"kg/(m·h)", "kg/(m·s)", 3600.0},
		// Revit conversion factors (from RevitUnits.csv, both directions)
		{"Pa-s", "N·s/m²", 1.0},
		{"N·s/m²", "Pa-s", 1.0},
		{"Pa-s", "kg/(m·s)", 1.0},
		{"kg/(m·s)", "Pa-s", 1.0},
		{"Pa-s", "kg/(m·h)", 0.0000846666666666667},
		{"kg/(m·h)", "Pa-s", 11811.0236220472},
		{"Pa-s", "cP", 3280.83989501312},
		{"cP", "Pa-s", 0.0003048},
		{"Pa-s", "lb·s/ft²", 0.0685217658567918},
		{"lb·s/ft²", "Pa-s", 14.5939029372064},
		{"Pa-s", "lbm/ft-h", 7936.64143865559},
		{"lbm/ft-h", "Pa-s", 0.000125997880555556},
		{"Pa-s", "lbm/ft-s", 2.20462262184878},
		{"lbm/ft-s", "Pa-s", 0.45359237},
		// Remove identity tests
	}
	testConversions(t, conversionTests)
}

func Test_DynamicViscosity_UnitSystems(t *testing.T) {
	tests := []unitSystemTest{
		{PascalSecond, SiSystem},
		{NewtonSecondPerSquareMeter, SiSystem},
		{KilogramPerMeterSecond, SiSystem},
		{KilogramPerMeterHour, SiSystem},
		{Centipoise, NoSystem},
		{PoundForceSecondPerSquareFoot, BiSystem},
		{PoundMassPerFootSecond, BiSystem},
		{PoundMassPerFootHour, BiSystem},
	}
	testUnitSystems(t, tests)
}
