package units

const MassPerTime UnitQuantity = "mass per time"

var (
	_massPerTime = Quantity(MassPerTime)

	// SI base unit: kilogram per second (Revit base with CF=1.0)
	KilogramPerSecond = mustCreateNewUnit("kilogram per second", "kg/s", _massPerTime, SI)

	// Other SI units
	KilogramPerMinute = mustCreateNewUnit("kilogram per minute", "kg/min", _massPerTime, SI)
	KilogramPerHour   = mustCreateNewUnit("kilogram per hour", "kg/h", _massPerTime, SI)

	// Imperial units
	PoundPerSecond = mustCreateNewUnit("pound per second", "lb/s", _massPerTime, BI)
	PoundPerMinute = mustCreateNewUnit("pound per minute", "lb/min", _massPerTime, BI)
	PoundPerHour   = mustCreateNewUnit("pound per hour", "lb/h", _massPerTime, BI)
)

func initMassPerTimeUnits() {
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
