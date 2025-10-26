package units

var (
	Power = Quantity("power")

	// metric
	Watt      = newUnit("watt", "W", Power, SI)
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

	VoltAmpere     = newUnit("volt-ampere", "V⋅A", Power, SI)
	KiloVoltAmpere = Kilo(VoltAmpere)
	MegaVoltAmpere = Mega(VoltAmpere)

	VoltAmpereReactive     = newUnit("volt-ampere reactive", "var", Power, SI)
	KiloVoltAmpereReactive = Kilo(VoltAmpereReactive)
	MegaVoltAmpereReactive = Mega(VoltAmpereReactive)

	BritishThermalUnitPerHour          = newUnit("British thermal unit per hour", "Btu/h", Power)
	BritishThermalUnitPerSecond        = newUnit("British thermal unit per second", "Btu/s", Power)
	CaloriePerSecond                   = newUnit("calorie per second", "cal/s", Power)
	Horsepower                         = newUnit("horsepower", "hp", Power)
	KiloCaloriePerSecond               = newUnit("kilocalorie per second", "kcal/s", Power)
	ThousandBritishThermalUnitsPerHour = newUnit("thousand British thermal units per hour", "MBH", Power)
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
