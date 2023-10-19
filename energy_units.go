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
	GigaWattHour = Giga(WattHour)
	MegaWattHour = Mega(WattHour)
	TeraWattHour = Tera(WattHour)
	PetaWattHour = Peta(WattHour)

	// other
	ElectronVolt = newUnit("electronvolt", "eV", Energy)
	Calorie      = newUnit("calorie", "cal", Energy)
)

func init() {
	// Non-metric to Metric conversions
	NewRatioConversion(ElectronVolt, Joule, 1.60218e-19)
	NewRatioConversion(Calorie, Joule, 4.184)
	NewRatioConversion(WattHour, Joule, 3600)
	NewRatioConversion(MegaWattHour, Joule, 3600*1e3)
	NewRatioConversion(KiloWattHour, Watt, 3600*1e6)
	NewRatioConversion(GigaWattHour, Watt, 3600*1e9)
	NewRatioConversion(TeraWattHour, Watt, 3600*1e12)
	NewRatioConversion(PetaWattHour, Watt, 3600*1e14)
}
