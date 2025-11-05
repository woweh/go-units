package units

const Density UnitQuantity = "density"

var (
	_density = Quantity(Density)

	// metric
	GramPerCubicCentimeter     = mustCreateNewUnit("gram per cubic centimeter", "g/cm³", _density, SI)
	KilogramPerCubicCentimeter = mustCreateNewUnit("kilogram per cubic centimeter", "kg/cm³", _density, SI)
	GramPerCubicMeter          = mustCreateNewUnit("gram per cubic meter", "g/m³", _density, SI)
	KilogramPerCubicMeter      = mustCreateNewUnit("kilogram per cubic meter", "kg/m³", _density, SI)
	GramPerMilliliter          = mustCreateNewUnit("gram per milliliter", "g/mL", _density, SI)
	GramPerLiter               = mustCreateNewUnit("gram per liter", "g/L", _density, SI)
	KilogramPerLiter           = mustCreateNewUnit("kilogram per liter", "kg/L", _density, SI)

	// imperial
	OuncePerCubicInch = mustCreateNewUnit("ounce per cubic inch", "oz/in³", _density, BI)
	OuncePerCubicFoot = mustCreateNewUnit("ounce per cubic foot", "oz/ft³", _density, BI)
	OuncePerGallon    = mustCreateNewUnit("ounce per gallon", "oz/gal", _density, BI)
	PoundPerCubicInch = mustCreateNewUnit("pound per cubic inch", "lb/in³", _density, BI)
	PoundPerCubicFoot = mustCreateNewUnit("pound per cubic foot", "lb/ft³", _density, BI)
	PoundPerGallon    = mustCreateNewUnit("pound per gallon", "lb/gal", _density, BI)
	SlugPerCubicFoot  = mustCreateNewUnit("slug per cubic foot", "slug/ft³", _density, BI)
	TonPerCubicYard   = mustCreateNewUnit("ton per cubic yard", "l ton/yd³", _density, BI)
)

func init() {
	// KilogramPerCubicMeter is the base unit of Density

	// Metric to Metric conversions
	NewRatioConversion(GramPerCubicCentimeter, KilogramPerCubicMeter, 1000)
	NewRatioConversion(KilogramPerCubicCentimeter, KilogramPerCubicMeter, 1000000)
	NewRatioConversion(GramPerCubicMeter, KilogramPerCubicMeter, 0.001)
	NewRatioConversion(GramPerMilliliter, KilogramPerCubicMeter, 1000)
	NewRatioConversion(GramPerLiter, KilogramPerCubicMeter, 1)
	NewRatioConversion(KilogramPerLiter, KilogramPerCubicMeter, 1000)

	// Imperial to Metric conversions
	NewRatioConversion(OuncePerCubicInch, KilogramPerCubicMeter, 48.9879) // Revit: 1 lb/in³ = 783.8076 kg/m³, 1 oz/in³ = 1/16 lb/in³
	NewRatioConversion(OuncePerCubicFoot, KilogramPerCubicMeter, 1.001153)
	NewRatioConversion(OuncePerGallon, KilogramPerCubicMeter, 6.236023)
	NewRatioConversion(PoundPerCubicInch, KilogramPerCubicMeter, 783.806) // Revit: 1 lb/in³ = 783.8076 kg/m³
	NewRatioConversion(PoundPerCubicFoot, KilogramPerCubicMeter, 16.018463)
	NewRatioConversion(PoundPerGallon, KilogramPerCubicMeter, 99.776372)
	NewRatioConversion(SlugPerCubicFoot, KilogramPerCubicMeter, 515.3788184)
	NewRatioConversion(TonPerCubicYard, KilogramPerCubicMeter, 1328.939)
}
