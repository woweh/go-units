package units

import (
	"testing"

	ns "github.com/woweh/go-units/numstr"
)

func Test_Ratio_Conversions(t *testing.T) {
	var conversionTests = []conversionTest{
		{"fraction", "percent", ns.Hundred},
		{"fraction", "permille", ns.Thousand},
		{"fraction", "partsPerMillion", ns.Million},
		{"fraction", "partsPerBillion", ns.Billion},
		{"fraction", "partsPerTrillion", ns.Trillion},
		{"percent", "fraction", ns.Hundredth},
		{"percent", "permille", ns.Ten},
		{"percent", "partsPerMillion", ns.TenThousand},
		{"percent", "partsPerBillion", "10000000"},
		{"percent", "partsPerTrillion", "10000000000"},
		{"permille", "fraction", ns.Thousandth},
		{"permille", "percent", ns.Tenth},
		{"permille", "partsPerMillion", ns.Thousand},
		{"permille", "partsPerBillion", ns.Million},
		{"permille", "partsPerTrillion", ns.Billion},
		{"partsPerMillion", "fraction", ns.Millionth},
		{"partsPerMillion", "percent", "0.0001"},
		{"partsPerMillion", "permille", ns.Thousandth},
		{"partsPerMillion", "partsPerBillion", ns.Thousand},
		{"partsPerMillion", "partsPerTrillion", ns.Million},
		{"partsPerBillion", "fraction", ns.Billionth},
		{"partsPerBillion", "percent", "0.0000001"},
		{"partsPerBillion", "permille", ns.Millionth},
		{"partsPerBillion", "partsPerMillion", ns.Thousandth},
		{"partsPerBillion", "partsPerTrillion", ns.Thousand},
		{"partsPerTrillion", "fraction", ns.Trillionth},
		{"partsPerTrillion", "percent", "0.0000000001"},
		{"partsPerTrillion", "permille", ns.Billionth},
		{"partsPerTrillion", "partsPerMillion", ns.Millionth},
		{"partsPerTrillion", "partsPerBillion", ns.Thousandth},
	}

	testConversions(t, conversionTests)
}
