package units

var (
	Mass = UnitOptionQuantity("mass")

	// metric
	Gram      = newUnit("gram", "g", Mass)
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
	Grain  = newUnit("grain", "gr", Mass, BI)
	Drachm = newUnit("drachm", "dr", Mass, BI)
	Ounce  = newUnit("ounce", "oz", Mass, BI)
	Pound  = newUnit("pound", "lb", Mass, BI)
	Stone  = newUnit("stone", "st", Mass, BI)
	Ton    = newUnit("ton", "t", Mass, BI)
	Slug   = newUnit("slug", "", Mass, BI)
)

func init() {
	NewRatioConversion(Grain, Gram, 0.06479891)
	NewRatioConversion(Drachm, Gram, 1.7718451953125)
	NewRatioConversion(Ounce, Gram, 28.349523125)
	NewRatioConversion(Pound, Gram, 453.59237)
	NewRatioConversion(Stone, Gram, 6350.29318)

	// 1 long ton
	NewRatioConversion(Ton, Gram, 1016046.9088)
	NewRatioConversion(Slug, Gram, 14593.90294)
}
