package units

import (
	"fmt"
	"math"
	"testing"
)

// conversionTest struct defines test data for a conversion, where a value of 1.0 in
// the `from` unit name is expected to equal `val` when converted to the `to` unit name
type conversionTest struct {
	from string
	to   string
	val  string
}

func testConversions(t *testing.T, convTests []conversionTest) {
	fmtOpts := FmtOptions{false, false, 12, true, " "}
	for _, cTest := range convTests {
		label := fmt.Sprintf("%s <-> %s conversion", cTest.from, cTest.to)
		t.Run(
			label, func(t *testing.T) {
				u1, err := Find(cTest.from)
				if err != nil {
					t.Fatalf("failed to find unit %s: %v", cTest.from, err)
				}
				u2, err := Find(cTest.to)
				if err != nil {
					t.Fatalf("failed to find unit %s: %v", cTest.to, err)
				}
				converted := MustConvertFloat(1.0, u1, u2)
				convStr := converted.Fmt(fmtOpts)
				if convStr != cTest.val {
					t.Fatalf("%s -> %s conversion failed: expected %s, got %s", cTest.from, cTest.to, cTest.val, convStr)
				}

				// test inverse conversion
				inverse := MustConvertFloat(converted.Float(), u2, u1)
				invFloat := roundFloat(inverse.Float(), 12)
				if invFloat != 1.0 {
					t.Fatalf("%s <- %s inverse conversion failed: expected 1.0, got %f", cTest.from, cTest.to, invFloat)
				}
			},
		)
	}
}

func roundFloat(f float64, precision uint) float64 {
	r := math.Pow(10, float64(precision))
	return math.Round(f*r) / r
}

func Test_ResolveConversion(t *testing.T) {
	from, err := Find("meter")
	if err != nil {
		t.Fatalf("failed to find unit meter: %v", err)
	}
	to, err := Find("foot")
	if err != nil {
		t.Fatalf("failed to find unit foot: %v", err)
	}
	path, err := ResolveConversion(from, to)
	if err != nil {
		t.Fatalf("ResolveConversion failed: %v", err)
	}
	if len(path) == 0 {
		t.Fatal("expected non-empty conversion path")
	}
	// Verify the path by applying conversions
	val := 1.0
	for _, c := range path {
		val = c.Fn(val)
	}
	expected := 3.280839895013123 // approximate 1 meter in feet
	if math.Abs(val-expected) > 1e-6 {
		t.Fatalf("conversion path incorrect: got %f, expected %f", val, expected)
	}
}

func Test_checkConversionIntegrity(t *testing.T) {
	// Call the checkConversionIntegrity function and check for errors
	err := CheckConversionIntegrity()
	if err != nil {
		t.Errorf("checkConversionIntegrity failed: %v", err)
	} else {
		t.Log("checkConversionIntegrity passed")
	}
}
