package units

var (
	// MomentOfInertia is the unit quantity for moment of inertia.
	MomentOfInertia = NewUnitQuantity("moment of inertia")

	// SI base unit: meter to the fourth power
	MeterToTheFourthPower      = MomentOfInertia.MustCreateUnit("meter to the fourth power", "m⁴", SI)
	CentimeterToTheFourthPower = MomentOfInertia.MustCreateUnit("centimeter to the fourth power", "cm⁴", SI)
	MillimeterToTheFourthPower = MomentOfInertia.MustCreateUnit("millimeter to the fourth power", "mm⁴", SI)

	// Imperial/US units
	FootToTheFourthPower = MomentOfInertia.MustCreateUnit("foot to the fourth power", "ft⁴", BI)
	InchToTheFourthPower = MomentOfInertia.MustCreateUnit("inch to the fourth power", "in⁴", BI)
)

func initMomentOfInertiaUnits() {
	// Derive conversion factors from length⁴
	NewRatioConversion(MeterToTheFourthPower, FootToTheFourthPower, momentOfInertiaFactor(Foot))
	NewRatioConversion(MeterToTheFourthPower, InchToTheFourthPower, momentOfInertiaFactor(Inch))
	NewRatioConversion(MeterToTheFourthPower, CentimeterToTheFourthPower, momentOfInertiaFactor(CentiMeter))
	NewRatioConversion(MeterToTheFourthPower, MillimeterToTheFourthPower, momentOfInertiaFactor(MilliMeter))

	MeterToTheFourthPower.AddAliases("meters to the fourth power", "metre to the fourth power", "metres to the fourth power", "m^4", "meter^4")
	CentimeterToTheFourthPower.AddAliases("centimeters to the fourth power", "centimetre to the fourth power", "centimetres to the fourth power", "cm^4", "centimeter^4")
	MillimeterToTheFourthPower.AddAliases("millimeters to the fourth power", "millimetre to the fourth power", "millimetres to the fourth power", "mm^4", "millimeter^4")
	FootToTheFourthPower.AddAliases("feet to the fourth power", "ft^4", "foot^4")
	InchToTheFourthPower.AddAliases("inches to the fourth power", "in^4", "inch^4")
}

// momentOfInertiaFactor computes the moment of inertia conversion factor from length units.
// moment of inertia = length⁴
func momentOfInertiaFactor(length Unit) float64 {
	factor := Meter.to(length)
	return factor * factor * factor * factor
}
