package units

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Temperature_Conversions(t *testing.T) {
	var conversionTests = []conversionTest{
		{from: "celsius", to: "celsius", val: "1"},
		{from: "celsius", to: "fahrenheit", val: "33.8"},
		{from: "celsius", to: "kelvin", val: "274.15"},
		{from: "celsius", to: "rankine", val: "493.47"},
		{from: "fahrenheit", to: "celsius", val: "-17.222222"},
		{from: "fahrenheit", to: "fahrenheit", val: "1"},
		{from: "fahrenheit", to: "kelvin", val: "255.927778"},
		{from: "fahrenheit", to: "rankine", val: "460.67"},
		{from: "kelvin", to: "celsius", val: "-272.15"},
		{from: "kelvin", to: "fahrenheit", val: "-457.87"},
		{from: "kelvin", to: "kelvin", val: "1"},
		{from: "kelvin", to: "rankine", val: "1.8"},
		{from: "rankine", to: "celsius", val: "-272.594444"},
		{from: "rankine", to: "fahrenheit", val: "-458.67"},
		{from: "rankine", to: "kelvin", val: "0.555556"},
		{from: "rankine", to: "rankine", val: "1"},
	}

	testConversions(t, conversionTests)
}

func Test_Temperature_Systems(t *testing.T) {
	si := SiSystem
	assert.Equal(t, si, Celsius.System())
	assert.Equal(t, si, Kelvin.System())
	us := UsSystem
	assert.Equal(t, us, Fahrenheit.System())
	assert.Equal(t, us, Rankine.System())
}
