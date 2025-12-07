package units

var (
	// Acceleration is a unit quantity for acceleration
	Acceleration = NewUnitQuantity("acceleration")

	// SI base unit: meter per second squared
	MeterPerSecondSquared     = Acceleration.MustCreateUnit("meter per second squared", "m/s²", SI)
	KilometerPerSecondSquared = Acceleration.MustCreateUnit("kilometer per second squared", "km/s²", SI)

	// Imperial/US units
	FootPerSecondSquared = Acceleration.MustCreateUnit("foot per second squared", "ft/s²", BI)
	InchPerSecondSquared = Acceleration.MustCreateUnit("inch per second squared", "in/s²", BI)
	MilePerSecondSquared = Acceleration.MustCreateUnit("mile per second squared", "mi/s²", BI)
)

func initAccelerationUnits() {
	// Derive conversion factors from length units
	// acceleration = length / time² where time is always seconds
	NewRatioConversion(MeterPerSecondSquared, FootPerSecondSquared, accelFactor(Foot))
	NewRatioConversion(MeterPerSecondSquared, InchPerSecondSquared, accelFactor(Inch))
	NewRatioConversion(MeterPerSecondSquared, KilometerPerSecondSquared, accelFactor(KiloMeter))
	NewRatioConversion(MeterPerSecondSquared, MilePerSecondSquared, accelFactor(Mile))

	MeterPerSecondSquared.AddAliases("meters per second squared", "metre per second squared", "metres per second squared", "m/sec²")
	KilometerPerSecondSquared.AddAliases("kilometers per second squared", "kilometre per second squared", "kilometres per second squared", "km/sec²")
	FootPerSecondSquared.AddAliases("feet per second squared", "ft/sec²", "ft/sec/sec")
	InchPerSecondSquared.AddAliases("inches per second squared", "in/sec²", "in/sec/sec")
	MilePerSecondSquared.AddAliases("miles per second squared", "mi/sec²")
}

// accelFactor computes the acceleration conversion factor from length units.
// Since all acceleration units use the same time unit (second²), only length varies.
// Returns how many of the target length unit equals 1 meter (for m/s² conversions).
func accelFactor(length Unit) float64 {
	return Meter.to(length)
}
