package units

var (
	Force = Quantity("force")

	Newton      = newUnit("newton", "N", Force, BaseSiUnit)
	CentiNewton = Centi(Newton)
	DeciNewton  = Deci(Newton)
	MilliNewton = Milli(Newton)
	MicroNewton = Micro(Newton)
	NanoNewton  = Nano(Newton)
	PicoNewton  = Pico(Newton)
	FemtoNewton = Femto(Newton)
	AttoNewton  = Atto(Newton)
	ZeptoNewton = Zepto(Newton)
	YoctoNewton = Yocto(Newton)
	DecaNewton  = Deca(Newton)
	HectoNewton = Hecto(Newton)
	KiloNewton  = Kilo(Newton)
	MegaNewton  = Mega(Newton)
	GigaNewton  = Giga(Newton)
	TeraNewton  = Tera(Newton)
	PetaNewton  = Peta(Newton)
	ExaNewton   = Exa(Newton)
	ZettaNewton = Zetta(Newton)
	YottaNewton = Yotta(Newton)

	PoundForce = newUnit("pound force", "lbf", Force, BI, Plural(PluralNone))

	Dyne = newUnit("dyne", "dyn", Force, CGS)

	Poundal = newUnit("poundal", "pdl", Force, US)

	KilogramForce = newUnit("kilogram-force", "kgf", Force, MKpS)
	TonneForce    = newUnit("tonne-force", "tf", Force, MKpS)

	Kip           = Kilo(PoundForce)
	ShortTonForce = newUnit("short ton force", "Tons", Force, BI)
)

func init() {
	// https://qudt.org/vocab/unit/LB_F  4.448222
	NewRatioConversion(PoundForce, Newton, 4.448222)

	// https://en.wikipedia.org/wiki/Dyne?oldid=494703827
	NewRatioConversion(Dyne, Newton, 1e-5)

	// https://en.wikipedia.org/w/index.php?title=Poundal&oldid=1168735176
	NewRatioConversion(Poundal, Newton, 0.138254954376)

	// https://en.wikipedia.org/w/index.php?title=Kilogram-force&oldid=1159132202
	NewRatioConversion(KilogramForce, Newton, 9.80665)
	KilogramForce.AddAliases("kilopond", "kilogramme-force")
	KilogramForce.AddSymbols("kp")

	// https://en.wikipedia.org/w/index.php?title=Ton-force&oldid=1170043394#Tonne-force
	NewRatioConversion(TonneForce, Newton, 9806.65)
	TonneForce.AddAliases("metric ton-force", "megagram-force", "megapond")
	TonneForce.AddSymbols("Mp")

	NewRatioConversion(ShortTonForce, Newton, 2000*4.448222)
}
