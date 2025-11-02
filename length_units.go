package units

var (
	Length = Quantity("length")

	// metric
	Meter      = mustCreateNewUnit("meter", "m", Length, SI)
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

	Angstrom = mustCreateNewUnit("angstrom", "Å", Length, BI)
	Inch     = mustCreateNewUnit("inch", "in", Length, BI, Plural("inches"))
	Foot     = mustCreateNewUnit("foot", "ft", Length, BI, Plural("feet"))
	Yard     = mustCreateNewUnit("yard", "yd", Length, BI)
	Mile     = mustCreateNewUnit("mile", "mi", Length, BI)
	League   = mustCreateNewUnit("league", "lea", Length, BI)
	Furlong  = mustCreateNewUnit("furlong", "fur", Length, BI)

	USSurveyFoot = mustCreateNewUnit("US survey foot", "USft", Length, BI)
)

func init() {
	NewRatioConversion(Angstrom, Meter, 0.0000000001)
	NewRatioConversion(Inch, Meter, 0.0254)
	NewRatioConversion(Foot, Meter, 0.3048)
	NewRatioConversion(Yard, Meter, 0.9144)
	NewRatioConversion(Mile, Meter, 1609.344)
	NewRatioConversion(League, Meter, 4828.032)
	NewRatioConversion(Furlong, Meter, 201.168)
	NewRatioConversion(USSurveyFoot, Meter, 1200.0/3937.0)

	Meter.AddAliases("metre")
	Angstrom.AddAliases("ångström", "angstroms", "ångströms")
	Inch.AddSymbols("in.", "″")
	Foot.AddSymbols("ft.", "′")
	Yard.AddSymbols("yd.")
	Mile.AddSymbols("mi.")
}
