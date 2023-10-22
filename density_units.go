package units

var (
	Density = Quantity("density")

	// metric
	GramPerCubicCentimeter     = newUnit("gram per cubic centimeter", "g/cm³", Density, SI)
	KilogramPerCubicCentimeter = newUnit("kilogram per cubic centimeter", "kg/cm³", Density, SI)
	GramPerCubicMeter          = newUnit("gram per cubic meter", "g/m³", Density, SI)
	KilogramPerCubicMeter      = newUnit("kilogram per cubic meter", "kg/m³", Density, SI)
	GramPerMilliliter          = newUnit("gram per milliliter", "g/mL", Density, SI)
	GramPerLiter               = newUnit("gram per liter", "g/L", Density, SI)
	KilogramPerLiter           = newUnit("kilogram per liter", "kg/L", Density, SI)

	// imperial
	OuncePerCubicInch = newUnit("ounce per cubic inch", "oz/in³", Density, BI)
	OuncePerCubicFoot = newUnit("ounce per cubic foot", "oz/ft³", Density, BI)
	OuncePerGallon    = newUnit("ounce per gallon", "oz/gal", Density, BI)
	PoundPerCubicInch = newUnit("pound per cubic inch", "lb/in³", Density, BI)
	PoundPerCubicFoot = newUnit("pound per cubic foot", "lb/ft³", Density, BI)
	PoundPerGallon    = newUnit("pound per gallon", "lb/gal", Density, BI)
	SlugPerCubicFoot  = newUnit("slug per cubic foot", "slug/ft³", Density, BI)
	TonPerCubicYard   = newUnit("ton per cubic yard", "l ton/yd³", Density, BI)
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
	NewRatioConversion(OuncePerCubicInch, KilogramPerCubicMeter, 1729.994)
	NewRatioConversion(OuncePerCubicFoot, KilogramPerCubicMeter, 1.001153)
	NewRatioConversion(OuncePerGallon, KilogramPerCubicMeter, 6.236023)
	NewRatioConversion(PoundPerCubicInch, KilogramPerCubicMeter, 27679.90471)
	NewRatioConversion(PoundPerCubicFoot, KilogramPerCubicMeter, 16.018463)
	NewRatioConversion(PoundPerGallon, KilogramPerCubicMeter, 99.776372)
	NewRatioConversion(SlugPerCubicFoot, KilogramPerCubicMeter, 515.3788184)
	NewRatioConversion(TonPerCubicYard, KilogramPerCubicMeter, 1328.939)
}
