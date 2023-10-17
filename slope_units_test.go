package units

import "testing"

var slopeConvTests = []conversionTest{
	{from: "slope value", to: "slope ratio", val: "1"},
	{from: "slope value", to: "inverse slope ratio", val: "1"},
	{from: "slope value", to: "slope degree", val: "45"},
	{from: "slope value", to: "slope percent", val: "100"},
	{from: "slope value", to: "slope permille", val: "1000"},
	{from: "slope percent", to: "slope degree", val: "0.572939"},
	{from: "slope percent", to: "slope permille", val: "10"},
	{from: "slope permille", to: "slope degree", val: "0.0572958"},
	{from: "slope degree", to: "slope percent", val: "1.745506"},
	{from: "slope degree", to: "slope permille", val: "17.455065"},
	{from: "slope degree", to: "slope value", val: "0.01745506"},
	{from: "slope degree", to: "slope ratio", val: "57.289962"},
}

func Test_SlopeConversions(t *testing.T) {
	testConversions(t, slopeConvTests)
}
