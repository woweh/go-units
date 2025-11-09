package units

import (
	"testing"
)

func Test_Energy_Conversions(t *testing.T) {
	conversionTests := []conversionTest{
		// SI base and metric
		{"joule", "calorie", 0.2390057361376673},
		{"joule", "electronvolt", 6.241507648655548e+18},
		{"joule", "watt-hour", 0.0002777777777777778},
		{"watt-hour", "joule", 3600},
		{"megawatt-hour", "joule", 3.6e+9},
		{"kilowatt-hour", "joule", 3.6e+6},
		{"joule", "megajoule", 1e-6},
		{"joule", "gigajoule", 1e-9},
		// Imperial/other
		{"joule", "british thermal unit", 0.0009478171203133172},
		{"british thermal unit", "joule", 1055.05585262},
		{"therm", "joule", 105505585.262},
		// Additional metric progressions (sampled)
		{"kilojoule", "joule", 1000},
		{"megajoule", "kilojoule", 1000},
		{"gigajoule", "megajoule", 1000},
		{"millijoule", "joule", 0.001},
		{"microjoule", "joule", 0.000001},
		// Electronvolt progressions (sampled)
		{"kiloelectronvolt", "electronvolt", 1000},
		{"megaelectronvolt", "kiloelectronvolt", 1000},
		{"gigaelectronvolt", "megaelectronvolt", 1000},
		// Calorie progressions
		{"kilocalorie", "calorie", 1000},
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
