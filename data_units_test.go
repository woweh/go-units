package units

import (
	"testing"
)

func Test_Data_Conversions(t *testing.T) {
	conversionTests := []conversionTest{
		// Binary (IEC) progressions
		{from: "kibibyte", to: "byte", exp: 1024},
		{from: "mebibyte", to: "kibibyte", exp: 1024},
		{from: "gibibyte", to: "mebibyte", exp: 1024},
		{from: "tebibyte", to: "gibibyte", exp: 1024},
		{from: "pebibyte", to: "tebibyte", exp: 1024},
		{from: "exbibyte", to: "pebibyte", exp: 1024},
		{from: "zebibyte", to: "exbibyte", exp: 1024},
		{from: "yobibyte", to: "zebibyte", exp: 1024},
		// Decimal (SI) progressions (bytes)
		{from: "kilobyte", to: "byte", exp: 1000},
		{from: "megabyte", to: "kilobyte", exp: 1000},
		{from: "gigabyte", to: "megabyte", exp: 1000},
		{from: "terabyte", to: "gigabyte", exp: 1000},
		// Decimal (SI) progressions (bits)
		{from: "kilobit", to: "bit", exp: 1000},
		{from: "megabit", to: "kilobit", exp: 1000},
		{from: "gigabit", to: "megabit", exp: 1000},
		// Cross-system and special conversions
		{from: "byte", to: "bit", exp: 8},
		{from: "nibble", to: "bit", exp: 4},
		{from: "kibibyte", to: "kilobit", exp: 8.192},
		{from: "mebibyte", to: "megabit", exp: 8.388608},
		{from: "kilobyte", to: "kibibyte", exp: 0.9765625},
		{from: "megabyte", to: "mebibyte", exp: 0.95367431640625},
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
