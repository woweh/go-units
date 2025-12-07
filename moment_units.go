package units

var (
	// Moment is the unit quantity for moment (torque).
	Moment = NewUnitQuantity("moment")

	// SI base unit: newton meter
	NewtonMeter = Moment.MustCreateUnit("newton meter", "N-m", SI)

	// SI units
	DekaNewtonMeter = Moment.MustCreateUnit("dekanewton meter", "daN-m", SI)
	KiloNewtonMeter = Moment.MustCreateUnit("kilonewton meter", "kN-m", SI)
	MegaNewtonMeter = Moment.MustCreateUnit("meganewton meter", "MN-m", SI)

	// Imperial units
	PoundForceFoot = Moment.MustCreateUnit("pound force foot", "lb-ft", BI)
	KipFoot        = Moment.MustCreateUnit("kip foot", "kip-ft", BI)

	// Other units
	KilogramForceMeter = Moment.MustCreateUnit("kilogram force meter", "kgf-m", MKpS)
	TonneForceMeter    = Moment.MustCreateUnit("tonne force meter", "Tf-m", MKpS)
)

func initMomentUnits() {
	NewRatioConversion(NewtonMeter, DekaNewtonMeter, 0.1)
	NewRatioConversion(NewtonMeter, KiloNewtonMeter, 0.001)
	NewRatioConversion(NewtonMeter, MegaNewtonMeter, 0.000001)

	// Calculated: 1 N-m = X lb-ft
	NewRatioConversion(NewtonMeter, PoundForceFoot, 1.0/momentFactor(PoundForce, Foot))

	// 1 kip-ft = 1000 lb-ft
	// Calculated: kip-ft to lb-ft
	NewRatioConversion(KipFoot, PoundForceFoot, momentFactor(Kip, Foot)/momentFactor(PoundForce, Foot))

	// 1 kgf-m = 9.80665 N-m
	// Calculated: kgf·m to N·m
	NewRatioConversion(KilogramForceMeter, NewtonMeter, momentFactor(KilogramForce, Meter))

	// 1 Tf-m = 9806.65 N-m
	// Calculated: Tf·m to N·m
	NewRatioConversion(TonneForceMeter, NewtonMeter, momentFactor(TonneForce, Meter))

	NewtonMeter.AddAliases("newton meters", "N*m", "Nm")
	DekaNewtonMeter.AddAliases("dekanewton meters")
	KiloNewtonMeter.AddAliases("kilonewton meters")
	MegaNewtonMeter.AddAliases("meganewton meters")
	PoundForceFoot.AddAliases("pound force feet", "pound-foot", "pound-feet", "lbf-ft")
	KipFoot.AddAliases("kip feet", "kip-feet")
	KilogramForceMeter.AddAliases("kilogram force meters")
	TonneForceMeter.AddAliases("tonne force meters")
}

// momentFactor calculates the conversion factor for moment (torque) units.
// Moment = force × distance
// Base unit: NewtonMeter = N·m
//
// Example: lbf·ft
// - forceRatio: how many N in 1 lbf
// - lengthRatio: how many m in 1 ft
// - momentFactor = forceRatio × lengthRatio
func momentFactor(force, length Unit) float64 {
	// How many Newtons in 1 unit of the target force
	forceRatio := force.to(Newton)
	// How many meters in 1 unit of target length
	lengthRatio := length.to(Meter)
	// Moment factor: force × length
	return forceRatio * lengthRatio
}
