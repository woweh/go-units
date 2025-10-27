package units

import (
	"testing"

	"github.com/alecthomas/assert/v2"
)

func Test_ElectricCurrent_UnitSystems(t *testing.T) {
	si := SiSystem
	assert.Equal(t, si, Ampere.System())
	assert.Equal(t, si, MilliAmpere.System())
	assert.Equal(t, si, MicroAmpere.System())
	assert.Equal(t, si, NanoAmpere.System())
	assert.Equal(t, si, PicoAmpere.System())
	assert.Equal(t, si, FemtoAmpere.System())
	assert.Equal(t, si, AttoAmpere.System())
	assert.Equal(t, si, ZeptoAmpere.System())
	assert.Equal(t, si, YoctoAmpere.System())
	assert.Equal(t, si, KiloAmpere.System())
	assert.Equal(t, si, MegaAmpere.System())
	assert.Equal(t, si, GigaAmpere.System())
	assert.Equal(t, si, TeraAmpere.System())
	assert.Equal(t, si, PetaAmpere.System())
	assert.Equal(t, si, ExaAmpere.System())
	assert.Equal(t, si, ZettaAmpere.System())
	assert.Equal(t, si, YottaAmpere.System())
}

func Test_ElectricCurrent_BaseUnits(t *testing.T) {
	assert.Equal(t, Ampere, Ampere.Base())
	assert.Equal(t, Ampere, MilliAmpere.Base())
	assert.Equal(t, Ampere, MicroAmpere.Base())
	assert.Equal(t, Ampere, NanoAmpere.Base())
	assert.Equal(t, Ampere, PicoAmpere.Base())
	assert.Equal(t, Ampere, FemtoAmpere.Base())
	assert.Equal(t, Ampere, AttoAmpere.Base())
	assert.Equal(t, Ampere, ZeptoAmpere.Base())
	assert.Equal(t, Ampere, YoctoAmpere.Base())
	assert.Equal(t, Ampere, KiloAmpere.Base())
	assert.Equal(t, Ampere, MegaAmpere.Base())
	assert.Equal(t, Ampere, GigaAmpere.Base())
	assert.Equal(t, Ampere, TeraAmpere.Base())
	assert.Equal(t, Ampere, PetaAmpere.Base())
	assert.Equal(t, Ampere, ExaAmpere.Base())
	assert.Equal(t, Ampere, ZettaAmpere.Base())
	assert.Equal(t, Ampere, YottaAmpere.Base())
}
