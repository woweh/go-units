package units

const Density UnitQuantity = "density"

var (
	_density = Quantity(Density)

	// metric
	GramPerCubicCentimeter      = mustCreateNewUnit("gram per cubic centimeter", "g/cm³", _density, SI)
	KilogramPerCubicCentimeter  = mustCreateNewUnit("kilogram per cubic centimeter", "kg/cm³", _density, SI)
	GramPerCubicMeter           = mustCreateNewUnit("gram per cubic meter", "g/m³", _density, SI)
	KilogramPerCubicMeter       = mustCreateNewUnit("kilogram per cubic meter", "kg/m³", _density, SI)
	GramPerMilliliter           = mustCreateNewUnit("gram per milliliter", "g/mL", _density, SI)
	GramPerLiter                = mustCreateNewUnit("gram per liter", "g/L", _density, SI)
	KilogramPerLiter            = mustCreateNewUnit("kilogram per liter", "kg/L", _density, SI)
	TonnePerCubicMeter          = mustCreateNewUnit("tonne per cubic meter", "t/m³", _density, SI)
	MilligramPerCubicMeter      = mustCreateNewUnit("milligram per cubic meter", "mg/m³", _density, SI)
	KilogramPerCubicDecimeter   = mustCreateNewUnit("kilogram per cubic decimeter", "kg/dm³", _density, SI)
	GramPerCubicDecimeter       = mustCreateNewUnit("gram per cubic decimeter", "g/dm³", _density, SI)
	MilligramPerCubicDecimeter  = mustCreateNewUnit("milligram per cubic decimeter", "mg/dm³", _density, SI)
	MilligramPerCubicCentimeter = mustCreateNewUnit("milligram per cubic centimeter", "mg/cm³", _density, SI)
	KilogramPerMilliliter       = mustCreateNewUnit("kilogram per milliliter", "kg/mL", _density, SI)
	MilligramPerMilliliter      = mustCreateNewUnit("milligram per milliliter", "mg/mL", _density, SI)
	MilligramPerLiter           = mustCreateNewUnit("milligram per liter", "mg/L", _density, SI)

	// imperial (grouped by type and sorted alphabetically)
	OuncePerCubicYard = mustCreateNewUnit("ounce per cubic yard", "oz/yd³", _density, BI)
	OuncePerCubicFoot = mustCreateNewUnit("ounce per cubic foot", "oz/ft³", _density, BI)
	OuncePerCubicInch = mustCreateNewUnit("ounce per cubic inch", "oz/in³", _density, BI)
	OuncePerGallon    = mustCreateNewUnit("ounce per gallon", "oz/gal", _density, BI)

	PoundPerCubicYard = mustCreateNewUnit("pound per cubic yard", "lb/yd³", _density, BI)
	PoundPerCubicFoot = mustCreateNewUnit("pound per cubic foot", "lb/ft³", _density, BI)
	PoundPerCubicInch = mustCreateNewUnit("pound per cubic inch", "lb/in³", _density, BI)
	PoundPerGallon    = mustCreateNewUnit("pound per gallon", "lb/gal", _density, BI)

	SlugPerCubicFoot = mustCreateNewUnit("slug per cubic foot", "slug/ft³", _density, BI)
	SlugPerCubicInch = mustCreateNewUnit("slug per cubic inch", "slug/in³", _density, BI)
	SlugPerCubicYard = mustCreateNewUnit("slug per cubic yard", "slug/yd³", _density, BI)

	TonPerCubicYard = mustCreateNewUnit("ton per cubic yard", "l ton/yd³", _density, BI)
)

// NOTE: Revit uses 1 lb/in³ = 783.80761536 kg/m³, but the correct value per NIST is 27,673.999 kg/m³.
// This code uses the highest possible accuracy for all density conversions.
// If you need to match Revit for interoperability, override the conversion factor for lb/in³ accordingly.

func initDensityUnits() {
	// KilogramPerCubicMeter is the base unit of Density

	// Metric conversions
	NewRatioConversion(KilogramPerCubicMeter, TonnePerCubicMeter, 0.001)
	NewRatioConversion(KilogramPerCubicMeter, GramPerCubicMeter, 1000)
	NewRatioConversion(KilogramPerCubicMeter, MilligramPerCubicMeter, 1000000)
	NewRatioConversion(KilogramPerCubicMeter, KilogramPerLiter, 0.001)
	NewRatioConversion(KilogramPerCubicMeter, GramPerLiter, 1)
	NewRatioConversion(KilogramPerCubicMeter, MilligramPerLiter, 1000)
	NewRatioConversion(KilogramPerCubicMeter, KilogramPerCubicDecimeter, 0.001)
	NewRatioConversion(KilogramPerCubicMeter, GramPerCubicDecimeter, 1)
	NewRatioConversion(KilogramPerCubicMeter, MilligramPerCubicDecimeter, 1000)
	NewRatioConversion(KilogramPerCubicMeter, KilogramPerCubicCentimeter, 0.000001)
	NewRatioConversion(KilogramPerCubicMeter, GramPerCubicCentimeter, 0.001)
	NewRatioConversion(KilogramPerCubicMeter, MilligramPerCubicCentimeter, 1)
	NewRatioConversion(KilogramPerCubicMeter, KilogramPerMilliliter, 0.000001)
	NewRatioConversion(KilogramPerCubicMeter, GramPerMilliliter, 0.001)
	NewRatioConversion(KilogramPerCubicMeter, MilligramPerMilliliter, 1)

	// Imperial conversions
	NewRatioConversion(KilogramPerCubicMeter, OuncePerCubicYard, 26.968879)
	NewRatioConversion(KilogramPerCubicMeter, OuncePerCubicFoot, 0.998847369)
	NewRatioConversion(KilogramPerCubicMeter, OuncePerCubicInch, 0.000578036672)
	NewRatioConversion(KilogramPerCubicMeter, OuncePerGallon, 0.13352647123)

	NewRatioConversion(KilogramPerCubicMeter, PoundPerCubicYard, 1.6855549)
	NewRatioConversion(KilogramPerCubicMeter, PoundPerCubicFoot, 0.062427960576)
	NewRatioConversion(KilogramPerCubicMeter, PoundPerCubicInch, 0.000036127292)
	NewRatioConversion(KilogramPerCubicMeter, PoundPerGallon, 0.008345404452)

	NewRatioConversion(KilogramPerCubicMeter, SlugPerCubicYard, 0.052388649)
	NewRatioConversion(KilogramPerCubicMeter, SlugPerCubicFoot, 0.001940320)
	NewRatioConversion(KilogramPerCubicMeter, SlugPerCubicInch, 0.0000011228706)

	NewRatioConversion(KilogramPerCubicMeter, TonPerCubicYard, 0.000842777)
}
