package units

import "testing"

func Test_Data_Conversions(t *testing.T) {
	var conversionTests = []conversionTest{
		// Base conversions (binary)
		{"kibibyte", "byte", "1024"},
		{"mebibyte", "kibibyte", "1024"},
		{"gibibyte", "mebibyte", "1024"},
		{"tebibyte", "gibibyte", "1024"},
		{"pebibyte", "tebibyte", "1024"},
		{"exbibyte", "pebibyte", "1024"},
		{"zebibyte", "exbibyte", "1024"},
		{"yobibyte", "zebibyte", "1024"},

		// Base conversions (decimal)
		{"kilobit", "bit", "1000"},
		{"megabit", "kilobit", "1000"},
		{"gigabit", "megabit", "1000"},
		{"terabit", "gigabit", "1000"},
		{"petabit", "terabit", "1000"},
		{"exabit", "petabit", "1000"},

		// Cross-system conversions
		{"byte", "bit", "8"},
		{"kibibyte", "kilobit", "8.192"},
		{"mebibyte", "megabit", "8.388608"},
		{"gibibyte", "gigabit", "8.589935"},
		{"tebibyte", "terabit", "8.796093"},
		{"pebibyte", "petabit", "9.007199"},

		// Uncommon but useful conversions
		{"yobibyte", "bit", "9671406556917033397649408"},
		{"exabit", "byte", "125000000000000000"},
		{"nibble", "bit", "4"},
	}

	testConversions(t, conversionTests)
}
