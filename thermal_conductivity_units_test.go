package units

import (
	"testing"
)

func Test_ThermalConductivity_Conversions(t *testing.T) {
	conversionTests := []conversionTest{
		// Revit conversions (from/to internal, see RevitUnits.csv)
		{from: "BTU/(h·ft·°F)", to: "W/(m·K)", exp: 1.730734666, tol: fPtr(1e-8)},
		{from: "W/(m·K)", to: "BTU/(h·ft·°F)", exp: 0.577789205, tol: fPtr(1e-6)},
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
