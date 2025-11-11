package units

// TemperatureDifference is a unit quantity for temperature difference
const TemperatureDifference UnitQuantity = "temperature difference"

/*
NOTE:
The symbols for temperature difference are the same as for temperature.
BUT, because this library requires unique symbols for units, we use the "Δ" prefix for all symbols.
*/

var (
	_temperatureDifference = Quantity(TemperatureDifference)

	// SI base units: Kelvin and Celsius intervals (identical scale for differences)
	KelvinInterval  = mustCreateNewUnit("kelvin interval", "ΔK", _temperatureDifference, SI)
	CelsiusInterval = mustCreateNewUnit("celsius interval", "Δ°C", _temperatureDifference, SI)

	// Imperial/US units
	FahrenheitInterval = mustCreateNewUnit("fahrenheit interval", "Δ°F", _temperatureDifference, BI)
	RankineInterval    = mustCreateNewUnit("rankine interval", "Δ°R", _temperatureDifference, BI)
)

func initTemperatureDifferenceUnits() {
	// Kelvin and Celsius have the same interval scale
	NewRatioConversion(KelvinInterval, CelsiusInterval, 1.0)

	// 1 K (or °C) interval = 1.8 °F (or °R) interval
	NewRatioConversion(KelvinInterval, FahrenheitInterval, 1.8)
	NewRatioConversion(KelvinInterval, RankineInterval, 1.8)

	KelvinInterval.AddAliases("kelvin interval", "K interval", "K difference", "Kelvin difference")
	KelvinInterval.AddSymbols("delta K", "deltaK")

	CelsiusInterval.AddAliases("Celsius interval", "°C interval", "C interval", "Celsius difference", "°C interval", "C difference")
	CelsiusInterval.AddSymbols("delta°C", "delta C", "deltaC")

	FahrenheitInterval.AddAliases("Fahrenheit interval", "°F interval", "F interval", "Fahrenheit difference", "°F difference", "F difference")
	FahrenheitInterval.AddSymbols("delta F", "deltaF")

	RankineInterval.AddAliases("Rankine interval", "°R interval", "R interval", "Rankine difference", "°R difference", "R difference")
	RankineInterval.AddSymbols("delta R", "deltaR")
}
