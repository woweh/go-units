package units

import "math"

var (
	// Frequency is the unit quantity for frequency.
	Frequency = NewUnitQuantity("frequency")

	Hertz      = Frequency.MustCreateUnit("hertz", "Hz", SI, Plural(PluralNone))
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

	RadianPerSecond = Frequency.MustCreateUnit("radian per second", "rad/s", SI)
	RadianPerMinute = Frequency.MustCreateUnit("radian per minute", "rad/min", SI)
	RadianPerHour   = Frequency.MustCreateUnit("radian per hour", "rad/h", SI)
	RadianPerDay    = Frequency.MustCreateUnit("radian per day", "rad/d", SI)

	DegreePerSecond = Frequency.MustCreateUnit("degree per second", "°/s", SI)
	DegreePerMinute = Frequency.MustCreateUnit("degree per minute", "°/min", SI)
	DegreePerHour   = Frequency.MustCreateUnit("degree per hour", "°/h", SI)
	DegreePerDay    = Frequency.MustCreateUnit("degree per day", "°/d", SI)

	RevolutionPerSecond = Frequency.MustCreateUnit("revolution per second", "rev/s", SI)
	RevolutionPerMinute = Frequency.MustCreateUnit("revolution per minute", "rev/min", SI)
	RevolutionPerHour   = Frequency.MustCreateUnit("revolution per hour", "rev/h", SI)
	RevolutionPerDay    = Frequency.MustCreateUnit("revolution per day", "rev/d", SI)
)

func initFrequencyUnits() {
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
