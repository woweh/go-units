package units

const Mass UnitQuantity = "mass"

var (
	_mass = Quantity(Mass)

	// metric
	Gram      = mustCreateNewUnit("gram", "g", _mass, SI)
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
	MetricTon = mustCreateNewUnit("metric ton", "t", _mass, SI)

	// imperial
	Grain  = mustCreateNewUnit("grain", "gr", _mass, BI)
	Drachm = mustCreateNewUnit("drachm", "dr", _mass, BI)
	Ounce  = mustCreateNewUnit("ounce", "oz", _mass, BI)
	Pound  = mustCreateNewUnit("pound", "lb", _mass, BI)
	Stone  = mustCreateNewUnit("stone", "st", _mass, BI)
	Ton    = mustCreateNewUnit("ton", "LT", _mass, BI)
	Slug   = mustCreateNewUnit("slug", "", _mass, BI)

	ShortTon = mustCreateNewUnit("short ton", "short ton", _mass, BI)
)

func init() {
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
