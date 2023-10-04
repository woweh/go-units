package units

var (
	Volume = UnitOptionQuantity("volume")

	// metric

	Liter = NewUnit(
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
	Quart      = NewUnit("quart", "qt", Volume, BI)
	Pint       = NewUnit("pint", "pt", Volume, BI)
	Gallon     = NewUnit("gallon", "gal", Volume, BI)
	FluidOunce = NewUnit("fluid ounce", "fl oz", Volume, BI, UnitOptionAliases("floz"))

	// US
	FluidQuart          = NewUnit("fluid quart", "fl qt", Volume, US)
	FluidPint           = NewUnit("fluid pint", "fl pt", Volume, US)
	FluidGallon         = NewUnit("fluid gallon", "", Volume, US)
	CustomaryFluidOunce = NewUnit("customary fluid ounce", "", Volume, US)
)

func init() {
	NewRatioConversion(Quart, Liter, 1.1365225)
	NewRatioConversion(Pint, Liter, 0.56826125)
	NewRatioConversion(Gallon, Liter, 4.54609)
	NewRatioConversion(FluidOunce, MilliLiter, 28.4130625)

	NewRatioConversion(FluidQuart, Liter, 0.946352946)
	NewRatioConversion(FluidPint, Liter, 0.473176473)
	NewRatioConversion(FluidGallon, Liter, 3.785411784)
	NewRatioConversion(CustomaryFluidOunce, MilliLiter, 29.5735295625)
}
