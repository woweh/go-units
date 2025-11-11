package units

import (
	"testing"
)

func Test_Stress_Conversions(t *testing.T) {
	conversionTests := []conversionTest{
		// Stress-specific units (not covered in pressure_units_test.go)
		{from: "ksf", to: "Pa", exp: 47880.26},
		{from: "Pa", to: "ksf", exp: 0.0000208854},
		{from: "ksi", to: "Pa", exp: 6894757.293168},
		{from: "Pa", to: "ksi", exp: 0.0000001450377},
		{from: "Tf/m²", to: "Pa", exp: 9806.65},
		{from: "Pa", to: "Tf/m²", exp: 0.0001019716},
		{from: "kgf/m²", to: "Pa", exp: 9.80665},
		{from: "Pa", to: "kgf/m²", exp: 0.1019716},
		{from: "N/mm²", to: "Pa", exp: 1000000},
		{from: "Pa", to: "N/mm²", exp: 0.000001},
		{from: "kN/cm²", to: "Pa", exp: 10000000},
		{from: "Pa", to: "kN/cm²", exp: 0.0000001},
		{from: "kN/mm²", to: "Pa", exp: 1000000000},
		{from: "Pa", to: "kN/mm²", exp: 0.000000001},
		{from: "MN/m²", to: "Pa", exp: 1000000},
		{from: "Pa", to: "MN/m²", exp: 0.000001},
		{from: "daN/m²", to: "Pa", exp: 10},
		{from: "Pa", to: "daN/m²", exp: 0.1},
	}
	testConversions(t, conversionTests)
}

func Test_Stress_UnitSystem(t *testing.T) {
	tests := []unitSystemTest{
		{KipPerSquareFoot, BiSystem},
		{KipPerSquareInch, BiSystem},
		{NewtonPerSquareMillimeter, SiSystem},
		{KiloNewtonPerSquareCentimeter, SiSystem},
		{KiloNewtonPerSquareMillimeter, SiSystem},
		{MegaNewtonPerSquareMeter, SiSystem},
		{DekaNewtonPerSquareMeter, SiSystem},
		{KilogramForcePerSquareMeter, MKpSSystem},
		{TonneForcePerSquareMeter, MKpSSystem},
	}
	testUnitSystems(t, tests)
}
