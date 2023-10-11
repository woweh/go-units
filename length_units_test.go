package units

import "testing"

func Test_LengthAliases(t *testing.T) {
	
	tests := []aliasTest{
		{Meter, "metre"},
		{Angstrom, "ångström"},
		{Angstrom, "angstroms"},
		{Angstrom, "ångströms"},
		{Inch, "inches"},
		{Inch, "in."},
		{Inch, "″"},
		{Inch, "\""},
		{Foot, "feet"},
		{Foot, "ft."},
		{Foot, "′"},
		{Yard, "yd."},
		{Mile, "mi."},
	}

	testAliases(t, tests)
}
