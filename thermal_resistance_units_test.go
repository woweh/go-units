package units

import (
	"testing"
)

func Test_ThermalResistance_Conversions(t *testing.T) {
	conversionTests := []conversionTest{
		// Revit conversions (from/to internal, see RevitUnits.csv)
		{from: "(m²·K)/W", to: "(h·ft²·°F)/BTU", exp: 5.67826334111349},
		{from: "(h·ft²·°F)/BTU", to: "(m²·K)/W", exp: 0.176110183682306},
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
