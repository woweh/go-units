package units

var (
	Velocity = Quantity("velocity")

	// SI base unit: meter per second
	MeterPerSecond      = mustCreateNewUnit("meter per second", "m/s", Velocity, SI)
	CentimeterPerMinute = mustCreateNewUnit("centimeter per minute", "cm/min", Velocity, SI)
	KilometerPerHour    = mustCreateNewUnit("kilometer per hour", "km/h", Velocity, SI)

	// Imperial/US units
	FootPerSecond = mustCreateNewUnit("foot per second", "ft/s", Velocity, BI)
	FootPerMinute = mustCreateNewUnit("foot per minute", "ft/min", Velocity, BI)
	MilePerHour   = mustCreateNewUnit("mile per hour", "mph", Velocity, BI)
)

func init() {
	// SI base unit: m/s
	// Conversions: 1 m/s = 3.28084 ft/s = 196.85 ft/min = 3.6 km/h = 2.237 mph
	NewRatioConversion(MeterPerSecond, FootPerSecond, 3.2808398950131)
	NewRatioConversion(MeterPerSecond, FootPerMinute, 196.85039370079)
	NewRatioConversion(MeterPerSecond, CentimeterPerMinute, 6000.0)
	NewRatioConversion(MeterPerSecond, KilometerPerHour, 3.6)
	NewRatioConversion(MeterPerSecond, MilePerHour, 2.2369362920544)

	MeterPerSecond.AddAliases("meters per second", "metre per second", "metres per second", "mps")
	CentimeterPerMinute.AddAliases("centimeters per minute", "centimetre per minute", "centimetres per minute")
	KilometerPerHour.AddAliases("kilometers per hour", "kilometre per hour", "kilometres per hour", "kph")
	FootPerSecond.AddAliases("feet per second", "fps")
	FootPerMinute.AddAliases("feet per minute", "fpm")
	MilePerHour.AddAliases("miles per hour")
}
