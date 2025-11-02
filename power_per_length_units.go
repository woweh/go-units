package units

var (
	PowerPerLength = Quantity("power per length")

	// Base unit: watt per meter (Revit base with CF=0.3048)
	WattPerMeter = mustCreateNewUnit("watt per meter", "W/m", PowerPerLength, SI)

	// Imperial unit
	WattPerFoot = mustCreateNewUnit("watt per foot", "W/ft", PowerPerLength)
)

func init() {
	// From RevitUnits.json:
	// W/m: CF = 0.3048
	// W/ft: CF = 0.09290304
	// Ratio: 0.3048 / 0.09290304 = 3.28084
	// This means: 1 W/m = 0.3048 W/ft
	// Or: 1 W/ft = 3.28084 W/m
	NewRatioConversion(WattPerFoot, WattPerMeter, 3.28084)

	WattPerMeter.AddAliases("watts per meter", "watt per metre", "watts per metre")
	WattPerFoot.AddAliases("watts per foot")
}
