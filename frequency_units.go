package units

import "math"

var (
	Frequency = Quantity("frequency")

	Hertz      = newUnit("hertz", "Hz", Frequency, SI, BaseUnit)
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

	RadianPerSecond = newUnit("radian per second", "rad/s", Frequency, SI)
	RadianPerMinute = newUnit("radian per minute", "rad/min", Frequency, SI)
	RadianPerHour   = newUnit("radian per hour", "rad/h", Frequency, SI)
	RadianPerDay    = newUnit("radian per day", "rad/d", Frequency, SI)

	DegreePerSecond = newUnit("degree per second", "°/s", Frequency, BI)
	DegreePerMinute = newUnit("degree per minute", "°/min", Frequency, BI)
	DegreePerHour   = newUnit("degree per hour", "°/h", Frequency, BI)
	DegreePerDay    = newUnit("degree per day", "°/d", Frequency, BI)

	RevolutionPerSecond = newUnit("revolution per second", "rev/s", Frequency)
	RevolutionPerMinute = newUnit("revolution per minute", "rev/min", Frequency)
	RevolutionPerHour   = newUnit("revolution per hour", "rev/h", Frequency)
	RevolutionPerDay    = newUnit("revolution per day", "rev/d", Frequency)
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
