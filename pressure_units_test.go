package units

import (
	"testing"
)

func Test_Pressure_Conversions(t *testing.T) {
	conversionTests := []conversionTest{
		// Pascal to metric progressions
		{from: "pascal", to: "kilopascal", exp: 0.001},
		{from: "pascal", to: "megapascal", exp: 1e-6},
		{from: "pascal", to: "hectopascal", exp: 0.01},
		{from: "kilopascal", to: "pascal", exp: 1000},
		{from: "megapascal", to: "kilopascal", exp: 1000},

		// Pascal to non-SI units
		{from: "pascal", to: "technical atmosphere", exp: 1.01971621297793e-5},
		{from: "pascal", to: "standard atmosphere", exp: 9.869232667160128e-6},
		{from: "pascal", to: "bar", exp: 1e-5},
		{from: "pascal", to: "barye", exp: 10},
		{from: "pascal", to: "newton per square meter", exp: 1},
		{from: "pascal", to: "pound-force per square inch", exp: 0.00014503773773020924},
		{from: "pascal", to: "torr", exp: 0.007500616827041698},

		// Technical atmosphere conversions
		{from: "technical atmosphere", to: "pascal", exp: 98066.5},
		{from: "technical atmosphere", to: "standard atmosphere", exp: 0.9678411053541968},
		{from: "technical atmosphere", to: "bar", exp: 0.980665},
		{from: "technical atmosphere", to: "pound-force per square inch", exp: 14.22334331221469},

		// Standard atmosphere conversions
		{from: "standard atmosphere", to: "pascal", exp: 101325},
		{from: "standard atmosphere", to: "bar", exp: 1.01325},
		{from: "standard atmosphere", to: "pound-force per square inch", exp: 14.6959487755134},
		{from: "standard atmosphere", to: "torr", exp: 760},

		// Bar conversions
		{from: "bar", to: "pascal", exp: 100000},
		{from: "bar", to: "millibar", exp: 1000},
		{from: "bar", to: "pound-force per square inch", exp: 14.503773773020924},
		{from: "barye", to: "pascal", exp: 0.1},

		// Water column units
		{from: "inch of Water Column", to: "pascal", exp: 249.0889},
		{from: "meter of Water Column", to: "pascal", exp: 9806.65},
		{from: "millimeter of Water Column", to: "pascal", exp: 9.80665},
		{from: "meter of Water Column", to: "inch of Water Column", exp: 39.37007874015748},

		// Mercury column units
		{from: "inch of Mercury", to: "pascal", exp: 3386.389},
		{from: "meter of Mercury", to: "pascal", exp: 133322.4},
		{from: "millimeter of Mercury", to: "pascal", exp: 133.3224},
		{from: "meter of Mercury", to: "millimeter of Mercury", exp: 1000},
		{from: "meter of Mercury", to: "inch of Mercury", exp: 39.37007874015748},

		// PSI conversions
		{from: "pound-force per square inch", to: "pascal", exp: 6894.757},
		{from: "pound-force per square inch", to: "bar", exp: 0.06894757},
		{from: "pound-force per square inch", to: "inch of Mercury", exp: 2.03602},

		// Torr conversions
		{from: "torr", to: "pascal", exp: 133.3224},
		{from: "torr", to: "millimeter of Mercury", exp: 1},

		// Additional units
		{from: "foot of Water Column", to: "pascal", exp: 2989.06692},
		{from: "inches of Water Column", to: "pascal", exp: 249.0889},
		{from: "pound-force per square inch gauge", to: "pascal", exp: 6894.757},
	}
	testConversions(t, conversionTests)
}

func Test_Pressure_UnitSystems(t *testing.T) {
	tests := []unitSystemTest{
		{Pascal, SiSystem},
		{KiloPascal, SiSystem},
		{MegaPascal, SiSystem},
		{HectoPascal, SiSystem},
		{Bar, NoSystem},
		{Barye, NoSystem},
		{Torr, NoSystem},
	}
	testUnitSystems(t, tests)
}
