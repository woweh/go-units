package units

import (
	"testing"

	"github.com/alecthomas/assert/v2"
)

func Test_ThermalConductivity_Conversions(t *testing.T) {
	conversionTests := []conversionTest{
		// Revit conversions (from/to internal, see RevitUnits.csv)
		{"BTU/(h·ft·°F)", "W/(m·K)", 5.67826334111349},
		{"W/(m·K)", "BTU/(h·ft·°F)", 0.176110183682306},
		// Library ratio
		{"BTU/(h·ft·°F)", "W/(m·K)", 1.730735},
		{"W/(m·K)", "BTU/(h·ft·°F)", 0.577789},
	}
	testConversions(t, conversionTests)
}

func Test_ThermalConductivity_UnitSystem(t *testing.T) {
	assert.Equal(t, SiSystem, WattPerMeterKelvin.System())
	assert.Equal(t, BiSystem, BritishThermalUnitPerHourFootFahrenheit.System())
}

// No metric factories for thermal conductivity, so no base unit tests are needed.

func Test_ThermalConductivity_Lookup_Names_and_Symbols(t *testing.T) {
	tests := lookUpTests{
		{WattPerMeterKelvin, "W/(m·K)"},
		{WattPerMeterKelvin, "watts per meter kelvin"},
		{WattPerMeterKelvin, "watt per metre kelvin"},
		{WattPerMeterKelvin, "watts per metre kelvin"},
		{WattPerMeterKelvin, "W/(m*K)"},
		{WattPerMeterKelvin, "W/m/K"},
		{BritishThermalUnitPerHourFootFahrenheit, "BTU/(h·ft·°F)"},
		{BritishThermalUnitPerHourFootFahrenheit, "British thermal unit per hour foot degree Fahrenheit"},
		{BritishThermalUnitPerHourFootFahrenheit, "British thermal units per hour foot degree Fahrenheit"},
		{BritishThermalUnitPerHourFootFahrenheit, "BTU/h/ft/F"},
		{BritishThermalUnitPerHourFootFahrenheit, "Btu/(h·ft·°F)"},
	}
	testLookupNamesAndSymbols(t, tests)
}
