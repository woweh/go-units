package units

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_AreaSystems(t *testing.T) {
	si := UnitSystem("metric")
	assert.Equal(t, si, SquareMilliMeter.System())
	assert.Equal(t, si, SquareCentiMeter.System())
	assert.Equal(t, si, SquareDeciMeter.System())
	assert.Equal(t, si, SquareDecaMeter.System())
	assert.Equal(t, si, SquareHectoMeter.System())
	assert.Equal(t, si, SquareKiloMeter.System())

	bi := UnitSystem("imperial")
	assert.Equal(t, bi, SquareMile.System())
	assert.Equal(t, bi, Acre.System())
	assert.Equal(t, bi, SquareInch.System())
	assert.Equal(t, bi, SquareFoot.System())
	assert.Equal(t, bi, SquareYard.System())
}

func Test_AreaNamesAndSymbols(t *testing.T) {

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
		{SquareFoot, "ft2"},
		{SquareFoot, "ft^2"},
		{SquareFoot, "ft**2"},
		{SquareYard, "yd2"},
		{SquareYard, "yd^2"},
		{SquareYard, "sq yd"},
		{SquareYard, "yd/-2"},
		{SquareMile, "mi2"},
		{SquareMile, "mi^2"},
		{SquareMile, "mi**2"},
		{SquareMile, "sq mi"},
		{Acre, "acres"},
	}

	testNamesAndSymbols(t, tests)
}
