package units

// Volume is a unit quantity for volume
const VolumeQuantity UnitQuantity = "volume"

var (
	_volume = Quantity(VolumeQuantity)

	// SI/metric units
	Liter      = mustCreateNewUnit("liter", "l", _volume, SI, Aliases("litre"))
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

	CubicMeter      = mustCreateNewUnit("cubic meter", "m³", _volume, SI)
	CubicKiloMeter  = mustCreateNewUnit("cubic kilometer", "km³", _volume, SI)
	CubicHectoMeter = mustCreateNewUnit("cubic hectometer", "hm³", _volume, SI)
	CubicDecaMeter  = mustCreateNewUnit("cubic decameter", "dam³", _volume, SI)
	CubicDeciMeter  = mustCreateNewUnit("cubic decimeter", "dm³", _volume, SI)
	CubicCentiMeter = mustCreateNewUnit("cubic centimeter", "cm³", _volume, SI)
	CubicMilliMeter = mustCreateNewUnit("cubic millimeter", "mm³", _volume, SI)

	// Imperial units
	Quart      = mustCreateNewUnit("quart", "qt", _volume, BI)
	Pint       = mustCreateNewUnit("pint", "pt", _volume, BI)
	Gallon     = mustCreateNewUnit("gallon", "gal", _volume, BI)
	FluidOunce = mustCreateNewUnit("fluid ounce", "fl oz", _volume, BI)
	CubicFoot  = mustCreateNewUnit("cubic foot", "ft³", _volume, BI)
	CubicYard  = mustCreateNewUnit("cubic yard", "yd³", _volume, BI)
	CubicInch  = mustCreateNewUnit("cubic inch", "in³", _volume, BI)
	CubicMile  = mustCreateNewUnit("cubic mile", "mi³", _volume, BI)
	AcreFoot   = mustCreateNewUnit("acre foot", "ac ft", _volume, BI)

	// US customary units
	FluidQuart          = mustCreateNewUnit("fluid quart", "fl qt", _volume, US)
	FluidPint           = mustCreateNewUnit("fluid pint", "fl pt", _volume, US)
	FluidGallon         = mustCreateNewUnit("fluid gallon", "", _volume, US)
	CustomaryFluidOunce = mustCreateNewUnit("customary fluid ounce", "", _volume, US)
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
