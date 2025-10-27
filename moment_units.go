package units

var (
	Moment = Quantity("moment")

	// Base unit: newton meter (Revit base with CF=0.09290304)
	NewtonMeter = newUnit("newton meter", "N-m", Moment, SI)

	// SI units
	DekaNewtonMeter = newUnit("dekanewton meter", "daN-m", Moment, SI)
	KiloNewtonMeter = newUnit("kilonewton meter", "kN-m", Moment, SI)
	MegaNewtonMeter = newUnit("meganewton meter", "MN-m", Moment, SI)

	// Imperial units
	PoundForceFoot = newUnit("pound force foot", "lb-ft", Moment, BI)
	KipFoot        = newUnit("kip foot", "kip-ft", Moment, BI)

	// Other units
	KilogramForceMeter = newUnit("kilogram force meter", "kgf-m", Moment)
	TonneForceMeter    = newUnit("tonne force meter", "Tf-m", Moment)
)

func init() {
	// From RevitUnits.json:
	// N-m: CF = 0.09290304
	// daN-m: CF = 0.009290304
	// kN-m: CF = 9.290304e-05
	// MN-m: CF = 9.290304e-08
	// lb-ft: CF = 0.06852176585679176
	// kip-ft: CF = 6.852176585679177e-05
	// kgf-m: CF = 0.0094734736122937
	// Tf-m: CF = 9.4734736122937e-06
	
	NewRatioConversion(NewtonMeter, DekaNewtonMeter, 0.1)
	NewRatioConversion(NewtonMeter, KiloNewtonMeter, 0.001)
	NewRatioConversion(NewtonMeter, MegaNewtonMeter, 0.000001)
	
	// Standard conversion: 1 N-m = 0.737562 lb-ft
	NewRatioConversion(NewtonMeter, PoundForceFoot, 0.737562)
	
	// 1 kip-ft = 1000 lb-ft
	NewRatioConversion(KipFoot, PoundForceFoot, 1000.0)
	
	// 1 kgf-m = 9.80665 N-m
	NewRatioConversion(KilogramForceMeter, NewtonMeter, 9.80665)
	
	// 1 Tf-m = 9806.65 N-m
	NewRatioConversion(TonneForceMeter, NewtonMeter, 9806.65)

	NewtonMeter.AddAliases("newton meters", "N*m", "Nm")
	DekaNewtonMeter.AddAliases("dekanewton meters")
	KiloNewtonMeter.AddAliases("kilonewton meters")
	MegaNewtonMeter.AddAliases("meganewton meters")
	PoundForceFoot.AddAliases("pound force feet", "pound-foot", "pound-feet", "lbf-ft")
	KipFoot.AddAliases("kip feet", "kip-feet")
	KilogramForceMeter.AddAliases("kilogram force meters")
	TonneForceMeter.AddAliases("tonne force meters")
}
