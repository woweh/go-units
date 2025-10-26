package units

var (
	Pressure = Quantity("pressure")

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
	At       = newUnit("technical atmosphere", "at", Pressure, BI)
	Atm      = newUnit("standard atmosphere", "atm", Pressure, BI)
	Bar      = newUnit("bar", "bar", Pressure, BI)
	CentiBar = Centi(Bar, BI)
	MilliBar = Milli(Bar, BI)
	MicroBar = Micro(Bar, BI)
	Barye    = newUnit("barye", "Ba", Pressure, BI)
	InH2O    = newUnit(
		"inch of Water Column", "inH2O", Pressure, BI, Plural("inches of Water Column"),
	)
	InHg = newUnit("inch of Mercury", "inHg", Pressure, BI, Plural("inches of Mercury"))
	MH2O = newUnit(
		"meter of Water Column", "mH2O", Pressure, BI,
		Plural("meters of Water Column"),
	)
	MilliMH2O                    = Milli(MH2O, BI, Plural("millimeters of Water Column"))
	CentiMH2O                    = Centi(MH2O, BI, Plural("centimeters of Water Column"))
	MHg                          = newUnit("meter of Mercury", "mmHg", Pressure, BI, Plural("meters of Mercury"))
	MilliMHg                     = Milli(MHg, BI, Plural("millimeters of Mercury"))
	CentiMHg                     = Centi(MHg, BI, Plural("centimeters of Mercury"))
	NewtonSqm                    = newUnit("newton per square meter", "N/m²", Pressure, BI)
	KiloNewtonSqm                = Kilo(NewtonSqm)
	Psi                          = newUnit("pound-force per square inch", "psi", Pressure, BI)
	Torr                         = newUnit("torr", "Torr", Pressure, BI)
	FootH2O                      = newUnit("foot of Water Column", "FT", Pressure, BI)
	InchesOfWater                = newUnit("inches of Water Column", "in-wg", Pressure, BI)
	PoundForcePerSquareInchGauge = newUnit("pound-force per square inch gauge", "psig", Pressure, BI)
)

func init() {
	// https://www.nist.gov/pml/special-publication-811/nist-guide-si-appendix-b-conversion-factors/nist-guide-si-appendix-b9#PRESSURE
	NewRatioConversion(At, Pascal, 9.80665e+04)
	NewRatioConversion(Atm, Pascal, 1.01325e+05)
	NewRatioConversion(Bar, Pascal, 1.0e+05)
	NewRatioConversion(Barye, Pascal, 0.1)
	// inch of water, conventional (inH2O)
	NewRatioConversion(InH2O, Pascal, 2.490889e+02)
	// inch of mercury, conventional (inHg)
	NewRatioConversion(InHg, Pascal, 3.386389e+03)
	// millimeter of water, conventional (mmH2O)
	NewRatioConversion(MilliMH2O, Pascal, 9.80665e+00)
	// millimeter of mercury, conventional (mmHg)
	NewRatioConversion(MilliMHg, Pascal, 1.333224e+02)
	NewRatioConversion(NewtonSqm, Pascal, 1)
	NewRatioConversion(Psi, Pascal, 6.894757e+03)
	NewRatioConversion(Torr, Pascal, 1.333224e+02)

	Psi.AddAliases("pound-force per square inch")
	Psi.AddSymbols("lbf/in²", "lbf/in^2")

	NewRatioConversion(FootH2O, Pascal, 2989.06692)
	NewRatioConversion(InchesOfWater, Pascal, 249.0889)
	NewRatioConversion(PoundForcePerSquareInchGauge, Pascal, 6894.757)
}
