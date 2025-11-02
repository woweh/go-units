package units

// Pressure is the UnitQuantity (string/name) for pressure units
const Pressure UnitQuantity = "pressure"

var (
	// _pressure is the UnitOption for pressure units, used in unit definitions
	_pressure = Quantity(Pressure)

	// SI unit metric
	Pascal      = mustCreateNewUnit("pascal", "Pa", _pressure, SI)
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
	At       = mustCreateNewUnit("technical atmosphere", "at", _pressure, BI)
	Atm      = mustCreateNewUnit("standard atmosphere", "atm", _pressure, BI)
	Bar      = mustCreateNewUnit("bar", "bar", _pressure, BI)
	CentiBar = Centi(Bar, BI)
	MilliBar = Milli(Bar, BI)
	MicroBar = Micro(Bar, BI)
	Barye    = mustCreateNewUnit("barye", "Ba", _pressure, BI)
	InH2O    = mustCreateNewUnit(
		"inch of Water Column", "inH2O", _pressure, BI, Plural("inches of Water Column"),
	)
	InHg                         = mustCreateNewUnit("inch of Mercury", "inHg", _pressure, BI, Plural("inches of Mercury"))
	MH2O                         = mustCreateNewUnit("meter of Water Column", "mH2O", _pressure, BI, Plural("meters of Water Column"))
	MilliMH2O                    = Milli(MH2O, BI, Plural("millimeters of Water Column"))
	CentiMH2O                    = Centi(MH2O, BI, Plural("centimeters of Water Column"))
	MHg                          = mustCreateNewUnit("meter of Mercury", "mmHg", _pressure, BI, Plural("meters of Mercury"))
	MilliMHg                     = Milli(MHg, BI, Plural("millimeters of Mercury"))
	CentiMHg                     = Centi(MHg, BI, Plural("centimeters of Mercury"))
	NewtonSqm                    = mustCreateNewUnit("newton per square meter", "N/m²", _pressure, SI)
	KiloNewtonSqm                = Kilo(NewtonSqm)
	Psi                          = mustCreateNewUnit("pound-force per square inch", "psi", _pressure, BI)
	Torr                         = mustCreateNewUnit("torr", "Torr", _pressure, BI)
	FootH2O                      = mustCreateNewUnit("foot of Water Column", "FT", _pressure, BI)
	InchesOfWater                = mustCreateNewUnit("inches of Water Column", "in-wg", _pressure, BI)
	PoundForcePerSquareInchGauge = mustCreateNewUnit("pound-force per square inch gauge", "psig", _pressure, BI)
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
