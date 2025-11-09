package units

import (
	"testing"
)

func Test_PowerPerLength_Conversions(t *testing.T) {
	conversionTests := []conversionTest{
		// Revit conversions
		{"W/m", "W/ft", 0.3048},
		{"W/ft", "W/m", 3.28083989501312},
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
