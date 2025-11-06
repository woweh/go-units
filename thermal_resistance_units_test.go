package units

import (
	"testing"

	"github.com/alecthomas/assert/v2"
)

func Test_ThermalResistance_Conversions(t *testing.T) {
	conversionTests := []conversionTest{
		// Revit conversions (from/to internal, see RevitUnits.csv)
		{"(m²·K)/W", "(h·ft²·°F)/BTU", 5.67826334111349},
		{"(h·ft²·°F)/BTU", "(m²·K)/W", 0.176110183682306},
	}
	testConversions(t, conversionTests)
}

func Test_ThermalResistance_UnitSystem(t *testing.T) {
	assert.Equal(t, SiSystem, SquareMeterKelvinPerWatt.System())
	assert.Equal(t, BiSystem, HourSquareFootFahrenheitPerBritishThermalUnit.System())
}

// No metric factories for thermal resistance, so no base unit tests are needed.

func Test_ThermalResistance_Lookup_Names_and_Symbols(t *testing.T) {
	tests := lookUpTests{
		// SI
		{SquareMeterKelvinPerWatt, "(m²·K)/W"},
		{SquareMeterKelvinPerWatt, "square meter kelvin per watt"},
		{SquareMeterKelvinPerWatt, "square meter kelvins per watt"},
		{SquareMeterKelvinPerWatt, "square metre kelvin per watt"},
		{SquareMeterKelvinPerWatt, "square metre kelvins per watt"},
		{SquareMeterKelvinPerWatt, "(m2*K)/W"},
		{SquareMeterKelvinPerWatt, "m2*K/W"},
		// Imperial/US
		{HourSquareFootFahrenheitPerBritishThermalUnit, "(h·ft²·°F)/BTU"},
		{HourSquareFootFahrenheitPerBritishThermalUnit, "hour square foot degree Fahrenheit per British thermal unit"},
		{HourSquareFootFahrenheitPerBritishThermalUnit, "hour square foot degrees Fahrenheit per British thermal unit"},
		{HourSquareFootFahrenheitPerBritishThermalUnit, "h*ft2*F/BTU"},
		{HourSquareFootFahrenheitPerBritishThermalUnit, "(h·ft²·°F)/Btu"},
	}
	testLookupNamesAndSymbols(t, tests)
}
