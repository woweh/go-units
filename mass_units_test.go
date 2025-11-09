package units

import (
	"testing"
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
	tests := []unitSystemTest{
		{Gram, SiSystem},
		{ExaGram, SiSystem},
		{PetaGram, SiSystem},
		{TeraGram, SiSystem},
		{GigaGram, SiSystem},
		{MegaGram, SiSystem},
		{KiloGram, SiSystem},
		{HectoGram, SiSystem},
		{DecaGram, SiSystem},
		{DeciGram, SiSystem},
		{CentiGram, SiSystem},
		{MilliGram, SiSystem},
		{MicroGram, SiSystem},
		{NanoGram, SiSystem},
		{PicoGram, SiSystem},
		{FemtoGram, SiSystem},
		{AttoGram, SiSystem},
		{MetricTon, SiSystem},
		{Grain, BiSystem},
		{Drachm, BiSystem},
		{Ounce, BiSystem},
		{Pound, BiSystem},
		{Stone, BiSystem},
		{Ton, BiSystem},
		{Slug, BiSystem},
		{ShortTon, BiSystem},
	}
	testUnitSystems(t, tests)
}
