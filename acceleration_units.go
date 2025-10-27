package units

var (
	Acceleration = Quantity("acceleration")

	// Base unit: foot per second squared (Revit base with CF=1.0)
	FootPerSecondSquared = newUnit("foot per second squared", "ft/s²", Acceleration, BI)

	// Other units
	InchPerSecondSquared      = newUnit("inch per second squared", "in/s²", Acceleration, BI)
	MeterPerSecondSquared     = newUnit("meter per second squared", "m/s²", Acceleration, SI)
	KilometerPerSecondSquared = newUnit("kilometer per second squared", "km/s²", Acceleration, SI)
	MilePerSecondSquared      = newUnit("mile per second squared", "mi/s²", Acceleration, BI)
)

func init() {
	// From RevitUnits.json:
	// ft/s² (base): CF = 1.0
	// in/s²: CF = 11.999999999999998
	// m/s²: CF = 0.3048
	// km/s²: CF = 0.00030480000000000004
	// mi/s²: CF = 0.00018939393939393937
	// This means: 1 ft/s² = 12 in/s², 1 ft/s² = 0.3048 m/s², etc.
	NewRatioConversion(FootPerSecondSquared, InchPerSecondSquared, 12.0)
	NewRatioConversion(FootPerSecondSquared, MeterPerSecondSquared, 0.3048)
	NewRatioConversion(FootPerSecondSquared, KilometerPerSecondSquared, 0.0003048)
	NewRatioConversion(FootPerSecondSquared, MilePerSecondSquared, 0.00018939393939393937)

	FootPerSecondSquared.AddAliases("feet per second squared", "ft/sec²", "ft/sec/sec")
	InchPerSecondSquared.AddAliases("inches per second squared", "in/sec²", "in/sec/sec")
	MeterPerSecondSquared.AddAliases("meters per second squared", "metre per second squared", "metres per second squared", "m/sec²")
	KilometerPerSecondSquared.AddAliases("kilometers per second squared", "kilometre per second squared", "kilometres per second squared", "km/sec²")
	MilePerSecondSquared.AddAliases("miles per second squared", "mi/sec²")
}
