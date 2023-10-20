package units

var (
	Energy = UnitOptionQuantity("energy")

	// metric
	Joule      = newUnit("joule", "J", Energy, SI)
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

	WattHour     = newUnit("watt-hour", "Wh", Energy)
	KiloWattHour = Kilo(WattHour)
	MegaWattHour = Mega(WattHour)
	GigaWattHour = Giga(WattHour)
	TeraWattHour = Tera(WattHour)
	PetaWattHour = Peta(WattHour)

	// other
	ElectronVolt     = newUnit("electronvolt", "eV", Energy)
	KiloElectronVolt = Kilo(ElectronVolt)
	MegaElectronVolt = Mega(ElectronVolt)
	GigaElectronVolt = Giga(ElectronVolt)

	Calorie     = newUnit("calorie", "cal", Energy)
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
	WattHour.AddAliases("Volt Ampere Hour", "Volt Ampere Reactive Hour")
	WattHour.AddSymbols("VAh", "varh", "V⋅A⋅hr", "V.A.h", "V.A{reactive}.h", "V⋅A{reactive}⋅hr")

	KiloWattHour.AddAliases("Kilovolt Ampere Hour", "Kilovolt Ampere Reactive Hour")
	KiloWattHour.AddSymbols("kVAh", "kvarh", "kV⋅A⋅hr", "kV.A.h", "kV.A{reactive}.h", "kV⋅A{reactive}⋅hr")

	MegaWattHour.AddAliases("Megavolt Ampere Hour", "Megavolt Ampere Reactive Hour")
	MegaWattHour.AddSymbols("MVAh", "Mvarh", "MV⋅A⋅hr", "MV.A.h", "MV.A{reactive}.h", "MV⋅A{reactive}⋅hr")

	GigaWattHour.AddAliases("Gigavolt Ampere Hour", "Gigavolt Ampere Reactive Hour")
	GigaWattHour.AddSymbols("GVAh", "Gvarh", "GV⋅A⋅hr", "GV.A.h", "GV.A{reactive}.h", "GV⋅A{reactive}⋅hr")

	TeraWattHour.AddAliases("Teravolt Ampere Hour", "Teravolt Ampere Reactive Hour")
	TeraWattHour.AddSymbols("TVAh", "Tvarh", "TV⋅A⋅hr", "TV.A.h", "TV.A{reactive}.h", "TV⋅A{reactive}⋅hr")

	PetaWattHour.AddAliases("Petavolt Ampere Hour", "Petavolt Ampere Reactive Hour")
	PetaWattHour.AddSymbols("PVAh", "Pvarh", "PV⋅A⋅hr", "PV.A.h", "PV.A{reactive}.h", "PV⋅A{reactive}⋅hr")
}
