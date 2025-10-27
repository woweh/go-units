package units

import (
	"github.com/alecthomas/assert/v2"
	"testing"
)

func Test_Data_Conversions(t *testing.T) {
	conversionTests := []conversionTest{
		// Base conversions (binary)
		{"kibibyte", "byte", "1024"},
		{"mebibyte", "kibibyte", "1024"},
		{"gibibyte", "mebibyte", "1024"},
		// Base conversions (decimal)
		{"kilobit", "bit", "1000"},
		{"megabit", "kilobit", "1000"},
		{"gigabit", "megabit", "1000"},
		// Cross-system conversions (sampled)
		{"byte", "bit", "8"},
		{"kibibyte", "kilobit", "8.192"},
		{"mebibyte", "megabit", "8.388608"},
		// Uncommon but useful conversions
		{"yobibyte", "bit", "9671406556917033397649408"},
		{"nibble", "bit", "4"},
	}
	testConversions(t, conversionTests)
}

func Test_Data_UnitSystems(t *testing.T) {
	assert.Equal(t, SiSystem, Byte.System())
	assert.Equal(t, SiSystem, KiloByte.System())
	assert.Equal(t, SiSystem, MegaByte.System())
	assert.Equal(t, SiSystem, GigaByte.System())
	assert.Equal(t, SiSystem, TeraByte.System())
	assert.Equal(t, SiSystem, PetaByte.System())
	assert.Equal(t, SiSystem, ExaByte.System())
	assert.Equal(t, SiSystem, ZettaByte.System())
	assert.Equal(t, SiSystem, YottaByte.System())
	assert.Equal(t, IecSystem, Kibibyte.System())
	assert.Equal(t, IecSystem, Mebibyte.System())
	assert.Equal(t, IecSystem, Gibibyte.System())
	assert.Equal(t, IecSystem, Tebibyte.System())
	assert.Equal(t, IecSystem, Pebibyte.System())
	assert.Equal(t, IecSystem, Exbibyte.System())
	assert.Equal(t, IecSystem, Zebibyte.System())
	assert.Equal(t, IecSystem, Yobibyte.System())
}

func Test_Data_BaseUnits(t *testing.T) {
	assert.Equal(t, Byte, Byte.Base())
	assert.Equal(t, Byte, KiloByte.Base())
	assert.Equal(t, Byte, MegaByte.Base())
	assert.Equal(t, Byte, GigaByte.Base())
	assert.Equal(t, Byte, TeraByte.Base())
	assert.Equal(t, Byte, PetaByte.Base())
	assert.Equal(t, Byte, ExaByte.Base())
	assert.Equal(t, Byte, ZettaByte.Base())
	assert.Equal(t, Byte, YottaByte.Base())
}
