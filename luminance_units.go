package units

var (
	Luminance = Quantity("luminance")

	// SI base unit: candela per square meter
	CandelaPerSquareMeter = newUnit("candela per square meter", "cd/m²", Luminance, SI)

	// Imperial/US units
	CandelaPerSquareFoot = newUnit("candela per square foot", "cd/ft²", Luminance, BI)
	Footlambert          = newUnit("footlambert", "ftL", Luminance, BI)
)

func init() {
	// SI base unit: cd/m²
	// Conversions: 1 cd/m² = 0.092903 cd/ft² = 0.291864 ftL
	NewRatioConversion(CandelaPerSquareMeter, CandelaPerSquareFoot, 0.09290304)
	NewRatioConversion(CandelaPerSquareMeter, Footlambert, 0.29186350801620855)

	CandelaPerSquareMeter.AddAliases("candelas per square meter", "candela per square metre", "candelas per square metre")
	CandelaPerSquareFoot.AddAliases("candelas per square foot")
	Footlambert.AddAliases("footlamberts", "foot-lambert", "foot-lamberts", "fL")
}
