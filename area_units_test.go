package units

import (
	"testing"
)

func Test_Area_Conversions(t *testing.T) {
	conversionTests := []conversionTest{
		// Revit and implementation conversions (all from RevitUnits.csv and area_units.go)
		{from: "m²", to: "ft²", exp: 10.7639104167097},
		{from: "ft²", to: "m²", exp: 0.09290304},
		{from: "m²", to: "in²", exp: 1550.0031000062},
		{from: "in²", to: "m²", exp: 0.00064516},
		{from: "m²", to: "yd²", exp: 1.1959900463011},
		{from: "yd²", to: "m²", exp: 0.83612736},
		{from: "km²", to: "m²", exp: 1000000},
		{from: "m²", to: "km²", exp: 0.000001},
		{from: "cm²", to: "m²", exp: 0.0001},
		{from: "m²", to: "cm²", exp: 10000},
		{from: "mm²", to: "m²", exp: 0.000001},
		{from: "m²", to: "mm²", exp: 1000000},
		{from: "mi²", to: "m²", exp: 2589988.110336},
		{from: "m²", to: "mi²", exp: 3.8610215854244586e-7},
		{from: "ac", to: "m²", exp: 4046.8564224},
		{from: "m²", to: "ac", exp: 0.00024710538146717},
		{from: "hectare", to: "m²", exp: 10000},
		{from: "are", to: "m²", exp: 100},
		// Imperial conversions (sampled)
		{from: "square mile", to: "acre", exp: 640},
		{from: "square yard", to: "square foot", exp: 9},
		{from: "square foot", to: "square inch", exp: 144},
		// Imperial to Metric (sampled)
		{from: "acre", to: "square meter", exp: 4046.8564224},
		{from: "square yard", to: "square meter", exp: 0.83612736},
		{from: "square foot", to: "square meter", exp: 0.09290304},
		{from: "square inch", to: "square centimeter", exp: 6.4516},
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
