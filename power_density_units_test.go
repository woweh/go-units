package units

import (
	"testing"

	"github.com/alecthomas/assert/v2"
)

func Test_PowerDensity_Conversions(t *testing.T) {
	conversionTests := []conversionTest{
		// Revit conversions from internal (W/m²)
		{"W/m²", "W/ft²", 0.09290304},
		{"W/ft²", "W/m²", 10.7639104167097},
		{"W/m²", "Btu/(h·ft²)", 0.316998330628151},
		{"Btu/(h·ft²)", "W/m²", 3.15459074506305},

		// Cross conversions
		{"W/ft²", "Btu/(h·ft²)", 3.41214},
		{"Btu/(h·ft²)", "W/ft²", 0.29307107017222},
	}
	testConversions(t, conversionTests)
}

func Test_PowerDensity_UnitSystems(t *testing.T) {
	assert.Equal(t, SiSystem, WattPerSquareMeter.System())
	assert.Equal(t, BiSystem, WattPerSquareFoot.System())
	assert.Equal(t, BiSystem, BritishThermalUnitPerHourSquareFoot.System())
}

func Test_Lookup_PowerDensity_Names_and_Symbols(t *testing.T) {
	tests := lookUpTests{
		{WattPerSquareMeter, "watt per square meter"},
		{WattPerSquareMeter, "W/m²"},
		{WattPerSquareMeter, "watts per square meter"},
		{WattPerSquareMeter, "watt per square metre"},
		{WattPerSquareMeter, "watts per square metre"},
		{WattPerSquareMeter, "W/m2"},
		{WattPerSquareFoot, "watt per square foot"},
		{WattPerSquareFoot, "W/ft²"},
		{WattPerSquareFoot, "watts per square foot"},
		{WattPerSquareFoot, "W/ft2"},
		{BritishThermalUnitPerHourSquareFoot, "British thermal unit per hour square foot"},
		{BritishThermalUnitPerHourSquareFoot, "Btu/(h·ft²)"},
		{BritishThermalUnitPerHourSquareFoot, "British thermal units per hour square foot"},
		{BritishThermalUnitPerHourSquareFoot, "BTU/h/ft²"},
		{BritishThermalUnitPerHourSquareFoot, "BTU/(h·ft²)"},
	}
	testLookupNamesAndSymbols(t, tests)
}
