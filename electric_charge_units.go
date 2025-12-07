package units

var (
	// ElectricCharge is the unit quantity for electric charge.
	ElectricCharge = NewUnitQuantity("electric charge")

	// SI unit metric
	Coulomb      = ElectricCharge.MustCreateUnit("coulomb", "C", SI)
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

	AmpereHour      = ElectricCharge.MustCreateUnit("ampere-hour", "A·h", SI, Symbols("A⋅h", "A*h", "A.h", "Ah", "AHr"))
	KiloAmpereHour  = Kilo(AmpereHour)
	MilliAmpereHour = Milli(AmpereHour)

	AmpereMinute = ElectricCharge.MustCreateUnit(
		"ampere-minute", "A·min", SI, Symbols("A⋅min", "A*min", "A.min", "Amin"),
	)
	KiloAmpereMinute  = Kilo(AmpereMinute)
	MilliAmpereMinute = Milli(AmpereMinute)

	AmpereSecond      = ElectricCharge.MustCreateUnit("ampere-second", "A·s", SI, Symbols("A⋅s", "A*s", "A.s", "As"))
	KiloAmpereSecond  = Kilo(AmpereSecond)
	MilliAmpereSecond = Milli(AmpereSecond)
)

func initElectricChargeUnits() {
	NewRatioConversion(AmpereHour, Coulomb, 3600)
	NewRatioConversion(AmpereMinute, Coulomb, 60)
	NewRatioConversion(AmpereSecond, Coulomb, 1)
}
