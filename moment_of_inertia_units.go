package units

var (
	MomentOfInertia = Quantity("moment of inertia")

	// SI base unit: meter to the fourth power
	MeterToTheFourthPower      = mustCreateNewUnit("meter to the fourth power", "m⁴", MomentOfInertia, SI)
	CentimeterToTheFourthPower = mustCreateNewUnit("centimeter to the fourth power", "cm⁴", MomentOfInertia, SI)
	MillimeterToTheFourthPower = mustCreateNewUnit("millimeter to the fourth power", "mm⁴", MomentOfInertia, SI)

	// Imperial/US units
	FootToTheFourthPower = mustCreateNewUnit("foot to the fourth power", "ft⁴", MomentOfInertia, BI)
	InchToTheFourthPower = mustCreateNewUnit("inch to the fourth power", "in⁴", MomentOfInertia, BI)
)

func init() {
	// SI base unit: m⁴
	// Conversions: 1 m⁴ = 115.86 ft⁴ = 2402510 in⁴ = 100000000 cm⁴ = 1e12 mm⁴
	NewRatioConversion(MeterToTheFourthPower, FootToTheFourthPower, 115.86176745901283)
	NewRatioConversion(MeterToTheFourthPower, InchToTheFourthPower, 2402509.6100288294)
	NewRatioConversion(MeterToTheFourthPower, CentimeterToTheFourthPower, 100000000.0)
	NewRatioConversion(MeterToTheFourthPower, MillimeterToTheFourthPower, 1000000000000.0)

	MeterToTheFourthPower.AddAliases("meters to the fourth power", "metre to the fourth power", "metres to the fourth power", "m^4", "meter^4")
	CentimeterToTheFourthPower.AddAliases("centimeters to the fourth power", "centimetre to the fourth power", "centimetres to the fourth power", "cm^4", "centimeter^4")
	MillimeterToTheFourthPower.AddAliases("millimeters to the fourth power", "millimetre to the fourth power", "millimetres to the fourth power", "mm^4", "millimeter^4")
	FootToTheFourthPower.AddAliases("feet to the fourth power", "ft^4", "foot^4")
	InchToTheFourthPower.AddAliases("inches to the fourth power", "in^4", "inch^4")
}
