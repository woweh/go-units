package units

import (
	"testing"
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

func Test_ThermalConductivity_UnitSystems(t *testing.T) {
	tests := []unitSystemTest{
		{WattPerMeterKelvin, SiSystem},
		{BritishThermalUnitPerHourFootFahrenheit, BiSystem},
	}
	testUnitSystems(t, tests)
}
