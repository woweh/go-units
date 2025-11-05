package units

// Voltage is a unit quantity for voltage
const VoltageQuantity UnitQuantity = "voltage"

var (
	_voltage = Quantity(VoltageQuantity)

	// SI base unit: volt
	Volt      = mustCreateNewUnit("volt", "V", _voltage, SI)
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
