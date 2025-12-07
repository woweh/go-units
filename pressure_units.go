package units

var (
	// Pressure is the unit quantity for pressure.
	Pressure = NewUnitQuantity("pressure")

	// SI unit metric
	Pascal      = Pressure.MustCreateUnit("pascal", "Pa", SI)
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
	At       = Pressure.MustCreateUnit("technical atmosphere", "at", BI)
	Atm      = Pressure.MustCreateUnit("standard atmosphere", "atm", BI)
	Bar      = Pressure.MustCreateUnit("bar", "bar")
	CentiBar = Centi(Bar, BI)
	MilliBar = Milli(Bar, BI)
	MicroBar = Micro(Bar, BI)
	Barye    = Pressure.MustCreateUnit("barye", "Ba")
	InH2O    = Pressure.MustCreateUnit(
		"inch of Water Column", "inH2O", BI, Plural("inches of Water Column"),
	)
	InHg                         = Pressure.MustCreateUnit("inch of Mercury", "inHg", BI, Plural("inches of Mercury"))
	MH2O                         = Pressure.MustCreateUnit("meter of Water Column", "mH2O", BI, Plural("meters of Water Column"))
	MilliMH2O                    = Milli(MH2O, BI, Plural("millimeters of Water Column"))
	CentiMH2O                    = Centi(MH2O, BI, Plural("centimeters of Water Column"))
	MHg                          = Pressure.MustCreateUnit("meter of Mercury", "mmHg", BI, Plural("meters of Mercury"))
	MilliMHg                     = Milli(MHg, BI, Plural("millimeters of Mercury"))
	CentiMHg                     = Centi(MHg, BI, Plural("centimeters of Mercury"))
	NewtonSqm                    = Pressure.MustCreateUnit("newton per square meter", "N/m²", SI)
	KiloNewtonSqm                = Kilo(NewtonSqm)
	Psi                          = Pressure.MustCreateUnit("pound-force per square inch", "psi", BI)
	Torr                         = Pressure.MustCreateUnit("torr", "Torr")
	FootH2O                      = Pressure.MustCreateUnit("foot of Water Column", "FT", BI)
	InchesOfWater                = Pressure.MustCreateUnit("inch of water gauge", "in-wg", BI)
	PoundForcePerSquareInchGauge = Pressure.MustCreateUnit("pound-force per square inch gauge", "psig", BI)
)

func initPressureUnits() {
	// https://www.nist.gov/pml/special-publication-811/nist-guide-si-appendix-b-conversion-factors/nist-guide-si-appendix-b9#PRESSURE

	// Empirical/standard pressure definitions (remain hardcoded)
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
	NewRatioConversion(Torr, Pascal, 1.333224e+02)
	NewRatioConversion(FootH2O, Pascal, 2989.06692)
	NewRatioConversion(InchesOfWater, Pascal, 249.0889)

	// Calculated pressure conversions (force/area)
	NewRatioConversion(NewtonSqm, Pascal, 1) // By definition: 1 Pa = 1 N/m²
	// Psi = lbf/in²
	NewRatioConversion(Psi, Pascal, pressureFactor(PoundForce, Inch))
	Psi.AddAliases("pound-force per square inch")
	Psi.AddSymbols("lbf/in²", "lbf/in^2")

	// Gauge pressure (same as absolute for conversion purposes)
	NewRatioConversion(PoundForcePerSquareInchGauge, Pascal, pressureFactor(PoundForce, Inch))
}

// pressureFactor calculates the conversion factor for pressure units.
// Pressure = force / area
// Base unit: Pascal = N/m²
//
// To calculate how many Pascals in 1 unit of target pressure (e.g., psi):
// - forceRatio: how many Newtons in 1 unit of target force
// - areaRatio: how many m² in 1 unit of target area
// - pressure factor = forceRatio / areaRatio
//
// Example: How many Pa in 1 psi?
// - 1 psi = 1 lbf/in²
// - 1 lbf = 4.448222 N
// - 1 in² = 0.00064516 m²
// - 1 psi = 4.448222 N / 0.00064516 m² = 6894.757 Pa
func pressureFactor(force, length Unit) float64 {
	// How many Newtons in 1 unit of the target force
	forceRatio := force.to(Newton)
	// How many m² in 1 unit of target area (length²)
	areaRatio := areaFactor(length)
	// Pressure factor: force per area
	return forceRatio / areaRatio
}
