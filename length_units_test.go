package units

import "testing"

func Test_Lookup_Length_Names_and_Symbols(t *testing.T) {

	tests := []lookUpTestUnit{
		{Meter, "metre"},
		{Angstrom, "ångström"},
		{Angstrom, "angstroms"},
		{Angstrom, "ångströms"},
		{Inch, "inches"},
		{Inch, "in."},
		{Inch, "″"},
		{Foot, "feet"},
		{Foot, "ft."},
		{Foot, "′"},
		{Yard, "yd."},
		{Mile, "mi."},
	}

	testLookupNamesAndSymbols(t, tests)
}
