package units

import (
	"testing"

	"github.com/alecthomas/assert/v2"
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
	assert.Equal(t, BiSystem, KipPerSquareFoot.System())
	assert.Equal(t, BiSystem, KipPerSquareInch.System())
	assert.Equal(t, SiSystem, NewtonPerSquareMillimeter.System())
	assert.Equal(t, SiSystem, KiloNewtonPerSquareCentimeter.System())
	assert.Equal(t, SiSystem, KiloNewtonPerSquareMillimeter.System())
	assert.Equal(t, SiSystem, MegaNewtonPerSquareMeter.System())
	assert.Equal(t, SiSystem, DekaNewtonPerSquareMeter.System())
	assert.Equal(t, MKpSSystem, KilogramForcePerSquareMeter.System())
	assert.Equal(t, MKpSSystem, TonneForcePerSquareMeter.System())
}

func Test_Stress_MetricFactory_BaseUnits(t *testing.T) {
	// Only test stress-specific metric factories if any (none in stress_units.go)
	// No test needed here as all metric factories for stress are covered in pressure_units_test.go
}

func Test_Stress_Lookup_Names_and_Symbols(t *testing.T) {
	tests := lookUpTests{
		{KipPerSquareFoot, "ksf"},
		{KipPerSquareFoot, "kip per square foot"},
		{KipPerSquareFoot, "kips per square foot"},
		{KipPerSquareInch, "ksi"},
		{KipPerSquareInch, "kip per square inch"},
		{KipPerSquareInch, "kips per square inch"},
		{NewtonPerSquareMillimeter, "N/mm²"},
		{NewtonPerSquareMillimeter, "newton per square millimeter"},
		{NewtonPerSquareMillimeter, "newtons per square millimeter"},
		{KiloNewtonPerSquareCentimeter, "kN/cm²"},
		{KiloNewtonPerSquareCentimeter, "kilonewton per square centimeter"},
		{KiloNewtonPerSquareCentimeter, "kilonewtons per square centimeter"},
		{KiloNewtonPerSquareMillimeter, "kN/mm²"},
		{KiloNewtonPerSquareMillimeter, "kilonewton per square millimeter"},
		{KiloNewtonPerSquareMillimeter, "kilonewtons per square millimeter"},
		{MegaNewtonPerSquareMeter, "MN/m²"},
		{MegaNewtonPerSquareMeter, "meganewton per square meter"},
		{MegaNewtonPerSquareMeter, "meganewtons per square meter"},
		{DekaNewtonPerSquareMeter, "daN/m²"},
		{DekaNewtonPerSquareMeter, "dekanewton per square meter"},
		{DekaNewtonPerSquareMeter, "dekanewtons per square meter"},
		{KilogramForcePerSquareMeter, "kgf/m²"},
		{KilogramForcePerSquareMeter, "kilogram force per square meter"},
		{KilogramForcePerSquareMeter, "kilograms force per square meter"},
		{TonneForcePerSquareMeter, "Tf/m²"},
		{TonneForcePerSquareMeter, "tonne force per square meter"},
		{TonneForcePerSquareMeter, "tonnes force per square meter"},
	}
	testLookupNamesAndSymbols(t, tests)
}
