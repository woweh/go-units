package units

// ThermalMass is a unit quantity for thermal mass
const ThermalMass UnitQuantity = "thermal mass"

var (
	_thermalMass = Quantity(ThermalMass)

	// SI base unit: joule per kelvin
	JoulePerKelvin = mustCreateNewUnit("joule per kelvin", "J/K", _thermalMass, SI)

	// Other SI unit
	KiloJoulePerKelvin = mustCreateNewUnit("kilojoule per kelvin", "kJ/K", _thermalMass, SI)

	// Imperial/US unit
	BritishThermalUnitPerFahrenheit = mustCreateNewUnit("British thermal unit per degree Fahrenheit", "BTU/°F", _thermalMass, BI)
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
