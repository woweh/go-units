package units

var (
	Area = UnitOptionQuantity("area")

	// metric
	SquareMilliMeter = NewUnit(
		"square millimeter", "mm²", Area, SI,
		UnitOptionAliases("square millimetre", "mm2", "mm^2", "mm**2", "㎟"),
	)
	SquareCentiMeter = NewUnit(
		"square centimeter", "cm²", Area, SI,
		UnitOptionAliases("square centimetre", "cm2", "cm^2", "cm**2", "㎠"),
	)
	SquareDeciMeter = NewUnit(
		"square decimeter", "dm²", Area, SI,
		UnitOptionAliases("square decimetre", "dm2", "dm^2", "dm**2"),
	)
	SquareMeter = NewUnit(
		"square meter", "m²", Area, SI,
		UnitOptionAliases("square metre", "m2", "m^2", "m**2", "㎡"),
	)
	SquareDecaMeter = NewUnit(
		"square decameter", "dam²", Area, SI,
		UnitOptionAliases("square decametre", "dam2", "dam^2", "dam**2", "are"),
	)
	SquareHectoMeter = NewUnit(
		"square hectometer", "hm²", Area, SI,
		UnitOptionAliases("square hectometre", "hm2", "hm^2", "hm**2", "hectare", "ha"),
	)
	SquareKiloMeter = NewUnit(
		"square kilometer", "km²", Area, SI,
		UnitOptionAliases("square kilometre", "km2", "km^2", "km**2", "㎢"),
	)

	// imperial
	SquareMile = NewUnit(
		"square mile", "mi²", Area, BI, UnitOptionAliases(
			"sqmi", "sqmi.", "sq mi", "sq mi.", "mi2", "mi^2", "mi**2", "sqmi", "sqmi.", "sq mi", "sq mi.",
			"square mi", "square mi.", "sq mile", "sq. mile",
		),
	)
	Acre       = NewUnit("acre", "ac", Area, BI, UnitOptionAliases("acres"))
	SquareInch = NewUnit(
		"square inch", "in²", Area, BI, UnitOptionAliases(
			"square inches", "square in", "square in.", "square ins", "square ins.", "in2", "in^2", "in**2",
			"sq in", "sq in.", "sq ins", "sq ins.", "sqin", "sqin.", "sqins", "□″", "sq inches", "sq inch", "inches/-2",
			"inch/-2", "in/-2", "inches2", "inch2", "\"2", "″2",
		),
	)
	SquareFoot = NewUnit(
		"square foot", "ft²", Area, BI, UnitOptionAliases(
			"square feet", "square ft", "square ft.", "square feet.", "ft2", "FT2", "ft^2", "ft**2", "sq ft",
			"sq ft.", "sqft", "sqft.", "sqft", "'2",
		),
	)
	SquareYard = NewUnit(
		"square yard", "yd²", Area, BI, UnitOptionAliases(
			"square yds", "square yd", "sq yards", "sq yard", "sq yds", "sq yd", "sq.yd.", "yards/-2", "yard/-2",
			"yds/-2", "yd/-2", "yards^2", "yard^2", "yds^2", "yd^2", "yards²", "yard²", "yds²", "yd²",
		),
	)
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

	// imperial
	NewRatioConversion(SquareInch, SquareMeter, 0.00064516)
	NewRatioConversion(SquareFoot, SquareMeter, 0.09290304)
	NewRatioConversion(SquareYard, SquareMeter, 0.83612736)
	NewRatioConversion(SquareMile, SquareMeter, 2589988.110336)
	NewRatioConversion(Acre, SquareMeter, 4046.8564224)
}
