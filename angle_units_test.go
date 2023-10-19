package units

import (
	"fmt"
	"math"
	"testing"
)

var angleConvTests = []conversionTest{
	{"turn", "degree", "360"},
	{"turn", "gon", "400"},
	{"turn", "radian", fmt.Sprintf("%f", 2*math.Pi)},
	{"turn", "milliradian", fmt.Sprintf("%f", 2*math.Pi*1000)},
	{"radian", "milliradian", "1000"},
	{"radian", "microradian", "1000000"},
	{"radian", "degree", fmt.Sprintf("%.5f", 180/math.Pi)},
	{"radian", "gon", fmt.Sprintf("%f", 200/math.Pi)},
	{"degree", "gon", fmt.Sprintf("%f", 10/9.0)},
	{"gon", "degree", "0.9"},
	{"gon", "milligon", "1000"},
	{"gon", "centigon", "100"},
	{"gon", "decigon", "10"},
	{"gon", "decagon", "0.1"},
	{"gon", "hectogon", "0.01"},
}

func Test_AngleConversions(t *testing.T) {
	testConversions(t, angleConvTests)
}
