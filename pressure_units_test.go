package units

import (
	"testing"
)

func Test_Pressure_Conversions(t *testing.T) {
	conversionTests := []conversionTest{
		// Pascal to metric progressions
		{"pascal", "kilopascal", 0.001},
		{"pascal", "megapascal", 1e-6},
		{"pascal", "hectopascal", 0.01},
		{"kilopascal", "pascal", 1000},
		{"megapascal", "kilopascal", 1000},

		// Pascal to non-SI units
		{"pascal", "technical atmosphere", 1.01971621297793e-5},
		{"pascal", "standard atmosphere", 9.869232667160128e-6},
		{"pascal", "bar", 1e-5},
		{"pascal", "barye", 10},
		{"pascal", "newton per square meter", 1},
		{"pascal", "pound-force per square inch", 0.00014503773773020924},
		{"pascal", "torr", 0.007500616827041698},

		// Technical atmosphere conversions
		{"technical atmosphere", "pascal", 98066.5},
		{"technical atmosphere", "standard atmosphere", 0.9678411053541968},
		{"technical atmosphere", "bar", 0.980665},
		{"technical atmosphere", "pound-force per square inch", 14.22334331221469},

		// Standard atmosphere conversions
		{"standard atmosphere", "pascal", 101325},
		{"standard atmosphere", "bar", 1.01325},
		{"standard atmosphere", "pound-force per square inch", 14.6959487755134},
		{"standard atmosphere", "torr", 760},

		// Bar conversions
		{"bar", "pascal", 100000},
		{"bar", "millibar", 1000},
		{"bar", "pound-force per square inch", 14.503773773020924},
		{"barye", "pascal", 0.1},

		// Water column units
		{"inch of Water Column", "pascal", 249.0889},
		{"meter of Water Column", "pascal", 9806.65},
		{"millimeter of Water Column", "pascal", 9.80665},
		{"meter of Water Column", "inch of Water Column", 39.37007874015748},

		// Mercury column units
		{"inch of Mercury", "pascal", 3386.389},
		{"meter of Mercury", "pascal", 133322.4},
		{"millimeter of Mercury", "pascal", 133.3224},
		{"meter of Mercury", "millimeter of Mercury", 1000},
		{"meter of Mercury", "inch of Mercury", 39.37007874015748},

		// PSI conversions
		{"pound-force per square inch", "pascal", 6894.757},
		{"pound-force per square inch", "bar", 0.06894757},
		{"pound-force per square inch", "inch of Mercury", 2.03602},

		// Torr conversions
		{"torr", "pascal", 133.3224},
		{"torr", "millimeter of Mercury", 1},

		// Additional units
		{"foot of Water Column", "pascal", 2989.06692},
		{"inches of Water Column", "pascal", 249.0889},
		{"pound-force per square inch gauge", "pascal", 6894.757},
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
