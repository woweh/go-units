package units

const ElectricalResistance UnitQuantity = "electrical resistance"
const ElectricalResistivity UnitQuantity = "electrical resistivity"

var (
	_electricalResistance  = Quantity(ElectricalResistance)
	_electricalResistivity = Quantity(ElectricalResistivity)

	Ohm      = mustCreateNewUnit("ohm", "Ω", _electricalResistance, SI, Symbols("Ω"))
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

	// Resistivity
	OhmMeter = mustCreateNewUnit("ohm meter", "ohm·m", _electricalResistivity, SI, Aliases("ohm-meter"))
)
