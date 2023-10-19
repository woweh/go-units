package units

var (
	Power = UnitOptionQuantity("power")

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

	VoltAmpere     = newUnit("volt-ampere", "VA", Power, SI)
	KiloVoltAmpere = Kilo(VoltAmpere)
	MegaVoltAmpere = Mega(VoltAmpere)
)

func init() {
	VoltAmpere.AddSymbols("V⋅A", "V*A", "V.A", "V A")
}
