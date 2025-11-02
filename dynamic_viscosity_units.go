package units

var (
	DynamicViscosity = Quantity("dynamic viscosity")

	// Base unit: pascal second (Revit base with CF=3.280839895013123)
	PascalSecond = mustCreateNewUnit("pascal second", "Pa-s", DynamicViscosity, SI)

	// SI equivalent units
	NewtonSecondPerSquareMeter = mustCreateNewUnit("newton second per square meter", "N·s/m²", DynamicViscosity, SI)
	KilogramPerMeterSecond     = mustCreateNewUnit("kilogram per meter second", "kg/(m·s)", DynamicViscosity, SI)
	KilogramPerMeterHour       = mustCreateNewUnit("kilogram per meter hour", "kg/(m·h)", DynamicViscosity, SI)

	// Common unit
	Centipoise = mustCreateNewUnit("centipoise", "cP", DynamicViscosity)

	// Imperial units
	PoundForceSecondPerSquareFoot = mustCreateNewUnit("pound force second per square foot", "lb·s/ft²", DynamicViscosity, BI)
	PoundMassPerFootSecond        = mustCreateNewUnit("pound mass per foot second", "lbm/ft-s", DynamicViscosity, BI)
	PoundMassPerFootHour          = mustCreateNewUnit("pound mass per foot hour", "lbm/ft-h", DynamicViscosity, BI)
)

func init() {
	// From RevitUnits.json (all relative to Pa-s):
	// Pa-s: CF = 3.280839895013123
	// N·s/m²: CF = 3.280839895013123 (same as Pa-s)
	// kg/(m·s): CF = 3.280839895013123 (same as Pa-s)
	// kg/(m·h): CF = 11811.023622047243
	// cP: CF = 3280.839895013123
	// lb·s/ft²: CF = 0.06852176585679176
	// lbm/ft-s: CF = 2.2046226218487757
	// lbm/ft-h: CF = 7936.6414386555925

	// Pa-s equivalents
	NewRatioConversion(PascalSecond, NewtonSecondPerSquareMeter, 1.0)
	NewRatioConversion(PascalSecond, KilogramPerMeterSecond, 1.0)

	// 1 kg/(m·h) = 3600 kg/(m·s)
	NewRatioConversion(KilogramPerMeterHour, KilogramPerMeterSecond, 3600.0)

	// 1 Pa-s = 1000 cP
	NewRatioConversion(PascalSecond, Centipoise, 1000.0)

	// From ratios: 3.280839895013123 / 0.06852176585679176 = 47.88
	NewRatioConversion(PascalSecond, PoundForceSecondPerSquareFoot, 47.88)

	// 1 lbm/ft-s = 1.4882 Pa-s (standard conversion)
	NewRatioConversion(PoundMassPerFootSecond, PascalSecond, 1.4882)

	// 1 lbm/ft-h = 3600 lbm/ft-s
	NewRatioConversion(PoundMassPerFootHour, PoundMassPerFootSecond, 3600.0)

	PascalSecond.AddAliases("pascal seconds", "Pa*s", "Pas")
	NewtonSecondPerSquareMeter.AddAliases("newton seconds per square meter", "newton second per square metre", "N*s/m²")
	KilogramPerMeterSecond.AddAliases("kilograms per meter second", "kilogram per metre second", "kg/m/s")
	KilogramPerMeterHour.AddAliases("kilograms per meter hour", "kilogram per metre hour", "kg/m/h")
	Centipoise.AddAliases("centipoises", "centipoise", "cp")
	PoundForceSecondPerSquareFoot.AddAliases("pound force seconds per square foot", "lbf*s/ft²")
	PoundMassPerFootSecond.AddAliases("pounds mass per foot second", "lbm/ft/s")
	PoundMassPerFootHour.AddAliases("pounds mass per foot hour", "lbm/ft/h")
}
