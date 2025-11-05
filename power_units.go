package units

// Power is a unit quantity for power
const Power UnitQuantity = "power"

var (
	_power = Quantity(Power)

	// SI base unit: watt
	Watt      = mustCreateNewUnit("watt", "W", _power, SI)
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

	VoltAmpere     = mustCreateNewUnit("volt-ampere", "V⋅A", _power, SI)
	KiloVoltAmpere = Kilo(VoltAmpere)
	MegaVoltAmpere = Mega(VoltAmpere)

	VoltAmpereReactive     = mustCreateNewUnit("volt-ampere reactive", "var", _power, SI)
	KiloVoltAmpereReactive = Kilo(VoltAmpereReactive)
	MegaVoltAmpereReactive = Mega(VoltAmpereReactive)

	// Imperial/US and other units
	BritishThermalUnitPerHour          = mustCreateNewUnit("British thermal unit per hour", "Btu/h", _power, BI)
	BritishThermalUnitPerSecond        = mustCreateNewUnit("British thermal unit per second", "Btu/s", _power, BI)
	CaloriePerSecond                   = mustCreateNewUnit("calorie per second", "cal/s", _power, SI)
	Horsepower                         = mustCreateNewUnit("horsepower", "hp", _power, BI)
	KiloCaloriePerSecond               = mustCreateNewUnit("kilocalorie per second", "kcal/s", _power, SI)
	ThousandBritishThermalUnitsPerHour = mustCreateNewUnit("thousand British thermal units per hour", "MBH", _power, BI)
)

func init() {
	NewRatioConversion(Watt, VoltAmpere, 1)
	NewRatioConversion(Watt, VoltAmpereReactive, 1)

	VoltAmpere.AddAliases("volt ampere")
	VoltAmpere.AddSymbols("VA", "V*A", "V.A", "V A")

	VoltAmpereReactive.AddAliases("volt ampere reactive")
	VoltAmpereReactive.AddSymbols("VAR", "VAr", "V⋅Ar", "V.A.r", "V A r", "V.A.r")

	NewRatioConversion(BritishThermalUnitPerHour, Watt, 0.2930710701722222)
	NewRatioConversion(BritishThermalUnitPerSecond, Watt, 293.0710701722222)
	NewRatioConversion(CaloriePerSecond, Watt, 4.184)
	NewRatioConversion(Horsepower, Watt, 745.6998715822702)
	NewRatioConversion(KiloCaloriePerSecond, Watt, 4184)
	NewRatioConversion(ThousandBritishThermalUnitsPerHour, Watt, 2930.710701722222)
}
