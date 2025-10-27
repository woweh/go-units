package units

import (
	"testing"

	"github.com/alecthomas/assert/v2"
)

func Test_Area_Conversions(t *testing.T) {
	conversionTests := []conversionTest{
		// Metric conversions (sampled, Revit/implementation)
		{"square kilometer", "square meter", "1000000"},
		{"square meter", "square centimeter", "10000"},
		{"square meter", "square millimeter", "1000000"},
		// Metric to Imperial (sampled)
		{"square meter", "square yard", "1.19599"},
		{"square meter", "square foot", "10.76391"},
		{"square centimeter", "square inch", "0.1550003"},
		// Imperial conversions (sampled)
		{"square mile", "acre", "640"},
		{"square yard", "square foot", "9"},
		{"square foot", "square inch", "144"},
		// Imperial to Metric (sampled)
		{"acre", "square meter", "4046.856422"},
		{"square yard", "square meter", "0.836127"},
		{"square foot", "square meter", "0.092903"},
		{"square inch", "square centimeter", "6.4516"},
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
	assert.Equal(t, SquareMeter, SquareHectoMeter.Base())
	assert.Equal(t, SquareMeter, SquareKiloMeter.Base())
}

func Test_Lookup_Area_Names_and_Symbols(t *testing.T) {
	tests := lookUpTests{
		{SquareMilliMeter, "square millimeter"},
		{SquareMilliMeter, "mm²"},
		{SquareCentiMeter, "square centimeter"},
		{SquareCentiMeter, "cm²"},
		{SquareMeter, "square meter"},
		{SquareMeter, "m²"},
		{SquareKiloMeter, "square kilometer"},
		{SquareKiloMeter, "km²"},
		{Acre, "acre"},
		{Acre, "ac"},
		{SquareFoot, "square foot"},
		{SquareFoot, "ft²"},
		{SquareInch, "square inch"},
		{SquareInch, "in²"},
		{SquareYard, "square yard"},
		{SquareYard, "yd²"},
	}
	testLookupNamesAndSymbols(t, tests)
}
