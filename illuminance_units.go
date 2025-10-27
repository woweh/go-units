package units

var (
	Illuminance = Quantity("illuminance")

	// Base unit: footcandle (Revit base with CF=1.0)
	Footcandle = newUnit("footcandle", "Ftc", Illuminance, BI)

	// SI unit
	Lux = newUnit("lux", "lx", Illuminance, SI)
)

func init() {
	// From RevitUnits.json:
	// Ftc (base): CF = 1.0
	// lx: CF = 10.763910416709722
	// This means: 1 Ftc = 10.763910416709722 lx
	NewRatioConversion(Footcandle, Lux, 10.763910416709722)

	Footcandle.AddAliases("footcandles", "foot-candle", "foot-candles", "fc")
}
