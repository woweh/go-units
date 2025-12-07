package units

var (
	// PowerPerLength is the unit quantity for power per length.
	PowerPerLength = NewUnitQuantity("power per length")

	// SI base unit: watt per meter
	WattPerMeter = PowerPerLength.MustCreateUnit("watt per meter", "W/m", SI)

	// Imperial unit
	WattPerFoot = PowerPerLength.MustCreateUnit("watt per foot", "W/ft", BI)
)

func initPowerPerLengthUnits() {
	// Derive from length units (power/length)
	NewRatioConversion(WattPerMeter, WattPerFoot, powerPerLengthFactor(Foot))

	WattPerMeter.AddAliases("watts per meter", "watt per metre", "watts per metre")
	WattPerFoot.AddAliases("watts per foot")
}

// powerPerLengthFactor computes the power per length conversion factor from length units.
// power per length = power / length, for same power unit: 1 / length
func powerPerLengthFactor(length Unit) float64 {
	return length.to(Meter)
}
