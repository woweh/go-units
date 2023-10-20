package units

var (
	ElectricCharge = UnitOptionQuantity("electric charge")

	// SI unit metric
	Coulomb      = newUnit("coulomb", "C", ElectricCharge, SI)
	ExaCoulomb   = Exa(Coulomb)
	PetaCoulomb  = Peta(Coulomb)
	TeraCoulomb  = Tera(Coulomb)
	GigaCoulomb  = Giga(Coulomb)
	MegaCoulomb  = Mega(Coulomb)
	KiloCoulomb  = Kilo(Coulomb)
	HectoCoulomb = Hecto(Coulomb)
	DecaCoulomb  = Deca(Coulomb)
	DeciCoulomb  = Deci(Coulomb)
	CentiCoulomb = Centi(Coulomb)
	MilliCoulomb = Milli(Coulomb)
	MicroCoulomb = Micro(Coulomb)
	NanoCoulomb  = Nano(Coulomb)
	PicoCoulomb  = Pico(Coulomb)
	FemtoCoulomb = Femto(Coulomb)
	AttoCoulomb  = Atto(Coulomb)

	AmpereHour      = newUnit("ampere-hour", "A·h", ElectricCharge, SI)
	KiloAmpereHour  = Kilo(AmpereHour)
	MilliAmpereHour = Milli(AmpereHour)

	AmpereMinute      = newUnit("ampere-minute", "A·min", ElectricCharge, SI)
	KiloAmpereMinute  = Kilo(AmpereMinute)
	MilliAmpereMinute = Milli(AmpereMinute)

	AmpereSecond      = newUnit("ampere-second", "A·s", ElectricCharge, SI)
	KiloAmpereSecond  = Kilo(AmpereSecond)
	MilliAmpereSecond = Milli(AmpereSecond)
)

func init() {
	NewRatioConversion(AmpereHour, Coulomb, 3600)
	AmpereHour.AddSymbols("A⋅h", "A*h", "A.h", "Ah", "AHr")
	KiloAmpereHour.AddSymbols("kA⋅h", "kA*h", "kA.h", "kAh", "kAHr")
	MilliAmpereHour.AddSymbols("mA⋅h", "mA*h", "mA.h", "mAh", "mAHr")

	NewRatioConversion(AmpereMinute, Coulomb, 60)
	AmpereMinute.AddSymbols("A⋅min", "A*min", "A.min", "Amin")
	KiloAmpereMinute.AddSymbols("kA⋅min", "kA*min", "kA.min", "kAmin")
	MilliAmpereMinute.AddSymbols("mA⋅min", "mA*min", "mA.min", "mAmin")

	NewRatioConversion(AmpereSecond, Coulomb, 1)
	AmpereSecond.AddSymbols("A⋅s", "A*s", "A.s", "As")
	KiloAmpereSecond.AddSymbols("kA⋅s", "kA*s", "kA.s", "kAs")
	MilliAmpereSecond.AddSymbols("mA⋅s", "mA*s", "mA.s", "mAs")
}
