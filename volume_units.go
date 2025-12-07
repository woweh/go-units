package units

var (
	// Volume is the unit quantity for volume.
	Volume = NewUnitQuantity("volume")

	// SI/metric units
	Liter      = Volume.MustCreateUnit("liter", "l", SI, Aliases("litre"))
	ExaLiter   = Exa(Liter)
	PetaLiter  = Peta(Liter)
	TeraLiter  = Tera(Liter)
	GigaLiter  = Giga(Liter)
	MegaLiter  = Mega(Liter)
	KiloLiter  = Kilo(Liter)
	HectoLiter = Hecto(Liter)
	DecaLiter  = Deca(Liter)
	DeciLiter  = Deci(Liter)
	CentiLiter = Centi(Liter)
	MilliLiter = Milli(Liter)
	MicroLiter = Micro(Liter)
	NanoLiter  = Nano(Liter)
	PicoLiter  = Pico(Liter)
	FemtoLiter = Femto(Liter)
	AttoLiter  = Atto(Liter)

	CubicMeter      = Volume.MustCreateUnit("cubic meter", "m³", SI)
	CubicKiloMeter  = Volume.MustCreateUnit("cubic kilometer", "km³", SI)
	CubicHectoMeter = Volume.MustCreateUnit("cubic hectometer", "hm³", SI)
	CubicDecaMeter  = Volume.MustCreateUnit("cubic decameter", "dam³", SI)
	CubicDeciMeter  = Volume.MustCreateUnit("cubic decimeter", "dm³", SI)
	CubicCentiMeter = Volume.MustCreateUnit("cubic centimeter", "cm³", SI)
	CubicMilliMeter = Volume.MustCreateUnit("cubic millimeter", "mm³", SI)

	// Imperial units
	Quart      = Volume.MustCreateUnit("quart", "qt", BI)
	Pint       = Volume.MustCreateUnit("pint", "pt", BI)
	Gallon     = Volume.MustCreateUnit("gallon", "gal", BI)
	FluidOunce = Volume.MustCreateUnit("fluid ounce", "fl oz", BI)
	CubicFoot  = Volume.MustCreateUnit("cubic foot", "ft³", BI)
	CubicYard  = Volume.MustCreateUnit("cubic yard", "yd³", BI)
	CubicInch  = Volume.MustCreateUnit("cubic inch", "in³", BI)
	CubicMile  = Volume.MustCreateUnit("cubic mile", "mi³", BI)
	AcreFoot   = Volume.MustCreateUnit("acre foot", "ac ft", BI)

	// US customary units
	FluidQuart          = Volume.MustCreateUnit("fluid quart", "fl qt", US)
	FluidPint           = Volume.MustCreateUnit("fluid pint", "fl pt", US)
	FluidGallon         = Volume.MustCreateUnit("fluid gallon", "", US)
	CustomaryFluidOunce = Volume.MustCreateUnit("customary fluid ounce", "", US)
)

func initVolumeUnits() {
	NewRatioConversion(CubicKiloMeter, TeraLiter, 1)
	NewRatioConversion(CubicHectoMeter, GigaLiter, 1)
	NewRatioConversion(CubicDecaMeter, MegaLiter, 1)
	NewRatioConversion(CubicMeter, KiloLiter, 1)
	NewRatioConversion(CubicDeciMeter, Liter, 1)
	NewRatioConversion(CubicCentiMeter, MilliLiter, 1)
	NewRatioConversion(CubicMilliMeter, MicroLiter, 1)

	NewRatioConversion(Quart, Liter, 1.1365225)
	NewRatioConversion(Pint, Liter, 0.56826125)
	NewRatioConversion(Gallon, Liter, 4.54609)
	NewRatioConversion(FluidOunce, MilliLiter, 28.4130625)

	NewRatioConversion(CubicInch, Liter, volumeFactor(Inch)*1000) // *1000 to convert m³ to L
	NewRatioConversion(CubicFoot, Liter, volumeFactor(Foot)*1000)
	NewRatioConversion(CubicYard, Liter, volumeFactor(Yard)*1000)
	NewRatioConversion(AcreFoot, Liter, 1233481.84)
	NewRatioConversion(CubicMile, Liter, 4168181825440.64)

	NewRatioConversion(FluidQuart, Liter, 0.946352946)
	NewRatioConversion(FluidPint, Liter, 0.473176473)
	NewRatioConversion(FluidGallon, Liter, 3.785411784)
	NewRatioConversion(CustomaryFluidOunce, MilliLiter, 29.5735295625)

	CubicMeter.AddAliases("cubic metre")
	CubicMeter.AddSymbols("m3", "m^3", "m**3", "cum", "cbm", "CBM", "M3")

	CubicDeciMeter.AddAliases("decimetre")
	CubicDeciMeter.AddSymbols("dm3", "dm^3", "dm**3")

	CubicKiloMeter.AddAliases("cubic kilometre")
	CubicKiloMeter.AddSymbols("km3", "km^3", "km**3")

	CubicHectoMeter.AddAliases("cubic hectometre")
	CubicHectoMeter.AddSymbols("hm3", "hm^3", "hm**3", "mcm")

	CubicDecaMeter.AddAliases("cubic decametre")
	CubicDecaMeter.AddSymbols("dam3", "dam^3", "dam**3")

	CubicCentiMeter.AddAliases("cubic centimetre")
	CubicCentiMeter.AddSymbols("cc", "ccm", "cm3", "cm^3", "cm**3")

	CubicMilliMeter.AddAliases("cubic millimetre")
	CubicMilliMeter.AddSymbols("mm3", "mm^3", "mm**3")

	FluidOunce.AddSymbols("floz", "fl. oz", "fl.oz", "oz. fl", "oz fl", "ozfl", "ozfl.")
	CubicFoot.AddSymbols("cuft", "cu ft", "cft", "ft^3", "ft**3", "ft3")
	CubicYard.AddSymbols("cuyd", "cu yd", "cyd", "yd^3", "yd**3", "yd3")
	CubicInch.AddSymbols("cuin", "cu in", "cin", "in^3", "in**3", "in3")
	CubicMile.AddSymbols("cumi", "cubmi", "mi^3", "mi**3", "mi3")
	AcreFoot.AddSymbols("acft", "acreft", "acre-ft", "acre ft", "ac-ft")
}

// volumeFactor computes the conversion factor from length units to cubic meters.
// volumeFactor = length³
func volumeFactor(length Unit) float64 {
	factor := length.to(Meter)
	return factor * factor * factor
}
