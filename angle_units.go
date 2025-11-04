package units

import (
	"fmt"
	"math"
)

var (
	Angle = Quantity("angle")

	Turn = mustCreateNewUnit("turn", "tr", Angle)

	Radian      = mustCreateNewUnit("radian", "rad", Angle, SI)
	MilliRadian = Milli(Radian)
	MicroRadian = Micro(Radian)

	// Degree (= decimal degree) is a unit of angle equal to 1/360 of a circle.
	Degree = mustCreateNewUnit("degree", "°", Angle)

	Gon      = mustCreateNewUnit("gon", "gon", Angle, SI, Symbols("grad"))
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

	// DMS (Degrees, Minutes, Seconds) as a unit
	DMS = mustCreateNewUnit("degree minute second", "DMS", Angle)
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

	// Register DMS <-> Degree conversions
	NewConversionFromFn(DMS, Degree, func(x float64) float64 {
		// x is encoded as D.MMSS, e.g. 12.3045 means 12°30'45''
		D := int(x)
		M := int((x - float64(D)) * 100)
		S := ((x-float64(D))*100 - float64(M)) * 100
		return float64(D) + float64(M)/60 + S/3600
	}, "DMS (D.MMSS) to degrees")

	NewConversionFromFn(Degree, DMS, func(x float64) float64 {
		D := int(x)
		M := int((x - float64(D)) * 60)
		S := ((x-float64(D))*60 - float64(M)) * 60
		return float64(D) + float64(M)/100 + S/10000
	}, "degrees to DMS (D.MMSS)")

	DMS.AddFormatFn(func(x float64) string {
		D := int(x)
		M := int((x - float64(D)) * 100)
		S := ((x-float64(D))*100 - float64(M)) * 100
		return fmt.Sprintf("%d°%02d'%02d''", D, M, int(S))
	})
}
