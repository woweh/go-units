package units

var (
	// MassPerTime is the unit quantity for mass per time.
	MassPerTime = NewUnitQuantity("mass per time")

	// SI base unit: kilogram per second (Revit base with CF=1.0)
	KilogramPerSecond = MassPerTime.MustCreateUnit("kilogram per second", "kg/s", SI)

	// Other SI units
	KilogramPerMinute = MassPerTime.MustCreateUnit("kilogram per minute", "kg/min", SI)
	KilogramPerHour   = MassPerTime.MustCreateUnit("kilogram per hour", "kg/h", SI)

	// Imperial units
	PoundPerSecond = MassPerTime.MustCreateUnit("pound per second", "lb/s", BI)
	PoundPerMinute = MassPerTime.MustCreateUnit("pound per minute", "lb/min", BI)
	PoundPerHour   = MassPerTime.MustCreateUnit("pound per hour", "lb/h", BI)
)

func initMassPerTimeUnits() {
	// Calculated mass flow rate conversions
	// Note: NewRatioConversion expects "how many target units in 1 source unit"
	// For kg/s to kg/min: 1 kg/s = 60 kg/min, so ratio = 1/massPerTimeFactor
	NewRatioConversion(KilogramPerSecond, KilogramPerMinute, 1.0/massPerTimeFactor(KiloGram, Minute))
	NewRatioConversion(KilogramPerSecond, KilogramPerHour, 1.0/massPerTimeFactor(KiloGram, Hour))
	NewRatioConversion(KilogramPerSecond, PoundPerSecond, 1.0/massPerTimeFactor(Pound, Second))
	NewRatioConversion(KilogramPerSecond, PoundPerMinute, 1.0/massPerTimeFactor(Pound, Minute))
	NewRatioConversion(KilogramPerSecond, PoundPerHour, 1.0/massPerTimeFactor(Pound, Hour))

	KilogramPerSecond.AddAliases("kilograms per second", "kg per second")
	KilogramPerMinute.AddAliases("kilograms per minute", "kg per minute")
	KilogramPerHour.AddAliases("kilograms per hour", "kg per hour")
	PoundPerSecond.AddAliases("pounds per second", "pounds mass per second", "lbm/s")
	PoundPerMinute.AddAliases("pounds per minute", "pounds mass per minute", "lbm/min")
	PoundPerHour.AddAliases("pounds per hour", "pounds mass per hour", "lbm/h")
}

// massPerTimeFactor calculates the conversion factor for mass flow rate units.
// MassPerTime = mass / time
// Base unit: KilogramPerSecond = kg/s
//
// Example: lb/min
// - massRatio: how many kg in 1 lb
// - timeRatio: how many s in 1 min
// - massPerTimeFactor = massRatio / timeRatio
func massPerTimeFactor(mass, time Unit) float64 {
	// How many kg in 1 unit of the target mass
	massRatio := mass.to(KiloGram)
	// How many seconds in 1 unit of target time
	timeRatio := time.to(Second)
	// Mass per time factor
	return massRatio / timeRatio
}
