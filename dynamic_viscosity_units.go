package units

const DynamicViscosity UnitQuantity = "dynamic viscosity"

var (
	_dynamicViscosity = Quantity(DynamicViscosity)

	// Base unit: pascal second (Revit base with CF=3.280839895013123)
	PascalSecond = mustCreateNewUnit("pascal second", "Pa-s", _dynamicViscosity, SI)

	// SI equivalent units
	NewtonSecondPerSquareMeter = mustCreateNewUnit("newton second per square meter", "N·s/m²", _dynamicViscosity, SI)
	KilogramPerMeterSecond     = mustCreateNewUnit("kilogram per meter second", "kg/(m·s)", _dynamicViscosity, SI)
	KilogramPerMeterHour       = mustCreateNewUnit("kilogram per meter hour", "kg/(m·h)", _dynamicViscosity, SI)

	// Common unit
	Centipoise = mustCreateNewUnit("centipoise", "cP", _dynamicViscosity)

	// Imperial units
	PoundForceSecondPerSquareFoot = mustCreateNewUnit("pound force second per square foot", "lb·s/ft²", _dynamicViscosity, BI)
	PoundMassPerFootSecond        = mustCreateNewUnit("pound mass per foot second", "lbm/ft-s", _dynamicViscosity, BI)
	PoundMassPerFootHour          = mustCreateNewUnit("pound mass per foot hour", "lbm/ft-h", _dynamicViscosity, BI)
)

func initDynamicViscosityUnits() {
	// Pa-s equivalents
	NewRatioConversion(PascalSecond, NewtonSecondPerSquareMeter, 1.0)
	NewRatioConversion(PascalSecond, KilogramPerMeterSecond, 1.0)

	// 1 kg/(m·h) = 3600 kg/(m·s)
	NewRatioConversion(KilogramPerMeterHour, KilogramPerMeterSecond, 3600.0)

	// 1 Pa-s = 1000 cP
	NewRatioConversion(PascalSecond, Centipoise, 1000.0)

	// From ratios: 3.280839895013123 / 0.06852176585679176 = 47.88
	NewRatioConversion(PascalSecond, PoundForceSecondPerSquareFoot, 0.0685217658567918)

	// 1 Pa-s = 2.20462262184878 lbm/ft-s (Revit)
	NewRatioConversion(PascalSecond, PoundMassPerFootSecond, 2.20462262184878)

	// 1 Pa-s = 7936.64143865559 lbm/ft-h (Revit)
	NewRatioConversion(PascalSecond, PoundMassPerFootHour, 7936.64143865559)

	PascalSecond.AddAliases("pascal seconds", "Pa*s", "Pas")
	NewtonSecondPerSquareMeter.AddAliases("newton seconds per square meter", "newton second per square metre", "N*s/m²")
	KilogramPerMeterSecond.AddAliases("kilograms per meter second", "kilogram per metre second", "kg/m/s")
	KilogramPerMeterHour.AddAliases("kilograms per meter hour", "kilogram per metre hour", "kg/m/h")
	Centipoise.AddAliases("centipoises", "centipoise", "cp")
	PoundForceSecondPerSquareFoot.AddAliases("pound force seconds per square foot", "lbf*s/ft²")
	PoundMassPerFootSecond.AddAliases("pounds mass per foot second", "lbm/ft/s")
	PoundMassPerFootHour.AddAliases("pounds mass per foot hour", "lbm/ft/h")
}
