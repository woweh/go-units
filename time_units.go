package units

var (
	Time = Quantity("time")

	Second      = mustCreateNewUnit("second", "s", Time, SI)
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

	Minute = mustCreateNewUnit("minute", "min", Time)
	Hour   = mustCreateNewUnit("hour", "hr", Time)
	Day    = mustCreateNewUnit("day", "d", Time)
	Month  = mustCreateNewUnit("month", "", Time)
	Year   = mustCreateNewUnit("year", "yr", Time)

	Decade     = mustCreateNewUnit("decade", "", Time)
	Century    = mustCreateNewUnit("century", "", Time)
	Millennium = mustCreateNewUnit("millennium", "", Time)

	// more esoteric time units
	PlanckTime = mustCreateNewUnit("planck time", "𝑡ₚ", Time)
	Fortnight  = mustCreateNewUnit("fortnight", "", Time)
	Score      = mustCreateNewUnit("score", "", Time)
)

func init() {
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
