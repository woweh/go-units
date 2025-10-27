package units

var (
	Acceleration = Quantity("acceleration")

	// SI base unit: meter per second squared
	MeterPerSecondSquared     = newUnit("meter per second squared", "m/s²", Acceleration, BaseSiUnit)
	KilometerPerSecondSquared = newUnit("kilometer per second squared", "km/s²", Acceleration, SI)

	// Imperial/US units
	FootPerSecondSquared = newUnit("foot per second squared", "ft/s²", Acceleration, BI)
	InchPerSecondSquared = newUnit("inch per second squared", "in/s²", Acceleration, BI)
	MilePerSecondSquared = newUnit("mile per second squared", "mi/s²", Acceleration, BI)
)

func init() {
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
