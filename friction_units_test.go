package units

import (
	"testing"
)

func Test_Friction_Conversions(t *testing.T) {
	var conversionTests = []conversionTest{
		// Metric / SI relations
		{"mH2O/m", "mmH2O/m", 1000.0}, // 1 mH2O/m = 1000 mmH2O/m
		{"mmH2O/m", "Pa/m", 9.80665},  // 1 mmH2O/m = 9.80665 Pa/m
		{"mH2O/m", "Pa/m", 9806.65},   // 1 mH2O/m = 9806.65 Pa/m

		// Imperial relations (representative)
		{"mmH2O/m", "FT/100ft", 0.1},   // 1 mmH2O/m = 0.1 FT/100ft (per definitions)
		{"FT/100ft", "Pa/m", 0.980665}, // 1 FT/100ft = 0.980665 Pa/m
		{"in-wg/100ft", "Pa/m", 11.76798},
	}

	testConversions(t, conversionTests)
}

func Test_Friction_UnitSystems(t *testing.T) {
	tests := []unitSystemTest{
		{PascalPerMeter, SiSystem},
		{MillimeterOfWaterPerMeter, SiSystem},
		{MeterOfWaterPerMeter, SiSystem},
		{FootOfWaterPer100Feet, BiSystem},
		{InchOfWaterPer100Feet, BiSystem},
	}
	testUnitSystems(t, tests)
}
