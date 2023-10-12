package units

var (
	Volume = UnitOptionQuantity("volume")

	// metric

	Liter      = newUnit("liter", "l", Volume, SI)
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

	// imperial
	Quart      = newUnit("quart", "qt", Volume, BI)
	Pint       = newUnit("pint", "pt", Volume, BI)
	Gallon     = newUnit("gallon", "gal", Volume, BI)
	FluidOunce = newUnit("fluid ounce", "fl oz", Volume, BI)
	CubicFoot  = newUnit("cubic foot", "ft³", Volume, BI)
	CubicYard  = newUnit("cubic yard", "yd³", Volume, BI)
	CubicInch  = newUnit("cubic inch", "in³", Volume, BI)
	CubicMile  = newUnit("cubic mile", "mi³", Volume, BI)
	AcreFoot   = newUnit("acre foot", "ac ft", Volume, BI)

	// US
	FluidQuart          = newUnit("fluid quart", "fl qt", Volume, US)
	FluidPint           = newUnit("fluid pint", "fl pt", Volume, US)
	FluidGallon         = newUnit("fluid gallon", "", Volume, US)
	CustomaryFluidOunce = newUnit("customary fluid ounce", "", Volume, US)
)

func init() {
	Liter.AddAliases("litre", "cubic decimeter", "cubic decimetre")
	Liter.AddSymbols("dm³", "dm3", "dm^3", "dm**3")

	TeraLiter.AddAliases("cubic kilometer", "cubic kilometre")
	TeraLiter.AddSymbols("km³", "km3", "km^3", "km**3")

	GigaLiter.AddAliases("cubic hectometer", "cubic hectometre")
	GigaLiter.AddSymbols("hm³", "hm3", "hm^3", "hm**3", "mcm")

	MegaLiter.AddAliases("cubic decameter", "cubic decametre")
	MegaLiter.AddSymbols("dam³", "dam3", "dam^3", "dam**3")

	KiloLiter.AddAliases("cubic meter", "cubic metre")
	KiloLiter.AddSymbols("cum", "cbm", "m³", "m3", "m^3", "m**3")

	MilliLiter.AddAliases("cubic centimeter", "cubic centimetre")
	MilliLiter.AddSymbols("cc", "ccm", "cm³", "cm3", "cm^3", "cm**3")

	MicroLiter.AddAliases("cubic millimeter", "cubic millimetre")
	MicroLiter.AddSymbols("mm³", "mm3", "mm^3", "mm**3")

	FluidOunce.AddSymbols("floz", "fl. oz", "fl.oz", "oz. fl", "oz fl", "ozfl", "ozfl.")
	CubicFoot.AddSymbols("cuft", "cu ft", "cft", "ft^3", "ft**3", "ft3")
	CubicYard.AddSymbols("cuyd", "cu yd", "cyd", "yd^3", "yd**3", "yd3")
	CubicInch.AddSymbols("cuin", "cu in", "cin", "in^3", "in**3", "in3")
	CubicMile.AddSymbols("cumi", "cubmi", "mi^3", "mi**3", "mi3")
	AcreFoot.AddSymbols("acft", "acreft", "acre-ft", "acre ft", "ac-ft")

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
}
