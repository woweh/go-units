package units

var (
	// Luminance is the unit quantity for luminance.
	Luminance = NewUnitQuantity("luminance")

	// SI base unit: candela per square meter
	CandelaPerSquareMeter = Luminance.MustCreateUnit("candela per square meter", "cd/m²", SI)

	// Imperial/US units
	CandelaPerSquareFoot = Luminance.MustCreateUnit("candela per square foot", "cd/ft²", BI)
	Footlambert          = Luminance.MustCreateUnit("footlambert", "ftL", BI)
)

func initLuminanceUnits() {
	// SI base unit: cd/m²
	// Conversions: 1 cd/m² = 0.09290304 cd/ft² = 0.29186351 ftL
	NewRatioConversion(CandelaPerSquareMeter, CandelaPerSquareFoot, 0.09290304)
	NewRatioConversion(CandelaPerSquareMeter, Footlambert, 0.29186351)

	CandelaPerSquareMeter.AddAliases("candelas per square meter", "candela per square metre", "candelas per square metre")
	CandelaPerSquareFoot.AddAliases("candelas per square foot")
	Footlambert.AddAliases("footlamberts", "foot-lambert", "foot-lamberts", "fL")
}

// luminanceFactor computes the luminance conversion factor from area units.
// luminance = candela / area
func luminanceFactor(areaLength Unit) float64 {
	return areaFactor(areaLength)
}
