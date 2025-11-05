package units

// Velocity is a unit quantity for velocity
const VelocityQuantity UnitQuantity = "velocity"

var (
	_velocity = Quantity(VelocityQuantity)

	// SI base unit: meter per second
	MeterPerSecond      = mustCreateNewUnit("meter per second", "m/s", _velocity, SI)
	CentimeterPerMinute = mustCreateNewUnit("centimeter per minute", "cm/min", _velocity, SI)
	KilometerPerHour    = mustCreateNewUnit("kilometer per hour", "km/h", _velocity, SI)

	// Imperial/US units
	FootPerSecond = mustCreateNewUnit("foot per second", "ft/s", _velocity, BI)
	FootPerMinute = mustCreateNewUnit("foot per minute", "ft/min", _velocity, BI)
	MilePerHour   = mustCreateNewUnit("mile per hour", "mph", _velocity, BI)
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
