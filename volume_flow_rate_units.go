package units

// VolumeFlowRate is a unit quantity for volume flow rate
const VolumeFlowRateQuantity UnitQuantity = "volume flow rate"

var (
	_volumeFlowRate = Quantity(VolumeFlowRateQuantity)

	// SI units
	CubicMeterPerSecond = mustCreateNewUnit(
		"cubic meter per second", "m³/s", _volumeFlowRate, SI,
		Aliases("cubic metre per second"), Symbols("m3/s", "m3s-1"),
	)
	CubicMeterPerMinute = mustCreateNewUnit(
		"cubic meter per minute", "m³/min", _volumeFlowRate, SI,
		Aliases("cubic metre per minute"), Symbols("m3/min", "m3m-1"),
	)
	CubicMeterPerHour = mustCreateNewUnit(
		"cubic meter per hour", "m³/h", _volumeFlowRate, SI,
		Aliases("cubic metre per hour"), Symbols("m3/h", "m3h-1"),
	)
	CubicMeterPerDay = mustCreateNewUnit(
		"cubic meter per day", "m³/d", _volumeFlowRate, SI,
		Aliases("cubic metre per day"), Symbols("m3/d", "m3d-1"),
	)
	CubicDecimeterPerSecond = mustCreateNewUnit(
		"cubic decimeter per second", "dm³/s", _volumeFlowRate, SI,
		Aliases("cubic decimetre per second"), Symbols("dm3/s", "dm3s-1"),
	)
	CubicDecimeterPerMinute = mustCreateNewUnit(
		"cubic decimeter per minute", "dm³/min", _volumeFlowRate, SI,
		Aliases("cubic decimetre per minute"), Symbols("dm3/min", "dm3m-1"),
	)
	CubicDecimeterPerHour = mustCreateNewUnit(
		"cubic decimeter per hour", "dm³/h", _volumeFlowRate, SI,
		Aliases("cubic decimetre per hour"), Symbols("dm3/h", "dm3h-1"),
	)
	CubicDecimeterPerDay = mustCreateNewUnit(
		"cubic decimeter per day", "dm³/d", _volumeFlowRate, SI,
		Aliases("cubic decimetre per day"), Symbols("dm3/d", "dm3d-1"),
	)
	CubicCentimeterPerSecond = mustCreateNewUnit(
		"cubic centimeter per second", "cm³/s", _volumeFlowRate, SI,
		Aliases("cubic centimetre per second"), Symbols("cm3/s", "cm3s-1"),
	)
	CubicCentimeterPerMinute = mustCreateNewUnit(
		"cubic centimeter per minute", "cm³/min", _volumeFlowRate, SI,
		Aliases("cubic centimetre per minute"), Symbols("cm3/min", "cm3m-1"),
	)
	CubicCentimeterPerHour = mustCreateNewUnit(
		"cubic centimeter per hour", "cm³/h", _volumeFlowRate, SI,
		Aliases("cubic centimetre per hour"), Symbols("cm3/h", "cm3h-1"),
	)
	CubicCentimeterPerDay = mustCreateNewUnit(
		"cubic centimeter per day", "cm³/d", _volumeFlowRate, SI,
		Aliases("cubic centimetre per day"), Symbols("cm3/d", "cm3d-1"),
	)

	// Imperial/US units
	CubicInchPerSecond = mustCreateNewUnit("cubic inch per second", "in³/s", _volumeFlowRate, BI)
	CubicInchPerMinute = mustCreateNewUnit("cubic inch per minute", "in³/min", _volumeFlowRate, BI)
	CubicInchPerHour   = mustCreateNewUnit("cubic inch per hour", "in³/h", _volumeFlowRate, BI)
	CubicInchPerDay    = mustCreateNewUnit("cubic inch per day", "in³/d", _volumeFlowRate, BI)

	CubicFootPerSecond = mustCreateNewUnit("cubic foot per second", "ft³/s", _volumeFlowRate, BI)
	CubicFootPerMinute = mustCreateNewUnit("cubic foot per minute", "ft³/min", _volumeFlowRate, BI)
	CubicFootPerHour   = mustCreateNewUnit("cubic foot per hour", "ft³/h", _volumeFlowRate, BI)
	CubicFootPerDay    = mustCreateNewUnit("cubic foot per day", "ft³/d", _volumeFlowRate, BI)

	CubicYardPerSecond = mustCreateNewUnit("cubic yard per second", "yd³/s", _volumeFlowRate, BI)
	CubicYardPerMinute = mustCreateNewUnit("cubic yard per minute", "yd³/min", _volumeFlowRate, BI)
	CubicYardPerHour   = mustCreateNewUnit("cubic yard per hour", "yd³/h", _volumeFlowRate, BI)
	CubicYardPerDay    = mustCreateNewUnit("cubic yard per day", "yd³/d", _volumeFlowRate, BI)
)

