package units

var (
	// DimensionlessRatio is the unit quantity for dimensionless ratio.
	DimensionlessRatio = NewUnitQuantity("dimensionless ratio")
	Fraction           = DimensionlessRatio.MustCreateUnit("fraction", "", SI, Plural(PluralNone))
	Percent            = DimensionlessRatio.MustCreateUnit("percent", "%", SI, Plural(PluralNone))
	Permille           = DimensionlessRatio.MustCreateUnit("permille", "‰", SI, Plural(PluralNone))
	PartsPerMillion    = DimensionlessRatio.MustCreateUnit("partsPerMillion", "ppm", SI, Plural(PluralNone))
	PartsPerBillion    = DimensionlessRatio.MustCreateUnit("partsPerBillion", "", SI, Plural(PluralNone))
	PartsPerTrillion   = DimensionlessRatio.MustCreateUnit("partsPerTrillion", "", SI, Plural(PluralNone))
)

func initDimensionlessRatioUnits() {
	NewRatioConversion(Fraction, Percent, 100)
	NewRatioConversion(Fraction, Permille, 1000)
	NewRatioConversion(Fraction, PartsPerMillion, 1000000)
	NewRatioConversion(Fraction, PartsPerBillion, 1000000000)
	NewRatioConversion(Fraction, PartsPerTrillion, 1000000000000)
}
