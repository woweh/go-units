package units

import (
	"testing"

	"github.com/alecthomas/assert/v2"
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
	assert.Equal(t, SiSystem, PascalSecond.System())
	assert.Equal(t, SiSystem, NewtonSecondPerSquareMeter.System())
	assert.Equal(t, SiSystem, KilogramPerMeterSecond.System())
	assert.Equal(t, SiSystem, KilogramPerMeterHour.System())
	assert.NotEqual(t, SiSystem, Centipoise.System())
	assert.Equal(t, BiSystem, PoundForceSecondPerSquareFoot.System())
	assert.Equal(t, BiSystem, PoundMassPerFootSecond.System())
	assert.Equal(t, BiSystem, PoundMassPerFootHour.System())
}

func Test_DynamicViscosity_BaseUnits(t *testing.T) {
	assert.Equal(t, PascalSecond, PascalSecond.Base())
	assert.Equal(t, PascalSecond, NewtonSecondPerSquareMeter.Base())
	assert.Equal(t, PascalSecond, KilogramPerMeterSecond.Base())
	assert.Equal(t, PascalSecond, KilogramPerMeterHour.Base())
	assert.Equal(t, PascalSecond, Centipoise.Base())
	assert.Equal(t, PascalSecond, PoundForceSecondPerSquareFoot.Base())
	assert.Equal(t, PascalSecond, PoundMassPerFootSecond.Base())
	assert.Equal(t, PascalSecond, PoundMassPerFootHour.Base())
}

func Test_Lookup_DynamicViscosity_Names_and_Symbols(t *testing.T) {
	tests := lookUpTests{
		{PascalSecond, "pascal second"},
		{PascalSecond, "Pa-s"},
		{PascalSecond, "pascal seconds"},
		{PascalSecond, "Pa*s"},
		{PascalSecond, "Pas"},
		{NewtonSecondPerSquareMeter, "newton second per square meter"},
		{NewtonSecondPerSquareMeter, "N·s/m²"},
		{NewtonSecondPerSquareMeter, "newton second per square metre"},
		{NewtonSecondPerSquareMeter, "N*s/m²"},
		{KilogramPerMeterSecond, "kilogram per meter second"},
		{KilogramPerMeterSecond, "kg/(m·s)"},
		{KilogramPerMeterSecond, "kilogram per metre second"},
		{KilogramPerMeterSecond, "kg/m/s"},
		{KilogramPerMeterHour, "kilogram per meter hour"},
		{KilogramPerMeterHour, "kg/(m·h)"},
		{KilogramPerMeterHour, "kilogram per metre hour"},
		{KilogramPerMeterHour, "kg/m/h"},
		{Centipoise, "centipoise"},
		{Centipoise, "cP"},
		{Centipoise, "centipoises"},
		{Centipoise, "cp"},
		{PoundForceSecondPerSquareFoot, "pound force second per square foot"},
		{PoundForceSecondPerSquareFoot, "lb·s/ft²"},
		{PoundForceSecondPerSquareFoot, "pound force seconds per square foot"},
		{PoundForceSecondPerSquareFoot, "lbf*s/ft²"},
		{PoundMassPerFootSecond, "pound mass per foot second"},
		{PoundMassPerFootSecond, "lbm/ft-s"},
		{PoundMassPerFootSecond, "pounds mass per foot second"},
		{PoundMassPerFootSecond, "lbm/ft/s"},
		{PoundMassPerFootHour, "pound mass per foot hour"},
		{PoundMassPerFootHour, "lbm/ft-h"},
		{PoundMassPerFootHour, "pounds mass per foot hour"},
		{PoundMassPerFootHour, "lbm/ft/h"},
	}
	testLookupNamesAndSymbols(t, tests)
}
