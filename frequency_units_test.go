package units

import "testing"

var frequencyConvTests = []conversionTest{
	{"hertz", "radian per second", "6.283185"},
	{"hertz", "radian per minute", "376.991118"},
	{"hertz", "radian per hour", "22619.467106"},
	{"hertz", "radian per day", "542867.21054"},
	{"hertz", "degree per second", "360"},
	{"hertz", "degree per minute", "21600"},
	{"hertz", "degree per hour", "1296000"},
	{"hertz", "degree per day", "31104000"},
	{"hertz", "revolution per second", "1"},
	{"hertz", "revolution per minute", "60"},
	{"hertz", "revolution per hour", "3600"},
	{"hertz", "revolution per day", "86400"},
	{"GigaHertz", "fresnel", "0.001"},
	{"TeraHertz", "fresnel", "1"},
	{"PetaHertz", "fresnel", "1000"},
	{"radian per second", "radian per minute", "60"},
	{"radian per second", "radian per hour", "3600"},
	{"radian per second", "radian per day", "86400"},
	{"radian per second", "degree per second", "57.29578"},
	{"radian per second", "degree per minute", "3437.746771"},
	{"radian per second", "degree per hour", "206264.806247"},
	{"radian per second", "degree per day", "4950355.34993"},
	{"radian per second", "revolution per second", "0.1591549"},
	{"radian per second", "revolution per minute", "9.549297"},
	{"radian per second", "revolution per hour", "572.957795"},
	{"radian per second", "revolution per day", "13750.987083"},
	{"radian per minute", "degree per minute", "57.29578"},
	{"radian per minute", "revolution per minute", "0.1591549"},
	{"radian per hour", "degree per hour", "57.29578"},
	{"radian per hour", "revolution per hour", "0.1591549"},
	{"radian per day", "degree per day", "57.29578"},
	{"radian per day", "revolution per day", "0.1591549"},
	{"degree per second", "degree per minute", "60"},
	{"degree per second", "degree per hour", "3600"},
	{"degree per second", "degree per day", "86400"},
	{"degree per second", "revolution per second", "0.002777778"},
	{"degree per minute", "revolution per minute", "0.002777778"},
	{"degree per hour", "revolution per hour", "0.002777778"},
	{"degree per day", "revolution per day", "0.002777778"},
	{"revolution per second", "revolution per minute", "60"},
	{"revolution per second", "revolution per hour", "3600"},
	{"revolution per second", "revolution per day", "86400"},
}

func Test_FrequencyConversions(t *testing.T) {
	testConversions(t, frequencyConvTests)
}
