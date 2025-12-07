package units

var (
	// Time is the unit quantity for time.
	Time = NewUnitQuantity("time")

	Second      = Time.MustCreateUnit("second", "s", SI)
	ExaSecond   = Exa(Second)
	PetaSecond  = Peta(Second)
	TeraSecond  = Tera(Second)
	GigaSecond  = Giga(Second)
	MegaSecond  = Mega(Second)
	KiloSecond  = Kilo(Second)
	HectoSecond = Hecto(Second)
	DecaSecond  = Deca(Second)
	DeciSecond  = Deci(Second)
	CentiSecond = Centi(Second)
	MilliSecond = Milli(Second)
	MicroSecond = Micro(Second)
	NanoSecond  = Nano(Second)
	PicoSecond  = Pico(Second)
	FemtoSecond = Femto(Second)
	AttoSecond  = Atto(Second)

	Minute = Time.MustCreateUnit("minute", "min")
	Hour   = Time.MustCreateUnit("hour", "hr")
	Day    = Time.MustCreateUnit("day", "d")
	Month  = Time.MustCreateUnit("month", "")
	Year   = Time.MustCreateUnit("year", "yr")

	Decade     = Time.MustCreateUnit("decade", "")
	Century    = Time.MustCreateUnit("century", "")
	Millennium = Time.MustCreateUnit("millennium", "")

	// more esoteric time units
	PlanckTime = Time.MustCreateUnit("planck time", "𝑡ₚ")
	Fortnight  = Time.MustCreateUnit("fortnight", "")
	Score      = Time.MustCreateUnit("score", "")
)

func initTimeUnits() {
	NewRatioConversion(Minute, Second, 60.0)
	NewRatioConversion(Hour, Second, 3600.0)
	NewRatioConversion(Day, Hour, 24.0)
	NewRatioConversion(Month, Day, 30.0)
	NewRatioConversion(Year, Day, 365.25)

	NewRatioConversion(Decade, Year, 10.0)
	NewRatioConversion(Century, Year, 100.0)
	NewRatioConversion(Millennium, Year, 1000.0)

	NewRatioConversion(PlanckTime, Second, 5.39e-44)
	NewRatioConversion(Fortnight, Day, 14)
	NewRatioConversion(Score, Year, 20.0)
}
