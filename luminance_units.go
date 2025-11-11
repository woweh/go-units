package units

const Luminance UnitQuantity = "luminance"

var (
	_luminance = Quantity(Luminance)

	// SI base unit: candela per square meter
	CandelaPerSquareMeter = mustCreateNewUnit("candela per square meter", "cd/m²", _luminance, SI)

	// Imperial/US units
	CandelaPerSquareFoot = mustCreateNewUnit("candela per square foot", "cd/ft²", _luminance, BI)
	Footlambert          = mustCreateNewUnit("footlambert", "ftL", _luminance, BI)
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
