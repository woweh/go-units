package units

var (
	Pressure = UnitOptionQuantity("pressure")

	// SI unit metric
	Pascal      = newUnit("pascal", "Pa", Pressure, SI)
	ExaPascal   = Exa(Pascal)
	PetaPascal  = Peta(Pascal)
	TeraPascal  = Tera(Pascal)
	GigaPascal  = Giga(Pascal)
	MegaPascal  = Mega(Pascal)
	KiloPascal  = Kilo(Pascal)
	HectoPascal = Hecto(Pascal)
	DecaPascal  = Deca(Pascal)
	DeciPascal  = Deci(Pascal)
	CentiPascal = Centi(Pascal)
	MilliPascal = Milli(Pascal)
	MicroPascal = Micro(Pascal)
	NanoPascal  = Nano(Pascal)
	PicoPascal  = Pico(Pascal)
	FemtoPascal = Femto(Pascal)
	AttoPascal  = Atto(Pascal)

	// Other
	At       = newUnit("technical atmosphere", "at", Pressure, BI, UnitOptionPlural(Custom, "technical atmospheres"))
	Atm      = newUnit("standard atmosphere", "atm", Pressure, BI, UnitOptionPlural(Custom, "standard atmospheres"))
	Bar      = newUnit("bar", "bar", Pressure, BI, UnitOptionPlural(Custom, "bars"))
	CentiBar = Centi(Bar)
	MilliBar = Milli(Bar)
	MicroBar = Micro(Bar)
	Barye    = newUnit("barye", "Ba", Pressure, BI, UnitOptionPlural(Custom, "baryes"))
	InH2O    = newUnit("inch of Water Column", "inH2O", Pressure, BI)
	InHg     = newUnit("inch of Mercury", "inHg", Pressure, BI)
	MH2O     = newUnit(
		"meter of Water Column", "mmH2O", Pressure, BI, UnitOptionPlural(Custom, "meters of Water Column"),
	)
	MmH2O  = Milli(MH2O)
	CmH2O  = Centi(MH2O)
	MHg    = newUnit("meter of Mercury", "mmHg", Pressure, BI, UnitOptionPlural(Custom, "meters of Mercury"))
	MmHg   = Milli(MHg)
	CmHg   = Centi(MHg)
	Newton = newUnit("newton per square meter", "N/m²", Pressure, BI)
	Psi    = newUnit("pound-force per square inch", "psi", Pressure, BI)
	Torr   = newUnit("torr", "Torr", Pressure, BI)
)

func init() {
	NewRatioConversion(At, Pascal, 98066.5)
	NewRatioConversion(Atm, Pascal, 101325.2738)
	NewRatioConversion(Bar, Pascal, 100000)
	NewRatioConversion(Barye, Pascal, 0.1)
	NewRatioConversion(InH2O, Pascal, 248.84)
	NewRatioConversion(InHg, Pascal, 3386.38815789)
	NewRatioConversion(MH2O, Pascal, 9806.65)
	NewRatioConversion(MHg, Pascal, 133.322368421)
	NewRatioConversion(Newton, Pascal, 1)
	NewRatioConversion(Psi, Pascal, 6894.757)
	NewRatioConversion(Torr, Pascal, 133.322368421)
}
