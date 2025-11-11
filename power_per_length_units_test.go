package units

import (
	"testing"
)

func Test_PowerPerLength_Conversions(t *testing.T) {
	conversionTests := []conversionTest{
		// Revit conversions
		{from: "W/m", to: "W/ft", exp: 0.3048},
		{from: "W/ft", to: "W/m", exp: 3.28083989501312},
	}
	testConversions(t, conversionTests)
}

func Test_PowerPerLength_UnitSystems(t *testing.T) {
	tests := []unitSystemTest{
		{WattPerMeter, SiSystem},
		{WattPerFoot, BiSystem},
	}
	testUnitSystems(t, tests)
}
