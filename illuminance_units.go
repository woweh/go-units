package units

const Illuminance UnitQuantity = "illuminance"

var (
	_illuminance = Quantity(Illuminance)

	// SI base unit: lux
	Lux = mustCreateNewUnit("lux", "lx", _illuminance, SI)

	// Imperial/US unit
	Footcandle = mustCreateNewUnit("footcandle", "Ftc", _illuminance, BI)
)

func init() {
	// SI base unit: lux
	// Conversion: 1 lx = 0.09290304 Ftc (1 Ftc = 10.76391 lx)
	NewRatioConversion(Lux, Footcandle, 0.09290304)

	Footcandle.AddAliases("footcandles", "foot-candle", "foot-candles", "fc")
}
