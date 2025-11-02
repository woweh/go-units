package units

var (
	Volume = Quantity("volume")

	// metric

	Liter      = mustCreateNewUnit("liter", "l", Volume, SI, Aliases("litre"))
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

	CubicMeter      = mustCreateNewUnit("cubic meter", "m³", Volume, SI)
	CubicKiloMeter  = mustCreateNewUnit("cubic kilometer", "km³", Volume, SI)
	CubicHectoMeter = mustCreateNewUnit("cubic hectometer", "hm³", Volume, SI)
	CubicDecaMeter  = mustCreateNewUnit("cubic decameter", "dam³", Volume, SI)
	CubicDeciMeter  = mustCreateNewUnit("cubic decimeter", "dm³", Volume, SI)
	CubicCentiMeter = mustCreateNewUnit("cubic centimeter", "cm³", Volume, SI)
	CubicMilliMeter = mustCreateNewUnit("cubic millimeter", "mm³", Volume, SI)

	// imperial
	Quart      = mustCreateNewUnit("quart", "qt", Volume, BI)
	Pint       = mustCreateNewUnit("pint", "pt", Volume, BI)
	Gallon     = mustCreateNewUnit("gallon", "gal", Volume, BI)
	FluidOunce = mustCreateNewUnit("fluid ounce", "fl oz", Volume, BI)
	CubicFoot  = mustCreateNewUnit("cubic foot", "ft³", Volume, BI)
	CubicYard  = mustCreateNewUnit("cubic yard", "yd³", Volume, BI)
	CubicInch  = mustCreateNewUnit("cubic inch", "in³", Volume, BI)
	CubicMile  = mustCreateNewUnit("cubic mile", "mi³", Volume, BI)
	AcreFoot   = mustCreateNewUnit("acre foot", "ac ft", Volume, BI)

	// US
	FluidQuart          = mustCreateNewUnit("fluid quart", "fl qt", Volume, US)
	FluidPint           = mustCreateNewUnit("fluid pint", "fl pt", Volume, US)
	FluidGallon         = mustCreateNewUnit("fluid gallon", "", Volume, US)
	CustomaryFluidOunce = mustCreateNewUnit("customary fluid ounce", "", Volume, US)
)

func init() {
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

	NewRatioConversion(CubicInch, Liter, 0.016387064)
	NewRatioConversion(CubicFoot, Liter, 28.316846592)
	NewRatioConversion(CubicYard, Liter, 764.554857984)
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
