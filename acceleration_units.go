package units

// Acceleration is a unit quantity for acceleration
const Acceleration UnitQuantity = "acceleration"

var (
	_acceleration = Quantity(Acceleration)

	// SI base unit: meter per second squared
	MeterPerSecondSquared     = mustCreateNewUnit("meter per second squared", "m/s²", _acceleration, SI)
	KilometerPerSecondSquared = mustCreateNewUnit("kilometer per second squared", "km/s²", _acceleration, SI)

	// Imperial/US units
	FootPerSecondSquared = mustCreateNewUnit("foot per second squared", "ft/s²", _acceleration, BI)
	InchPerSecondSquared = mustCreateNewUnit("inch per second squared", "in/s²", _acceleration, BI)
	MilePerSecondSquared = mustCreateNewUnit("mile per second squared", "mi/s²", _acceleration, BI)
)

func initAccelerationUnits() {
	// SI base unit: m/s²
	// Conversions: 1 m/s² = 3.28084 ft/s² = 39.37008 in/s²
	NewRatioConversion(MeterPerSecondSquared, FootPerSecondSquared, 3.2808398950131)
	NewRatioConversion(MeterPerSecondSquared, InchPerSecondSquared, 39.370078740157)
	NewRatioConversion(MeterPerSecondSquared, KilometerPerSecondSquared, 0.001)
	NewRatioConversion(MeterPerSecondSquared, MilePerSecondSquared, 0.00062137119223733)

	MeterPerSecondSquared.AddAliases("meters per second squared", "metre per second squared", "metres per second squared", "m/sec²")
	KilometerPerSecondSquared.AddAliases("kilometers per second squared", "kilometre per second squared", "kilometres per second squared", "km/sec²")
	FootPerSecondSquared.AddAliases("feet per second squared", "ft/sec²", "ft/sec/sec")
	InchPerSecondSquared.AddAliases("inches per second squared", "in/sec²", "in/sec/sec")
	MilePerSecondSquared.AddAliases("miles per second squared", "mi/sec²")
}
