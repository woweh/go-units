package units

import (
	"testing"
)

func Test_Area_Conversions(t *testing.T) {
	conversionTests := []conversionTest{
		// Revit and implementation conversions (all from RevitUnits.csv and area_units.go)
		{"m²", "ft²", 10.7639104167097},
		{"ft²", "m²", 0.09290304},
		{"m²", "in²", 1550.0031000062},
		{"in²", "m²", 0.00064516},
		{"m²", "yd²", 1.1959900463011},
		{"yd²", "m²", 0.83612736},
		{"km²", "m²", 1000000},
		{"m²", "km²", 0.000001},
		{"cm²", "m²", 0.0001},
		{"m²", "cm²", 10000},
		{"mm²", "m²", 0.000001},
		{"m²", "mm²", 1000000},
		{"mi²", "m²", 2589988.110336},
		{"m²", "mi²", 3.8610215854244586e-7},
		{"ac", "m²", 4046.8564224},
		{"m²", "ac", 0.00024710538146717},
		{"hectare", "m²", 10000},
		{"are", "m²", 100},
		// Imperial conversions (sampled)
		{"square mile", "acre", 640},
		{"square yard", "square foot", 9},
		{"square foot", "square inch", 144},
		// Imperial to Metric (sampled)
		{"acre", "square meter", 4046.8564224},
		{"square yard", "square meter", 0.83612736},
		{"square foot", "square meter", 0.09290304},
		{"square inch", "square centimeter", 6.4516},
	}
	testConversions(t, conversionTests)
}

func Test_Area_UnitSystems(t *testing.T) {
	tests := []unitSystemTest{
		{SquareMilliMeter, SiSystem},
		{SquareCentiMeter, SiSystem},
		{SquareDeciMeter, SiSystem},
		{SquareMeter, SiSystem},
		{SquareDecaMeter, SiSystem},
		{SquareHectoMeter, SiSystem},
		{SquareKiloMeter, SiSystem},
		{Are, SiSystem},
		{Hectare, SiSystem},
		{SquareMile, BiSystem},
		{Acre, BiSystem},
		{SquareInch, BiSystem},
		{SquareFoot, BiSystem},
		{SquareYard, BiSystem},
	}
	testUnitSystems(t, tests)
}
