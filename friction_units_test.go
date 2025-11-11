package units

import (
	"testing"
)

func Test_Friction_Conversions(t *testing.T) {
	var conversionTests = []conversionTest{
		// Metric / SI relations
		{from: "mH2O/m", to: "mmH2O/m", exp: 1000.0}, // 1 mH2O/m = 1000 mmH2O/m
		{from: "mmH2O/m", to: "Pa/m", exp: 9.80665},  // 1 mmH2O/m = 9.80665 Pa/m
		{from: "mH2O/m", to: "Pa/m", exp: 9806.65},   // 1 mH2O/m = 9806.65 Pa/m

		// Imperial relations (representative)
		{from: "mmH2O/m", to: "FT/100ft", exp: 0.1},   // 1 mmH2O/m = 0.1 FT/100ft (per definitions)
		{from: "FT/100ft", to: "Pa/m", exp: 0.980665}, // 1 FT/100ft = 0.980665 Pa/m
		{from: "in-wg/100ft", to: "Pa/m", exp: 11.76798},
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
