package units

var (
	ThermalMass = Quantity("thermal mass")

	// Base unit: joule per kelvin (Revit base with CF=0.09290304)
	JoulePerKelvin = newUnit("joule per kelvin", "J/K", ThermalMass, SI)

	// Other units
	KiloJoulePerKelvin                        = newUnit("kilojoule per kelvin", "kJ/K", ThermalMass, SI)
	BritishThermalUnitPerFahrenheit = newUnit("British thermal unit per degree Fahrenheit", "BTU/°F", ThermalMass)
)

func init() {
	// From RevitUnits.json:
	// J/K: CF = 0.09290304
	// kJ/K: CF = 9.290304e-05
	// BTU/°F: CF = 4.891949546730718e-05
	
	// 1 kJ/K = 1000 J/K
	NewRatioConversion(KiloJoulePerKelvin, JoulePerKelvin, 1000.0)
	
	// From the ratios: 0.09290304 / 4.891949546730718e-05 = 1899.1
	// This means: 1 J/K = 0.0005267 BTU/°F
	// Or: 1 BTU/°F = 1899.1 J/K
	NewRatioConversion(BritishThermalUnitPerFahrenheit, JoulePerKelvin, 1899.1)

	JoulePerKelvin.AddAliases("joules per kelvin", "J per K")
	KiloJoulePerKelvin.AddAliases("kilojoules per kelvin", "kJ per K")
	BritishThermalUnitPerFahrenheit.AddAliases("British thermal units per degree Fahrenheit", "BTU per F", "Btu/°F")
}
