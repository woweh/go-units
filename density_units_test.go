package units

import (
	"testing"
)

func Test_Density_Conversions(t *testing.T) {
	conversionTests := []conversionTest{
		// SI base conversions
		{from: "kg/m³", to: "t/m³", exp: 0.001},
		{from: "kg/m³", to: "g/m³", exp: 1000},
		{from: "kg/m³", to: "mg/m³", exp: 1000000},
		{from: "kg/m³", to: "kg/L", exp: 0.001},
		{from: "kg/m³", to: "g/L", exp: 1},
		{from: "kg/m³", to: "mg/L", exp: 1000},
		{from: "kg/m³", to: "kg/dm³", exp: 0.001},
		{from: "kg/m³", to: "g/dm³", exp: 1},
		{from: "kg/m³", to: "mg/dm³", exp: 1000},
		{from: "kg/m³", to: "kg/cm³", exp: 0.000001},
		{from: "kg/m³", to: "g/cm³", exp: 0.001},
		{from: "kg/m³", to: "mg/cm³", exp: 1},
		{from: "kg/m³", to: "kg/mL", exp: 0.000001},
		{from: "kg/m³", to: "g/mL", exp: 0.001},
		{from: "kg/m³", to: "mg/mL", exp: 1},

		// Imperial conversions
		{from: "kg/m³", to: "oz/ft³", exp: 0.998847369},
		{from: "kg/m³", to: "oz/in³", exp: 0.000578036672},
		{from: "kg/m³", to: "oz/yd³", exp: 26.968878969},
		{from: "kg/m³", to: "oz/gal", exp: 0.16035856},
		{from: "kg/m³", to: "lb/ft³", exp: 0.062427960576},
		{from: "kg/m³", to: "lb/in³", exp: 0.000036127292},
		{from: "kg/m³", to: "lb/yd³", exp: 1.685554936},
		{from: "kg/m³", to: "lb/gal", exp: 0.010022412},
		{from: "kg/m³", to: "slug/ft³", exp: 0.001940320},
		{from: "kg/m³", to: "slug/in³", exp: 0.0000011228706},
		{from: "kg/m³", to: "slug/yd³", exp: 0.052388649},
		{from: "kg/m³", to: "l ton/yd³", exp: 0.000752480},

		// Metric to metric
		{from: "g/cm³", to: "kg/m³", exp: 1000},
		{from: "kg/cm³", to: "kg/m³", exp: 1000000},
		{from: "g/m³", to: "kg/m³", exp: 0.001},
		{from: "g/mL", to: "kg/m³", exp: 1000},
		{from: "g/L", to: "kg/m³", exp: 1},
		{from: "kg/L", to: "kg/m³", exp: 1000},

		// Imperial to metric
		{from: "oz/in³", to: "kg/m³", exp: 1729.994044392},
		{from: "oz/ft³", to: "kg/m³", exp: 1.001153961},
		{from: "oz/gal", to: "kg/m³", exp: 6.236025068},
		{from: "lb/in³", to: "kg/m³", exp: 27679.904710203},
		{from: "lb/ft³", to: "kg/m³", exp: 16.018463374},
		{from: "lb/gal", to: "kg/m³", exp: 99.776381175},
		{from: "slug/ft³", to: "kg/m³", exp: 515.378818492},
		{from: "l ton/yd³", to: "kg/m³", exp: 1328.939, tol: fPtr(1)},

		// Cross-system
		{from: "oz/in³", to: "slug/ft³", exp: 3.356742618},
		{from: "g/cm³", to: "lb/ft³", exp: 62.42796057600, tol: fPtr(1e-6)},
		{from: "kg/L", to: "lb/gal", exp: 10.022412},
		{from: "kg/m³", to: "oz/in³", exp: 0.000578037},
		{from: "oz/in³", to: "kg/m³", exp: 1729.994044392},
	}
	testConversions(t, conversionTests)
}

func Test_Density_UnitSystems(t *testing.T) {
	tests := []unitSystemTest{
		// Metric units
		{GramPerCubicCentimeter, SiSystem},
		{KilogramPerCubicCentimeter, SiSystem},
		{GramPerCubicMeter, SiSystem},
		{KilogramPerCubicMeter, SiSystem},
		{GramPerMilliliter, SiSystem},
		{GramPerLiter, SiSystem},
		{KilogramPerLiter, SiSystem},
		{TonnePerCubicMeter, SiSystem},
		{MilligramPerCubicMeter, SiSystem},
		{KilogramPerCubicDecimeter, SiSystem},
		{GramPerCubicDecimeter, SiSystem},
		{MilligramPerCubicDecimeter, SiSystem},
		{MilligramPerCubicCentimeter, SiSystem},
		{KilogramPerMilliliter, SiSystem},
		{MilligramPerMilliliter, SiSystem},
		{MilligramPerLiter, SiSystem},

		// Imperial units
		{OuncePerCubicFoot, BiSystem},
		{OuncePerCubicInch, BiSystem},
		{OuncePerCubicYard, BiSystem},
		{OuncePerGallon, BiSystem},
		{PoundPerCubicFoot, BiSystem},
		{PoundPerCubicInch, BiSystem},
		{PoundPerCubicYard, BiSystem},
		{PoundPerGallon, BiSystem},
		{SlugPerCubicFoot, BiSystem},
		{SlugPerCubicInch, BiSystem},
		{SlugPerCubicYard, BiSystem},
		{TonPerCubicYard, BiSystem},
	}
	testUnitSystems(t, tests)
}
