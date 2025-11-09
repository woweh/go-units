package units

import (
	"testing"
)

func Test_SpecificHeat_Conversions(t *testing.T) {
	conversionTests := []conversionTest{
		// Library and Revit conversions (from internal and to internal)
		{"J/(g·°C)", "J/(kg·°C)", 1000},
		{"J/(kg·°C)", "J/(g·°C)", 0.001},
		{"BTU/(lb·°F)", "J/(kg·°C)", 4186.8},
		{"J/(kg·°C)", "BTU/(lb·°F)", 0.0002388459},
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
