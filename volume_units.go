package units

var (
	Volume = UnitOptionQuantity("volume")

	// metric

	Liter = newUnit(
		"liter", "l", Volume, SI,
		UnitOptionAliases("litre", "cubic decimeter", "cubic decimetre", "dm³", "dm3", "dm^3", "dm**3"),
	)
	ExaLiter  = Exa(Liter)
	PetaLiter = Peta(Liter)
	TeraLiter = Tera(
		Liter, UnitOptionAliases("cubic kilometer", "cubic kilometre", "km³", "km3", "km^3", "km**3", "㎦"),
	)
	GigaLiter = Giga(
		Liter, UnitOptionAliases("cubic hectometer", "cubic hectometre", "hm³", "hm3", "hm^3", "hm**3", "mcm"),
	)
	MegaLiter = Mega(Liter, UnitOptionAliases("cubic decameter", "cubic decametre", "dam³", "dam3", "dam^3", "dam**3"))
	KiloLiter = Kilo(
		Liter, UnitOptionAliases("cubic meter", "cubic metre", "cum", "cbm", "m³", "m3", "m^3", "m**3", "㎥"),
	)
	HectoLiter = Hecto(Liter)
	DecaLiter  = Deca(Liter)
	DeciLiter  = Deci(Liter)
	CentiLiter = Centi(Liter)
	MilliLiter = Milli(
		Liter,
		UnitOptionAliases("cubic centimeter", "cubic centimetre", "cc", "ccm", "cm³", "cm3", "cm^3", "cm**3", "㎤"),
	)
	MicroLiter = Micro(
		Liter, UnitOptionAliases("cubic millimeter", "cubic millimetre", "mm³", "mm3", "mm^3", "mm**3", "㎣"),
	)
	NanoLiter  = Nano(Liter)
	PicoLiter  = Pico(Liter)
	FemtoLiter = Femto(Liter)
	AttoLiter  = Atto(Liter)

	// imperial
	Quart      = newUnit("quart", "qt", Volume, BI)
	Pint       = newUnit("pint", "pt", Volume, BI)
	Gallon     = newUnit("gallon", "gal", Volume, BI)
	FluidOunce = newUnit("fluid ounce", "fl oz", Volume, BI, UnitOptionAliases("floz"))
	CubicFoot  = newUnit(
		"cubic foot", "ft³", Volume, BI, UnitOptionAliases("cu ft", "cuft", "cft", "ft^3", "ft**3", "ft3"),
	)
	CubicYard = newUnit(
		"cubic yard", "yd³", Volume, BI, UnitOptionAliases("cu yd", "cuyd", "cyd", "yd^3", "yd**3", "yd3"),
	)
	CubicInch = newUnit(
		"cubic inch", "in³", Volume, BI, UnitOptionAliases("cu in", "cuin", "cin", "in^3", "in**3", "in3"),
	)
	CubicMile = newUnit(
		"cubic mile", "mi³", Volume, BI, UnitOptionAliases("cu mi", "cumi", "cubmi", "mi^3", "mi**3", "mi3"),
	)
	AcreFoot = newUnit("acre foot", "ac ft", Volume, BI, UnitOptionAliases("acre-ft", "acreft", "ac-ft", "acft"))

	// US
	FluidQuart          = newUnit("fluid quart", "fl qt", Volume, US)
	FluidPint           = newUnit("fluid pint", "fl pt", Volume, US)
	FluidGallon         = newUnit("fluid gallon", "", Volume, US)
	CustomaryFluidOunce = newUnit("customary fluid ounce", "", Volume, US)
)

func init() {
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
