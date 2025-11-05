package units

const Length UnitQuantity = "length"

var (
	_length = Quantity(Length)

	// metric
	Meter      = mustCreateNewUnit("meter", "m", _length, SI)
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

	Angstrom = mustCreateNewUnit("angstrom", "Å", _length, SI)

	Inch    = mustCreateNewUnit("inch", "in", _length, BI, Plural("inches"))
	Foot    = mustCreateNewUnit("foot", "ft", _length, BI, Plural("feet"))
	Yard    = mustCreateNewUnit("yard", "yd", _length, BI)
	Mile    = mustCreateNewUnit("mile", "mi", _length, BI)
	League  = mustCreateNewUnit("league", "lea", _length, BI)
	Furlong = mustCreateNewUnit("furlong", "fur", _length, BI)

	USSurveyFoot = mustCreateNewUnit("US survey foot", "USft", _length, US)
)

func init() {
	NewRatioConversion(Angstrom, Meter, 1e-10)
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
