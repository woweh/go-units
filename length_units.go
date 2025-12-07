package units

var (
	// Length is the unit quantity for length.
	Length = NewUnitQuantity("length")

	// metric
	Meter       = Length.MustCreateUnit("meter", "m", SI)
	QuettaMeter = Quetta(Meter)
	RonnaMeter  = Ronna(Meter)
	ZettaMeter  = Zetta(Meter)
	ExaMeter    = Exa(Meter)
	PetaMeter   = Peta(Meter)
	TeraMeter   = Tera(Meter)
	GigaMeter   = Giga(Meter)
	MegaMeter   = Mega(Meter)
	KiloMeter   = Kilo(Meter)
	HectoMeter  = Hecto(Meter)
	DecaMeter   = Deca(Meter)
	DeciMeter   = Deci(Meter)
	CentiMeter  = Centi(Meter)
	MilliMeter  = Milli(Meter)
	MicroMeter  = Micro(Meter)
	NanoMeter   = Nano(Meter)
	PicoMeter   = Pico(Meter)
	FemtoMeter  = Femto(Meter)
	AttoMeter   = Atto(Meter)
	ZeptoMeter  = Zepto(Meter)
	YoctoMeter  = Yocto(Meter)
	RontoMeter  = Ronto(Meter)
	QuectoMeter = Quecto(Meter)

	Angstrom = Length.MustCreateUnit("angstrom", "Å", SI)

	Inch    = Length.MustCreateUnit("inch", "in", BI, Plural("inches"))
	Foot    = Length.MustCreateUnit("foot", "ft", BI, Plural("feet"))
	Yard    = Length.MustCreateUnit("yard", "yd", BI)
	Mile    = Length.MustCreateUnit("mile", "mi", BI)
	League  = Length.MustCreateUnit("league", "lea", BI)
	Furlong = Length.MustCreateUnit("furlong", "fur", BI)

	USSurveyFoot = Length.MustCreateUnit("US survey foot", "USft", US)
)

func initLengthUnits() {
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
