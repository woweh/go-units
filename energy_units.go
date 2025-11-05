package units

const Energy UnitQuantity = "energy"

var (
	_energy = Quantity(Energy)

	// metric
	Joule      = mustCreateNewUnit("joule", "J", _energy, SI)
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

	WattHour = mustCreateNewUnit(
		"watt-hour", "Wh", _energy, SI,
		Aliases("volt ampere hour", "volt ampere reactive hour", "volt ampere hour (reactive)"),
		Symbols("VAh", "varh", "V⋅A⋅hr", "V.A.h", "V.A{reactive}.h", "V⋅A{reactive}⋅hr"),
	)
	KiloWattHour = Kilo(WattHour)
	MegaWattHour = Mega(WattHour)
	GigaWattHour = Giga(WattHour)
	TeraWattHour = Tera(WattHour)
	PetaWattHour = Peta(WattHour)

	// other
	ElectronVolt     = mustCreateNewUnit("electronvolt", "eV", _energy, SI)
	KiloElectronVolt = Kilo(ElectronVolt)
	MegaElectronVolt = Mega(ElectronVolt)
	GigaElectronVolt = Giga(ElectronVolt)

	Calorie     = mustCreateNewUnit("calorie", "cal", _energy, SI)
	KiloCalorie = Kilo(Calorie)

	BritishThermalUnit = mustCreateNewUnit("British thermal unit", "Btu", _energy, BI)
	Therm              = mustCreateNewUnit("therm", "therm", _energy, BI)
)

func init() {
	// https://www.nist.gov/pml/special-publication-811/nist-guide-si-appendix-b-conversion-factors/nist-guide-si-appendix-b9#ENERGY
	// electronvolt (eV)	joule (J)	1.602177e-19
	NewRatioConversion(ElectronVolt, Joule, 1.602177e-19)
	ElectronVolt.AddAliases("electron volt")

	NewRatioConversion(Calorie, Joule, 4.184)
	Calorie.AddAliases("Thermochemical Calorie", "calorie (th)")

	// 1 watt second = 1 joule
	Joule.AddAliases("watt second")
	Joule.AddSymbols("W⋅s")

	// 1 watt hour = 1 Volt Ampere Hour = 1 Volt Ampere Reactive Hour = 3600 joules
	NewRatioConversion(WattHour, Joule, 3600)

	NewRatioConversion(BritishThermalUnit, Joule, 1055.05585262)
	NewRatioConversion(Therm, Joule, 105505585.262)
}
