package units

const DimensionlessRatio UnitQuantity = "dimensionless ratio"

var (
	_dimensionlessRatio = Quantity(DimensionlessRatio)

	Fraction         = mustCreateNewUnit("fraction", "", _dimensionlessRatio, SI, Plural(PluralNone))
	Percent          = mustCreateNewUnit("percent", "%", _dimensionlessRatio, SI, Plural(PluralNone))
	Permille         = mustCreateNewUnit("permille", "‰", _dimensionlessRatio, SI, Plural(PluralNone))
	PartsPerMillion  = mustCreateNewUnit("partsPerMillion", "ppm", _dimensionlessRatio, SI, Plural(PluralNone))
	PartsPerBillion  = mustCreateNewUnit("partsPerBillion", "", _dimensionlessRatio, SI, Plural(PluralNone))
	PartsPerTrillion = mustCreateNewUnit("partsPerTrillion", "", _dimensionlessRatio, SI, Plural(PluralNone))
)

func initDimensionlessRatioUnits() {
	NewRatioConversion(Fraction, Percent, 100)
	NewRatioConversion(Fraction, Permille, 1000)
	NewRatioConversion(Fraction, PartsPerMillion, 1000000)
	NewRatioConversion(Fraction, PartsPerBillion, 1000000000)
	NewRatioConversion(Fraction, PartsPerTrillion, 1000000000000)
}
