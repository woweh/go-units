package units

var (
	// Area is the unit quantity for area.
	Area = NewUnitQuantity("area")

	// metric
	SquareMilliMeter = Area.MustCreateUnit("square millimeter", "mm²", SI)
	SquareCentiMeter = Area.MustCreateUnit("square centimeter", "cm²", SI)
	SquareDeciMeter  = Area.MustCreateUnit("square decimeter", "dm²", SI)
	SquareMeter      = Area.MustCreateUnit("square meter", "m²", SI)
	SquareDecaMeter  = Area.MustCreateUnit("square decameter", "dam²", SI)
	SquareHectoMeter = Area.MustCreateUnit("square hectometer", "hm²", SI)
	SquareKiloMeter  = Area.MustCreateUnit("square kilometer", "km²", SI)

	// SI aliases as distinct units
	Are     = Area.MustCreateUnit("are", "are", SI)
	Hectare = Area.MustCreateUnit("hectare", "ha", SI)

	// imperial
	SquareMile = Area.MustCreateUnit("square mile", "mi²", BI)
	Acre       = Area.MustCreateUnit("acre", "ac", BI)
	SquareInch = Area.MustCreateUnit("square inch", "in²", BI)
	SquareFoot = Area.MustCreateUnit("square foot", "ft²", BI)
	SquareYard = Area.MustCreateUnit("square yard", "yd²", BI)
)

func initAreaUnits() {
	// Derive conversion factors from length units squared
	// area = length²
	NewRatioConversion(SquareMilliMeter, SquareMeter, areaFactor(MilliMeter))
	NewRatioConversion(SquareCentiMeter, SquareMeter, areaFactor(CentiMeter))
	NewRatioConversion(SquareDeciMeter, SquareMeter, areaFactor(DeciMeter))
	NewRatioConversion(SquareMeter, SquareMeter, 1.0)
	NewRatioConversion(SquareDecaMeter, SquareMeter, areaFactor(DecaMeter))
	NewRatioConversion(SquareHectoMeter, SquareMeter, areaFactor(HectoMeter))
	NewRatioConversion(SquareKiloMeter, SquareMeter, areaFactor(KiloMeter))

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

	// SI aliases as distinct units
	NewRatioConversion(Are, SquareDecaMeter, 1.0)
	NewRatioConversion(Hectare, SquareHectoMeter, 1.0)

	// imperial
	NewRatioConversion(SquareInch, SquareMeter, areaFactor(Inch))
	NewRatioConversion(SquareFoot, SquareMeter, areaFactor(Foot))
	NewRatioConversion(SquareYard, SquareMeter, areaFactor(Yard))
	NewRatioConversion(SquareMile, SquareMeter, areaFactor(Mile))

	// Acre is defined as 43,560 square feet
	NewRatioConversion(Acre, SquareMeter, 43560*areaFactor(Foot))

	SquareInch.AddAliases(
		"square inches", "square in", "square in.", "square ins", "square ins.", "sq inches", "sq inch",
		"inches/-2", "inch/-2", "inches2", "inch2",
	)
	SquareInch.AddSymbols("in2", "in^2", "in**2", "in/-2", "sq in", "sq in.", "sq ins", "sq ins.", "sqin", "sqin.", "sqins", "□″", "″2")

	SquareFoot.AddAliases("square feet", "square ft", "square ft.", "square feet.")
	SquareFoot.AddSymbols("ft2", "ft^2", "ft**2", "sq ft", "sq ft.", "sqft", "sqft.", "SF", "sqft", "'2")

	SquareYard.AddAliases("square yds", "square yd", "sq yards", "sq yard", "yards/-2", "yard/-2", "yards^2", "yard^2")
	SquareYard.AddSymbols("yds^2", "yd^2", "yd2", "yards²", "yard²", "yds²", "yds/-2", "yd/-2", "sq yds", "sq yd", "sq.yd.")

	SquareMile.AddAliases("square mi", "square mi.", "sq mile", "sq. mile")
	SquareMile.AddSymbols("mi2", "mi^2", "mi**2", "sqmi", "sqmi.", "sq mi", "sq mi.", "sq.mi.", "sq.mi")

	Acre.AddAliases("acres")
}

// areaFactor computes the area conversion factor from length units to square meters.
// For example, areaFactor(Foot) returns 0.09290304 (1 ft² = 0.09290304 m²).
func areaFactor(length Unit) float64 {
	factor := length.to(Meter)
	return factor * factor
}
