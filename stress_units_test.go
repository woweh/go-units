package units

import (
	"testing"
)

func Test_Stress_Conversions(t *testing.T) {
	conversionTests := []conversionTest{
		// Stress-specific units (not covered in pressure_units_test.go)
		{"ksf", "Pa", 47880.26},
		{"Pa", "ksf", 0.0000208854},
		{"ksi", "Pa", 6894757.293168},
		{"Pa", "ksi", 0.0000001450377},
		{"Tf/m²", "Pa", 9806.65},
		{"Pa", "Tf/m²", 0.0001019716},
		{"kgf/m²", "Pa", 9.80665},
		{"Pa", "kgf/m²", 0.1019716},
		{"N/mm²", "Pa", 1000000},
		{"Pa", "N/mm²", 0.000001},
		{"kN/cm²", "Pa", 10000000},
		{"Pa", "kN/cm²", 0.0000001},
		{"kN/mm²", "Pa", 1000000000},
		{"Pa", "kN/mm²", 0.000000001},
		{"MN/m²", "Pa", 1000000},
		{"Pa", "MN/m²", 0.000001},
		{"daN/m²", "Pa", 10},
		{"Pa", "daN/m²", 0.1},
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
