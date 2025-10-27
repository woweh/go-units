package units

import (
    "testing"

	ns "github.com/woweh/go-units/numstr"
)

func Test_Ratio_Conversions(t *testing.T) {
    var conversionTests = []conversionTest{
        // Base conversions from fraction
        {"fraction", "percent", ns.Hundred},
        {"fraction", "permille", ns.Thousand},
        {"fraction", "partsPerMillion", ns.Million},
        {"fraction", "partsPerBillion", ns.Billion},
        {"fraction", "partsPerTrillion", ns.Trillion},

        // Base conversions to fraction
        {"percent", "fraction", ns.Hundredth},
        {"permille", "fraction", ns.Thousandth},
        {"partsPerMillion", "fraction", ns.Millionth},
        {"partsPerBillion", "fraction", ns.Billionth},
        {"partsPerTrillion", "fraction", ns.Trillionth},

        // Additional important conversions
        {"percent", "permille", ns.Ten},
        {"permille", "partsPerMillion", ns.Thousand},
        {"partsPerMillion", "partsPerBillion", ns.Thousand},
        {"partsPerBillion", "partsPerTrillion", ns.Thousand},

        // Explicit conversions for verification
        {"percent", "partsPerMillion", ns.TenThousand},
        {"percent", "partsPerBillion", "10000000"},
        {"percent", "partsPerTrillion", "10000000000"},
        {"partsPerMillion", "percent", "0.0001"},
        {"partsPerBillion", "percent", "0.0000001"},
        {"partsPerTrillion", "percent", "0.0000000001"},
    }

    testConversions(t, conversionTests)
}
