package units

import (
	"testing"

	"github.com/alecthomas/assert/v2"
)

func Test_CoefficientOfHeatTransfer_Conversions(t *testing.T) {
	conversionTests := []conversionTest{
		// Base conversions (Revit/implementation)
		{"BTU/(h·ft²·°F)", "W/(m²·K)", "5.67826"},
		{"W/(m²·K)", "BTU/(h·ft²·°F)", "0.1761103"},
		// Identity
		{"W/(m²·K)", "W/(m²·K)", "1"},
		{"BTU/(h·ft²·°F)", "BTU/(h·ft²·°F)", "1"},
	}
	testConversions(t, conversionTests)
}

func Test_CoefficientOfHeatTransfer_UnitSystems(t *testing.T) {
	assert.Equal(t, SiSystem, WattPerSquareMeterKelvin.System())
	assert.NotEqual(t, SiSystem, BritishThermalUnitPerHourSquareFootFahrenheit.System())
}

func Test_CoefficientOfHeatTransfer_BaseUnits(t *testing.T) {
	assert.Equal(t, WattPerSquareMeterKelvin, WattPerSquareMeterKelvin.Base())
	assert.Equal(t, WattPerSquareMeterKelvin, BritishThermalUnitPerHourSquareFootFahrenheit.Base())
}

func Test_Lookup_CoefficientOfHeatTransfer_Names_and_Symbols(t *testing.T) {
	tests := lookUpTests{
		{WattPerSquareMeterKelvin, "watt per square meter kelvin"},
		{WattPerSquareMeterKelvin, "W/(m²·K)"},
		{WattPerSquareMeterKelvin, "W/(m2*K)"},
		{WattPerSquareMeterKelvin, "W/m2/K"},
		{BritishThermalUnitPerHourSquareFootFahrenheit, "British thermal unit per hour square foot degree Fahrenheit"},
		{BritishThermalUnitPerHourSquareFootFahrenheit, "BTU/(h·ft²·°F)"},
		{BritishThermalUnitPerHourSquareFootFahrenheit, "BTU/h/ft2/F"},
	}
	testLookupNamesAndSymbols(t, tests)
}
