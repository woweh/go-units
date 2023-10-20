package units

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Area_Conversions(t *testing.T) {
	var conversionTests = []conversionTest{
		{"square kilometer", "square meter", "1000000"},
		{"square hectometer", "square meter", "10000"},
		{"square decameter", "square meter", "100"},
		{"square meter", "square decimeter", "100"},
		{"square meter", "square centimeter", "10000"},
		{"square meter", "square millimeter", "1000000"},
		{"m2", "mm2", "1000000"},
		{"square decimeter", "square centimeter", "100"},
		{"square decimeter", "square millimeter", "10000"},
		{"square centimeter", "square millimeter", "100"},
		{"square mile", "square kilometer", "2.589988"},
		{"acre", "square kilometer", "0.00404686"},
		{"square yard", "square meter", "0.836127"},
		{"square foot", "square meter", "0.092903"},
		{"square inch", "square meter", "0.00064516"},
		{"square inch", "square foot", "0.00694444"},
		{"square yard", "square foot", "9"},
		{"square mile", "acre", "640"},
		{"square yard", "square inch", "1296"},
		{"square mile", "square yard", "3097600"},
		{"square mile", "square foot", "27878400"},
		{"square mile", "square inch", "4014489600"},
		{"square foot", "square inch", "144"},
		{"square yard", "square inch", "1296"},
		{"square mile", "square inch", "4014489600"},
	}

	testConversions(t, conversionTests)
}

func Test_Area_Systems(t *testing.T) {
	si := SiSystem
	assert.Equal(t, si, SquareMilliMeter.System())
	assert.Equal(t, si, SquareCentiMeter.System())
	assert.Equal(t, si, SquareDeciMeter.System())
	assert.Equal(t, si, SquareDecaMeter.System())
	assert.Equal(t, si, SquareHectoMeter.System())
	assert.Equal(t, si, SquareKiloMeter.System())

	bi := BiSystem
	assert.Equal(t, bi, SquareMile.System())
	assert.Equal(t, bi, Acre.System())
	assert.Equal(t, bi, SquareInch.System())
	assert.Equal(t, bi, SquareFoot.System())
	assert.Equal(t, bi, SquareYard.System())
}

func Test_Lookup_Area_Names_and_Symbols(t *testing.T) {

	tests := []lookUpTestUnit{
		{nil, "some unit key that does not exist"},
		{nil, "MM2"}, // symbols are case-sensitive
		{nil, "dAm^2"},
		{SquareMilliMeter, "mm²"},
		{SquareMilliMeter, "mm2"},
		{SquareMilliMeter, "mm^2"},
		{SquareCentiMeter, "cm2"},
		{SquareCentiMeter, "cm^2"},
		{SquareDeciMeter, "dm2"},
		{SquareDeciMeter, "dm^2"},
		{SquareMeter, "m2"},
		{SquareMeter, "m^2"},
		{SquareDecaMeter, "dam2"},
		{SquareDecaMeter, "dam^2"},
		{SquareHectoMeter, "hm2"},
		{SquareHectoMeter, "hm^2"},
		{SquareHectoMeter, "hectare"},
		{SquareHectoMeter, "ha"},
		{SquareKiloMeter, "km2"},
		{SquareKiloMeter, "km^2"},
		{SquareInch, "sq in"},
		{SquareInch, "sq in."},
		{SquareInch, "in2"},
		{SquareInch, "in^2"},
		{SquareInch, "in**2"},
		{SquareInch, "in/-2"},
		{SquareInch, "sq inch"},
		{SquareInch, "square inches"},
		{SquareInch, "square ins"},
		{SquareInch, "sq inches"},
		{SquareFoot, "square feet"},
		{SquareFoot, "square ft"},
		{SquareFoot, "ft2"},
		{SquareFoot, "ft^2"},
		{SquareFoot, "ft**2"},
		{SquareFoot, "sqft"},
		{SquareYard, "yd2"},
		{SquareYard, "yd^2"},
		{SquareYard, "sq yd"},
		{SquareYard, "yd/-2"},
		{SquareYard, "square yds"},
		{SquareYard, "sq yards"},
		{SquareMile, "mi2"},
		{SquareMile, "mi^2"},
		{SquareMile, "mi**2"},
		{SquareMile, "sq mi"},
		{SquareMile, "square mi"},
		{SquareMile, "sq. mile"},
		{Acre, "acres"},
	}

	testLookupNamesAndSymbols(t, tests)
}
