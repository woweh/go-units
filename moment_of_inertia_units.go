package units

var (
	MomentOfInertia = Quantity("moment of inertia")

	// Base unit: foot to the fourth power (Revit base with CF=1.0)
	FootToTheFourthPower = newUnit("foot to the fourth power", "ft⁴", MomentOfInertia, BI)

	// Other units
	InchToTheFourthPower       = newUnit("inch to the fourth power", "in⁴", MomentOfInertia, BI)
	MeterToTheFourthPower      = newUnit("meter to the fourth power", "m⁴", MomentOfInertia, SI)
	CentimeterToTheFourthPower = newUnit("centimeter to the fourth power", "cm⁴", MomentOfInertia, SI)
	MillimeterToTheFourthPower = newUnit("millimeter to the fourth power", "mm⁴", MomentOfInertia, SI)
)

func init() {
	// From RevitUnits.json:
	// ft⁴ (base): CF = 1.0
	// in⁴: CF = 20735.999999999996
	// m⁴: CF = 0.008630974841241602
	// cm⁴: CF = 863097.4841241601
	// mm⁴: CF = 8630974841.2416
	// This means: 1 ft⁴ = 20736 in⁴, etc.
	NewRatioConversion(FootToTheFourthPower, InchToTheFourthPower, 20736.0)
	NewRatioConversion(FootToTheFourthPower, MeterToTheFourthPower, 0.008630974841241602)
	NewRatioConversion(FootToTheFourthPower, CentimeterToTheFourthPower, 863097.4841241601)
	NewRatioConversion(FootToTheFourthPower, MillimeterToTheFourthPower, 8630974841.2416)

	FootToTheFourthPower.AddAliases("feet to the fourth power", "ft^4", "foot^4")
	InchToTheFourthPower.AddAliases("inches to the fourth power", "in^4", "inch^4")
	MeterToTheFourthPower.AddAliases("meters to the fourth power", "metre to the fourth power", "metres to the fourth power", "m^4", "meter^4")
	CentimeterToTheFourthPower.AddAliases("centimeters to the fourth power", "centimetre to the fourth power", "centimetres to the fourth power", "cm^4", "centimeter^4")
	MillimeterToTheFourthPower.AddAliases("millimeters to the fourth power", "millimetre to the fourth power", "millimetres to the fourth power", "mm^4", "millimeter^4")
}
