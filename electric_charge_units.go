package units

const ElectricCharge UnitQuantity = "electric charge"

var (
	_electricCharge = Quantity(ElectricCharge)

	// SI unit metric
	Coulomb      = mustCreateNewUnit("coulomb", "C", _electricCharge, SI)
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

	AmpereHour      = mustCreateNewUnit("ampere-hour", "A·h", _electricCharge, SI, Symbols("A⋅h", "A*h", "A.h", "Ah", "AHr"))
	KiloAmpereHour  = Kilo(AmpereHour)
	MilliAmpereHour = Milli(AmpereHour)

	AmpereMinute = mustCreateNewUnit(
		"ampere-minute", "A·min", _electricCharge, SI, Symbols("A⋅min", "A*min", "A.min", "Amin"),
	)
	KiloAmpereMinute  = Kilo(AmpereMinute)
	MilliAmpereMinute = Milli(AmpereMinute)

	AmpereSecond      = mustCreateNewUnit("ampere-second", "A·s", _electricCharge, SI, Symbols("A⋅s", "A*s", "A.s", "As"))
	KiloAmpereSecond  = Kilo(AmpereSecond)
	MilliAmpereSecond = Milli(AmpereSecond)
)

func initElectricChargeUnits() {
	NewRatioConversion(AmpereHour, Coulomb, 3600)
	NewRatioConversion(AmpereMinute, Coulomb, 60)
	NewRatioConversion(AmpereSecond, Coulomb, 1)
}
