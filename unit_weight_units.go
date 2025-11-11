package units

// UnitWeight is a unit quantity for unit weight
const UnitWeightQuantity UnitQuantity = "unit weight"

var (
	_unitWeight = Quantity(UnitWeightQuantity)

	// SI base unit: kilonewton per cubic meter
	KiloNewtonPerCubicMeter = mustCreateNewUnit("kilonewton per cubic meter", "kN/m³", _unitWeight, SI)

	// Imperial base unit: pound force per cubic foot
	PoundForcePerCubicFoot = mustCreateNewUnit("pound force per cubic foot", "lbf/ft³", _unitWeight, BI)

	// Imperial unit
	KipPerCubicInch = mustCreateNewUnit("kip per cubic inch", "kip/in³", _unitWeight, BI)
)

func initUnitWeightUnits() {
	// SI base unit: kN/m³
	// 1 kN/m³ = 6.36588 lbf/ft³
	NewRatioConversion(KiloNewtonPerCubicMeter, PoundForcePerCubicFoot, 6.36588)

	// 1 kip/in³ = 1728000 lbf/ft³
	NewRatioConversion(KipPerCubicInch, PoundForcePerCubicFoot, 1728000.0)

	KiloNewtonPerCubicMeter.AddAliases("kilonewtons per cubic meter", "kilonewton per cubic metre", "kilonewtons per cubic metre", "kN/m3")
	PoundForcePerCubicFoot.AddAliases("pounds force per cubic foot", "lb/ft³", "lbf/ft3", "pcf")
	KipPerCubicInch.AddAliases("kips per cubic inch", "kip/in3")
}
