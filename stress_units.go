package units

// Note: Stress uses the same units as Pressure in physics
// Many stress units are already defined in pressure_units.go
// This file adds stress-specific units and aliases for structural engineering

// Stress is a unit quantity for stress
var (
	// Stress is the unit quantity for stress.
	Stress = AliasOf(Pressure)

	// Imperial/US units specific to stress
	KipPerSquareFoot = Stress.MustCreateUnit("kip per square foot", "ksf", BI)
	KipPerSquareInch = Stress.MustCreateUnit("kip per square inch", "ksi", BI)

	// Other metric units specific to stress
	NewtonPerSquareMillimeter     = Stress.MustCreateUnit("newton per square millimeter", "N/mm²", SI)
	KiloNewtonPerSquareCentimeter = Stress.MustCreateUnit("kilonewton per square centimeter", "kN/cm²", SI)
	KiloNewtonPerSquareMillimeter = Stress.MustCreateUnit("kilonewton per square millimeter", "kN/mm²", SI)
	MegaNewtonPerSquareMeter      = Stress.MustCreateUnit("meganewton per square meter", "MN/m²", SI)
	DekaNewtonPerSquareMeter      = Stress.MustCreateUnit("dekanewton per square meter", "daN/m²", SI)
	KilogramForcePerSquareMeter   = Stress.MustCreateUnit("kilogram force per square meter", "kgf/m²", MKpS)
	TonneForcePerSquareMeter      = Stress.MustCreateUnit("tonne force per square meter", "Tf/m²", MKpS)
)

func initStressUnits() {
	// Stress-specific conversions (using Pascal from pressure_units.go as base)
	// Stress = force/area, so we can calculate using pressureFactor

	// Imperial/US stress units
	// 1 ksf = 1 kip/ft² (kip = 1000 lbf)
	NewRatioConversion(KipPerSquareFoot, Pascal, pressureFactor(Kip, Foot))
	// 1 ksi = 1 kip/in²
	NewRatioConversion(KipPerSquareInch, Pascal, pressureFactor(Kip, Inch))

	// Metric stress units (calculated from force/area)
	NewRatioConversion(NewtonPerSquareMillimeter, Pascal, pressureFactor(Newton, MilliMeter))
	NewRatioConversion(KiloNewtonPerSquareCentimeter, Pascal, pressureFactor(KiloNewton, CentiMeter))
	NewRatioConversion(KiloNewtonPerSquareMillimeter, Pascal, pressureFactor(KiloNewton, MilliMeter))
	NewRatioConversion(MegaNewtonPerSquareMeter, Pascal, pressureFactor(MegaNewton, Meter))
	NewRatioConversion(DekaNewtonPerSquareMeter, Pascal, pressureFactor(DecaNewton, Meter))

	// Force-based stress units (using standard force definitions)
	// 1 kgf/m² = 9.80665 Pa
	NewRatioConversion(KilogramForcePerSquareMeter, Pascal, pressureFactor(KilogramForce, Meter))
	// 1 Tf/m² = 9806.65 Pa
	NewRatioConversion(TonneForcePerSquareMeter, Pascal, pressureFactor(TonneForce, Meter))

	// Aliases
	KipPerSquareFoot.AddAliases("kips per square foot")
	KipPerSquareInch.AddAliases("kips per square inch")
	NewtonPerSquareMillimeter.AddAliases("newtons per square millimeter", "newton per square millimetre", "newtons per square millimetre")
	KiloNewtonPerSquareCentimeter.AddAliases("kilonewtons per square centimeter", "kilonewton per square centimetre", "kilonewtons per square centimetre")
	KiloNewtonPerSquareMillimeter.AddAliases("kilonewtons per square millimeter", "kilonewton per square millimetre", "kilonewtons per square millimetre")
	MegaNewtonPerSquareMeter.AddAliases("meganewtons per square meter", "meganewton per square metre", "meganewtons per square metre")
	DekaNewtonPerSquareMeter.AddAliases("dekanewtons per square meter", "dekanewton per square metre", "dekanewtons per square metre")
	KilogramForcePerSquareMeter.AddAliases("kilograms force per square meter", "kilogram force per square metre", "kilograms force per square metre")
	TonneForcePerSquareMeter.AddAliases("tonnes force per square meter", "tonne force per square metre", "tonnes force per square metre")
}
