package units

import (
	"testing"
)

func Test_Mass_Conversions(t *testing.T) {
	conversionTests := []conversionTest{
		// Revit conversions from internal (kg)
		{from: "kg", to: "lb", exp: 2.20462262184878},
		{from: "lb", to: "kg", exp: 0.45359237},
		{from: "kg", to: "t", exp: 0.001},
		{from: "t", to: "kg", exp: 1000},
		{from: "kg", to: "Tons", exp: 0.00110231131092439},
		{from: "Tons", to: "kg", exp: 907.18474},

		// Imperial unit conversions
		{from: "lb", to: "oz", exp: 16},
		{from: "oz", to: "lb", exp: 0.0625},
		{from: "lb", to: "gr", exp: 7000},
		{from: "gr", to: "lb", exp: 0.00014285714285714},
		{from: "lb", to: "dr", exp: 256},
		{from: "dr", to: "lb", exp: 0.00390625},
		{from: "lb", to: "st", exp: 0.071428571428571},
		{from: "st", to: "lb", exp: 14},
		{from: "lb", to: "LT", exp: 0.00044642857142857},
		{from: "LT", to: "lb", exp: 2240},

		// Metric conversions (sample of key units to non-metric)
		{from: "g", to: "lb", exp: 0.0022046226218488},
		{from: "lb", to: "g", exp: 453.59237},
		{from: "mg", to: "lb", exp: 0.0000022046226218488},
		{from: "kg", to: "oz", exp: 35.27396194958},

		// Metric ton and short ton
		{from: "t", to: "lb", exp: 2204.6226218488},
		{from: "Tons", to: "lb", exp: 2000},
		{from: "t", to: "Tons", exp: 1.1023113109244},
		{from: "Tons", to: "t", exp: 0.90718474},
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
