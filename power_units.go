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
)

func init() {
	NewRatioConversion(Watt, VoltAmpere, 1)
	NewRatioConversion(Watt, VoltAmpereReactive, 1)

	VoltAmpere.AddAliases("volt ampere")
	VoltAmpere.AddSymbols("VA", "V*A", "V.A", "V A")

	VoltAmpereReactive.AddAliases("volt ampere reactive")
	VoltAmpereReactive.AddSymbols("VAR", "VAr", "V⋅Ar", "V.A.r", "V A r", "V.A.r")
}
