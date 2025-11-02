package units

var (
	TemperatureDifference = Quantity("temperature difference")

	// Base units: Kelvin and Celsius use the same scale for differences
	// (no offset, only ratio matters for temperature intervals)
	KelvinInterval  = mustCreateNewUnit("kelvin interval", "K", TemperatureDifference, SI)
	CelsiusInterval = mustCreateNewUnit("celsius interval", "delta°C", TemperatureDifference, SI)

	// Imperial units
	FahrenheitInterval = mustCreateNewUnit("fahrenheit interval", "delta°F", TemperatureDifference)
	RankineInterval    = mustCreateNewUnit("rankine interval", "delta°R", TemperatureDifference)
)

func init() {
	// From RevitUnits.json:
	// K: CF = 1.0
	// delta°C: CF = 1.0
	// delta°F: CF = 1.7999999999999998
	// delta°R: CF = 1.7999999999999998

	// Kelvin and Celsius have the same interval scale
	NewRatioConversion(KelvinInterval, CelsiusInterval, 1.0)

	// 1 K (or °C) interval = 1.8 °F (or °R) interval
	NewRatioConversion(KelvinInterval, FahrenheitInterval, 1.8)
	NewRatioConversion(KelvinInterval, RankineInterval, 1.8)

	KelvinInterval.AddAliases("kelvin", "K interval", "delta K")
	CelsiusInterval.AddAliases("celsius interval", "Celsius interval", "°C interval", "C interval", "delta C", "deltaC")
	FahrenheitInterval.AddAliases("fahrenheit interval", "Fahrenheit interval", "°F interval", "F interval", "delta F", "deltaF")
	RankineInterval.AddAliases("rankine interval", "Rankine interval", "°R interval", "R interval", "delta R", "deltaR")
}
