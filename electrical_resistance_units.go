package units

var (
	ElectricalResistance = Quantity("electrical resistance")

	Ohm      = newUnit("ohm", "Ω", ElectricalResistance, SI, Symbols("Ω"))
	DecaOhm  = Deca(Ohm)
	HectoOhm = Hecto(Ohm)
	KiloOhm  = Kilo(Ohm)
	MegaOhm  = Mega(Ohm)
	GigaOhm  = Giga(Ohm)
	TeraOhm  = Tera(Ohm)
	PetaOhm  = Peta(Ohm)
	ExaOhm   = Exa(Ohm)
	ZettaOhm = Zetta(Ohm)
	YottaOhm = Yotta(Ohm)
	DeciOhm  = Deci(Ohm)
	CentiOhm = Centi(Ohm)
	MilliOhm = Milli(Ohm)
	MicroOhm = Micro(Ohm)
	NanoOhm  = Nano(Ohm)
	PicoOhm  = Pico(Ohm)
	FemtoOhm = Femto(Ohm)
	AttoOhm  = Atto(Ohm)
	ZeptoOhm = Zepto(Ohm)
	YoctoOhm = Yocto(Ohm)
)
