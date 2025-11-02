package units

var (
	Friction = Quantity("friction")

	// Base unit: millimeter of water column per meter (Revit base with CF=1.0976133966960913)
	MillimeterOfWaterPerMeter = mustCreateNewUnit("millimeter of water column per meter", "mmH2O/m", Friction)

	// Other units
	MeterOfWaterPerMeter  = mustCreateNewUnit("meter of water column per meter", "mH2O/m", Friction)
	PascalPerMeter        = mustCreateNewUnit("pascal per meter", "Pa/m", Friction, SI)
	FootOfWaterPer100Feet = mustCreateNewUnit("foot of water per 100 feet", "FT/100ft", Friction)
	InchOfWaterPer100Feet = mustCreateNewUnit("inch of water per 100 feet", "in-wg/100ft", Friction)
)

func init() {
	// From RevitUnits.json:
	// mmH2O/m: CF = 1.0976133966960913
	// mH2O/m: CF = 0.0010976133966960912
	// Pa/m: CF = 10.763910416709722
	// FT/100ft: CF = 0.10976453154631757
	// in-wg/100ft: CF = 1.318453582628646

	// 1 mH2O/m = 1000 mmH2O/m
	NewRatioConversion(MeterOfWaterPerMeter, MillimeterOfWaterPerMeter, 1000.0)

	// From ratios: 10.763910416709722 / 1.0976133966960913 = 9.80665
	// 1 mmH2O/m = 9.80665 Pa/m
	NewRatioConversion(MillimeterOfWaterPerMeter, PascalPerMeter, 9.80665)

	// From ratios: 0.10976453154631757 / 1.0976133966960913 = 0.1
	// 1 mmH2O/m = 0.1 FT/100ft
	NewRatioConversion(MillimeterOfWaterPerMeter, FootOfWaterPer100Feet, 0.1)

	// From ratios: 1.318453582628646 / 1.0976133966960913 = 1.2
	// 1 mmH2O/m = 1.2 in-wg/100ft
	NewRatioConversion(MillimeterOfWaterPerMeter, InchOfWaterPer100Feet, 1.2)

	MillimeterOfWaterPerMeter.AddAliases("millimeters of water column per meter", "millimetre of water column per metre", "mmH2O per m")
	MeterOfWaterPerMeter.AddAliases("meters of water column per meter", "metre of water column per metre", "mH2O per m")
	PascalPerMeter.AddAliases("pascals per meter", "pascal per metre", "pascals per metre")
	FootOfWaterPer100Feet.AddAliases("feet of water per 100 feet", "FT per 100ft")
	InchOfWaterPer100Feet.AddAliases("inches of water per 100 feet", "in-wg per 100ft")
}
