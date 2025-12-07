package units

import (
	"testing"
)

func Test_MomentOfInertia_Conversions(t *testing.T) {
	conversionTests := []conversionTest{
		// Revit conversions from internal (ft⁴)
		{from: "ft⁴", to: "cm⁴", exp: 863097.48412416, tol: fPtr(1e-4)},
		{from: "cm⁴", to: "ft⁴", exp: 1.15861767458952e-06},
		{from: "ft⁴", to: "in⁴", exp: 20736, tol: fPtr(1e-6)},
		{from: "in⁴", to: "ft⁴", exp: 4.82253086419753e-05},
		{from: "ft⁴", to: "m⁴", exp: 0.0086309748412416},
		{from: "m⁴", to: "ft⁴", exp: 115.861767458952},
		{from: "ft⁴", to: "mm⁴", exp: 8630974841.2416, tol: fPtr(1e-2)},
		{from: "mm⁴", to: "ft⁴", exp: 1.15861767458952e-10},

		// Cross conversions (metric units)
		{from: "m⁴", to: "cm⁴", exp: 100000000},
		{from: "cm⁴", to: "m⁴", exp: 1e-08},
		{from: "m⁴", to: "mm⁴", exp: 1000000000000},
		{from: "mm⁴", to: "m⁴", exp: 1e-12},

		// Cross conversions (imperial)
		{from: "m⁴", to: "in⁴", exp: 2402509.6100288294},
		{from: "in⁴", to: "m⁴", exp: 4.16231426e-07},
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