func initVolumeFlowRateUnits() {
	// use CubicMeterPerSecond as the base unit
	NewRatioConversion(CubicMeterPerSecond, CubicMeterPerMinute, 60)
	NewRatioConversion(CubicMeterPerSecond, CubicMeterPerHour, 3600)
	NewRatioConversion(CubicMeterPerSecond, CubicMeterPerDay, 86400)

	// 1 CubicMeter = 1e3 CubicDecimeter
	r := 1e3
	NewRatioConversion(CubicMeterPerSecond, CubicDecimeterPerSecond, r)
	NewRatioConversion(CubicMeterPerMinute, CubicDecimeterPerMinute, r)
	NewRatioConversion(CubicMeterPerHour, CubicDecimeterPerHour, r)
	NewRatioConversion(CubicMeterPerDay, CubicDecimeterPerDay, r)

	// 1 CubicMeter = 1e6 CubicCentimeter
	r = 1e6
	NewRatioConversion(CubicMeterPerSecond, CubicCentimeterPerSecond, r)
	NewRatioConversion(CubicMeterPerMinute, CubicCentimeterPerMinute, r)
	NewRatioConversion(CubicMeterPerHour, CubicCentimeterPerHour, r)
	NewRatioConversion(CubicMeterPerDay, CubicCentimeterPerDay, r)

	// 1 inch = 0.0254 meter, 1 in3 =  in3 0.0254*0.0254*0.0254 m3
	// 1 m3 = 1 / (0.0254*0.0254*0.0254) in3
	// 1 CubicMeter ~ 61023.7441 in3
	r = 1 / (0.0254 * 0.0254 * 0.0254)
	NewRatioConversion(CubicMeterPerSecond, CubicInchPerSecond, r)
	NewRatioConversion(CubicMeterPerMinute, CubicInchPerMinute, r)
	NewRatioConversion(CubicMeterPerHour, CubicInchPerHour, r)
	NewRatioConversion(CubicMeterPerDay, CubicInchPerDay, r)

	// 1 foot = 0.3048 meter => 1 / (0.3048*0.3048*0.3048)
	// 1 CubicMeter ~ 35.314666721488 ft3
	r = 1 / (0.3048 * 0.3048 * 0.3048)
	NewRatioConversion(CubicMeterPerSecond, CubicFootPerSecond, r)
	NewRatioConversion(CubicMeterPerMinute, CubicFootPerMinute, r)
	NewRatioConversion(CubicMeterPerHour, CubicFootPerHour, r)
	NewRatioConversion(CubicMeterPerDay, CubicFootPerDay, r)

	// 1 yard = 0.9144 meter => 1 / (0.9144*0.9144*0.9144)
	// 1 CubicMeter ~ 1.3079506193144 yd3
	r = 1 / (0.9144 * 0.9144 * 0.9144)
	NewRatioConversion(CubicMeterPerSecond, CubicYardPerSecond, r)
	NewRatioConversion(CubicMeterPerMinute, CubicYardPerMinute, r)
	NewRatioConversion(CubicMeterPerHour, CubicYardPerHour, r)
	NewRatioConversion(CubicMeterPerDay, CubicYardPerDay, r)
}
