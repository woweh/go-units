package units

var (
	Energy = Quantity("energy")

	// metric
	Joule      = newUnit("joule", "J", Energy, SI, BaseUnit)
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

	WattHour = newUnit(
		"watt-hour", "Wh", Energy, SI,
		Aliases("volt ampere hour", "volt ampere reactive hour", "volt ampere hour (reactive)"),
		Symbols("VAh", "varh", "V⋅A⋅hr", "V.A.h", "V.A{reactive}.h", "V⋅A{reactive}⋅hr"),
		BaseUnit,
	)
	KiloWattHour = Kilo(WattHour)
	MegaWattHour = Mega(WattHour)
	GigaWattHour = Giga(WattHour)
	TeraWattHour = Tera(WattHour)
	PetaWattHour = Peta(WattHour)

	// other
	ElectronVolt     = newUnit("electronvolt", "eV", Energy, BaseUnit)
	KiloElectronVolt = Kilo(ElectronVolt)
	MegaElectronVolt = Mega(ElectronVolt)
	GigaElectronVolt = Giga(ElectronVolt)

	Calorie     = newUnit("calorie", "cal", Energy, BaseUnit)
	KiloCalorie = Kilo(Calorie)
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
}
