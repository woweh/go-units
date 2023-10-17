package units

var (
	Power = UnitOptionQuantity("power")

	// metric
	Watt      = newUnit("watt", "W", Power, SI)
	KiloWatt  = Kilo(Watt)
	MegaWatt  = Mega(Watt)
	GigaWatt  = Giga(Watt)
	TeraWatt  = Tera(Watt)
	PetaWatt  = Peta(Watt)
	ExaWatt   = Exa(Watt)
	ZettaWatt = Zetta(Watt)
	YottaWatt = Yotta(Watt)
)
