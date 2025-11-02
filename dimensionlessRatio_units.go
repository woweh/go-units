package units

var (
	DimensionlessRatio = Quantity("dimensionless ratio")

	Fraction         = mustCreateNewUnit("fraction", "", DimensionlessRatio, Plural(PluralNone))
	Percent          = mustCreateNewUnit("percent", "%", DimensionlessRatio, Plural(PluralNone))
	Permille         = mustCreateNewUnit("permille", "‰", DimensionlessRatio, Plural(PluralNone))
	PartsPerMillion  = mustCreateNewUnit("partsPerMillion", "ppm", DimensionlessRatio, Plural(PluralNone))
	PartsPerBillion  = mustCreateNewUnit("partsPerBillion", "", DimensionlessRatio, Plural(PluralNone))
	PartsPerTrillion = mustCreateNewUnit("partsPerTrillion", "", DimensionlessRatio, Plural(PluralNone))
)

func init() {
	NewRatioConversion(Fraction, Percent, 100)
	NewRatioConversion(Fraction, Permille, 1000)
	NewRatioConversion(Fraction, PartsPerMillion, 1000000)
	NewRatioConversion(Fraction, PartsPerBillion, 1000000000)
	NewRatioConversion(Fraction, PartsPerTrillion, 1000000000000)
}
