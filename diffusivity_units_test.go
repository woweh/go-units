package units

import (
	"testing"

	"github.com/alecthomas/assert/v2"
)

func Test_Diffusivity_Conversions(t *testing.T) {
	conversionTests := []conversionTest{
		{"m²/s", "ft²/s", 10.7639104167097},
		{"ft²/s", "m²/s", 0.09290304},
	}
	testConversions(t, conversionTests)
}

func Test_Diffusivity_UnitSystems(t *testing.T) {
	assert.Equal(t, SiSystem, SquareMeterPerSecond.System())
	assert.Equal(t, BiSystem, SquareFootPerSecond.System())
}

func Test_Diffusivity_BaseUnits(t *testing.T) {
	assert.Equal(t, SquareMeterPerSecond, SquareMeterPerSecond.Base())
	assert.Equal(t, SquareMeterPerSecond, SquareFootPerSecond.Base())
}

func Test_Lookup_Diffusivity_Names_and_Symbols(t *testing.T) {
	tests := lookUpTests{
		{SquareMeterPerSecond, "square meter per second"},
		{SquareMeterPerSecond, "m²/s"},
		{SquareMeterPerSecond, "square meters per second"},
		{SquareMeterPerSecond, "square metre per second"},
		{SquareMeterPerSecond, "square metres per second"},
		{SquareMeterPerSecond, "m2/s"},
		{SquareFootPerSecond, "square foot per second"},
		{SquareFootPerSecond, "ft²/s"},
		{SquareFootPerSecond, "square feet per second"},
		{SquareFootPerSecond, "ft2/s"},
	}
	testLookupNamesAndSymbols(t, tests)
}
