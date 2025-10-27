package units

import (
	"testing"
)

func Test_SpecificHeat_Conversions(t *testing.T) {
	var conversionTests = []conversionTest{
		// Base conversions (adjusted for library precision)
		{"J/(g·°C)", "J/(kg·°C)", "1000"},
		{"BTU/(lb·°F)", "J/(kg·°C)", "4186.8"},
		{"J/(kg·°C)", "BTU/(lb·°F)", "0.0002388459"},

		// Identity
		{"J/(kg·°C)", "J/(kg·°C)", "1"},
		{"BTU/(lb·°F)", "BTU/(lb·°F)", "1"},
	}

	testConversions(t, conversionTests)
}
