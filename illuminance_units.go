package units

var (
	// Illuminance is the unit quantity for illuminance.
	Illuminance = NewUnitQuantity("illuminance")

	// SI base unit: lux
	Lux = Illuminance.MustCreateUnit("lux", "lx", SI)

	// Imperial/US unit
	Footcandle = Illuminance.MustCreateUnit("footcandle", "Ftc", BI)
)

func initIlluminanceUnits() {
	// SI base unit: lux
	// Conversion: 1 lx = 0.09290304 Ftc (1 Ftc = 10.76391 lx)
	NewRatioConversion(Lux, Footcandle, 0.09290304)

	Footcandle.AddAliases("footcandles", "foot-candle", "foot-candles", "fc")
}

// illuminanceFactor computes the illuminance conversion factor from area units.
// illuminance = lumen / area
func illuminanceFactor(areaLength Unit) float64 {
	return areaFactor(areaLength)
}
