package units

var (
	// VolumeFlowRateQuantity is the unit quantity for volume flow rate.
	VolumeFlowRateQuantity = NewUnitQuantity("volume flow rate")

	// SI units
	CubicMeterPerSecond = VolumeFlowRateQuantity.MustCreateUnit(
		"cubic meter per second", "m³/s", SI,
		Aliases("cubic metre per second"), Symbols("m3/s", "m3s-1"),
	)
	CubicMeterPerMinute = VolumeFlowRateQuantity.MustCreateUnit(
		"cubic meter per minute", "m³/min", SI,
		Aliases("cubic metre per minute"), Symbols("m3/min", "m3m-1"),
	)
	CubicMeterPerHour = VolumeFlowRateQuantity.MustCreateUnit(
		"cubic meter per hour", "m³/h", SI,
		Aliases("cubic metre per hour"), Symbols("m3/h", "m3h-1"),
	)
	CubicMeterPerDay = VolumeFlowRateQuantity.MustCreateUnit(
		"cubic meter per day", "m³/d", SI,
		Aliases("cubic metre per day"), Symbols("m3/d", "m3d-1"),
	)
	CubicDecimeterPerSecond = VolumeFlowRateQuantity.MustCreateUnit(
		"cubic decimeter per second", "dm³/s", SI,
		Aliases("cubic decimetre per second"), Symbols("dm3/s", "dm3s-1"),
	)
	CubicDecimeterPerMinute = VolumeFlowRateQuantity.MustCreateUnit(
		"cubic decimeter per minute", "dm³/min", SI,
		Aliases("cubic decimetre per minute"), Symbols("dm3/min", "dm3m-1"),
	)
	CubicDecimeterPerHour = VolumeFlowRateQuantity.MustCreateUnit(
		"cubic decimeter per hour", "dm³/h", SI,
		Aliases("cubic decimetre per hour"), Symbols("dm3/h", "dm3h-1"),
	)
	CubicDecimeterPerDay = VolumeFlowRateQuantity.MustCreateUnit(
		"cubic decimeter per day", "dm³/d", SI,
		Aliases("cubic decimetre per day"), Symbols("dm3/d", "dm3d-1"),
	)
	CubicCentimeterPerSecond = VolumeFlowRateQuantity.MustCreateUnit(
		"cubic centimeter per second", "cm³/s", SI,
		Aliases("cubic centimetre per second"), Symbols("cm3/s", "cm3s-1"),
	)
	CubicCentimeterPerMinute = VolumeFlowRateQuantity.MustCreateUnit(
		"cubic centimeter per minute", "cm³/min", SI,
		Aliases("cubic centimetre per minute"), Symbols("cm3/min", "cm3m-1"),
	)
	CubicCentimeterPerHour = VolumeFlowRateQuantity.MustCreateUnit(
		"cubic centimeter per hour", "cm³/h", SI,
		Aliases("cubic centimetre per hour"), Symbols("cm3/h", "cm3h-1"),
	)
	CubicCentimeterPerDay = VolumeFlowRateQuantity.MustCreateUnit(
		"cubic centimeter per day", "cm³/d", SI,
		Aliases("cubic centimetre per day"), Symbols("cm3/d", "cm3d-1"),
	)

	// Imperial/US units
	CubicInchPerSecond = VolumeFlowRateQuantity.MustCreateUnit("cubic inch per second", "in³/s", BI)
	CubicInchPerMinute = VolumeFlowRateQuantity.MustCreateUnit("cubic inch per minute", "in³/min", BI)
	CubicInchPerHour   = VolumeFlowRateQuantity.MustCreateUnit("cubic inch per hour", "in³/h", BI)
	CubicInchPerDay    = VolumeFlowRateQuantity.MustCreateUnit("cubic inch per day", "in³/d", BI)

	CubicFootPerSecond = VolumeFlowRateQuantity.MustCreateUnit("cubic foot per second", "ft³/s", BI)
	CubicFootPerMinute = VolumeFlowRateQuantity.MustCreateUnit("cubic foot per minute", "ft³/min", BI)
	CubicFootPerHour   = VolumeFlowRateQuantity.MustCreateUnit("cubic foot per hour", "ft³/h", BI)
	CubicFootPerDay    = VolumeFlowRateQuantity.MustCreateUnit("cubic foot per day", "ft³/d", BI)

	CubicYardPerSecond = VolumeFlowRateQuantity.MustCreateUnit("cubic yard per second", "yd³/s", BI)
	CubicYardPerMinute = VolumeFlowRateQuantity.MustCreateUnit("cubic yard per minute", "yd³/min", BI)
	CubicYardPerHour   = VolumeFlowRateQuantity.MustCreateUnit("cubic yard per hour", "yd³/h", BI)
	CubicYardPerDay    = VolumeFlowRateQuantity.MustCreateUnit("cubic yard per day", "yd³/d", BI)
)

