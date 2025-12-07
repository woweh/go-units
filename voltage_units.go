package units

// Voltage is a unit quantity for voltage
var (
	// VoltageQuantity is the unit quantity for voltage.
	VoltageQuantity = NewUnitQuantity("voltage")

	// SI base unit: volt
	Volt      = VoltageQuantity.MustCreateUnit("volt", "V", SI)
	YottaVolt = Yotta(Volt)
	ZettaVolt = Zetta(Volt)
	ExaVolt   = Exa(Volt)
	PetaVolt  = Peta(Volt)
	TeraVolt  = Tera(Volt)
	GigaVolt  = Giga(Volt)
	MegaVolt  = Mega(Volt)
	KiloVolt  = Kilo(Volt)
	HectoVolt = Hecto(Volt)
	DecaVolt  = Deca(Volt)
	DeciVolt  = Deci(Volt)
	CentiVolt = Centi(Volt)
	MilliVolt = Milli(Volt)
	MicroVolt = Micro(Volt)
	NanoVolt  = Nano(Volt)
	PicoVolt  = Pico(Volt)
)
