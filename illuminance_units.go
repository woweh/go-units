package units

var (
	Illuminance = Quantity("illuminance")

	// SI base unit: lux
	Lux = newUnit("lux", "lx", Illuminance, SI)

	// Imperial/US unit
	Footcandle = newUnit("footcandle", "Ftc", Illuminance, BI)
)

func init() {
	// SI base unit: lux
	// Conversion: 1 lx = 0.092903 Ftc
	NewRatioConversion(Lux, Footcandle, 0.09290304)

	Footcandle.AddAliases("footcandles", "foot-candle", "foot-candles", "fc")
}