func initVolumeFlowRateUnits() {
	// use CubicMeterPerSecond as the base unit
	// Note: NewRatioConversion expects "how many target units in 1 source unit"
	// For m³/s to m³/min: 1 m³/s = 60 m³/min, so ratio = volumeFlowRateFactor
	NewRatioConversion(CubicMeterPerSecond, CubicMeterPerMinute, volumeFlowRateFactor(Meter, Minute))
	NewRatioConversion(CubicMeterPerSecond, CubicMeterPerHour, volumeFlowRateFactor(Meter, Hour))
	NewRatioConversion(CubicMeterPerSecond, CubicMeterPerDay, volumeFlowRateFactor(Meter, Day))

	// Decimeter-based volumes
	NewRatioConversion(CubicMeterPerSecond, CubicDecimeterPerSecond, volumeFlowRateFactor(DeciMeter, Second))
	NewRatioConversion(CubicMeterPerMinute, CubicDecimeterPerMinute, volumeFlowRateFactor(DeciMeter, Minute))
	NewRatioConversion(CubicMeterPerHour, CubicDecimeterPerHour, volumeFlowRateFactor(DeciMeter, Hour))
	NewRatioConversion(CubicMeterPerDay, CubicDecimeterPerDay, volumeFlowRateFactor(DeciMeter, Day))

	// Centimeter-based volumes
	NewRatioConversion(CubicMeterPerSecond, CubicCentimeterPerSecond, volumeFlowRateFactor(CentiMeter, Second))
	NewRatioConversion(CubicMeterPerMinute, CubicCentimeterPerMinute, volumeFlowRateFactor(CentiMeter, Minute))
	NewRatioConversion(CubicMeterPerHour, CubicCentimeterPerHour, volumeFlowRateFactor(CentiMeter, Hour))
	NewRatioConversion(CubicMeterPerDay, CubicCentimeterPerDay, volumeFlowRateFactor(CentiMeter, Day))

	// Imperial inch-based volumes
	NewRatioConversion(CubicMeterPerSecond, CubicInchPerSecond, volumeFlowRateFactor(Inch, Second))
	NewRatioConversion(CubicMeterPerMinute, CubicInchPerMinute, volumeFlowRateFactor(Inch, Minute))
	NewRatioConversion(CubicMeterPerHour, CubicInchPerHour, volumeFlowRateFactor(Inch, Hour))
	NewRatioConversion(CubicMeterPerDay, CubicInchPerDay, volumeFlowRateFactor(Inch, Day))

	// Imperial foot-based volumes
	NewRatioConversion(CubicMeterPerSecond, CubicFootPerSecond, volumeFlowRateFactor(Foot, Second))
	NewRatioConversion(CubicMeterPerMinute, CubicFootPerMinute, volumeFlowRateFactor(Foot, Minute))
	NewRatioConversion(CubicMeterPerHour, CubicFootPerHour, volumeFlowRateFactor(Foot, Hour))
	NewRatioConversion(CubicMeterPerDay, CubicFootPerDay, volumeFlowRateFactor(Foot, Day))

	// Imperial yard-based volumes
	NewRatioConversion(CubicMeterPerSecond, CubicYardPerSecond, volumeFlowRateFactor(Yard, Second))
	NewRatioConversion(CubicMeterPerMinute, CubicYardPerMinute, volumeFlowRateFactor(Yard, Minute))
	NewRatioConversion(CubicMeterPerHour, CubicYardPerHour, volumeFlowRateFactor(Yard, Hour))
	NewRatioConversion(CubicMeterPerDay, CubicYardPerDay, volumeFlowRateFactor(Yard, Day))
}

// volumeFlowRateFactor calculates the conversion factor for volume flow rate units.
// VolumeFlowRate = volume / time
// Base unit: CubicMeterPerSecond = m³/s
//
// Example: ft³/min
// - volumeRatio: how many m³ in 1 ft³
// - timeRatio: how many s in 1 min
// - volumeFlowRateFactor = volumeRatio / timeRatio
func volumeFlowRateFactor(length, time Unit) float64 {
	// Return how many target units (length^3 / time) there are in 1 m³/s.
	// 1 m³/s = (1 m³) / (1 s). For a target unit (L^3 / T):
	// ratio = (seconds in 1 target time) / (m³ in 1 target volume)
	volRatio := volumeFactor(length) // m³ per 1 target volume
	timeRatio := time.to(Second)     // seconds per 1 target time
	return timeRatio / volRatio
}
