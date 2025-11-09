package units

import (
	"testing"
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
	tests := []unitSystemTest{
		{Watt, SiSystem},
		{MilliWatt, SiSystem},
		{KiloWatt, SiSystem},
		{MegaWatt, SiSystem},
		{VoltAmpere, SiSystem},
		{MegaVoltAmpere, SiSystem},
		{VoltAmpereReactive, SiSystem},
		{BritishThermalUnitPerHour, BiSystem},
		{BritishThermalUnitPerSecond, BiSystem},
		{CaloriePerSecond, NoSystem},
		{Horsepower, NoSystem},
		{ThousandBritishThermalUnitsPerHour, BiSystem},
	}
	testUnitSystems(t, tests)
}
