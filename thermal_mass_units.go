package units

var (
	// ThermalMass is the unit quantity for thermal mass.
	ThermalMass = NewUnitQuantity("thermal mass")

	// SI base unit: joule per kelvin
	JoulePerKelvin = ThermalMass.MustCreateUnit("joule per kelvin", "J/K", SI)

	// Other SI unit
	KiloJoulePerKelvin = ThermalMass.MustCreateUnit("kilojoule per kelvin", "kJ/K", SI)

	// Imperial/US unit
	BritishThermalUnitPerFahrenheit = ThermalMass.MustCreateUnit("British thermal unit per degree Fahrenheit", "BTU/°F", BI)
)

func initThermalMassUnits() {
	// 1 kJ/K = 1000 J/K
	NewRatioConversion(KiloJoulePerKelvin, JoulePerKelvin, 1000.0)

	// 1 BTU/°F = 1899.1 J/K
	NewRatioConversion(BritishThermalUnitPerFahrenheit, JoulePerKelvin, 1899.1)

	JoulePerKelvin.AddAliases("joules per kelvin", "J per K")
	KiloJoulePerKelvin.AddAliases("kilojoules per kelvin", "kJ per K")
	BritishThermalUnitPerFahrenheit.AddAliases("British thermal units per degree Fahrenheit", "BTU per F", "Btu/°F")
}
