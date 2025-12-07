package units

var (
	// Power is the unit quantity for power.
	Power = NewUnitQuantity("power")

	// SI base unit: watt
	Watt      = Power.MustCreateUnit("watt", "W", SI)
	DeciWatt  = Deci(Watt)
	CentiWatt = Centi(Watt)
	MilliWatt = Milli(Watt)
	MicroWatt = Micro(Watt)
	NanoWatt  = Nano(Watt)
	PicoWatt  = Pico(Watt)
	FemtoWatt = Femto(Watt)
	AttoWatt  = Atto(Watt)
	ZeptoWatt = Zepto(Watt)
	YoctoWatt = Yocto(Watt)
	DecaWatt  = Deca(Watt)
	HectoWatt = Hecto(Watt)
	KiloWatt  = Kilo(Watt)
	MegaWatt  = Mega(Watt)
	GigaWatt  = Giga(Watt)
	TeraWatt  = Tera(Watt)
	PetaWatt  = Peta(Watt)
	ExaWatt   = Exa(Watt)
	ZettaWatt = Zetta(Watt)
	YottaWatt = Yotta(Watt)

	VoltAmpere     = Power.MustCreateUnit("volt-ampere", "V⋅A", SI)
	KiloVoltAmpere = Kilo(VoltAmpere)
	MegaVoltAmpere = Mega(VoltAmpere)

	VoltAmpereReactive     = Power.MustCreateUnit("volt-ampere reactive", "var", SI)
	KiloVoltAmpereReactive = Kilo(VoltAmpereReactive)
	MegaVoltAmpereReactive = Mega(VoltAmpereReactive)

	// Imperial/US and other units
	BritishThermalUnitPerHour          = Power.MustCreateUnit("British thermal unit per hour", "Btu/h", BI)
	BritishThermalUnitPerSecond        = Power.MustCreateUnit("British thermal unit per second", "Btu/s", BI)
	CaloriePerSecond                   = Power.MustCreateUnit("calorie per second", "cal/s")
	Horsepower                         = Power.MustCreateUnit("horsepower", "hp")
	KiloCaloriePerSecond               = Power.MustCreateUnit("kilocalorie per second", "kcal/s", SI)
	ThousandBritishThermalUnitsPerHour = Power.MustCreateUnit("thousand British thermal units per hour", "MBH", BI)
)

func initPowerUnits() {
	// Electrical power units (equivalent by definition)
	NewRatioConversion(Watt, VoltAmpere, 1)
	NewRatioConversion(Watt, VoltAmpereReactive, 1)

	VoltAmpere.AddAliases("volt ampere")
	VoltAmpere.AddSymbols("VA", "V*A", "V.A", "V A")

	VoltAmpereReactive.AddAliases("volt ampere reactive")
	VoltAmpereReactive.AddSymbols("VAR", "VAr", "V⋅Ar", "V.A.r", "V A r", "V.A.r")

	// NOTE: Revit uses different BTU and calorie definitions than standard thermochemical values
	// For interoperability with Revit, these conversions use Revit-specific factors
	// Standard calculation would be: powerFactor(BritishThermalUnit, Hour)
	// but Revit BTU appears to differ from thermochemical BTU (1055.056 J)
	NewRatioConversion(BritishThermalUnitPerHour, Watt, 0.316998330628151)
	NewRatioConversion(BritishThermalUnitPerSecond, Watt, 11356.526682227)
	NewRatioConversion(CaloriePerSecond, Watt, 45.0663401326803)
	NewRatioConversion(KiloCaloriePerSecond, Watt, 45066.3401326803)
	NewRatioConversion(ThousandBritishThermalUnitsPerHour, Watt, 316.998330628151)

	// Horsepower: Revit uses a specific definition that differs from mechanical horsepower (745.7 W)
	NewRatioConversion(Horsepower, Watt, 8026.6466154635)
}

// powerFactor calculates the conversion factor for power units.
// Power = energy / time
// Base unit: Watt = J/s = kg⋅m²/s³
//
// Example: BTU/hour
// - energyRatio: how many J in 1 BTU
// - timeRatio: how many s in 1 hour
// - powerFactor = energyRatio / timeRatio
func powerFactor(energy, time Unit) float64 {
	// How many Joules in 1 unit of the target energy
	energyRatio := energy.to(Joule)
	// How many seconds in 1 unit of target time
	timeRatio := time.to(Second)
	// Power factor: energy per time
	return energyRatio / timeRatio
}
