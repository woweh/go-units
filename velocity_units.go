package units

// Velocity is a unit quantity for velocity
var (
	// VelocityQuantity is the unit quantity for velocity.
	VelocityQuantity = NewUnitQuantity("velocity")

	// SI base unit: meter per second
	MeterPerSecond      = VelocityQuantity.MustCreateUnit("meter per second", "m/s", SI)
	CentimeterPerMinute = VelocityQuantity.MustCreateUnit("centimeter per minute", "cm/min", SI)
	KilometerPerHour    = VelocityQuantity.MustCreateUnit("kilometer per hour", "km/h", SI)

	// Imperial/US units
	FootPerSecond = VelocityQuantity.MustCreateUnit("foot per second", "ft/s", BI)
	FootPerMinute = VelocityQuantity.MustCreateUnit("foot per minute", "ft/min", BI)
	MilePerHour   = VelocityQuantity.MustCreateUnit("mile per hour", "mph", BI)
)

func initVelocityUnits() {
	// Derive conversion factors from length and time units
	// velocity = length / time
	NewRatioConversion(MeterPerSecond, FootPerSecond, velocityFactor(Foot, Second))
	NewRatioConversion(MeterPerSecond, FootPerMinute, velocityFactor(Foot, Minute))
	NewRatioConversion(MeterPerSecond, CentimeterPerMinute, velocityFactor(CentiMeter, Minute))
	NewRatioConversion(MeterPerSecond, KilometerPerHour, velocityFactor(KiloMeter, Hour))
	NewRatioConversion(MeterPerSecond, MilePerHour, velocityFactor(Mile, Hour))

	MeterPerSecond.AddAliases("meters per second", "metre per second", "metres per second", "mps")
	CentimeterPerMinute.AddAliases("centimeters per minute", "centimetre per minute", "centimetres per minute")
	KilometerPerHour.AddAliases("kilometers per hour", "kilometre per hour", "kilometres per hour", "kph")
	FootPerSecond.AddAliases("feet per second", "fps")
	FootPerMinute.AddAliases("feet per minute", "fpm")
	MilePerHour.AddAliases("miles per hour")
}

// velocityFactor computes the velocity conversion factor from length and time units.
// velocity = length / time
func velocityFactor(length, time Unit) float64 {
	return Meter.to(length) / Second.to(time)
}
