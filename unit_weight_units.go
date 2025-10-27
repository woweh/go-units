package units

var (
	UnitWeight = Quantity("unit weight")

	// Base unit: pound force per cubic foot (calculated from Revit data)
	PoundForcePerCubicFoot = newUnit("pound force per cubic foot", "lbf/ft³", UnitWeight, BI)

	// SI units
	KiloNewtonPerCubicMeter = newUnit("kilonewton per cubic meter", "kN/m³", UnitWeight, SI)

	// Imperial units
	KipPerCubicInch = newUnit("kip per cubic inch", "kip/in³", UnitWeight)
)

func init() {
	// From RevitUnits.json:
	// lbf/ft³: CF = 0.06852176585679176
	// kN/m³: CF = 0.010763910416709722
	// kip/in³: CF = 3.9653799685643397e-08
	
	// Using lbf/ft³ as base
	// From ratios: 0.06852176585679176 / 0.010763910416709722 = 6.36588
	// 1 lbf/ft³ = 0.157087 kN/m³
	// Or: 1 kN/m³ = 6.36588 lbf/ft³
	NewRatioConversion(KiloNewtonPerCubicMeter, PoundForcePerCubicFoot, 6.36588)
	
	// From ratios: 0.06852176585679176 / 3.9653799685643397e-08 = 1.728e6
	// 1 lbf/ft³ = 5.787e-7 kip/in³
	// Or: 1 kip/in³ = 1728000 lbf/ft³
	NewRatioConversion(KipPerCubicInch, PoundForcePerCubicFoot, 1728000.0)

	PoundForcePerCubicFoot.AddAliases("pounds force per cubic foot", "lb/ft³", "lbf/ft3", "pcf")
	KiloNewtonPerCubicMeter.AddAliases("kilonewtons per cubic meter", "kilonewton per cubic metre", "kilonewtons per cubic metre", "kN/m3")
	KipPerCubicInch.AddAliases("kips per cubic inch", "kip/in3")
}
