package units

import (
	"testing"
)

func Test_DynamicViscosity_Conversions(t *testing.T) {
	conversionTests := []conversionTest{
		// SI base and equivalents (Revit factors)
		{from: "Pa-s", to: "N·s/m²", exp: 1.0},
		{from: "Pa-s", to: "kg/(m·s)", exp: 1.0},
		{from: "Pa-s", to: "cP", exp: 1000.0},
		{from: "kg/(m·h)", to: "kg/(m·s)", exp: 3600.0},
		// Revit conversion factors (from RevitUnits.csv, both directions)
		{from: "Pa-s", to: "N·s/m²", exp: 1.0},
		{from: "N·s/m²", to: "Pa-s", exp: 1.0},
		{from: "Pa-s", to: "kg/(m·s)", exp: 1.0},
		{from: "kg/(m·s)", to: "Pa-s", exp: 1.0},
		{from: "Pa-s", to: "kg/(m·h)", exp: 0.0000846666666666667},
		{from: "kg/(m·h)", to: "Pa-s", exp: 11811.0236220472},
		{from: "Pa-s", to: "cP", exp: 3280.83989501312},
		{from: "cP", to: "Pa-s", exp: 0.0003048},
		{from: "Pa-s", to: "lb·s/ft²", exp: 0.0685217658567918},
		{from: "lb·s/ft²", to: "Pa-s", exp: 14.5939029372064},
		{from: "Pa-s", to: "lbm/ft-h", exp: 7936.64143865559},
		{from: "lbm/ft-h", to: "Pa-s", exp: 0.000125997880555556},
		{from: "Pa-s", to: "lbm/ft-s", exp: 2.20462262184878},
		{from: "lbm/ft-s", to: "Pa-s", exp: 0.45359237},
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
