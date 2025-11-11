package units

import (
	"fmt"
	"math"
)

const Angle UnitQuantity = "angle"

var (
	_angle = Quantity(Angle)

	Turn = mustCreateNewUnit("turn", "tr", _angle)

	Radian      = mustCreateNewUnit("radian", "rad", _angle, SI)
	MilliRadian = Milli(Radian)
	MicroRadian = Micro(Radian)

	// Degree (= decimal degree) is a unit of angle equal to 1/360 of a circle.
	Degree = mustCreateNewUnit("degree", "°", _angle, SI)

	Gon      = mustCreateNewUnit("gon", "gon", _angle, SI, Symbols("grad"))
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
	DMS = mustCreateNewUnit("degree minute second", "DMS", _angle)
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

	// DMS <-> Degree conversions
	NewConversionFromFn(DMS, Degree, DMS2Degree, "DMS to decimal degrees")

	// Degree <-> DMS conversions
	NewConversionFromFn(Degree, DMS, Degree2DMS, "decimal degrees to DMS")

	// Define DMS formatting
	DMS.AddFormatFn(FormatDMS)
}

// DMS2Degree converts DMS (D.MMSS) to decimal degrees.
func DMS2Degree(x float64) float64 {
	D := int(x)
	M := int((x - float64(D)) * 100)
	S := ((x-float64(D))*100 - float64(M)) * 100
	return float64(D) + float64(M)/60 + S/3600
}

// Degree2DMS converts decimal degrees to DMS (D.MMSS).
func Degree2DMS(x float64) float64 {
	D := int(x)
	M := int((x - float64(D)) * 60)
	S := ((x-float64(D))*60 - float64(M)) * 60
	return float64(D) + float64(M)/100 + S/10000
}

// FormatDMS formats a DMS value (D.MMSS) as a string (D°MM'SS”).
func FormatDMS(x float64) string {
	D := int(x)
	M := int((x - float64(D)) * 100)
	S := ((x-float64(D))*100 - float64(M)) * 100
	return fmt.Sprintf("%d°%02d'%02d''", D, M, int(S))
}
