package units

// Time is a unit quantity for time
const Time UnitQuantity = "time"

var (
	_time = Quantity(Time)

	Second      = mustCreateNewUnit("second", "s", _time, SI)
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

	Minute = mustCreateNewUnit("minute", "min", _time)
	Hour   = mustCreateNewUnit("hour", "hr", _time)
	Day    = mustCreateNewUnit("day", "d", _time)
	Month  = mustCreateNewUnit("month", "", _time)
	Year   = mustCreateNewUnit("year", "yr", _time)

	Decade     = mustCreateNewUnit("decade", "", _time)
	Century    = mustCreateNewUnit("century", "", _time)
	Millennium = mustCreateNewUnit("millennium", "", _time)

	// more esoteric time units
	PlanckTime = mustCreateNewUnit("planck time", "𝑡ₚ", _time)
	Fortnight  = mustCreateNewUnit("fortnight", "", _time)
	Score      = mustCreateNewUnit("score", "", _time)
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
