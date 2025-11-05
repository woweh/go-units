package units

import (
	"testing"

	"github.com/alecthomas/assert/v2"
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
	si := SiSystem
	assert.Equal(t, si, Joule.System())
	assert.Equal(t, si, KiloJoule.System())
	assert.Equal(t, si, MegaJoule.System())
	assert.Equal(t, si, GigaJoule.System())
	assert.Equal(t, si, TeraJoule.System())
	assert.Equal(t, si, PetaJoule.System())
	assert.Equal(t, si, ExaJoule.System())
	assert.Equal(t, si, ZettaJoule.System())
	assert.Equal(t, si, YottaJoule.System())
	assert.Equal(t, si, MilliJoule.System())
	assert.Equal(t, si, MicroJoule.System())
	assert.Equal(t, si, NanoJoule.System())
	assert.Equal(t, si, PicoJoule.System())
	assert.Equal(t, si, FemtoJoule.System())
	assert.Equal(t, si, AttoJoule.System())
	assert.Equal(t, si, WattHour.System())
	assert.Equal(t, si, KiloWattHour.System())
	assert.Equal(t, si, MegaWattHour.System())
	assert.Equal(t, si, GigaWattHour.System())
	assert.Equal(t, si, TeraWattHour.System())
	assert.Equal(t, si, PetaWattHour.System())
	assert.Equal(t, si, ElectronVolt.System())
	assert.Equal(t, si, KiloElectronVolt.System())
	assert.Equal(t, si, MegaElectronVolt.System())
	assert.Equal(t, si, GigaElectronVolt.System())
	assert.Equal(t, si, Calorie.System())
	assert.Equal(t, si, KiloCalorie.System())
	bi := BiSystem
	assert.Equal(t, bi, BritishThermalUnit.System())
	assert.Equal(t, bi, Therm.System())
}

func Test_Energy_BaseUnits(t *testing.T) {
	// Only a few representative metric units
	assert.Equal(t, Joule, Joule.Base())
	assert.Equal(t, Joule, KiloJoule.Base())
	assert.Equal(t, Joule, MegaJoule.Base())
	assert.Equal(t, Joule, WattHour.Base())
	assert.Equal(t, Joule, ElectronVolt.Base())
	assert.Equal(t, Joule, Calorie.Base())
}

func Test_Lookup_Energy_Names_and_Symbols(t *testing.T) {
	tests := lookUpTests{
		{Joule, "joule"}, {Joule, "J"}, {Joule, "watt second"}, {Joule, "W⋅s"},
		{KiloJoule, "kilojoule"}, {KiloJoule, "kJ"},
		{MegaJoule, "megajoule"}, {MegaJoule, "MJ"},
		{GigaJoule, "gigajoule"}, {GigaJoule, "GJ"},
		{MilliJoule, "millijoule"}, {MilliJoule, "mJ"},
		{MicroJoule, "microjoule"}, {MicroJoule, "μJ"}, {MicroJoule, "uJ"},
		{WattHour, "watt-hour"}, {WattHour, "Wh"}, {WattHour, "volt ampere hour"}, {WattHour, "VAh"}, {WattHour, "varh"},
		{KiloWattHour, "kilowatt-hour"}, {KiloWattHour, "kWh"},
		{MegaWattHour, "megawatt-hour"}, {MegaWattHour, "MWh"},
		{ElectronVolt, "electronvolt"}, {ElectronVolt, "eV"}, {ElectronVolt, "electron volt"},
		{KiloElectronVolt, "kiloelectronvolt"}, {KiloElectronVolt, "keV"},
		{MegaElectronVolt, "megaelectronvolt"}, {MegaElectronVolt, "MeV"},
		{GigaElectronVolt, "gigaelectronvolt"}, {GigaElectronVolt, "GeV"},
		{Calorie, "calorie"}, {Calorie, "cal"}, {Calorie, "Thermochemical Calorie"}, {Calorie, "calorie (th)"},
		{KiloCalorie, "kilocalorie"}, {KiloCalorie, "kcal"},
		{BritishThermalUnit, "British thermal unit"}, {BritishThermalUnit, "Btu"},
		{Therm, "therm"},
	}
	testLookupNamesAndSymbols(t, tests)
}
