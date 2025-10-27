package units

import (
	"testing"
)

func Test_PowerPerLength_Conversions(t *testing.T) {
	var conversionTests = []conversionTest{
		{"W/ft", "W/m", "3.28084"},
		{"W/m", "W/ft", "0.3048"},
		{"W/m", "W/m", "1"},
	}

	testConversions(t, conversionTests)
}
