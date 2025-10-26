package units

import (
	"fmt"
	"math"
	"testing"
)

func Test_Angle_Conversions(t *testing.T) {
	var conversionTests = []conversionTest{
		// Base conversions
		{"turn", "radian", fmt.Sprintf("%f", 2*math.Pi)},
		{"radian", "degree", fmt.Sprintf("%.5f", 180/math.Pi)},
		{"degree", "gon", fmt.Sprintf("%f", 10/9.0)},

		// Additional important conversions
		{"turn", "degree", "360"},
		{"turn", "gon", "400"},
		{"radian", "milliradian", "1000"},
		{"radian", "microradian", "1000000"},

		// Gon subdivisions
		{"gon", "milligon", "1000"},
		{"gon", "centigon", "100"},
		{"gon", "decigon", "10"},
		{"gon", "decagon", "0.1"},
		{"gon", "hectogon", "0.01"},

		// Derived conversion
		{"turn", "milliradian", fmt.Sprintf("%f", 2*math.Pi*1000)},
	}

	testConversions(t, conversionTests)
}
