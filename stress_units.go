package units

// Note: Stress uses the same units as Pressure in physics
// Many stress units are already defined in pressure_units.go
// This file adds stress-specific units and aliases for structural engineering

var (
	Stress = Quantity("stress")

	// Imperial/US units specific to stress
	KipPerSquareFoot = newUnit("kip per square foot", "ksf", Stress, BI)
	KipPerSquareInch = newUnit("kip per square inch", "ksi", Stress, BI)

	// Other metric units specific to stress
	NewtonPerSquareMillimeter     = newUnit("newton per square millimeter", "N/mm²", Stress, SI)
	KiloNewtonPerSquareCentimeter = newUnit("kilonewton per square centimeter", "kN/cm²", Stress, SI)
	KiloNewtonPerSquareMillimeter = newUnit("kilonewton per square millimeter", "kN/mm²", Stress, SI)
	MegaNewtonPerSquareMeter      = newUnit("meganewton per square meter", "MN/m²", Stress, SI)
	DekaNewtonPerSquareMeter      = newUnit("dekanewton per square meter", "daN/m²", Stress, SI)
	KilogramForcePerSquareMeter   = newUnit("kilogram force per square meter", "kgf/m²", Stress)
	TonneForcePerSquareMeter      = newUnit("tonne force per square meter", "Tf/m²", Stress)
)

func init() {
	// Stress-specific conversions (using Pascal from pressure_units.go as base)
	
	// 1 ksf = 1000 psf = 47880.26 Pa
	NewRatioConversion(KipPerSquareFoot, Pascal, 47880.26)
	
	// 1 ksi = 1000 psi = 6894757.29 Pa
	NewRatioConversion(KipPerSquareInch, Pascal, 6894757.293168)
	
	// Other metric units
	NewRatioConversion(NewtonPerSquareMillimeter, Pascal, 1000000.0)
	NewRatioConversion(KiloNewtonPerSquareCentimeter, Pascal, 10000000.0)
	NewRatioConversion(KiloNewtonPerSquareMillimeter, Pascal, 1000000000.0)
	NewRatioConversion(MegaNewtonPerSquareMeter, Pascal, 1000000.0)
	NewRatioConversion(DekaNewtonPerSquareMeter, Pascal, 10.0)
	
	// 1 kgf/m² = 9.80665 Pa
	NewRatioConversion(KilogramForcePerSquareMeter, Pascal, 9.80665)
	
	// 1 Tf/m² = 9806.65 Pa
	NewRatioConversion(TonneForcePerSquareMeter, Pascal, 9806.65)

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
