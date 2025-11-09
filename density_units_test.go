package units

import (
	"testing"
)

func Test_Density_Conversions(t *testing.T) {
	conversionTests := []conversionTest{
		// Metric to metric (sampled, not all)
		{"g/cm³", "kg/m³", 1000},
		{"kg/cm³", "kg/m³", 1000000},
		{"g/m³", "kg/m³", 0.001},
		{"g/mL", "kg/m³", 1000},
		{"g/L", "kg/m³", 1},
		{"kg/L", "kg/m³", 1000},
		// Imperial to metric (from implementation)
		{"oz/in³", "kg/m³", 48.9879},
		{"oz/ft³", "kg/m³", 1.001153},
		{"oz/gal", "kg/m³", 6.236023},
		{"lb/in³", "kg/m³", 783.806},
		{"lb/ft³", "kg/m³", 16.018463},
		{"lb/gal", "kg/m³", 99.776372},
		{"slug/ft³", "kg/m³", 515.3788184},
		{"l ton/yd³", "kg/m³", 1328.939},
		// Cross-system (sampled, not identity)
		{"oz/in³", "slug/ft³", 0.095059}, // 48.9879 / 515.3788184
		{"g/cm³", "lb/ft³", 62.42796},    // 1000 / 16.018463
		{"kg/L", "lb/gal", 10.022412},    // 1000 / 99.776372
	}
	testConversions(t, conversionTests)
}

func Test_Density_UnitSystems(t *testing.T) {
	tests := []unitSystemTest{
		{GramPerCubicCentimeter, SiSystem},
		{KilogramPerCubicCentimeter, SiSystem},
		{GramPerCubicMeter, SiSystem},
		{KilogramPerCubicMeter, SiSystem},
		{GramPerMilliliter, SiSystem},
		{GramPerLiter, SiSystem},
		{KilogramPerLiter, SiSystem},
		{OuncePerCubicInch, BiSystem},
		{OuncePerCubicFoot, BiSystem},
		{OuncePerGallon, BiSystem},
		{PoundPerCubicInch, BiSystem},
		{PoundPerCubicFoot, BiSystem},
		{PoundPerGallon, BiSystem},
		{SlugPerCubicFoot, BiSystem},
		{TonPerCubicYard, BiSystem},
	}
	testUnitSystems(t, tests)
}
