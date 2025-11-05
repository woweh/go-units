package units

// Moment is a unit quantity for moment (torque)
const Moment UnitQuantity = "moment"

var (
	_moment = Quantity(Moment)

	// SI base unit: newton meter
	NewtonMeter = mustCreateNewUnit("newton meter", "N-m", _moment, SI)

	// SI units
	DekaNewtonMeter = mustCreateNewUnit("dekanewton meter", "daN-m", _moment, SI)
	KiloNewtonMeter = mustCreateNewUnit("kilonewton meter", "kN-m", _moment, SI)
	MegaNewtonMeter = mustCreateNewUnit("meganewton meter", "MN-m", _moment, SI)

	// Imperial units
	PoundForceFoot = mustCreateNewUnit("pound force foot", "lb-ft", _moment, BI)
	KipFoot        = mustCreateNewUnit("kip foot", "kip-ft", _moment, BI)

	// Other units
	KilogramForceMeter = mustCreateNewUnit("kilogram force meter", "kgf-m", _moment, MKpS)
	TonneForceMeter    = mustCreateNewUnit("tonne force meter", "Tf-m", _moment, MKpS)
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
