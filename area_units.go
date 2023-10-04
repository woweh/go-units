package units

var (
	Area = UnitOptionQuantity("area")

	// metric
	SquareMilliMeter = NewUnit("square millimeter", "mm²", Area, SI)
	SquareCentiMeter = NewUnit("square centimeter", "cm²", Area, SI)
	SquareDeciMeter  = NewUnit("square decimeter", "dm²", Area, SI)
	SquareMeter      = NewUnit("square meter", "m²", Area, SI)
	SquareDecaMeter  = NewUnit("square decameter", "dam²", Area, SI)
	SquareHectoMeter = NewUnit("square hectometer", "hm²", Area, SI)
	SquareKiloMeter  = NewUnit("square kilometer", "km²", Area, SI)

	// imperial
	SquareMile = NewUnit("square mile", "mi²", Area, BI)
	Acre       = NewUnit("acre", "ac", Area, BI)
	SquareInch = NewUnit("square inch", "in²", Area, BI)
	SquareFoot = NewUnit("square foot", "ft²", Area, BI)
	SquareYard = NewUnit("square yard", "yd²", Area, BI)
)

func init() {
	// metric
	SquareMeter.AddAliases("square metre", "m2", "m^2", "m**2", "㎡")

	SquareMilliMeter.AddAliases("square millimetre", "mm2", "mm^2", "mm**2", "㎟")
	NewRatioConversion(SquareMeter, SquareMilliMeter, 1000000.0)
	SquareCentiMeter.AddAliases("square centimetre", "cm2", "cm^2", "cm**2", "㎠")
	NewRatioConversion(SquareMeter, SquareCentiMeter, 10000.0)
	SquareDeciMeter.AddAliases("square decimetre", "dm2", "dm^2", "dm**2")
	NewRatioConversion(SquareMeter, SquareDeciMeter, 100.0)
	SquareDecaMeter.AddAliases("square decametre", "dam2", "dam^2", "dam**2", "are")
	NewRatioConversion(SquareMeter, SquareDecaMeter, 0.01)
	SquareHectoMeter.AddAliases("square hectometre", "hm2", "hm^2", "hm**2", "hectare", "ha")
	NewRatioConversion(SquareMeter, SquareHectoMeter, 0.0001)
	SquareKiloMeter.AddAliases("square kilometre", "km2", "km^2", "km**2", "㎢")
	NewRatioConversion(SquareMeter, SquareKiloMeter, 0.000001)

	// imperial
	SquareMile.AddAliases(
		"sqmi", "sqmi.", "sq mi", "sq mi.", "mi2", "mi^2", "mi**2", "sqmi", "sqmi.", "sq mi", "sq mi.",
		"square mi", "square mi.", "sq mile", "sq. mile",
	)
	NewRatioConversion(SquareKiloMeter, SquareMile, 0.38610215854245)

	Acre.AddAliases("acres")
	NewRatioConversion(SquareKiloMeter, Acre, 247.10538146717)

	SquareInch.AddAliases(
		"square inches", "square in", "square in.", "square ins", "square ins.", "in2", "in^2", "in**2",
		"sq in", "sq in.", "sq ins", "sq ins.", "sqin", "sqin.", "sqins", "□″", "sq inches", "sq inch", "inches/-2",
		"inch/-2", "in/-2", "inches2", "inch2", "\"2", "″2",
	)
	NewRatioConversion(SquareMeter, SquareInch, 1550.0031000062)

	SquareFoot.AddAliases(
		"square feet", "square ft", "square ft.", "square feet.", "ft2", "FT2", "ft^2", "ft**2", "sq ft",
		"sq ft.", "sqft", "sqft.", "sqft", "'2",
	)
	NewRatioConversion(SquareMeter, SquareFoot, 10.76391041671)

	SquareYard.AddAliases(
		"square yds", "square yd", "sq yards", "sq yard", "sq yds", "sq yd", "sq.yd.", "yards/-2", "yard/-2",
		"yds/-2", "yd/-2", "yards^2", "yard^2", "yds^2", "yd^2", "yards²", "yard²", "yds²", "yd²",
	)
	NewRatioConversion(SquareMeter, SquareYard, 1.1959900463011)
}
