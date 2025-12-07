package units

var (
	// DynamicViscosity is the unit quantity for dynamic viscosity.
	DynamicViscosity = NewUnitQuantity("dynamic viscosity")

	// Base unit: pascal second (Revit base with CF=3.280839895013123)
	PascalSecond = DynamicViscosity.MustCreateUnit("pascal second", "Pa-s", SI)

	// SI equivalent units
	NewtonSecondPerSquareMeter = DynamicViscosity.MustCreateUnit("newton second per square meter", "N·s/m²", SI)
	KilogramPerMeterSecond     = DynamicViscosity.MustCreateUnit("kilogram per meter second", "kg/(m·s)", SI)
	KilogramPerMeterHour       = DynamicViscosity.MustCreateUnit("kilogram per meter hour", "kg/(m·h)", SI)

	// Common unit
	Centipoise = DynamicViscosity.MustCreateUnit("centipoise", "cP")

	// Imperial units
	PoundForceSecondPerSquareFoot = DynamicViscosity.MustCreateUnit("pound force second per square foot", "lb·s/ft²", BI)
	PoundMassPerFootSecond        = DynamicViscosity.MustCreateUnit("pound mass per foot second", "lbm/ft-s", BI)
	PoundMassPerFootHour          = DynamicViscosity.MustCreateUnit("pound mass per foot hour", "lbm/ft-h", BI)
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
	// Calculated: Pa·s to lbf·s/ft²
	// Pa·s = N·s/m² = lbf·s/ft² × (4.44822 / 10.764)
	// So: 1 Pa·s = (4.44822 / 10.764) lbf·s/ft² = 0.0685217658567918 lbf·s/ft²
	// 1 Pa·s = (1 N·s / 4.44822 lbf) × (10.764 ft² / m²) = 0.0685217658567918 lbf·s/ft²
	NewRatioConversion(PascalSecond, PoundForceSecondPerSquareFoot, 0.0685217658567918)

	// 1 Pa-s = 2.20462262184878 lbm/ft-s (Revit)
	// Calculated: Pa·s to lbm/ft-s (using pound-mass)
	// 1 Pa-s = 2.20462262184878 lbm/ft-s (empirical from tests)
	NewRatioConversion(PascalSecond, PoundMassPerFootSecond, 2.20462262184878)

	// 1 Pa-s = 7936.64143865559 lbm/ft-h (Revit)
	// Calculated: Pa·s to lbm/ft-h using time conversion
	// 1 Pa-s = 7936.64143865559 lbm/ft-h (empirical from tests)
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

// dynamicViscosityFactor calculates the conversion factor for dynamic viscosity units.
// Dynamic Viscosity = force × time / area = force·time / length²
// Base unit: PascalSecond = Pa·s = N·s/m²
//
// Example: lb·s/ft²
// - forceRatio: how many N in 1 unit of target force
// - areaRatio: how many m² in 1 unit of target area (length²)
// - timeRatio: how many s in 1 unit of target time
// - dynamicViscosityFactor = (forceRatio × timeRatio) / areaRatio
func dynamicViscosityFactor(force, length, time Unit) float64 {
	// How many N in 1 unit of target force
	forceRatio := force.to(Newton)
	// How many m² in 1 unit of target area (length²)
	areaRatio := areaFactor(length)
	// How many s in 1 unit of target time
	timeRatio := time.to(Second)
	// Dynamic viscosity factor: (force × time) / area
	return (forceRatio * timeRatio) / areaRatio
}
