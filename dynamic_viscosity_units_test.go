package units

import (
	"testing"

	"github.com/alecthomas/assert/v2"
)

func Test_DynamicViscosity_Conversions(t *testing.T) {
	conversionTests := []conversionTest{
		// SI base and equivalents
		{"Pa-s", "N·s/m²", "1"},
		{"Pa-s", "cP", "1000"},
		{"kg/(m·h)", "kg/(m·s)", "3600"},
		// Imperial to SI (sampled)
		{"lb·s/ft²", "Pa-s", "0.020885"},
		{"lbm/ft-s", "Pa-s", "1.4882"},
		// Identity
		{"Pa-s", "Pa-s", "1"},
	}
	testConversions(t, conversionTests)
}

func Test_DynamicViscosity_UnitSystems(t *testing.T) {
	assert.Equal(t, SiSystem, PascalSecond.System())
	assert.Equal(t, SiSystem, NewtonSecondPerSquareMeter.System())
	assert.Equal(t, SiSystem, KilogramPerMeterSecond.System())
	assert.Equal(t, SiSystem, KilogramPerMeterHour.System())
	assert.NotEqual(t, SiSystem, Centipoise.System())
	assert.Equal(t, BiSystem, PoundForceSecondPerSquareFoot.System())
	assert.Equal(t, BiSystem, PoundMassPerFootSecond.System())
	assert.Equal(t, BiSystem, PoundMassPerFootHour.System())
}

func Test_DynamicViscosity_BaseUnits(t *testing.T) {
	assert.Equal(t, PascalSecond, PascalSecond.Base())
	assert.Equal(t, PascalSecond, NewtonSecondPerSquareMeter.Base())
	assert.Equal(t, PascalSecond, KilogramPerMeterSecond.Base())
	assert.Equal(t, PascalSecond, KilogramPerMeterHour.Base())
}
