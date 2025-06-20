package units

var (
	DimensionlessRatio = Quantity("dimensionless ratio")

	Fraction         = newUnit("fraction", "", DimensionlessRatio, Plural(PluralNone))
	Percent          = newUnit("percent", "%", DimensionlessRatio, Plural(PluralNone))
	Permille         = newUnit("permille", "‰", DimensionlessRatio, Plural(PluralNone))
	PartsPerMillion  = newUnit("partsPerMillion", "ppm", DimensionlessRatio, Plural(PluralNone))
	PartsPerBillion  = newUnit("partsPerBillion", "", DimensionlessRatio, Plural(PluralNone))
	PartsPerTrillion = newUnit("partsPerTrillion", "", DimensionlessRatio, Plural(PluralNone))
)

func init() {
	NewRatioConversion(Fraction, Percent, 100)
	NewRatioConversion(Fraction, Permille, 1000)
	NewRatioConversion(Fraction, PartsPerMillion, 1000000)
	NewRatioConversion(Fraction, PartsPerBillion, 1000000000)
	NewRatioConversion(Fraction, PartsPerTrillion, 1000000000000)
}
