package units

var (
	Velocity = Quantity("velocity")

	// Base unit: foot per second (Revit base with CF=1.0)
	FootPerSecond = newUnit("foot per second", "ft/s", Velocity, BI)

	// Other units
	FootPerMinute       = newUnit("foot per minute", "ft/min", Velocity, BI)
	MeterPerSecond      = newUnit("meter per second", "m/s", Velocity, SI)
	CentimeterPerMinute = newUnit("centimeter per minute", "cm/min", Velocity, SI)
	KilometerPerHour    = newUnit("kilometer per hour", "km/h", Velocity, SI)
	MilePerHour         = newUnit("mile per hour", "mph", Velocity, BI)
)

func init() {
	// From RevitUnits.json:
	// ft/s (base): CF = 1.0
	// ft/min: CF = 60.0
	// m/s: CF = 0.3048
	// cm/min: CF = 1828.8000000000002
	// km/h: CF = 1.09728
	// mph: CF = 0.6818181818181817
	// This means: 1 ft/s = 60 ft/min, 1 ft/s = 0.3048 m/s, etc.
	NewRatioConversion(FootPerSecond, FootPerMinute, 60.0)
	NewRatioConversion(FootPerSecond, MeterPerSecond, 0.3048)
	NewRatioConversion(FootPerSecond, CentimeterPerMinute, 1828.8000000000002)
	NewRatioConversion(FootPerSecond, KilometerPerHour, 1.09728)
	NewRatioConversion(FootPerSecond, MilePerHour, 0.6818181818181817)

	FootPerSecond.AddAliases("feet per second", "fps")
	FootPerMinute.AddAliases("feet per minute", "fpm")
	MeterPerSecond.AddAliases("meters per second", "metre per second", "metres per second", "mps")
	CentimeterPerMinute.AddAliases("centimeters per minute", "centimetre per minute", "centimetres per minute")
	KilometerPerHour.AddAliases("kilometers per hour", "kilometre per hour", "kilometres per hour", "kph")
	MilePerHour.AddAliases("miles per hour")
}
