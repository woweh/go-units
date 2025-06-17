package units

var (
	ElectricCharge = Quantity("electric charge")

	// SI unit metric
	Coulomb      = newUnit("coulomb", "C", ElectricCharge, SI, BaseUnit)
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

	AmpereHour      = newUnit("ampere-hour", "A·h", ElectricCharge, SI, Symbols("A⋅h", "A*h", "A.h", "Ah", "AHr"), BaseUnit)
	KiloAmpereHour  = Kilo(AmpereHour)
	MilliAmpereHour = Milli(AmpereHour)

	AmpereMinute = newUnit(
		"ampere-minute", "A·min", ElectricCharge, SI, BaseUnit, Symbols("A⋅min", "A*min", "A.min", "Amin"),
	)
	KiloAmpereMinute  = Kilo(AmpereMinute)
	MilliAmpereMinute = Milli(AmpereMinute)

	AmpereSecond      = newUnit("ampere-second", "A·s", ElectricCharge, SI, BaseUnit, Symbols("A⋅s", "A*s", "A.s", "As"))
	KiloAmpereSecond  = Kilo(AmpereSecond)
	MilliAmpereSecond = Milli(AmpereSecond)
)

func init() {
	NewRatioConversion(AmpereHour, Coulomb, 3600)
	NewRatioConversion(AmpereMinute, Coulomb, 60)
	NewRatioConversion(AmpereSecond, Coulomb, 1)
}
