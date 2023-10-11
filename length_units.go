package units

var (
	Length = UnitOptionQuantity("length")

	// metric
	Meter      = newUnit("meter", "m", Length, SI)
	ExaMeter   = Exa(Meter)
	PetaMeter  = Peta(Meter)
	TeraMeter  = Tera(Meter)
	GigaMeter  = Giga(Meter)
	MegaMeter  = Mega(Meter)
	KiloMeter  = Kilo(Meter)
	HectoMeter = Hecto(Meter)
	DecaMeter  = Deca(Meter)
	DeciMeter  = Deci(Meter)
	CentiMeter = Centi(Meter)
	MilliMeter = Milli(Meter)
	MicroMeter = Micro(Meter)
	NanoMeter  = Nano(Meter)
	PicoMeter  = Pico(Meter)
	FemtoMeter = Femto(Meter)
	AttoMeter  = Atto(Meter)

	Angstrom = newUnit("angstrom", "Å", Length, BI)
	Inch     = newUnit("inch", "in", Length, BI, UnitOptionPlural("inches"))
	Foot     = newUnit("foot", "ft", Length, BI, UnitOptionPlural("feet"))
	Yard     = newUnit("yard", "yd", Length, BI)
	Mile     = newUnit("mile", "mi", Length, BI)
	League   = newUnit("league", "lea", Length, BI)
	Furlong  = newUnit("furlong", "fur", Length, BI)
)

func init() {
	NewRatioConversion(Angstrom, Meter, 0.0000000001)
	NewRatioConversion(Inch, Meter, 0.0254)
	NewRatioConversion(Foot, Meter, 0.3048)
	NewRatioConversion(Yard, Meter, 0.9144)
	NewRatioConversion(Mile, Meter, 1609.344)
	NewRatioConversion(League, Meter, 4828.032)
	NewRatioConversion(Furlong, Meter, 201.168)

	Meter.AddAliases("metre")
	Angstrom.AddAliases("ångström", "angstroms", "ångströms")
	Inch.AddAliases("in.", "″", "\"")
	Foot.AddAliases("ft.", "′")
	Yard.AddAliases("yd.")
	Mile.AddAliases("mi.")
}
