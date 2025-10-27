package units

var (
	Luminance = Quantity("luminance")

	// Base unit: candela per square foot (Revit base with CF=1.0)
	CandelaPerSquareFoot = newUnit("candela per square foot", "cd/ft²", Luminance, BI)

	// SI and other units
	CandelaPerSquareMeter = newUnit("candela per square meter", "cd/m²", Luminance, SI)
	Footlambert           = newUnit("footlambert", "ftL", Luminance, BI)
)

func init() {
	// From RevitUnits.json:
	// cd/ft² (base): CF = 1.0
	// cd/m²: CF = 10.763910416709722
	// ftL: CF = 3.1415926535897896
	// This means: 1 cd/ft² = 10.763910416709722 cd/m²
	//             1 cd/ft² = 3.1415926535897896 ftL
	NewRatioConversion(CandelaPerSquareFoot, CandelaPerSquareMeter, 10.763910416709722)
	NewRatioConversion(CandelaPerSquareFoot, Footlambert, 3.1415926535897896)

	CandelaPerSquareFoot.AddAliases("candelas per square foot")
	CandelaPerSquareMeter.AddAliases("candelas per square meter", "candela per square metre", "candelas per square metre")
	Footlambert.AddAliases("footlamberts", "foot-lambert", "foot-lamberts", "fL")
}
