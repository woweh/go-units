package units

import (
	"fmt"
	"math"
	"testing"

	"github.com/alecthomas/assert/v2"
)

func Test_Angle_Conversions(t *testing.T) {
	conversionTests := []conversionTest{
		// Base conversions (Revit/implementation)
		{"turn", "radian", fmt.Sprintf("%f", 2*math.Pi)},
		{"radian", "degree", fmt.Sprintf("%.5f", 180/math.Pi)},
		{"degree", "gon", fmt.Sprintf("%f", 10/9.0)},

		// Additional important conversions
		{"turn", "degree", "360"},
		{"turn", "gon", "400"},
		{"radian", "milliradian", "1000"},
		{"radian", "microradian", "1000000"},

		// Gon subdivisions (sampled)
		{"gon", "milligon", "1000"},
		{"gon", "centigon", "100"},

		// Derived conversion
		{"turn", "milliradian", fmt.Sprintf("%f", 2*math.Pi*1000)},
	}
	testConversions(t, conversionTests)
}

func Test_Angle_UnitSystems(t *testing.T) {
	assert.Equal(t, SiSystem, Radian.System())
	assert.Equal(t, SiSystem, MilliRadian.System())
	assert.Equal(t, SiSystem, MicroRadian.System())
	assert.Equal(t, SiSystem, Gon.System())
	assert.Equal(t, SiSystem, MilliGon.System())
	assert.NotEqual(t, SiSystem, Degree.System())
	assert.NotEqual(t, SiSystem, Turn.System())
}

func Test_Angle_BaseUnits(t *testing.T) {
	assert.Equal(t, Radian, Radian.Base())
	assert.Equal(t, Radian, MilliRadian.Base())
	assert.Equal(t, Radian, MicroRadian.Base())
	assert.Equal(t, Gon, MilliGon.Base())
	assert.Equal(t, Gon, CentiGon.Base())
}

func Test_Lookup_Angle_Names_and_Symbols(t *testing.T) {
	tests := lookUpTests{
		{Turn, "turn"},
		{Turn, "tr"},
		{Radian, "radian"},
		{Radian, "rad"},
		{Degree, "degree"},
		{Degree, "°"},
		{Gon, "gon"},
		{Gon, "grad"},
		{MilliRadian, "milliradian"},
		{MilliGon, "milligon"},
	}
	testLookupNamesAndSymbols(t, tests)
}
