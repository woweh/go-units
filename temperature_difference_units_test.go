package units

import (
	"testing"
)

func Test_TemperatureDifference_Conversions(t *testing.T) {
	var conversionTests = []conversionTest{
		// Base conversions
		{"K", "delta°C", "1"},
		{"K", "delta°F", "1.8"},
		{"K", "delta°R", "1.8"},
		{"delta°C", "delta°F", "1.8"},

		// Identity
		{"K", "K", "1"},
		{"delta°C", "delta°C", "1"},
		{"delta°F", "delta°F", "1"},
	}

	testConversions(t, conversionTests)
}
