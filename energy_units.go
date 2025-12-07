package units

var (
	// Energy is the unit quantity for energy.
	Energy = NewUnitQuantity("energy")

	// metric
	Joule      = Energy.MustCreateUnit("joule", "J", SI)
	KiloJoule  = Kilo(Joule)
	MegaJoule  = Mega(Joule)
	GigaJoule  = Giga(Joule)
	TeraJoule  = Tera(Joule)
	PetaJoule  = Peta(Joule)
	ExaJoule   = Exa(Joule)
	ZettaJoule = Zetta(Joule)
	YottaJoule = Yotta(Joule)
	MilliJoule = Milli(Joule)
	MicroJoule = Micro(Joule)
	NanoJoule  = Nano(Joule)
	PicoJoule  = Pico(Joule)
	FemtoJoule = Femto(Joule)
	AttoJoule  = Atto(Joule)

	WattHour = Energy.MustCreateUnit(
		"watt-hour", "Wh", SI,
		Aliases("volt ampere hour", "volt ampere reactive hour", "volt ampere hour (reactive)"),
		Symbols("VAh", "varh", "V⋅A⋅hr", "V.A.h", "V.A{reactive}.h", "V⋅A{reactive}⋅hr"),
	)
	KiloWattHour = Kilo(WattHour)
	MegaWattHour = Mega(WattHour)
	GigaWattHour = Giga(WattHour)
	TeraWattHour = Tera(WattHour)
	PetaWattHour = Peta(WattHour)

	// other
	ElectronVolt     = Energy.MustCreateUnit("electronvolt", "eV")
	KiloElectronVolt = Energy.MustCreateUnit("kiloelectronvolt", "keV")
	MegaElectronVolt = Energy.MustCreateUnit("megaelectronvolt", "MeV")
	GigaElectronVolt = Energy.MustCreateUnit("gigaelectronvolt", "GeV")

	Calorie     = Energy.MustCreateUnit("calorie", "cal")
	KiloCalorie = Energy.MustCreateUnit("kilocalorie", "kcal")

	BritishThermalUnit = Energy.MustCreateUnit("British thermal unit", "Btu")
	Therm              = Energy.MustCreateUnit("therm", "therm")
)

func initEnergyUnits() {
	// https://www.nist.gov/pml/special-publication-811/nist-guide-si-appendix-b-conversion-factors/nist-guide-si-appendix-b9#ENERGY

	// Empirical/standard energy definitions (remain hardcoded)
	// electronvolt (eV)	joule (J)	1.602177e-19
	NewRatioConversion(ElectronVolt, Joule, 1.602177e-19)
	ElectronVolt.AddAliases("electron volt")
	NewRatioConversion(KiloElectronVolt, ElectronVolt, 1000)
	NewRatioConversion(MegaElectronVolt, ElectronVolt, 1e6)
	NewRatioConversion(GigaElectronVolt, ElectronVolt, 1e9)

	NewRatioConversion(Calorie, Joule, 4.184)
	Calorie.AddAliases("Thermochemical Calorie", "calorie (th)")
	NewRatioConversion(KiloCalorie, Calorie, 1000)

	NewRatioConversion(BritishThermalUnit, Joule, 1055.05585262)
	NewRatioConversion(Therm, Joule, 105505585.262)

	// Calculated energy conversions
	// 1 watt second = 1 joule (by definition: W = J/s, so W⋅s = J)
	Joule.AddAliases("watt second")
	Joule.AddSymbols("W⋅s")

	// 1 watt hour = 3600 joules (1 W × 3600 s = 3600 J)
	// Note: While we could calculate this from Watt and Hour, those are in different
	// unit quantities, so we keep the standard conversion factor
	NewRatioConversion(WattHour, Joule, 3600)
}

// energyFactor calculates the conversion factor for energy units.
// Energy = force × distance
// Base unit: Joule = N⋅m = kg⋅m²/s²
//
// Example: foot-pound force (ft⋅lbf)
// - 1 ft⋅lbf = 1 lbf × 1 ft
// - forceRatio: how many N in 1 lbf
// - lengthRatio: how many m in 1 ft
// - energyFactor = forceRatio × lengthRatio
func energyFactor(force, length Unit) float64 {
	// How many Newtons in 1 unit of the target force
	forceRatio := force.to(Newton)
	// How many meters in 1 unit of target length
	lengthRatio := length.to(Meter)
	// Energy factor: force × distance
	return forceRatio * lengthRatio
}
