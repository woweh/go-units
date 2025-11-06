package units

import (
	"testing"

	"github.com/alecthomas/assert/v2"
)

func Test_Power_Conversions(t *testing.T) {
	conversionTests := []conversionTest{
		// SI base to metric
		{"watt", "milliwatt", 1000},
		{"watt", "kilowatt", 0.001},
		{"watt", "megawatt", 1e-6},
		{"kilowatt", "watt", 1000},
		{"megawatt", "kilowatt", 1000},

		// SI base to volt-ampere/VAR
		{"watt", "volt-ampere", 1},
		{"watt", "volt-ampere reactive", 1},
		{"kilowatt", "kilovolt-ampere", 1},

		// SI base to Imperial/US (internal conversions)
		{"watt", "British thermal unit per hour", 3.412141633127942},
		{"British thermal unit per hour", "watt", 0.2930710701722222},
		{"watt", "British thermal unit per second", 0.0009478171203133172},
		{"British thermal unit per second", "watt", 1055.05585262},
		{"watt", "calorie per second", 0.2390057361376673},
		{"calorie per second", "watt", 4.184},
		{"watt", "horsepower", 0.0013410220895950279},
		{"horsepower", "watt", 745.6998715822702},
		{"watt", "kilocalorie per second", 0.0002390057361376673},
		{"kilocalorie per second", "watt", 4184},
		{"watt", "thousand British thermal units per hour", 0.003412141633127942},
		{"thousand British thermal units per hour", "watt", 293.0710701722222},

		// Revit conversions (from RevitUnits.csv, ToInternal column)
		{"British thermal unit per hour", "watt", 0.316998330628151},
		{"British thermal unit per second", "watt", 0.0000880550918411529},
		{"calorie per second", "watt", 0.0221895098882201},
		{"horsepower", "watt", 0.00012458502883053},
		{"kilocalorie per second", "watt", 0.0000221895098882201},
		{"kilovolt-ampere", "watt", 0.0000929030400000000},
		{"kilowatt", "watt", 0.0000929030400000000},
		{"thousand British thermal units per hour", "watt", 0.000316998330628151},
		{"volt-ampere", "watt", 0.09290304},
	}
	testConversions(t, conversionTests)
}

func Test_Power_UnitSystems(t *testing.T) {
	si := SiSystem
	bi := BiSystem
	assert.Equal(t, si, Watt.System())
	assert.Equal(t, si, MilliWatt.System())
	assert.Equal(t, si, KiloWatt.System())
	assert.Equal(t, si, MegaWatt.System())
	assert.Equal(t, si, VoltAmpere.System())
	assert.Equal(t, si, KiloVoltAmpere.System())
	assert.Equal(t, si, MegaVoltAmpere.System())
	assert.Equal(t, si, VoltAmpereReactive.System())
	assert.Equal(t, si, KiloVoltAmpereReactive.System())
	assert.Equal(t, si, MegaVoltAmpereReactive.System())
	assert.Equal(t, si, CaloriePerSecond.System())
	assert.Equal(t, si, KiloCaloriePerSecond.System())
	assert.Equal(t, bi, BritishThermalUnitPerHour.System())
	assert.Equal(t, bi, BritishThermalUnitPerSecond.System())
	assert.Equal(t, bi, Horsepower.System())
	assert.Equal(t, bi, ThousandBritishThermalUnitsPerHour.System())
}

func Test_Power_BaseUnits(t *testing.T) {
	// Only a few representative metric units
	assert.Equal(t, Watt, MilliWatt.Base())
	assert.Equal(t, Watt, KiloWatt.Base())
	assert.Equal(t, Watt, MegaWatt.Base())
	assert.Equal(t, VoltAmpere, KiloVoltAmpere.Base())
	assert.Equal(t, VoltAmpereReactive, KiloVoltAmpereReactive.Base())
}

func Test_Lookup_Power_Names_and_Symbols(t *testing.T) {
	tests := lookUpTests{
		{Watt, "watt"}, {Watt, "W"},
		{KiloWatt, "kilowatt"}, {KiloWatt, "kW"},
		{MegaWatt, "megawatt"}, {MegaWatt, "MW"},
		{MilliWatt, "milliwatt"}, {MilliWatt, "mW"},
		{VoltAmpere, "volt-ampere"}, {VoltAmpere, "volt ampere"}, {VoltAmpere, "VA"}, {VoltAmpere, "V*A"}, {VoltAmpere, "V.A"}, {VoltAmpere, "V A"},
		{KiloVoltAmpere, "kilovolt-ampere"}, {KiloVoltAmpere, "kVA"},
		{VoltAmpereReactive, "volt-ampere reactive"}, {VoltAmpereReactive, "volt ampere reactive"}, {VoltAmpereReactive, "VAR"}, {VoltAmpereReactive, "VAr"}, {VoltAmpereReactive, "V⋅Ar"}, {VoltAmpereReactive, "V.A.r"}, {VoltAmpereReactive, "V A r"},
		{KiloVoltAmpereReactive, "kilovolt-ampere reactive"}, {KiloVoltAmpereReactive, "kVAR"},
		{BritishThermalUnitPerHour, "British thermal unit per hour"}, {BritishThermalUnitPerHour, "Btu/h"},
		{BritishThermalUnitPerSecond, "British thermal unit per second"}, {BritishThermalUnitPerSecond, "Btu/s"},
		{CaloriePerSecond, "calorie per second"}, {CaloriePerSecond, "cal/s"},
		{KiloCaloriePerSecond, "kilocalorie per second"}, {KiloCaloriePerSecond, "kcal/s"},
		{Horsepower, "horsepower"}, {Horsepower, "hp"},
		{ThousandBritishThermalUnitsPerHour, "thousand British thermal units per hour"}, {ThousandBritishThermalUnitsPerHour, "MBH"},
	}
	testLookupNamesAndSymbols(t, tests)
}
