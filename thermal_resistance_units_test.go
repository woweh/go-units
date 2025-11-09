package units

import (
	"testing"
)

func Test_ThermalResistance_Conversions(t *testing.T) {
	conversionTests := []conversionTest{
		// Revit conversions (from/to internal, see RevitUnits.csv)
		{"(m²·K)/W", "(h·ft²·°F)/BTU", 5.67826334111349},
		{"(h·ft²·°F)/BTU", "(m²·K)/W", 0.176110183682306},
	}
	testConversions(t, conversionTests)
}

func Test_ThermalResistance_UnitSystems(t *testing.T) {
	tests := []unitSystemTest{
		{SquareMeterKelvinPerWatt, SiSystem},
		{HourSquareFootFahrenheitPerBritishThermalUnit, BiSystem},
	}
	testUnitSystems(t, tests)
}
