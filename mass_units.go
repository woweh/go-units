package units

var (
	// Mass is the unit quantity for mass.
	Mass = NewUnitQuantity("mass")

	// metric
	Gram      = Mass.MustCreateUnit("gram", "g", SI)
	ExaGram   = Exa(Gram)
	PetaGram  = Peta(Gram)
	TeraGram  = Tera(Gram)
	GigaGram  = Giga(Gram)
	MegaGram  = Mega(Gram)
	KiloGram  = Kilo(Gram)
	HectoGram = Hecto(Gram)
	DecaGram  = Deca(Gram)
	DeciGram  = Deci(Gram)
	CentiGram = Centi(Gram)
	MilliGram = Milli(Gram)
	MicroGram = Micro(Gram)
	NanoGram  = Nano(Gram)
	PicoGram  = Pico(Gram)
	FemtoGram = Femto(Gram)
	AttoGram  = Atto(Gram)

	// metric ton (tonne) as a distinct unit
	MetricTon = Mass.MustCreateUnit("metric ton", "t", SI)

	// imperial
	Grain  = Mass.MustCreateUnit("grain", "gr", BI)
	Drachm = Mass.MustCreateUnit("drachm", "dr", BI)
	Ounce  = Mass.MustCreateUnit("ounce", "oz", BI)
	Pound  = Mass.MustCreateUnit("pound", "lb", BI)
	Stone  = Mass.MustCreateUnit("stone", "st", BI)
	Ton    = Mass.MustCreateUnit("ton", "LT", BI)
	Slug   = Mass.MustCreateUnit("slug", "", BI)

	ShortTon = Mass.MustCreateUnit("short ton", "short ton", BI)
)

func initMassUnits() {
	// metric ton (tonne) as a distinct unit
	NewRatioConversion(MetricTon, Gram, 1e6) // == MegaGram
	MetricTon.AddAliases("tonne", "tonnes", "metric tons")

	NewRatioConversion(Grain, Gram, 0.06479891)
	NewRatioConversion(Drachm, Gram, 1.7718451953125)
	NewRatioConversion(Ounce, Gram, 28.349523125)
	NewRatioConversion(Pound, Gram, 453.59237)
	NewRatioConversion(Stone, Gram, 6350.29318)

	// 1 long ton
	// https://en.wikipedia.org/wiki/Ton
	// https://en.wikipedia.org/wiki/Long_ton
	NewRatioConversion(Ton, Gram, 1016046.9088)
	Ton.AddAliases("long ton", "long tons", "imperial ton", "displacement ton")
	NewRatioConversion(Slug, Gram, 14593.90294)
	NewRatioConversion(ShortTon, Gram, 907184.74)
	ShortTon.AddAliases("US ton", "short tons")
	ShortTon.AddSymbols("Tons")
}
