package units

var (
	MassPerTime = Quantity("mass per time")

	// Base unit: kilogram per second (Revit base with CF=1.0)
	KilogramPerSecond = mustCreateNewUnit("kilogram per second", "kg/s", MassPerTime, SI)

	// Other SI units
	KilogramPerMinute = mustCreateNewUnit("kilogram per minute", "kg/min", MassPerTime, SI)
	KilogramPerHour   = mustCreateNewUnit("kilogram per hour", "kg/h", MassPerTime, SI)

	// Imperial units
	PoundPerSecond = mustCreateNewUnit("pound per second", "lb/s", MassPerTime, BI)
	PoundPerMinute = mustCreateNewUnit("pound per minute", "lb/min", MassPerTime, BI)
	PoundPerHour   = mustCreateNewUnit("pound per hour", "lb/h", MassPerTime, BI)
)

func init() {
	// From RevitUnits.json:
	// kg/s: CF = 1.0
	// kg/min: CF = 60.0
	// kg/h: CF = 3600.0
	// lb/s: CF = 2.2046226218487757
	// lb/min: CF = 132.27735731092656
	// lb/h: CF = 7936.6414386555925

	NewRatioConversion(KilogramPerSecond, KilogramPerMinute, 60.0)
	NewRatioConversion(KilogramPerSecond, KilogramPerHour, 3600.0)
	NewRatioConversion(KilogramPerSecond, PoundPerSecond, 2.2046226218487757)
	NewRatioConversion(KilogramPerSecond, PoundPerMinute, 132.27735731092656)
	NewRatioConversion(KilogramPerSecond, PoundPerHour, 7936.6414386555925)

	KilogramPerSecond.AddAliases("kilograms per second", "kg per second")
	KilogramPerMinute.AddAliases("kilograms per minute", "kg per minute")
	KilogramPerHour.AddAliases("kilograms per hour", "kg per hour")
	PoundPerSecond.AddAliases("pounds per second", "pounds mass per second", "lbm/s")
	PoundPerMinute.AddAliases("pounds per minute", "pounds mass per minute", "lbm/min")
	PoundPerHour.AddAliases("pounds per hour", "pounds mass per hour", "lbm/h")
}
