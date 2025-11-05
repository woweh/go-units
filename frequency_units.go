package units

import "math"

const Frequency UnitQuantity = "frequency"

var (
	_frequency = Quantity(Frequency)

	Hertz      = mustCreateNewUnit("hertz", "Hz", _frequency, SI, Plural(PluralNone))
	DecaHertz  = Deca(Hertz)
	HectoHertz = Hecto(Hertz)
	KiloHertz  = Kilo(Hertz)
	MegaHertz  = Mega(Hertz)
	GigaHertz  = Giga(Hertz)
	TeraHertz  = Tera(Hertz)
	PetaHertz  = Peta(Hertz)
	ExaHertz   = Exa(Hertz)
	ZettaHertz = Zetta(Hertz)
	YottaHertz = Yotta(Hertz)
	DeciHertz  = Deci(Hertz)
	CentiHertz = Centi(Hertz)
	MilliHertz = Milli(Hertz)
	MicroHertz = Micro(Hertz)
	NanoHertz  = Nano(Hertz)
	PicoHertz  = Pico(Hertz)
	FemtoHertz = Femto(Hertz)
	AttoHertz  = Atto(Hertz)
	ZeptoHertz = Zepto(Hertz)
	YoctoHertz = Yocto(Hertz)

	RadianPerSecond = mustCreateNewUnit("radian per second", "rad/s", _frequency, SI)
	RadianPerMinute = mustCreateNewUnit("radian per minute", "rad/min", _frequency, SI)
	RadianPerHour   = mustCreateNewUnit("radian per hour", "rad/h", _frequency, SI)
	RadianPerDay    = mustCreateNewUnit("radian per day", "rad/d", _frequency, SI)

	DegreePerSecond = mustCreateNewUnit("degree per second", "°/s", _frequency, SI)
	DegreePerMinute = mustCreateNewUnit("degree per minute", "°/min", _frequency, SI)
	DegreePerHour   = mustCreateNewUnit("degree per hour", "°/h", _frequency, SI)
	DegreePerDay    = mustCreateNewUnit("degree per day", "°/d", _frequency, SI)

	RevolutionPerSecond = mustCreateNewUnit("revolution per second", "rev/s", _frequency, SI)
	RevolutionPerMinute = mustCreateNewUnit("revolution per minute", "rev/min", _frequency, SI)
	RevolutionPerHour   = mustCreateNewUnit("revolution per hour", "rev/h", _frequency, SI)
	RevolutionPerDay    = mustCreateNewUnit("revolution per day", "rev/d", _frequency, SI)
)

func init() {
	Hertz.AddAliases("cycles per second", "cycles/second", "1-per-second", "Frames-per-second")
	Hertz.AddSymbols("cps", "1/s", "fps")

	TeraHertz.AddSymbols("fresnel")

	// 1 Hertz = 1 cycle/revolution per second
	NewRatioConversion(Hertz, RevolutionPerSecond, 1)
	NewRatioConversion(Hertz, RevolutionPerMinute, 60)
	NewRatioConversion(Hertz, RevolutionPerHour, 3600)
	NewRatioConversion(Hertz, RevolutionPerDay, 86400)

	// 1 Hertz = 1 cycle (= 360 degrees) per second
	NewRatioConversion(Hertz, DegreePerSecond, 360)
	NewRatioConversion(Hertz, DegreePerMinute, 360*60)
	NewRatioConversion(Hertz, DegreePerHour, 360*3600)
	NewRatioConversion(Hertz, DegreePerDay, 360*86400)

	// 1 Hertz = 1 cycle (= 2 * π radians) per second
	tau := 2 * math.Pi
	NewRatioConversion(Hertz, RadianPerSecond, tau)
	NewRatioConversion(Hertz, RadianPerMinute, tau*60)
	NewRatioConversion(Hertz, RadianPerHour, tau*3600)
	NewRatioConversion(Hertz, RadianPerDay, tau*86400)
}
