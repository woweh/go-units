package units

import (
	"fmt"
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

// conversionTest struct defines test data for a conversion, where a value of 1.0 in
// the `from` unit name is expected to equal `val` when converted to the `to` unit name
type conversionTest struct {
	from string
	to   string
	val  string
}

func testConversions(t *testing.T, convTests []conversionTest) {
	fmtOpts := FmtOptions{false, false, 6}
	for _, cTest := range convTests {
		label := fmt.Sprintf("%s <-> %s conversion", cTest.from, cTest.to)
		t.Run(
			label, func(t *testing.T) {
				u1, err := Find(cTest.from)
				if err != nil {
					t.Fatal(err.Error())
				}
				u2, err := Find(cTest.to)
				if err != nil {
					t.Fatal(err.Error())
				}
				converted := MustConvertFloat(1.0, u1, u2)
				convStr := converted.Fmt(fmtOpts)
				assert.Equal(
					t, cTest.val, convStr,
					"%s -> %s conversion test failed", cTest.from, cTest.to,
				)
				t.Logf("%s -> %s conversion: %s", cTest.from, cTest.to, convStr)

				// test inverse conversion
				inverse := MustConvertFloat(converted.Float(), u2, u1)
				assert.Equal(
					t, 1.0, roundFloat(inverse.Float(), 12),
					"%s <- %s inverse conversion test failed", cTest.from, cTest.to,
				)
			},
		)
	}
}

func roundFloat(f float64, precision uint) float64 {
	r := math.Pow(10, float64(precision))
	return math.Round(f*r) / r
}
