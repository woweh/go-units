package units

// TemperatureDifference is a unit quantity for temperature difference
var (
	// TemperatureDifference is the unit quantity for temperature difference.
	TemperatureDifference = NewUnitQuantity("temperature difference")

	// SI base units: Kelvin and Celsius intervals (identical scale for differences)
	KelvinInterval  = TemperatureDifference.MustCreateUnit("kelvin interval", "ΔK", SI)
	CelsiusInterval = TemperatureDifference.MustCreateUnit("celsius interval", "Δ°C", SI)

	// Imperial/US units
	FahrenheitInterval = TemperatureDifference.MustCreateUnit("fahrenheit interval", "Δ°F", BI)
	RankineInterval    = TemperatureDifference.MustCreateUnit("rankine interval", "Δ°R", BI)
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
	FahrenheitInterval.AddSymbols("delta°F", "delta F", "deltaF")

	RankineInterval.AddAliases("Rankine interval", "°R interval", "R interval", "Rankine difference", "°R difference", "R difference")
	RankineInterval.AddSymbols("delta°R", "delta R", "deltaR")
}
