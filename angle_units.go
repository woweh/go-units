package units

import "math"

var (
	Angle = Quantity("angle")

	Turn = newUnit("turn", "tr", Angle)

	Radian      = newUnit("radian", "rad", Angle, SI, BaseUnit)
	MilliRadian = Milli(Radian)
	MicroRadian = Micro(Radian)

	// Degree (= decimal degree) is a unit of angle equal to 1/360 of a circle.
	Degree = newUnit("degree", "°", Angle)

	Gon      = newUnit("gon", "gon", Angle, BaseUnit, Symbols("grad"))
	DeciGon  = Deci(Gon)
	CentiGon = Centi(Gon)
	MilliGon = Milli(Gon)
	MicroGon = Micro(Gon)
	NanoGon  = Nano(Gon)
	PicoGon  = Pico(Gon)
	FemtoGon = Femto(Gon)
	AttoGon  = Atto(Gon)
	ZeptoGon = Zepto(Gon)
	YoctoGon = Yocto(Gon)
	DecaGon  = Deca(Gon)
	HectoGon = Hecto(Gon)
	KiloGon  = Kilo(Gon)
	MegaGon  = Mega(Gon)
	GigaGon  = Giga(Gon)
	TeraGon  = Tera(Gon)
	PetaGon  = Peta(Gon)
)

func init() {
	NewRatioConversion(Turn, Radian, 2*math.Pi)
	NewRatioConversion(Turn, Degree, 360)
	NewRatioConversion(Turn, Gon, 400)

	Gon.AddAliases("gradian", "gradians", "grads")

	Degree.AddAliases("degrees", "degree of arc", "degrees of arc", "arc degree", "arcdegree")
	Degree.AddSymbols("deg")

	Turn.AddAliases("revolution", "revolutions", "revs", "cycle", "cycles")
	Turn.AddSymbols("pla", "rev", "cyc")
}
