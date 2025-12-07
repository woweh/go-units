package units

var (
	// Density is the unit quantity for density.
	Density = NewUnitQuantity("density")

	// metric
	GramPerCubicCentimeter      = Density.MustCreateUnit("gram per cubic centimeter", "g/cm³", SI)
	KilogramPerCubicCentimeter  = Density.MustCreateUnit("kilogram per cubic centimeter", "kg/cm³", SI)
	GramPerCubicMeter           = Density.MustCreateUnit("gram per cubic meter", "g/m³", SI)
	KilogramPerCubicMeter       = Density.MustCreateUnit("kilogram per cubic meter", "kg/m³", SI)
	GramPerMilliliter           = Density.MustCreateUnit("gram per milliliter", "g/mL", SI)
	GramPerLiter                = Density.MustCreateUnit("gram per liter", "g/L", SI)
	KilogramPerLiter            = Density.MustCreateUnit("kilogram per liter", "kg/L", SI)
	TonnePerCubicMeter          = Density.MustCreateUnit("tonne per cubic meter", "t/m³", SI)
	MilligramPerCubicMeter      = Density.MustCreateUnit("milligram per cubic meter", "mg/m³", SI)
	KilogramPerCubicDecimeter   = Density.MustCreateUnit("kilogram per cubic decimeter", "kg/dm³", SI)
	GramPerCubicDecimeter       = Density.MustCreateUnit("gram per cubic decimeter", "g/dm³", SI)
	MilligramPerCubicDecimeter  = Density.MustCreateUnit("milligram per cubic decimeter", "mg/dm³", SI)
	MilligramPerCubicCentimeter = Density.MustCreateUnit("milligram per cubic centimeter", "mg/cm³", SI)
	KilogramPerMilliliter       = Density.MustCreateUnit("kilogram per milliliter", "kg/mL", SI)
	MilligramPerMilliliter      = Density.MustCreateUnit("milligram per milliliter", "mg/mL", SI)
	MilligramPerLiter           = Density.MustCreateUnit("milligram per liter", "mg/L", SI)

	// imperial (grouped by type and sorted alphabetically)
	OuncePerCubicYard = Density.MustCreateUnit("ounce per cubic yard", "oz/yd³", BI)
	OuncePerCubicFoot = Density.MustCreateUnit("ounce per cubic foot", "oz/ft³", BI)
	OuncePerCubicInch = Density.MustCreateUnit("ounce per cubic inch", "oz/in³", BI)
	OuncePerGallon    = Density.MustCreateUnit("ounce per gallon", "oz/gal", BI)

	PoundPerCubicYard = Density.MustCreateUnit("pound per cubic yard", "lb/yd³", BI)
	PoundPerCubicFoot = Density.MustCreateUnit("pound per cubic foot", "lb/ft³", BI)
	PoundPerCubicInch = Density.MustCreateUnit("pound per cubic inch", "lb/in³", BI)
	PoundPerGallon    = Density.MustCreateUnit("pound per gallon", "lb/gal", BI)

	SlugPerCubicFoot = Density.MustCreateUnit("slug per cubic foot", "slug/ft³", BI)
	SlugPerCubicInch = Density.MustCreateUnit("slug per cubic inch", "slug/in³", BI)
	SlugPerCubicYard = Density.MustCreateUnit("slug per cubic yard", "slug/yd³", BI)

	TonPerCubicYard = Density.MustCreateUnit("ton per cubic yard", "l ton/yd³", BI)
)

// NOTE: Revit uses 1 lb/in³ = 783.80761536 kg/m³, but the correct value per NIST is 27,673.999 kg/m³.
// This code uses the highest possible accuracy for all density conversions.
// If you need to match Revit for interoperability, override the conversion factor for lb/in³ accordingly.

