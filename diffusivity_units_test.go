package units

import (
	"testing"

	"github.com/alecthomas/assert/v2"
)

func Test_Diffusivity_Conversions(t *testing.T) {
	conversionTests := []conversionTest{
		// SI base to Imperial
		{"m²/s", "ft²/s", "10.76391"},
		{"ft²/s", "m²/s", "0.092903"},
		// Identity
		{"m²/s", "m²/s", "1"},
		{"ft²/s", "ft²/s", "1"},
	}
	testConversions(t, conversionTests)
}

func Test_Diffusivity_UnitSystems(t *testing.T) {
	assert.Equal(t, SiSystem, SquareMeterPerSecond.System())
	assert.Equal(t, BiSystem, SquareFootPerSecond.System())
}

func Test_Diffusivity_BaseUnits(t *testing.T) {
	assert.Equal(t, SquareMeterPerSecond, SquareMeterPerSecond.Base())
}
