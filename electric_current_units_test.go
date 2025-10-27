package units

import (
	"testing"

	"github.com/alecthomas/assert/v2"
)

func Test_ElectricCurrent_Systems(t *testing.T) {
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
