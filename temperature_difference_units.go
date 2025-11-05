package units

// TemperatureDifference is a unit quantity for temperature difference
const TemperatureDifference UnitQuantity = "temperature difference"

var (
	_temperatureDifference = Quantity(TemperatureDifference)

	// SI base units: Kelvin and Celsius intervals (identical scale for differences)
	KelvinInterval  = mustCreateNewUnit("kelvin interval", "K", _temperatureDifference, SI)
	CelsiusInterval = mustCreateNewUnit("celsius interval", "delta°C", _temperatureDifference, SI)

	// Imperial/US units
	FahrenheitInterval = mustCreateNewUnit("fahrenheit interval", "delta°F", _temperatureDifference, BI)
	RankineInterval    = mustCreateNewUnit("rankine interval", "delta°R", _temperatureDifference, BI)
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
