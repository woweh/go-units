package units

var (
	// ElectricCurrent is the unit quantity for electric current.
	ElectricCurrent = NewUnitQuantity("electric current")

	// metric
	Ampere      = ElectricCurrent.MustCreateUnit("ampere", "A", SI)
	MilliAmpere = Milli(Ampere)
	MicroAmpere = Micro(Ampere)
	NanoAmpere  = Nano(Ampere)
	PicoAmpere  = Pico(Ampere)
	FemtoAmpere = Femto(Ampere)
	AttoAmpere  = Atto(Ampere)
	ZeptoAmpere = Zepto(Ampere)
	YoctoAmpere = Yocto(Ampere)
	KiloAmpere  = Kilo(Ampere)
	MegaAmpere  = Mega(Ampere)
	GigaAmpere  = Giga(Ampere)
	TeraAmpere  = Tera(Ampere)
	PetaAmpere  = Peta(Ampere)
	ExaAmpere   = Exa(Ampere)
	ZettaAmpere = Zetta(Ampere)
	YottaAmpere = Yotta(Ampere)
)
