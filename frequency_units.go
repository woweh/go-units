package units

import "math"

var (
	Frequency = Quantity("frequency")

	Hertz      = mustCreateNewUnit("hertz", "Hz", Frequency, SI, Plural("hertz"))
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

	RadianPerSecond = mustCreateNewUnit("radian per second", "rad/s", Frequency, SI)
	RadianPerMinute = mustCreateNewUnit("radian per minute", "rad/min", Frequency, SI)
	RadianPerHour   = mustCreateNewUnit("radian per hour", "rad/h", Frequency, SI)
	RadianPerDay    = mustCreateNewUnit("radian per day", "rad/d", Frequency, SI)

	DegreePerSecond = mustCreateNewUnit("degree per second", "°/s", Frequency, BI)
	DegreePerMinute = mustCreateNewUnit("degree per minute", "°/min", Frequency, BI)
	DegreePerHour   = mustCreateNewUnit("degree per hour", "°/h", Frequency, BI)
	DegreePerDay    = mustCreateNewUnit("degree per day", "°/d", Frequency, BI)

	RevolutionPerSecond = mustCreateNewUnit("revolution per second", "rev/s", Frequency)
	RevolutionPerMinute = mustCreateNewUnit("revolution per minute", "rev/min", Frequency)
	RevolutionPerHour   = mustCreateNewUnit("revolution per hour", "rev/h", Frequency)
	RevolutionPerDay    = mustCreateNewUnit("revolution per day", "rev/d", Frequency)
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
