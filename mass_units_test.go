package units

import (
	"testing"

	"github.com/alecthomas/assert/v2"
)

func Test_Mass_Conversions(t *testing.T) {
	conversionTests := []conversionTest{
		// Revit conversions from internal (kg)
		{"kg", "lb", 2.20462262184878},
		{"lb", "kg", 0.45359237},
		{"kg", "t", 0.001},
		{"t", "kg", 1000},
		{"kg", "Tons", 0.00110231131092439},
		{"Tons", "kg", 907.18474},

		// Imperial unit conversions
		{"lb", "oz", 16},
		{"oz", "lb", 0.0625},
		{"lb", "gr", 7000},
		{"gr", "lb", 0.00014285714285714},
		{"lb", "dr", 256},
		{"dr", "lb", 0.00390625},
		{"lb", "st", 0.071428571428571},
		{"st", "lb", 14},
		{"lb", "LT", 0.00044642857142857},
		{"LT", "lb", 2240},

		// Metric conversions (sample of key units to non-metric)
		{"g", "lb", 0.0022046226218488},
		{"lb", "g", 453.59237},
		{"mg", "lb", 0.0000022046226218488},
		{"kg", "oz", 35.27396194958},

		// Metric ton and short ton
		{"t", "lb", 2204.6226218488},
		{"Tons", "lb", 2000},
		{"t", "Tons", 1.1023113109244},
		{"Tons", "t", 0.90718474},
	}
	testConversions(t, conversionTests)
}

func Test_Mass_UnitSystems(t *testing.T) {
	// Metric units
	assert.Equal(t, SiSystem, Gram.System())
	assert.Equal(t, SiSystem, ExaGram.System())
	assert.Equal(t, SiSystem, PetaGram.System())
	assert.Equal(t, SiSystem, TeraGram.System())
	assert.Equal(t, SiSystem, GigaGram.System())
	assert.Equal(t, SiSystem, MegaGram.System())
	assert.Equal(t, SiSystem, KiloGram.System())
	assert.Equal(t, SiSystem, HectoGram.System())
	assert.Equal(t, SiSystem, DecaGram.System())
	assert.Equal(t, SiSystem, DeciGram.System())
	assert.Equal(t, SiSystem, CentiGram.System())
	assert.Equal(t, SiSystem, MilliGram.System())
	assert.Equal(t, SiSystem, MicroGram.System())
	assert.Equal(t, SiSystem, NanoGram.System())
	assert.Equal(t, SiSystem, PicoGram.System())
	assert.Equal(t, SiSystem, FemtoGram.System())
	assert.Equal(t, SiSystem, AttoGram.System())
	assert.Equal(t, SiSystem, MetricTon.System())

	// Imperial units
	assert.Equal(t, BiSystem, Grain.System())
	assert.Equal(t, BiSystem, Drachm.System())
	assert.Equal(t, BiSystem, Ounce.System())
	assert.Equal(t, BiSystem, Pound.System())
	assert.Equal(t, BiSystem, Stone.System())
	assert.Equal(t, BiSystem, Ton.System())
	assert.Equal(t, BiSystem, Slug.System())
	assert.Equal(t, BiSystem, ShortTon.System())
}

func Test_Mass_MetricFactories(t *testing.T) {
	// Test that metric factory-created units have Gram as base
	assert.Equal(t, Gram, ExaGram.Base())
	assert.Equal(t, Gram, PetaGram.Base())
	assert.Equal(t, Gram, TeraGram.Base())
	assert.Equal(t, Gram, GigaGram.Base())
	assert.Equal(t, Gram, MegaGram.Base())
	assert.Equal(t, Gram, KiloGram.Base())
	assert.Equal(t, Gram, HectoGram.Base())
	assert.Equal(t, Gram, DecaGram.Base())
	assert.Equal(t, Gram, DeciGram.Base())
	assert.Equal(t, Gram, CentiGram.Base())
	assert.Equal(t, Gram, MilliGram.Base())
	assert.Equal(t, Gram, MicroGram.Base())
	assert.Equal(t, Gram, NanoGram.Base())
	assert.Equal(t, Gram, PicoGram.Base())
	assert.Equal(t, Gram, FemtoGram.Base())
	assert.Equal(t, Gram, AttoGram.Base())
}

func Test_Lookup_Mass_Names_and_Symbols(t *testing.T) {
	tests := lookUpTests{
		// Gram
		{Gram, "gram"},
		{Gram, "g"},

		// Metric prefixes (sample)
		{KiloGram, "kilogram"},
		{KiloGram, "kg"},
		{MilliGram, "milligram"},
		{MilliGram, "mg"},
		{MicroGram, "microgram"},
		{MicroGram, "μg"},

		// Metric ton
		{MetricTon, "metric ton"},
		{MetricTon, "t"},
		{MetricTon, "tonne"},
		{MetricTon, "tonnes"},
		{MetricTon, "metric tons"},

		// Imperial units
		{Grain, "grain"},
		{Grain, "gr"},
		{Drachm, "drachm"},
		{Drachm, "dr"},
		{Ounce, "ounce"},
		{Ounce, "oz"},
		{Pound, "pound"},
		{Pound, "lb"},
		{Stone, "stone"},
		{Stone, "st"},
		{Ton, "ton"},
		{Ton, "LT"},
		{Ton, "long ton"},
		{Ton, "long tons"},
		{Ton, "imperial ton"},
		{Ton, "displacement ton"},
		{Slug, "slug"},
		{ShortTon, "short ton"},
		{ShortTon, "Tons"},
		{ShortTon, "US ton"},
		{ShortTon, "short tons"},
	}
	testLookupNamesAndSymbols(t, tests)
}

