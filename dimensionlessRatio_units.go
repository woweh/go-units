package units

var (
	DimlessRatio = UnitOptionQuantity("dimensionlessRatio")

	Fraction         = newUnit("fraction", "", DimlessRatio, UnitOptionPlural(None, ""))
	Percent          = newUnit("percent", "", DimlessRatio, UnitOptionPlural(None, ""))
	Permille         = newUnit("permille", "", DimlessRatio, UnitOptionPlural(None, ""))
	PartsPerMillion  = newUnit("partsPerMillion", "", DimlessRatio, UnitOptionPlural(None, ""))
	PartsPerBillion  = newUnit("partsPerBillion", "", DimlessRatio, UnitOptionPlural(None, ""))
	PartsPerTrillion = newUnit("partsPerTrillion", "", DimlessRatio, UnitOptionPlural(None, ""))
)

func init() {
	NewRatioConversion(Fraction, Percent, 100)
	NewRatioConversion(Fraction, Permille, 1000)
	NewRatioConversion(Fraction, PartsPerMillion, 1000000)
	NewRatioConversion(Fraction, PartsPerBillion, 1000000000)
	NewRatioConversion(Fraction, PartsPerTrillion, 1000000000000)
}
