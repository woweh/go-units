package units

import (
	"testing"
)

func Test_ThermalConductivity_Conversions(t *testing.T) {
	conversionTests := []conversionTest{
		// Revit conversions (from/to internal, see RevitUnits.csv)
		{from: "BTU/(h·ft·°F)", to: "W/(m·K)", exp: 5.67826334111349},
		{from: "W/(m·K)", to: "BTU/(h·ft·°F)", exp: 0.176110183682306},
		// Library ratio
		{from: "BTU/(h·ft·°F)", to: "W/(m·K)", exp: 1.730735},
		{from: "W/(m·K)", to: "BTU/(h·ft·°F)", exp: 0.577789},
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