func initDensityUnits() {
	// KilogramPerCubicMeter is the base unit of Density

	// Metric conversions with cubic volume units
	NewRatioConversion(KilogramPerCubicMeter, TonnePerCubicMeter, densityFactor(MetricTon, Meter))
	NewRatioConversion(KilogramPerCubicMeter, GramPerCubicMeter, densityFactor(Gram, Meter))
	NewRatioConversion(KilogramPerCubicMeter, MilligramPerCubicMeter, densityFactor(MilliGram, Meter))
	NewRatioConversion(KilogramPerCubicMeter, KilogramPerCubicDecimeter, densityFactor(KiloGram, DeciMeter))
	NewRatioConversion(KilogramPerCubicMeter, GramPerCubicDecimeter, densityFactor(Gram, DeciMeter))
	NewRatioConversion(KilogramPerCubicMeter, MilligramPerCubicDecimeter, densityFactor(MilliGram, DeciMeter))
	NewRatioConversion(KilogramPerCubicMeter, KilogramPerCubicCentimeter, densityFactor(KiloGram, CentiMeter))
	NewRatioConversion(KilogramPerCubicMeter, GramPerCubicCentimeter, densityFactor(Gram, CentiMeter))
	NewRatioConversion(KilogramPerCubicMeter, MilligramPerCubicCentimeter, densityFactor(MilliGram, CentiMeter))

	// Metric conversions with liter-based volumes (liter = dm³)
	NewRatioConversion(KilogramPerCubicMeter, KilogramPerLiter, densityFactor(KiloGram, DeciMeter))
	NewRatioConversion(KilogramPerCubicMeter, GramPerLiter, densityFactor(Gram, DeciMeter))
	NewRatioConversion(KilogramPerCubicMeter, MilligramPerLiter, densityFactor(MilliGram, DeciMeter))
	NewRatioConversion(KilogramPerCubicMeter, KilogramPerMilliliter, densityFactor(KiloGram, CentiMeter))
	NewRatioConversion(KilogramPerCubicMeter, GramPerMilliliter, densityFactor(Gram, CentiMeter))
	NewRatioConversion(KilogramPerCubicMeter, MilligramPerMilliliter, densityFactor(MilliGram, CentiMeter))

	// Imperial conversions with cubic length units
	NewRatioConversion(KilogramPerCubicMeter, OuncePerCubicYard, densityFactor(Ounce, Yard))
	NewRatioConversion(KilogramPerCubicMeter, OuncePerCubicFoot, densityFactor(Ounce, Foot))
	NewRatioConversion(KilogramPerCubicMeter, OuncePerCubicInch, densityFactor(Ounce, Inch))
	NewRatioConversion(KilogramPerCubicMeter, PoundPerCubicYard, densityFactor(Pound, Yard))
	NewRatioConversion(KilogramPerCubicMeter, PoundPerCubicFoot, densityFactor(Pound, Foot))
	NewRatioConversion(KilogramPerCubicMeter, PoundPerCubicInch, densityFactor(Pound, Inch))
	NewRatioConversion(KilogramPerCubicMeter, SlugPerCubicYard, densityFactor(Slug, Yard))
	NewRatioConversion(KilogramPerCubicMeter, SlugPerCubicFoot, densityFactor(Slug, Foot))
	NewRatioConversion(KilogramPerCubicMeter, SlugPerCubicInch, densityFactor(Slug, Inch))
	NewRatioConversion(KilogramPerCubicMeter, TonPerCubicYard, densityFactor(Ton, Yard))

	// Conversions involving gallons remain hardcoded (gallon is empirical definition)
	NewRatioConversion(KilogramPerCubicMeter, PoundPerGallon, 0.010022412)
}

// densityFactor calculates the conversion factor for density units.
// Density = mass / volume
// Base unit: KilogramPerCubicMeter = Kilogram / Meter³
//
// To convert from kg/m³ to another density unit (e.g., g/cm³):
//   - We need to scale both mass and volume
//   - massRatio: how many target mass units in 1 kg
//     Example: 1 kg = 1000 g, so massRatio = 1000
//   - volumeRatio: how many m³ in 1 target volume unit
//     Example: 1 cm³ = 0.000001 m³, so volumeRatio = 0.000001
//
// The density conversion accounts for:
// - More mass units (multiply by massRatio)
// - Smaller volume units (multiply by volumeRatio for the reciprocal effect)
//
// Example: 1 kg/m³ to g/cm³
// - 1 kg = 1000 g (massRatio = 1000)
// - 1 cm³ = 0.000001 m³ (volumeRatio = 0.000001)
// - In a cm³ volume, we have 0.000001 of the mass from 1 m³
// - That 0.000001 kg = 0.001 g
// - So 1 kg/m³ = 1000 × 0.000001 = 0.001 g/cm³
func densityFactor(mass, length Unit) float64 {
	// How many of the target mass unit in 1 kg
	massRatio := KiloGram.to(mass)
	// How many m³ in 1 unit of target volume (length³)
	volumeRatio := volumeFactor(length)
	// Density factor accounts for both mass and volume scaling
	return massRatio * volumeRatio
}
