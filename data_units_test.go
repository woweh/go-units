package units

import (
	"testing"

	"github.com/alecthomas/assert/v2"
)

func Test_Data_Conversions(t *testing.T) {
	conversionTests := []conversionTest{
		// Binary (IEC) progressions
		{"kibibyte", "byte", 1024},
		{"mebibyte", "kibibyte", 1024},
		{"gibibyte", "mebibyte", 1024},
		{"tebibyte", "gibibyte", 1024},
		{"pebibyte", "tebibyte", 1024},
		{"exbibyte", "pebibyte", 1024},
		{"zebibyte", "exbibyte", 1024},
		{"yobibyte", "zebibyte", 1024},
		// Decimal (SI) progressions (bytes)
		{"kilobyte", "byte", 1000},
		{"megabyte", "kilobyte", 1000},
		{"gigabyte", "megabyte", 1000},
		{"terabyte", "gigabyte", 1000},
		// Decimal (SI) progressions (bits)
		{"kilobit", "bit", 1000},
		{"megabit", "kilobit", 1000},
		{"gigabit", "megabit", 1000},
		// Cross-system and special conversions
		{"byte", "bit", 8},
		{"nibble", "bit", 4},
		{"kibibyte", "kilobit", 8.192},
		{"mebibyte", "megabit", 8.388608},
		{"kilobyte", "kibibyte", 0.9765625},
		{"megabyte", "mebibyte", 0.95367431640625},
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
	assert.Equal(t, SiSystem, Bit.System())
	assert.Equal(t, SiSystem, KiloBit.System())
	assert.Equal(t, SiSystem, MegaBit.System())
	assert.Equal(t, SiSystem, GigaBit.System())
	assert.Equal(t, SiSystem, TeraBit.System())
	assert.Equal(t, SiSystem, PetaBit.System())
	assert.Equal(t, SiSystem, ExaBit.System())
	// Nibble does not have a UnitSystem
}

func Test_Data_BaseUnits(t *testing.T) {
	// Only a few representative metric and IEC units
	assert.Equal(t, Byte, Byte.Base())
	assert.Equal(t, Byte, KiloByte.Base())
	assert.Equal(t, Byte, MegaByte.Base())
	assert.Equal(t, Byte, Kibibyte.Base())
	assert.Equal(t, Byte, Mebibyte.Base())
}

func Test_Lookup_Data_Names_and_Symbols(t *testing.T) {
	tests := lookUpTests{
		{Byte, "byte"},
		{Byte, "B"},
		{KiloByte, "kilobyte"},
		{KiloByte, "KB"},
		{MegaByte, "megabyte"},
		{MegaByte, "MB"},
		{GigaByte, "gigabyte"},
		{GigaByte, "GB"},
		{TeraByte, "terabyte"},
		{TeraByte, "TB"},
		{Kibibyte, "kibibyte"},
		{Kibibyte, "KiB"},
		{Mebibyte, "mebibyte"},
		{Mebibyte, "MiB"},
		{Gibibyte, "gibibyte"},
		{Gibibyte, "GiB"},
		{Tebibyte, "tebibyte"},
		{Tebibyte, "TiB"},
		{Bit, "bit"},
		{Bit, "b"},
		{KiloBit, "kilobit"},
		{KiloBit, "kb"},
		{MegaBit, "megabit"},
		{MegaBit, "Mb"},
		{Nibble, "nibble"},
	}
	testLookupNamesAndSymbols(t, tests)
}
