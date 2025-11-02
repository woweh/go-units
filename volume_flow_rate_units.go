package units

var (
	VolumeFlowRate = Quantity("volume flow rate")

	CubicMeterPerSecond = mustCreateNewUnit(
		"cubic meter per second", "m³/s", VolumeFlowRate, SI,
		Aliases("cubic metre per second"), Symbols("m3/s", "m3s-1"),
	)
	CubicMeterPerMinute = mustCreateNewUnit(
		"cubic meter per minute", "m³/min", VolumeFlowRate, SI,
		Aliases("cubic metre per minute"), Symbols("m3/min", "m3m-1"),
	)
	CubicMeterPerHour = mustCreateNewUnit(
		"cubic meter per hour", "m³/h", VolumeFlowRate, SI,
		Aliases("cubic metre per hour"), Symbols("m3/h", "m3h-1"),
	)
	CubicMeterPerDay = mustCreateNewUnit(
		"cubic meter per day", "m³/d", VolumeFlowRate, SI,
		Aliases("cubic metre per day"), Symbols("m3/d", "m3d-1"),
	)
	CubicDecimeterPerSecond = mustCreateNewUnit(
		"cubic decimeter per second", "dm³/s", VolumeFlowRate, SI,
		Aliases("cubic decimetre per second"), Symbols("dm3/s", "dm3s-1"),
	)
	CubicDecimeterPerMinute = mustCreateNewUnit(
		"cubic decimeter per minute", "dm³/min", VolumeFlowRate, SI,
		Aliases("cubic decimetre per minute"), Symbols("dm3/min", "dm3m-1"),
	)
	CubicDecimeterPerHour = mustCreateNewUnit(
		"cubic decimeter per hour", "dm³/h", VolumeFlowRate, SI,
		Aliases("cubic decimetre per hour"), Symbols("dm3/h", "dm3h-1"),
	)
	CubicDecimeterPerDay = mustCreateNewUnit(
		"cubic decimeter per day", "dm³/d", VolumeFlowRate, SI,
		Aliases("cubic decimetre per day"), Symbols("dm3/d", "dm3d-1"),
	)

	CubicCentimeterPerSecond = mustCreateNewUnit(
		"cubic centimeter per second", "cm³/s", VolumeFlowRate, SI,
		Aliases("cubic centimetre per second"), Symbols("cm3/s", "cm3s-1"),
	)
	CubicCentimeterPerMinute = mustCreateNewUnit(
		"cubic centimeter per minute", "cm³/min", VolumeFlowRate, SI,
		Aliases("cubic centimetre per minute"), Symbols("cm3/min", "cm3m-1"),
	)
	CubicCentimeterPerHour = mustCreateNewUnit(
		"cubic centimeter per hour", "cm³/h", VolumeFlowRate, SI,
		Aliases("cubic centimetre per hour"), Symbols("cm3/h", "cm3h-1"),
	)
	CubicCentimeterPerDay = mustCreateNewUnit(
		"cubic centimeter per day", "cm³/d", VolumeFlowRate, SI,
		Aliases("cubic centimetre per day"), Symbols("cm3/d", "cm3d-1"),
	)

	CubicInchPerSecond = mustCreateNewUnit("cubic inch per second", "in³/s", VolumeFlowRate)
	CubicInchPerMinute = mustCreateNewUnit("cubic inch per minute", "in³/min", VolumeFlowRate)
	CubicInchPerHour   = mustCreateNewUnit("cubic inch per hour", "in³/h", VolumeFlowRate)
	CubicInchPerDay    = mustCreateNewUnit("cubic inch per day", "in³/d", VolumeFlowRate)

	CubicFootPerSecond = mustCreateNewUnit("cubic foot per second", "ft³/s", VolumeFlowRate)
	CubicFootPerMinute = mustCreateNewUnit("cubic foot per minute", "ft³/min", VolumeFlowRate)
	CubicFootPerHour   = mustCreateNewUnit("cubic foot per hour", "ft³/h", VolumeFlowRate)
	CubicFootPerDay    = mustCreateNewUnit("cubic foot per day", "ft³/d", VolumeFlowRate)

	CubicYardPerSecond = mustCreateNewUnit("cubic yard per second", "yd³/s", VolumeFlowRate)
	CubicYardPerMinute = mustCreateNewUnit("cubic yard per minute", "yd³/min", VolumeFlowRate)
	CubicYardPerHour   = mustCreateNewUnit("cubic yard per hour", "yd³/h", VolumeFlowRate)
	CubicYardPerDay    = mustCreateNewUnit("cubic yard per day", "yd³/d", VolumeFlowRate)
)

func init() {
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
