package units

import (
	"testing"

	"github.com/alecthomas/assert/v2"
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
	assert.Equal(t, SiSystem, WattPerSquareMeterKelvin.System())
	assert.Equal(t, BiSystem, BritishThermalUnitPerHourSquareFootFahrenheit.System())
}

func Test_CoefficientOfHeatTransfer_BaseUnits(t *testing.T) {
	assert.Equal(t, WattPerSquareMeterKelvin, WattPerSquareMeterKelvin.Base())
	assert.Equal(t, WattPerSquareMeterKelvin, BritishThermalUnitPerHourSquareFootFahrenheit.Base())
}

func Test_Lookup_CoefficientOfHeatTransfer_Names_and_Symbols(t *testing.T) {
	tests := lookUpTests{
		// SI/metric
		{WattPerSquareMeterKelvin, "watt per square meter kelvin"},
		{WattPerSquareMeterKelvin, "watts per square meter kelvin"},
		{WattPerSquareMeterKelvin, "watt per square metre kelvin"},
		{WattPerSquareMeterKelvin, "watts per square metre kelvin"},
		{WattPerSquareMeterKelvin, "W/(m²·K)"},
		{WattPerSquareMeterKelvin, "W/(m2*K)"},
		{WattPerSquareMeterKelvin, "W/m2/K"},
		// Imperial/US
		{BritishThermalUnitPerHourSquareFootFahrenheit, "British thermal unit per hour square foot degree Fahrenheit"},
		{BritishThermalUnitPerHourSquareFootFahrenheit, "British thermal units per hour square foot degree Fahrenheit"},
		{BritishThermalUnitPerHourSquareFootFahrenheit, "BTU/(h·ft²·°F)"},
		{BritishThermalUnitPerHourSquareFootFahrenheit, "BTU/h/ft2/F"},
		{BritishThermalUnitPerHourSquareFootFahrenheit, "Btu/(h·ft²·°F)"},
	}
	testLookupNamesAndSymbols(t, tests)
}
