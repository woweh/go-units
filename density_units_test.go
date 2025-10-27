package units

import (
	"testing"

	"github.com/alecthomas/assert/v2"
)

func Test_Density_Conversions(t *testing.T) {
	conversionTests := []conversionTest{
		// Metric conversions
		{"g/cm³", "kg/cm³", "0.001"},
		{"kg/cm³", "kg/m³", "1000000"},
		{"g/m³", "kg/m³", "0.001"},
		{"g/mL", "kg/m³", "1000"},
		{"g/L", "kg/m³", "1"},
		{"kg/L", "kg/m³", "1000"},
		// Imperial to Metric (sampled)
		{"oz/in³", "kg/m³", "1729.994"},
		{"lb/ft³", "kg/m³", "16.018463"},
		{"slug/ft³", "kg/m³", "515.378818"},
		// Cross-system (sampled)
		{"oz/in³", "slug/ft³", "3.356743"},
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
	assert.Equal(t, KilogramPerCubicMeter, GramPerCubicCentimeter.Base())
	assert.Equal(t, KilogramPerCubicMeter, KilogramPerCubicCentimeter.Base())
	assert.Equal(t, KilogramPerCubicMeter, GramPerCubicMeter.Base())
	assert.Equal(t, KilogramPerCubicMeter, KilogramPerCubicMeter.Base())
	assert.Equal(t, KilogramPerCubicMeter, GramPerMilliliter.Base())
	assert.Equal(t, KilogramPerCubicMeter, GramPerLiter.Base())
	assert.Equal(t, KilogramPerCubicMeter, KilogramPerLiter.Base())
}
