package units

import (
	"testing"
)

func Test_SpecificHeat_Conversions(t *testing.T) {
	conversionTests := []conversionTest{
		// Library and Revit conversions (from internal and to internal)
		{from: "J/(g·°C)", to: "J/(kg·°C)", exp: 1000},
		{from: "J/(kg·°C)", to: "J/(g·°C)", exp: 0.001},
		{from: "BTU/(lb·°F)", to: "J/(kg·°C)", exp: 4186.8},
		{from: "J/(kg·°C)", to: "BTU/(lb·°F)", exp: 0.0002388459},
	}
	testConversions(t, conversionTests)
}

func Test_SpecificHeat_UnitSystems(t *testing.T) {
	tests := []unitSystemTest{
		{JoulePerKilogramCelsius, SiSystem},
		{JoulePerGramCelsius, SiSystem},
		{BritishThermalUnitPerPoundFahrenheit, BiSystem},
	}
	testUnitSystems(t, tests)
}
