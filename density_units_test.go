package units

import (
	"testing"

	"github.com/alecthomas/assert/v2"
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
	si := SiSystem
	assert.Equal(t, si, GramPerCubicCentimeter.System())
	assert.Equal(t, si, KilogramPerCubicCentimeter.System())
	assert.Equal(t, si, GramPerCubicMeter.System())
	assert.Equal(t, si, KilogramPerCubicMeter.System())
	assert.Equal(t, si, GramPerMilliliter.System())
	assert.Equal(t, si, GramPerLiter.System())
	assert.Equal(t, si, KilogramPerLiter.System())

	bi := BiSystem
	assert.Equal(t, bi, OuncePerCubicInch.System())
	assert.Equal(t, bi, OuncePerCubicFoot.System())
	assert.Equal(t, bi, OuncePerGallon.System())
	assert.Equal(t, bi, PoundPerCubicInch.System())
	assert.Equal(t, bi, PoundPerCubicFoot.System())
	assert.Equal(t, bi, PoundPerGallon.System())
	assert.Equal(t, bi, SlugPerCubicFoot.System())
	assert.Equal(t, bi, TonPerCubicYard.System())
}

func Test_Density_BaseUnits(t *testing.T) {
	// Only a few representative metric and imperial units
	assert.Equal(t, KilogramPerCubicMeter, GramPerCubicCentimeter.Base())
	assert.Equal(t, KilogramPerCubicMeter, KilogramPerCubicMeter.Base())
	assert.Equal(t, KilogramPerCubicMeter, OuncePerCubicInch.Base())
	assert.Equal(t, KilogramPerCubicMeter, PoundPerCubicFoot.Base())
}

func Test_Lookup_Density_Names_and_Symbols(t *testing.T) {
	tests := lookUpTests{
		{GramPerCubicCentimeter, "gram per cubic centimeter"},
		{GramPerCubicCentimeter, "g/cm³"},
		{KilogramPerCubicCentimeter, "kilogram per cubic centimeter"},
		{KilogramPerCubicCentimeter, "kg/cm³"},
		{GramPerCubicMeter, "gram per cubic meter"},
		{GramPerCubicMeter, "g/m³"},
		{KilogramPerCubicMeter, "kilogram per cubic meter"},
		{KilogramPerCubicMeter, "kg/m³"},
		{GramPerMilliliter, "gram per milliliter"},
		{GramPerMilliliter, "g/mL"},
		{GramPerLiter, "gram per liter"},
		{GramPerLiter, "g/L"},
		{KilogramPerLiter, "kilogram per liter"},
		{KilogramPerLiter, "kg/L"},
		{OuncePerCubicInch, "ounce per cubic inch"},
		{OuncePerCubicInch, "oz/in³"},
		{OuncePerCubicFoot, "ounce per cubic foot"},
		{OuncePerCubicFoot, "oz/ft³"},
		{OuncePerGallon, "ounce per gallon"},
		{OuncePerGallon, "oz/gal"},
		{PoundPerCubicInch, "pound per cubic inch"},
		{PoundPerCubicInch, "lb/in³"},
		{PoundPerCubicFoot, "pound per cubic foot"},
		{PoundPerCubicFoot, "lb/ft³"},
		{PoundPerGallon, "pound per gallon"},
		{PoundPerGallon, "lb/gal"},
		{SlugPerCubicFoot, "slug per cubic foot"},
		{SlugPerCubicFoot, "slug/ft³"},
		{TonPerCubicYard, "ton per cubic yard"},
		{TonPerCubicYard, "l ton/yd³"},
	}
	testLookupNamesAndSymbols(t, tests)
}
