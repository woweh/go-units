package units

import (
	"testing"

	"github.com/alecthomas/assert/v2"
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
	si := SiSystem
	assert.Equal(t, si, SquareMilliMeter.System())
	assert.Equal(t, si, SquareCentiMeter.System())
	assert.Equal(t, si, SquareDeciMeter.System())
	assert.Equal(t, si, SquareMeter.System())
	assert.Equal(t, si, SquareDecaMeter.System())
	assert.Equal(t, si, SquareHectoMeter.System())
	assert.Equal(t, si, SquareKiloMeter.System())
	assert.Equal(t, si, Are.System())
	assert.Equal(t, si, Hectare.System())
	bi := BiSystem
	assert.Equal(t, bi, SquareMile.System())
	assert.Equal(t, bi, Acre.System())
	assert.Equal(t, bi, SquareInch.System())
	assert.Equal(t, bi, SquareFoot.System())
	assert.Equal(t, bi, SquareYard.System())
}

func Test_Area_BaseUnits(t *testing.T) {
	assert.Equal(t, SquareMeter, SquareMilliMeter.Base())
	assert.Equal(t, SquareMeter, SquareCentiMeter.Base())
	assert.Equal(t, SquareMeter, SquareDeciMeter.Base())
	assert.Equal(t, SquareMeter, SquareDecaMeter.Base())
}

func Test_Lookup_Area_Names_and_Symbols(t *testing.T) {
	tests := lookUpTests{
		// SI/metric
		{SquareMilliMeter, "square millimeter"}, {SquareMilliMeter, "mm²"}, {SquareMilliMeter, "square millimetre"}, {SquareMilliMeter, "mm2"}, {SquareMilliMeter, "mm^2"}, {SquareMilliMeter, "mm**2"},
		{SquareCentiMeter, "square centimeter"}, {SquareCentiMeter, "cm²"}, {SquareCentiMeter, "square centimetre"}, {SquareCentiMeter, "cm2"}, {SquareCentiMeter, "cm^2"}, {SquareCentiMeter, "cm**2"},
		{SquareDeciMeter, "square decimeter"}, {SquareDeciMeter, "dm²"}, {SquareDeciMeter, "square decimetre"}, {SquareDeciMeter, "dm2"}, {SquareDeciMeter, "dm^2"}, {SquareDeciMeter, "dm**2"},
		{SquareMeter, "square meter"}, {SquareMeter, "m²"}, {SquareMeter, "square metre"}, {SquareMeter, "m2"}, {SquareMeter, "m^2"}, {SquareMeter, "m**2"},
		{SquareDecaMeter, "square decameter"}, {SquareDecaMeter, "dam²"}, {SquareDecaMeter, "square decametre"}, {SquareDecaMeter, "dam2"}, {SquareDecaMeter, "dam^2"}, {SquareDecaMeter, "dam**2"}, {SquareDecaMeter, "are"},
		{SquareHectoMeter, "square hectometer"}, {SquareHectoMeter, "hm²"}, {SquareHectoMeter, "square hectometre"}, {SquareHectoMeter, "hm2"}, {SquareHectoMeter, "hm^2"}, {SquareHectoMeter, "hm**2"}, {SquareHectoMeter, "hectare"}, {SquareHectoMeter, "ha"},
		{SquareKiloMeter, "square kilometer"}, {SquareKiloMeter, "km²"}, {SquareKiloMeter, "square kilometre"}, {SquareKiloMeter, "km2"}, {SquareKiloMeter, "km^2"}, {SquareKiloMeter, "km**2"},
		{Are, "are"},
		{Hectare, "hectare"}, {Hectare, "ha"},
		// Imperial/US
		{SquareMile, "square mile"}, {SquareMile, "mi²"}, {SquareMile, "square mi"}, {SquareMile, "square mi."}, {SquareMile, "sq mile"}, {SquareMile, "sq. mile"}, {SquareMile, "mi2"}, {SquareMile, "mi^2"}, {SquareMile, "mi**2"}, {SquareMile, "sqmi"}, {SquareMile, "sqmi."}, {SquareMile, "sq mi"}, {SquareMile, "sq mi."}, {SquareMile, "sq.mi."}, {SquareMile, "sq.mi"},
		{Acre, "acre"}, {Acre, "ac"}, {Acre, "acres"},
		{SquareInch, "square inch"}, {SquareInch, "in²"}, {SquareInch, "square inches"}, {SquareInch, "square in"}, {SquareInch, "square in."}, {SquareInch, "square ins"}, {SquareInch, "square ins."}, {SquareInch, "sq inches"}, {SquareInch, "sq inch"}, {SquareInch, "inches/-2"}, {SquareInch, "inch/-2"}, {SquareInch, "inches2"}, {SquareInch, "inch2"}, {SquareInch, "in2"}, {SquareInch, "in^2"}, {SquareInch, "in**2"}, {SquareInch, "in/-2"}, {SquareInch, "sq in"}, {SquareInch, "sq in."}, {SquareInch, "sq ins"}, {SquareInch, "sq ins."}, {SquareInch, "sqin"}, {SquareInch, "sqin."}, {SquareInch, "sqins"}, {SquareInch, "□″"}, {SquareInch, "″2"},
		{SquareFoot, "square foot"}, {SquareFoot, "ft²"}, {SquareFoot, "square feet"}, {SquareFoot, "square ft"}, {SquareFoot, "square ft."}, {SquareFoot, "square feet."}, {SquareFoot, "ft2"}, {SquareFoot, "ft^2"}, {SquareFoot, "ft**2"}, {SquareFoot, "sq ft"}, {SquareFoot, "sq ft."}, {SquareFoot, "sqft"}, {SquareFoot, "sqft."}, {SquareFoot, "'2"},
		{SquareYard, "square yard"}, {SquareYard, "yd²"}, {SquareYard, "square yds"}, {SquareYard, "square yd"}, {SquareYard, "sq yards"}, {SquareYard, "sq yard"}, {SquareYard, "yards/-2"}, {SquareYard, "yard/-2"}, {SquareYard, "yards^2"}, {SquareYard, "yard^2"}, {SquareYard, "yds^2"}, {SquareYard, "yd^2"}, {SquareYard, "yd2"}, {SquareYard, "yards²"}, {SquareYard, "yard²"}, {SquareYard, "yds²"}, {SquareYard, "yds/-2"}, {SquareYard, "yd/-2"}, {SquareYard, "sq yds"}, {SquareYard, "sq yd"}, {SquareYard, "sq.yd."},
	}
	testLookupNamesAndSymbols(t, tests)
}
