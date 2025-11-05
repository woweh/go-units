package units

// ThermalInsulance is a unit quantity for thermal insulance
const ThermalInsulance UnitQuantity = "thermal insulance"

var (
	_thermalInsulance = Quantity(ThermalInsulance)

	// SI base unit: kelvin square meter per watt (K·m²/W)
	KelvinSquareMeterPerWatt = mustCreateNewUnit(
		"kelvin square meter per watt", "K·m²/W", _thermalInsulance, SI,
		Aliases("R-value", "RSI", "RSI-value"),
		Symbols("K*m2/W", "°C⋅m²/W", "°C*m2/W", "m2.K.W-1", "m²·K/W", "m2*K/W"),
	)

	// Imperial/US unit: degree Fahrenheit hour square foot per British thermal unit
	DegreeFahrenheitHourSquareFootPerBtu = mustCreateNewUnit(
		"degree Fahrenheit hour square foot per British thermal unit", "°F⋅hr⋅ft²/Btu", _thermalInsulance, BI,
		Aliases(
			"degree Fahrenheit hour square foot per British thermal unitIT",
			"degree Fahrenheit hour square foot per British thermal unitth",
			"°F · h · ft2/BtuIT", "°F*hr*ft2/Btu", "°F⋅ft²⋅h/BTU", "°F⋅ft2⋅h/BTU", "°F*ft2*h/BTU",
		),
	)
)

func init() {
	// SI base unit: K·m²/W
	// 1 I-P R-value = 0.1761102 SI R-value (NIST, Revit)
	NewRatioConversion(DegreeFahrenheitHourSquareFootPerBtu, KelvinSquareMeterPerWatt, 0.1761102)

	// https://en.wikipedia.org/wiki/R-value_(insulation)#Units:_metric_(SI)_vs._inch-pound_(I-P) :
	/*
		Units: metric (SI) vs. inch-pound (I-P)
		The SI (metric) unit of R-value is
		  kelvin square-metre per watt (K⋅m2/W or, equally, °C⋅m2/W),
		whereas the I-P (inch-pound) unit is
		  degree Fahrenheit square-foot hour per British thermal unit (°F⋅ft2⋅h/BTU).
		For R-values there is no difference between U.S. and Imperial units, so the same I-P unit is used in both.
		Some sources use "RSI" when referring to R-values in SI units.
		R-values expressed in I-P units are approximately 5.68 times as large as R-values expressed in SI units. For example, a window that is R-2 in the I-P system is about RSI 0.35, since 2/5.68 ≈ 0.35.
		In countries where the SI system is generally in use, the R-values will also normally be given in SI units. This includes the United Kingdom, Australia, and New Zealand.
		I-P values are commonly given in the United States and Canada, though in Canada normally both I-P and RSI values are listed.
		Because the units are usually not explicitly stated, one must decide from context which units are being used. In this regard, it helps to keep in mind that I-P R-values are 5.68 times larger than the corresponding SI R-values.
		More precisely,
		  R-value (in I-P) ≈ RSI-value (in SI) × 5.678263
		  RSI-value (in SI) ≈ R-value (in I-P) × 0.1761102
	*/

	// https://qudt.org/vocab/unit/DEG_F-HR-FT2-PER-BTU_TH
	// qudt:conversionMultiplier: 1.8969

	// https://www.nist.gov/pml/special-publication-811/nist-guide-si-appendix-b-conversion-factors/nist-guide-si-appendix-b9
	/*
		ToUnit convert...
			from: degree Fahrenheit hour square foot per British thermal unitIT(°F · h · ft2/BtuIT)
			to: square meter kelvin per watt (m2 · K/W)
			Multiply by: 1.761102 E-01

			from: degree Fahrenheit hour square foot per British thermal unitth(°F · h · ft2/Btuth)
			to: square meter kelvin per watt (m2 · K/W)
			Multiply by: 1.762280 E-01
	*/
}
