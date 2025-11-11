package units

import (
	"testing"
)

func Test_CoefficientOfHeatTransfer_Conversions(t *testing.T) {
	conversionTests := []conversionTest{
		// Revit and internal conversion factors
		{from: "BTU/(h·ft²·°F)", to: "W/(m²·K)", exp: 5.67826334111349},
		{from: "W/(m²·K)", to: "BTU/(h·ft²·°F)", exp: 0.176110183682306},
		// SI base unit identity (from Revit, 1:1)
		{from: "W/(m²·K)", to: "W/(m²·K)", exp: 1},
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
