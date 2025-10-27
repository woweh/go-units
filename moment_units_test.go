package units

import (
	"testing"
)

func Test_Moment_Conversions(t *testing.T) {
	var conversionTests = []conversionTest{
		{"N-m", "daN-m", "0.1"},
		{"N-m", "kN-m", "0.001"},
		{"kip-ft", "lb-ft", "1000"},
		{"N-m", "N-m", "1"},
	}

	testConversions(t, conversionTests)
}
