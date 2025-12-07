package units

import (
	"testing"
)

func Test_Energy_Conversions(t *testing.T) {
	conversionTests := []conversionTest{
		// SI base and metric
		{from: "joule", to: "calorie", exp: 0.2390057361376673},
		{from: "joule", to: "electronvolt", exp: 6.241507648655548e+18, tol: fPtr(1e+12)},
		{from: "joule", to: "watt-hour", exp: 0.0002777777777777778},
		{from: "watt-hour", to: "joule", exp: 3600},
		{from: "megawatt-hour", to: "joule", exp: 3.6e+9},
		{from: "kilowatt-hour", to: "joule", exp: 3.6e+6},
		{from: "joule", to: "megajoule", exp: 1e-6},
		{from: "joule", to: "gigajoule", exp: 1e-9},
		// Imperial/other
		{from: "joule", to: "british thermal unit", exp: 0.0009478171203133172},
		{from: "british thermal unit", to: "joule", exp: 1055.05585262},
		{from: "therm", to: "joule", exp: 105505585.262},
		// Additional metric progressions (sampled)
		{from: "kilojoule", to: "joule", exp: 1000},
		{from: "megajoule", to: "kilojoule", exp: 1000},
		{from: "gigajoule", to: "megajoule", exp: 1000},
		{from: "millijoule", to: "joule", exp: 0.001},
		{from: "microjoule", to: "joule", exp: 0.000001},
		// Electronvolt progressions (sampled)
		{from: "kiloelectronvolt", to: "electronvolt", exp: 1000},
		{from: "megaelectronvolt", to: "kiloelectronvolt", exp: 1000},
		{from: "gigaelectronvolt", to: "megaelectronvolt", exp: 1000},
		// Calorie progressions
		{from: "kilocalorie", to: "calorie", exp: 1000},
	}
	testConversions(t, conversionTests)
}

func Test_Energy_UnitSystems(t *testing.T) {
	tests := []unitSystemTest{
		{Joule, SiSystem},
		{KiloJoule, SiSystem},
		{MegaJoule, SiSystem},
		{GigaJoule, SiSystem},
		{TeraJoule, SiSystem},
		{PetaJoule, SiSystem},
		{ExaJoule, SiSystem},
		{ZettaJoule, SiSystem},
		{YottaJoule, SiSystem},
		{MilliJoule, SiSystem},
		{MicroJoule, SiSystem},
		{NanoJoule, SiSystem},
		{PicoJoule, SiSystem},
		{FemtoJoule, SiSystem},
		{AttoJoule, SiSystem},
		{WattHour, SiSystem},
		{KiloWattHour, SiSystem},
		{MegaWattHour, SiSystem},
		{GigaWattHour, SiSystem},
		{TeraWattHour, SiSystem},
		{Calorie, NoSystem},
		{KiloCalorie, NoSystem},
		{BritishThermalUnit, NoSystem},
		{Therm, NoSystem},
		{ElectronVolt, NoSystem},
		{KiloElectronVolt, NoSystem},
		{MegaElectronVolt, NoSystem},
		{GigaElectronVolt, NoSystem},
	}
	testUnitSystems(t, tests)
}
