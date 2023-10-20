package units

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Density_Conversions(t *testing.T) {
	var conversionTests = []conversionTest{
		{"g/cm³", "kg/cm³", "0.001"},
		{"kg/cm³", "kg/m³", "1000000"},
		{"g/m³", "kg/m³", "0.001"},
		{"g/mL", "kg/m³", "1000"},
		{"g/L", "kg/m³", "1"},
		{"kg/L", "kg/m³", "1000"},
		{"oz/in³", "kg/m³", "1729.994"},
		{"oz/ft³", "kg/m³", "1.001153"},
		{"oz/gal", "kg/m³", "6.236023"},
		{"lb/in³", "kg/m³", "27679.90471"},
		{"lb/ft³", "kg/m³", "16.018463"},
		{"lb/gal", "kg/m³", "99.776372"},
		{"slug/ft³", "kg/m³", "515.378818"},
		{"l ton/yd³", "kg/m³", "1328.939"},
		{"oz/in³", "l ton/yd³", "1.301786"},
		{"oz/in³", "slug/ft³", "3.356743"},
	}

	testConversions(t, conversionTests)
}

func Test_Density_System(t *testing.T) {
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
