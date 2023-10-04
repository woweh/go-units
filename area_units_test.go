package units

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAreaSystems(t *testing.T) {
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
