package units

var (
	// Force is the unit quantity for force.
	Force = NewUnitQuantity("force")

	Newton      = Force.MustCreateUnit("newton", "N", SI)
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

	PoundForce = Force.MustCreateUnit("pound force", "lbf", BI, Plural(PluralNone))

	Dyne = Force.MustCreateUnit("dyne", "dyn", CGS)

	Poundal = Force.MustCreateUnit("poundal", "pdl")

	KilogramForce = Force.MustCreateUnit("kilogram-force", "kgf", MKpS)
	TonneForce    = Force.MustCreateUnit("tonne-force", "tf", MKpS)

	Kip           = Kilo(PoundForce)
	ShortTonForce = Force.MustCreateUnit("short ton force", "stf", BI)
)

func initForceUnits() {
	// Standard force definitions (remain hardcoded per standards)
	// https://qudt.org/vocab/unit/LB_F  4.448222
	NewRatioConversion(PoundForce, Newton, 4.448222)

	// https://en.wikipedia.org/wiki/Dyne?oldid=494703827
	NewRatioConversion(Dyne, Newton, 1e-5)

	// Poundal can be calculated: 1 pdl = 1 lb⋅ft/s²
	// Newton = kg⋅m/s², so we need the inverse of the force factor
	// https://en.wikipedia.org/w/index.php?title=Poundal&oldid=1168735176
	NewRatioConversion(Poundal, Newton, 1.0/forceFactor(Pound, Foot))

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

// forceFactor calculates the conversion factor for force units.
// Force = mass × acceleration
// Base unit: Newton = kg⋅m/s²
//
// Example: To convert N to poundal (lb⋅ft/s²):
// - massRatio: how many pounds in 1 kg
// - accelRatio: how many ft/s² in 1 m/s²
// - forceFactor = massRatio × accelRatio
func forceFactor(mass, length Unit) float64 {
	// How many of the target mass unit in 1 kg
	massRatio := KiloGram.to(mass)
	// How many of the target acceleration unit in 1 m/s²
	accelRatio := accelFactor(length)
	// Force factor accounts for both mass and acceleration scaling
	return massRatio * accelRatio
}
