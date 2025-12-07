package units

var (
	// UnitWeight is the unit quantity for unit weight.
	UnitWeight = NewUnitQuantity("unit weight")

	// SI base unit: kilonewton per cubic meter
	KiloNewtonPerCubicMeter = UnitWeight.MustCreateUnit("kilonewton per cubic meter", "kN/m³", SI)

	// Imperial base unit: pound force per cubic foot
	PoundForcePerCubicFoot = UnitWeight.MustCreateUnit("pound force per cubic foot", "lbf/ft³", BI)

	// Imperial unit
	KipPerCubicInch = UnitWeight.MustCreateUnit("kip per cubic inch", "kip/in³", BI)
)

func initUnitWeightUnits() {
	// SI base unit: kN/m³
	// Calculated: 1 kN/m³ = X lbf/ft³
	NewRatioConversion(KiloNewtonPerCubicMeter, PoundForcePerCubicFoot, 1.0/unitWeightFactor(PoundForce, Foot))

	// 1 kip/in³ = 1728000 lbf/ft³
	// Calculated: kip/in³ to lbf/ft³
	NewRatioConversion(KipPerCubicInch, PoundForcePerCubicFoot, unitWeightFactor(Kip, Inch)/unitWeightFactor(PoundForce, Foot))

	KiloNewtonPerCubicMeter.AddAliases("kilonewtons per cubic meter", "kilonewton per cubic metre", "kilonewtons per cubic metre", "kN/m3")
	PoundForcePerCubicFoot.AddAliases("pounds force per cubic foot", "lb/ft³", "lbf/ft3", "pcf")
	KipPerCubicInch.AddAliases("kips per cubic inch", "kip/in3")
}

// unitWeightFactor calculates the conversion factor for unit weight units.
// Unit Weight = force / volume = force / length³
// Base unit: KiloNewtonPerCubicMeter = kN/m³
//
// Example: lbf/ft³
// - forceRatio: how many kN in 1 lbf
// - volumeRatio: how many m³ in 1 ft³
// - unitWeightFactor = forceRatio / volumeRatio
func unitWeightFactor(force, length Unit) float64 {
	// How many kN in 1 unit of the target force
	forceRatio := force.to(KiloNewton)
	// How many m³ in 1 unit of target volume (length³)
	volumeRatio := volumeFactor(length)
	// Unit weight factor: force per volume
	return forceRatio / volumeRatio
}
