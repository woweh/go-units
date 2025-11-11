package units

const Friction UnitQuantity = "friction"

var (
	_friction = Quantity(Friction)

	// SI base unit: pascal per meter (Pa/m)
	PascalPerMeter = mustCreateNewUnit("pascal per meter", "Pa/m", _friction, SI)

	// Metric (hydraulic engineering)
	MillimeterOfWaterPerMeter = mustCreateNewUnit("millimeter of water column per meter", "mmH2O/m", _friction)
	MeterOfWaterPerMeter      = mustCreateNewUnit("meter of water column per meter", "mH2O/m", _friction)

	// Imperial
	FootOfWaterPer100Feet = mustCreateNewUnit("foot of water per 100 feet", "FT/100ft", _friction, BI)
	InchOfWaterPer100Feet = mustCreateNewUnit("inch of water per 100 feet", "in-wg/100ft", _friction, BI)
)

func initFrictionUnits() {
	// SI base unit: Pa/m
	// Revit base: mmH2O/m (CF=1.0976133966960913), but SI base is Pa/m
	// mmH2O/m: 1 mmH2O = 9.80665 Pa, so 1 mmH2O/m = 9.80665 Pa/m
	// mH2O/m: 1 mH2O = 9806.65 Pa, so 1 mH2O/m = 9806.65 Pa/m

	NewRatioConversion(MillimeterOfWaterPerMeter, PascalPerMeter, 9.80665)
	NewRatioConversion(MeterOfWaterPerMeter, PascalPerMeter, 9806.65)

	// 1 mH2O/m = 1000 mmH2O/m
	NewRatioConversion(MeterOfWaterPerMeter, MillimeterOfWaterPerMeter, 1000.0)

	// Imperial conversions (from Revit CFs)
	// FT/100ft: 1 FT/100ft = 0.1 mmH2O/m = 0.980665 Pa/m
	NewRatioConversion(FootOfWaterPer100Feet, PascalPerMeter, 0.980665)
	// in-wg/100ft: 1 in-wg/100ft = 1.2 mmH2O/m = 11.76798 Pa/m
	NewRatioConversion(InchOfWaterPer100Feet, PascalPerMeter, 11.76798)

	MillimeterOfWaterPerMeter.AddAliases("millimeters of water column per meter", "millimetre of water column per metre", "mmH2O per m")
	MeterOfWaterPerMeter.AddAliases("meters of water column per meter", "metre of water column per metre", "mH2O per m")
	PascalPerMeter.AddAliases("pascals per meter", "pascal per metre", "pascals per metre")
	FootOfWaterPer100Feet.AddAliases("feet of water per 100 feet", "FT per 100ft")
	InchOfWaterPer100Feet.AddAliases("inches of water per 100 feet", "in-wg per 100ft")
}
