package units

import (
	"testing"

	ns "github.com/woweh/go-units/numericstring"
)

var ratioConvTests = []conversionTest{
	{from: "fraction", to: "percent", val: ns.Hundred},
	{from: "fraction", to: "permille", val: ns.Thousand},
	{from: "fraction", to: "partsPerMillion", val: ns.Million},
	{from: "fraction", to: "partsPerBillion", val: ns.Billion},
	{from: "fraction", to: "partsPerTrillion", val: ns.Trillion},
}

func Test_RatioConversions(t *testing.T) {
	testConversions(t, ratioConvTests)
}
