package units

import (
	"testing"
)

func Test_CoefficientOfHeatTransfer_Conversions(t *testing.T) {
	conversionTests := []conversionTest{
		// Revit and internal conversion factors
		{"BTU/(h·ft²·°F)", "W/(m²·K)", 5.67826334111349},
		{"W/(m²·K)", "BTU/(h·ft²·°F)", 0.176110183682306},
		// SI base unit identity (from Revit, 1:1)
		{"W/(m²·K)", "W/(m²·K)", 1},
	}
	testConversions(t, conversionTests)
}

func Test_CoefficientOfHeatTransfer_UnitSystems(t *testing.T) {
	tests := []unitSystemTest{
		{WattPerSquareMeterKelvin, SiSystem},
		{BritishThermalUnitPerHourSquareFootFahrenheit, BiSystem},
	}
	testUnitSystems(t, tests)
}
