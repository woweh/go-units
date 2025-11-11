package units

import (
	"testing"
)

func Test_Power_Conversions(t *testing.T) {
	conversionTests := []conversionTest{
		// SI base to metric
		{from: "watt", to: "milliwatt", exp: 1000},
		{from: "watt", to: "kilowatt", exp: 0.001},
		{from: "watt", to: "megawatt", exp: 1e-6},
		{from: "kilowatt", to: "watt", exp: 1000},
		{from: "megawatt", to: "kilowatt", exp: 1000},

		// SI base to volt-ampere/VAR
		{from: "watt", to: "volt-ampere", exp: 1},
		{from: "watt", to: "volt-ampere reactive", exp: 1},
		{from: "kilowatt", to: "kilovolt-ampere", exp: 1},

		// SI base to Imperial/US (internal conversions)
		{from: "watt", to: "British thermal unit per hour", exp: 3.412141633127942},
		{from: "British thermal unit per hour", to: "watt", exp: 0.2930710701722222},
		{from: "watt", to: "British thermal unit per second", exp: 0.0009478171203133172},
		{from: "British thermal unit per second", to: "watt", exp: 1055.05585262},
		{from: "watt", to: "calorie per second", exp: 0.2390057361376673},
		{from: "calorie per second", to: "watt", exp: 4.184},
		{from: "watt", to: "horsepower", exp: 0.0013410220895950279},
		{from: "horsepower", to: "watt", exp: 745.6998715822702},
		{from: "watt", to: "kilocalorie per second", exp: 0.0002390057361376673},
		{from: "kilocalorie per second", to: "watt", exp: 4184},
		{from: "watt", to: "thousand British thermal units per hour", exp: 0.003412141633127942},
		{from: "thousand British thermal units per hour", to: "watt", exp: 293.0710701722222},

		// Revit conversions (from RevitUnits.csv, ToInternal column)
		{from: "British thermal unit per hour", to: "watt", exp: 0.316998330628151},
		{from: "British thermal unit per second", to: "watt", exp: 0.0000880550918411529},
		{from: "calorie per second", to: "watt", exp: 0.0221895098882201},
		{from: "horsepower", to: "watt", exp: 0.00012458502883053},
		{from: "kilocalorie per second", to: "watt", exp: 0.0000221895098882201},
		{from: "kilovolt-ampere", to: "watt", exp: 0.0000929030400000000},
		{from: "kilowatt", to: "watt", exp: 0.0000929030400000000},
		{from: "thousand British thermal units per hour", to: "watt", exp: 0.000316998330628151},
		{from: "volt-ampere", to: "watt", exp: 0.09290304},
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
