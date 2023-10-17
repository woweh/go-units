package units

import "testing"

var ratioConvTests = []conversionTest{
	{from: "fraction", to: "percent", val: "100"},
	{from: "fraction", to: "permille", val: "1000"},
	{from: "fraction", to: "partsPerMillion", val: "1000000"},
	{from: "fraction", to: "partsPerBillion", val: "1000000000"},
	{from: "fraction", to: "partsPerTrillion", val: "1000000000000"},
}

func Test_RatioConversions(t *testing.T) {
	testConversions(t, ratioConvTests)
}
