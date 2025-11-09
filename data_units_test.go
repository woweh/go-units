package units

import (
	"testing"
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
	tests := []unitSystemTest{
		{Byte, SiSystem},
		{KiloByte, SiSystem},
		{MegaByte, SiSystem},
		{GigaByte, SiSystem},
		{TeraByte, SiSystem},
		{PetaByte, SiSystem},
		{ExaByte, SiSystem},
		{ZettaByte, SiSystem},
		{YottaByte, SiSystem},
		{Kibibyte, IecSystem},
		{Mebibyte, IecSystem},
		{Gibibyte, IecSystem},
		{Tebibyte, IecSystem},
		{Pebibyte, IecSystem},
		{Exbibyte, IecSystem},
		{Zebibyte, IecSystem},
		{Yobibyte, IecSystem},
		{Bit, SiSystem},
		{KiloBit, SiSystem},
		{MegaBit, SiSystem},
		{GigaBit, SiSystem},
		{TeraBit, SiSystem},
		{PetaBit, SiSystem},
		{ExaBit, SiSystem},
		{Nibble, NoSystem},
	}
	testUnitSystems(t, tests)
}
