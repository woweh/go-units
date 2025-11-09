package units

import (
	"testing"
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
	tests := []unitSystemTest{
		{MeterToTheFourthPower, SiSystem},
		{CentimeterToTheFourthPower, SiSystem},
		{MillimeterToTheFourthPower, SiSystem},
		{FootToTheFourthPower, BiSystem},
		{InchToTheFourthPower, BiSystem},
	}
	testUnitSystems(t, tests)
}
