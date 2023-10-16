package units

var (
	Area = UnitOptionQuantity("area")

	// metric
	SquareMilliMeter = newUnit("square millimeter", "mm²", Area, SI)
	SquareCentiMeter = newUnit("square centimeter", "cm²", Area, SI)
	SquareDeciMeter  = newUnit("square decimeter", "dm²", Area, SI)
	SquareMeter      = newUnit("square meter", "m²", Area, SI)
	SquareDecaMeter  = newUnit("square decameter", "dam²", Area, SI)
	SquareHectoMeter = newUnit("square hectometer", "hm²", Area, SI)
	SquareKiloMeter  = newUnit("square kilometer", "km²", Area, SI)

	// imperial
	SquareMile = newUnit("square mile", "mi²", Area, BI)
	Acre       = newUnit("acre", "ac", Area, BI)
	SquareInch = newUnit("square inch", "in²", Area, BI)
	SquareFoot = newUnit("square foot", "ft²", Area, BI)
	SquareYard = newUnit("square yard", "yd²", Area, BI)
)

func init() {
	// https://www.calculatorsoup.com/calculators/conversions/area.php

	// metric
	NewRatioConversion(SquareMilliMeter, SquareMeter, 0.000001)
	NewRatioConversion(SquareCentiMeter, SquareMeter, 0.0001)
	NewRatioConversion(SquareDeciMeter, SquareMeter, 0.01)
	NewRatioConversion(SquareMeter, SquareMeter, 1.0)
	NewRatioConversion(SquareDecaMeter, SquareMeter, 100.0)
	NewRatioConversion(SquareHectoMeter, SquareMeter, 10000.0)
	NewRatioConversion(SquareKiloMeter, SquareMeter, 1000000.0)

	SquareMilliMeter.AddAliases("square millimetre")
	SquareMilliMeter.AddSymbols("mm2", "mm^2", "mm**2")

	SquareCentiMeter.AddAliases("square centimetre")
	SquareCentiMeter.AddSymbols("cm2", "cm^2", "cm**2")

	SquareDeciMeter.AddAliases("square decimetre")
	SquareDeciMeter.AddSymbols("dm2", "dm^2", "dm**2")

	SquareMeter.AddAliases("square metre")
	SquareMeter.AddSymbols("m2", "m^2", "m**2")

	SquareDecaMeter.AddAliases("square decametre")
	SquareDecaMeter.AddSymbols("dam2", "dam^2", "dam**2", "are")

	SquareHectoMeter.AddAliases("square hectometre", "hectare")
	SquareHectoMeter.AddSymbols("hm2", "hm^2", "hm**2", "ha")

	SquareKiloMeter.AddAliases("square kilometre")
	SquareKiloMeter.AddSymbols("km2", "km^2", "km**2")

	// imperial
	NewRatioConversion(SquareInch, SquareMeter, 0.00064516)
	NewRatioConversion(SquareFoot, SquareMeter, 0.09290304)
	NewRatioConversion(SquareYard, SquareMeter, 0.83612736)
	NewRatioConversion(SquareMile, SquareMeter, 2589988.110336)
	NewRatioConversion(Acre, SquareMeter, 4046.8564224)

	SquareInch.AddAliases(
		"square inches", "square in", "square in.", "square ins", "square ins.", "sq inches", "sq inch",
		"inches/-2", "inch/-2", "inches2", "inch2",
	)
	SquareInch.AddSymbols(
		"in2", "in^2", "in**2", "in/-2", "sq in", "sq in.", "sq ins", "sq ins.", "sqin", "sqin.", "sqins", "□″", "″2",
	)

	SquareFoot.AddAliases("square feet", "square ft", "square ft.", "square feet.")
	SquareFoot.AddSymbols("ft2", "ft^2", "ft**2", "sq ft", "sq ft.", "sqft", "sqft.", "sqft", "'2")

	SquareYard.AddAliases("square yds", "square yd", "sq yards", "sq yard", "yards/-2", "yard/-2", "yards^2", "yard^2")
	SquareYard.AddSymbols(
		"yds^2", "yd^2", "yd2", "yards²", "yard²", "yds²", "yds/-2", "yd/-2", "sq yds", "sq yd", "sq.yd.",
	)

	SquareMile.AddAliases("square mi", "square mi.", "sq mile", "sq. mile")
	SquareMile.AddSymbols("mi2", "mi^2", "mi**2", "sqmi", "sqmi.", "sq mi", "sq mi.", "sq.mi.", "sq.mi")

	Acre.AddAliases("acres")
}
