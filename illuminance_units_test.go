package units

import (
	"testing"
)

func Test_Illuminance_Conversions(t *testing.T) {
	var conversionTests = []conversionTest{
		// Revit conversion factors (Footcandle is Revit internal unit)
		// 1 Ftc -> 10.7639104167097 lx (from Revit: CF from internal)
		{"Ftc", "lx", 10.7639104167097},
		// 1 lx -> 0.09290304 Ftc (from Revit: CF to internal)
		{"lx", "Ftc", 0.09290304},
	}

	testConversions(t, conversionTests)
}

func Test_Illuminance_UnitSystems(t *testing.T) {
	tests := []unitSystemTest{
		{Lux, SiSystem},
		{Footcandle, BiSystem},
	}
	testUnitSystems(t, tests)
}
