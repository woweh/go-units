package units

import (
	"testing"

	"github.com/alecthomas/assert/v2"
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
	si := SiSystem
	bi := BiSystem

	// Pascal family (SI)
	assert.Equal(t, si, Pascal.System())
	assert.Equal(t, si, KiloPascal.System())
	assert.Equal(t, si, MegaPascal.System())
	assert.Equal(t, si, HectoPascal.System())
	assert.Equal(t, si, MilliPascal.System())
	assert.Equal(t, si, NewtonSqm.System())
	assert.Equal(t, si, KiloNewtonSqm.System())

	// Bar family (BI but metric-based)
	assert.Equal(t, bi, Bar.System())
	assert.Equal(t, bi, MilliBar.System())
	assert.Equal(t, bi, CentiBar.System())
	assert.Equal(t, bi, MicroBar.System())

	// Atmospheres (BI)
	assert.Equal(t, bi, At.System())
	assert.Equal(t, bi, Atm.System())
	assert.Equal(t, bi, Barye.System())

	// Water and Mercury columns (BI)
	assert.Equal(t, bi, InH2O.System())
	assert.Equal(t, bi, InHg.System())
	assert.Equal(t, bi, MH2O.System())
	assert.Equal(t, bi, MilliMH2O.System())
	assert.Equal(t, bi, CentiMH2O.System())
	assert.Equal(t, bi, MHg.System())
	assert.Equal(t, bi, MilliMHg.System())
	assert.Equal(t, bi, CentiMHg.System())

	// Imperial/US (BI)
	assert.Equal(t, bi, Psi.System())
	assert.Equal(t, bi, Torr.System())
	assert.Equal(t, bi, FootH2O.System())
	assert.Equal(t, bi, InchesOfWater.System())
	assert.Equal(t, bi, PoundForcePerSquareInchGauge.System())
}

func Test_Pressure_BaseUnits(t *testing.T) {
	// Pascal metric family
	assert.Equal(t, Pascal, KiloPascal.Base())
	assert.Equal(t, Pascal, MegaPascal.Base())
	assert.Equal(t, Pascal, MilliPascal.Base())

	// Bar metric family
	assert.Equal(t, Bar, MilliBar.Base())
	assert.Equal(t, Bar, CentiBar.Base())
	assert.Equal(t, Bar, MicroBar.Base())

	// Water column metric family
	assert.Equal(t, MH2O, MilliMH2O.Base())
	assert.Equal(t, MH2O, CentiMH2O.Base())

	// Mercury column metric family
	assert.Equal(t, MHg, MilliMHg.Base())
	assert.Equal(t, MHg, CentiMHg.Base())

	// Newton per square meter metric family
	assert.Equal(t, NewtonSqm, KiloNewtonSqm.Base())
}

func Test_Lookup_Pressure_Names_and_Symbols(t *testing.T) {
	tests := lookUpTests{
		{Pascal, "pascal"}, {Pascal, "Pa"},
		{KiloPascal, "kilopascal"}, {KiloPascal, "kPa"},
		{MegaPascal, "megapascal"}, {MegaPascal, "MPa"},
		{HectoPascal, "hectopascal"}, {HectoPascal, "hPa"},
		{MilliPascal, "millipascal"}, {MilliPascal, "mPa"},
		{At, "technical atmosphere"}, {At, "at"},
		{Atm, "standard atmosphere"}, {Atm, "atm"},
		{Bar, "bar"}, {Bar, "bar"},
		{MilliBar, "millibar"}, {MilliBar, "mbar"},
		{Barye, "barye"}, {Barye, "Ba"},
		{InH2O, "inch of Water Column"}, {InH2O, "inH2O"}, {InH2O, "inches of Water Column"},
		{InHg, "inch of Mercury"}, {InHg, "inHg"}, {InHg, "inches of Mercury"},
		{MH2O, "meter of Water Column"}, {MH2O, "mH2O"}, {MH2O, "meters of Water Column"},
		{MilliMH2O, "millimeter of Water Column"}, {MilliMH2O, "mmH2O"}, {MilliMH2O, "millimeters of Water Column"},
		{MHg, "meter of Mercury"}, {MHg, "mmHg"}, {MHg, "meters of Mercury"},
		{MilliMHg, "millimeter of Mercury"}, {MilliMHg, "millimeters of Mercury"},
		{NewtonSqm, "newton per square meter"}, {NewtonSqm, "N/m²"},
		{KiloNewtonSqm, "kilonewton per square meter"}, {KiloNewtonSqm, "kN/m²"},
		{Psi, "pound-force per square inch"}, {Psi, "psi"}, {Psi, "lbf/in²"}, {Psi, "lbf/in^2"},
		{Torr, "torr"}, {Torr, "Torr"},
		{FootH2O, "foot of Water Column"}, {FootH2O, "FT"},
		{InchesOfWater, "inches of Water Column"}, {InchesOfWater, "in-wg"},
		{PoundForcePerSquareInchGauge, "pound-force per square inch gauge"}, {PoundForcePerSquareInchGauge, "psig"},
	}
	testLookupNamesAndSymbols(t, tests)
}
