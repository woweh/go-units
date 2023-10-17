package units

import "testing"

var voltageConvTests = []conversionTest{
	{from: "volt", to: "yottavolt", val: "0.000000000000000000000001"},
	{from: "volt", to: "zettavolt", val: "0.000000000000000000001"},
	{from: "volt", to: "exavolt", val: "0.000000000000000001"},
	{from: "volt", to: "petavolt", val: "0.000000000000001"},
	{from: "volt", to: "teravolt", val: "0.000000000001"},
	{from: "volt", to: "gigavolt", val: "0.000000001"},
	{from: "volt", to: "megavolt", val: "0.000001"},
	{from: "volt", to: "kilovolt", val: "0.001"},
	{from: "volt", to: "hectovolt", val: "0.01"},
	{from: "volt", to: "decavolt", val: "0.1"},
	{from: "volt", to: "decivolt", val: "10"},
	{from: "volt", to: "centivolt", val: "100"},
	{from: "volt", to: "millivolt", val: "1000"},
	{from: "volt", to: "microvolt", val: "1000000"},
	{from: "volt", to: "nanovolt", val: "1000000000"},
	{from: "volt", to: "picovolt", val: "1000000000000"},
}

func Test_VoltageConversions(t *testing.T) {
	testConversions(t, voltageConvTests)
}
