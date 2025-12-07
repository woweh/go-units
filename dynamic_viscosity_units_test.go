package units

import (
	"testing"
)

func Test_DynamicViscosity_Conversions(t *testing.T) {
	conversionTests := []conversionTest{
		// SI base and equivalents
		{from: "Pa-s", to: "N·s/m²", exp: 1.0},
		{from: "Pa-s", to: "kg/(m·s)", exp: 1.0},
		{from: "Pa-s", to: "cP", exp: 1000.0},
		{from: "cP", to: "Pa-s", exp: 0.001},
		{from: "kg/(m·h)", to: "kg/(m·s)", exp: 3600.0},
		{from: "kg/(m·s)", to: "kg/(m·h)", exp: 0.00027777777777777},
		// Imperial units
		{from: "Pa-s", to: "lb·s/ft²", exp: 0.0685217658567918},
		{from: "lb·s/ft²", to: "Pa-s", exp: 14.5939029372064, tol: fPtr(1e-8)},
		{from: "Pa-s", to: "lbm/ft-s", exp: 2.20462262184878, tol: fPtr(1e-9)},
		{from: "lbm/ft-s", to: "Pa-s", exp: 0.45359237},
		{from: "Pa-s", to: "lbm/ft-h", exp: 7936.64143865559, tol: fPtr(1e-6)},
		{from: "lbm/ft-h", to: "Pa-s", exp: 0.000125997880555556},
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
