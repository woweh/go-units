package units

import (
	"testing"
)

func Test_Density_Conversions(t *testing.T) {
	conversionTests := []conversionTest{
		// Metric to metric (sampled, not all)
		{from: "g/cm³", to: "kg/m³", exp: 1000},
		{from: "kg/cm³", to: "kg/m³", exp: 1000000},
		{from: "g/m³", to: "kg/m³", exp: 0.001},
		{from: "g/mL", to: "kg/m³", exp: 1000},
		{from: "g/L", to: "kg/m³", exp: 1},
		{from: "kg/L", to: "kg/m³", exp: 1000},
		// Imperial to metric (from implementation)
		{from: "oz/in³", to: "kg/m³", exp: 48.9879},
		{from: "oz/ft³", to: "kg/m³", exp: 1.001153},
		{from: "oz/gal", to: "kg/m³", exp: 6.236023},
		{from: "lb/in³", to: "kg/m³", exp: 783.806},
		{from: "lb/ft³", to: "kg/m³", exp: 16.018463},
		{from: "lb/gal", to: "kg/m³", exp: 99.776372},
		{from: "slug/ft³", to: "kg/m³", exp: 515.3788184},
		{from: "l ton/yd³", to: "kg/m³", exp: 1328.939},
		// Cross-system (sampled, not identity)
		{from: "oz/in³", to: "slug/ft³", exp: 0.095059}, // 48.9879 / 515.3788184
		{from: "g/cm³", to: "lb/ft³", exp: 62.42796},    // 1000 / 16.018463
		{from: "kg/L", to: "lb/gal", exp: 10.022412},    // 1000 / 99.776372
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
