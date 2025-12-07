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

		// SI base to Imperial/US (using Revit conversion factors)
		{from: "watt", to: "British thermal unit per hour", exp: 3.15459074506305, tol: fPtr(1e-8)},
		{from: "British thermal unit per hour", to: "watt", exp: 0.316998330628151, tol: fPtr(1e-9)},
		{from: "watt", to: "British thermal unit per second", exp: 8.80550918411529e-05, tol: fPtr(1e-14)},
		{from: "British thermal unit per second", to: "watt", exp: 11356.526682227, tol: fPtr(1e-6)},
		{from: "watt", to: "calorie per second", exp: 0.0221895098882201, tol: fPtr(1e-12)},
		{from: "calorie per second", to: "watt", exp: 45.0663401326803, tol: fPtr(1e-8)},
		{from: "watt", to: "horsepower", exp: 0.00012458502883053, tol: fPtr(1e-14)},
		{from: "horsepower", to: "watt", exp: 8026.6466154635, tol: fPtr(1e-7)},
		{from: "watt", to: "kilocalorie per second", exp: 2.21895098882201e-05, tol: fPtr(1e-14)},
		{from: "kilocalorie per second", to: "watt", exp: 45066.3401326803, tol: fPtr(1e-6)},
		{from: "watt", to: "thousand British thermal units per hour", exp: 0.00315459074506305, tol: fPtr(1e-11)},
		{from: "thousand British thermal units per hour", to: "watt", exp: 316.998330628151, tol: fPtr(1e-6)},
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
