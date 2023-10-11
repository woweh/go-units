package units

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_AreaSystems(t *testing.T) {
	SI := "metric"
	assert.Equal(t, SI, SquareMilliMeter.System())
	assert.Equal(t, SI, SquareCentiMeter.System())
	assert.Equal(t, SI, SquareDeciMeter.System())
	assert.Equal(t, SI, SquareDecaMeter.System())
	assert.Equal(t, SI, SquareHectoMeter.System())
	assert.Equal(t, SI, SquareKiloMeter.System())

	BI := "imperial"
	assert.Equal(t, BI, SquareMile.System())
	assert.Equal(t, BI, Acre.System())
	assert.Equal(t, BI, SquareInch.System())
	assert.Equal(t, BI, SquareFoot.System())
	assert.Equal(t, BI, SquareYard.System())
}

func Test_AreaAliases(t *testing.T) {

	tests := []aliasTest{
		{nil, "some unit alias that does not exist"},
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

	testAliases(t, tests)
}
