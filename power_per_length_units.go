package units

// PowerPerLength is a unit quantity for power per length
const PowerPerLength UnitQuantity = "power per length"

var (
	_powerPerLength = Quantity(PowerPerLength)

	// SI base unit: watt per meter
	WattPerMeter = mustCreateNewUnit("watt per meter", "W/m", _powerPerLength, SI)

	// Imperial unit
	WattPerFoot = mustCreateNewUnit("watt per foot", "W/ft", _powerPerLength, BI)
)

func initPowerPerLengthUnits() {
	// SI base unit: W/m
	// Conversion: 1 W/m = 0.3048 W/ft (1 W/ft = 3.28084 W/m)
	NewRatioConversion(WattPerMeter, WattPerFoot, 0.3048)

	WattPerMeter.AddAliases("watts per meter", "watt per metre", "watts per metre")
	WattPerFoot.AddAliases("watts per foot")
}
