package units

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVolumeSystems(t *testing.T) {
	si := UnitSystem("metric")
	assert.Equal(t, si, Liter.System())
	assert.Equal(t, si, ExaLiter.System())
	assert.Equal(t, si, PetaLiter.System())
	assert.Equal(t, si, TeraLiter.System())
	assert.Equal(t, si, GigaLiter.System())
	assert.Equal(t, si, MegaLiter.System())
	assert.Equal(t, si, KiloLiter.System())
	assert.Equal(t, si, HectoLiter.System())
	assert.Equal(t, si, DecaLiter.System())
	assert.Equal(t, si, DeciLiter.System())
	assert.Equal(t, si, CentiLiter.System())
	assert.Equal(t, si, MilliLiter.System())
	assert.Equal(t, si, MicroLiter.System())
	assert.Equal(t, si, NanoLiter.System())
	assert.Equal(t, si, PicoLiter.System())
	assert.Equal(t, si, FemtoLiter.System())
	assert.Equal(t, si, AttoLiter.System())

	bi := UnitSystem("imperial")
	assert.Equal(t, bi, Quart.System())
	assert.Equal(t, bi, Pint.System())
	assert.Equal(t, bi, Gallon.System())
	assert.Equal(t, bi, FluidOunce.System())

	us := UnitSystem("us")
	assert.Equal(t, us, FluidQuart.System())
	assert.Equal(t, us, FluidPint.System())
	assert.Equal(t, us, FluidGallon.System())
	assert.Equal(t, us, CustomaryFluidOunce.System())
}
