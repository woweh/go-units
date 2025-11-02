package units

var (
	Mass = Quantity("mass")

	// metric
	Gram      = mustCreateNewUnit("gram", "g", Mass, SI)
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

	// imperial
	Grain  = mustCreateNewUnit("grain", "gr", Mass, BI)
	Drachm = mustCreateNewUnit("drachm", "dr", Mass, BI)
	Ounce  = mustCreateNewUnit("ounce", "oz", Mass, BI)
	Pound  = mustCreateNewUnit("pound", "lb", Mass, BI)
	Stone  = mustCreateNewUnit("stone", "st", Mass, BI)
	Ton    = mustCreateNewUnit("ton", "LT", Mass, BI)
	Slug   = mustCreateNewUnit("slug", "", Mass, BI)

	ShortTon = mustCreateNewUnit("short ton", "short ton", Mass, BI)
)

func init() {
	MegaGram.AddAliases("tonne", "tonnes", "metric ton", "metric tons")
	MegaGram.AddAliases("t")

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
