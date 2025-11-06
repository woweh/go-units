package units

import (
	"testing"

	"github.com/alecthomas/assert/v2"
)

func Test_MomentOfInertia_Conversions(t *testing.T) {
	conversionTests := []conversionTest{
		// Revit conversions from internal (ft⁴)
		{"ft⁴", "cm⁴", 863097.48412416},
		{"cm⁴", "ft⁴", 1.15861767458952e-06},
		{"ft⁴", "in⁴", 20736},
		{"in⁴", "ft⁴", 4.82253086419753e-05},
		{"ft⁴", "m⁴", 0.0086309748412416},
		{"m⁴", "ft⁴", 115.861767458952},
		{"ft⁴", "mm⁴", 8630974841.2416},
		{"mm⁴", "ft⁴", 1.15861767458952e-10},

		// Cross conversions (metric units)
		{"m⁴", "cm⁴", 100000000},
		{"cm⁴", "m⁴", 1e-08},
		{"m⁴", "mm⁴", 1000000000000},
		{"mm⁴", "m⁴", 1e-12},

		// Cross conversions (imperial)
		{"m⁴", "in⁴", 2402509.6100288294},
		{"in⁴", "m⁴", 4.16231426e-07},
	}
	testConversions(t, conversionTests)
}

func Test_MomentOfInertia_UnitSystems(t *testing.T) {
	assert.Equal(t, SiSystem, MeterToTheFourthPower.System())
	assert.Equal(t, SiSystem, CentimeterToTheFourthPower.System())
	assert.Equal(t, SiSystem, MillimeterToTheFourthPower.System())
	assert.Equal(t, BiSystem, FootToTheFourthPower.System())
	assert.Equal(t, BiSystem, InchToTheFourthPower.System())
}

func Test_Lookup_MomentOfInertia_Names_and_Symbols(t *testing.T) {
	tests := lookUpTests{
		{MeterToTheFourthPower, "meter to the fourth power"},
		{MeterToTheFourthPower, "m⁴"},
		{MeterToTheFourthPower, "meters to the fourth power"},
		{MeterToTheFourthPower, "metre to the fourth power"},
		{MeterToTheFourthPower, "metres to the fourth power"},
		{MeterToTheFourthPower, "m^4"},
		{MeterToTheFourthPower, "meter^4"},
		{CentimeterToTheFourthPower, "centimeter to the fourth power"},
		{CentimeterToTheFourthPower, "cm⁴"},
		{CentimeterToTheFourthPower, "centimeters to the fourth power"},
		{CentimeterToTheFourthPower, "centimetre to the fourth power"},
		{CentimeterToTheFourthPower, "centimetres to the fourth power"},
		{CentimeterToTheFourthPower, "cm^4"},
		{CentimeterToTheFourthPower, "centimeter^4"},
		{MillimeterToTheFourthPower, "millimeter to the fourth power"},
		{MillimeterToTheFourthPower, "mm⁴"},
		{MillimeterToTheFourthPower, "millimeters to the fourth power"},
		{MillimeterToTheFourthPower, "millimetre to the fourth power"},
		{MillimeterToTheFourthPower, "millimetres to the fourth power"},
		{MillimeterToTheFourthPower, "mm^4"},
		{MillimeterToTheFourthPower, "millimeter^4"},
		{FootToTheFourthPower, "foot to the fourth power"},
		{FootToTheFourthPower, "ft⁴"},
		{FootToTheFourthPower, "feet to the fourth power"},
		{FootToTheFourthPower, "ft^4"},
		{FootToTheFourthPower, "foot^4"},
		{InchToTheFourthPower, "inch to the fourth power"},
		{InchToTheFourthPower, "in⁴"},
		{InchToTheFourthPower, "inches to the fourth power"},
		{InchToTheFourthPower, "in^4"},
		{InchToTheFourthPower, "inch^4"},
	}
	testLookupNamesAndSymbols(t, tests)
}
