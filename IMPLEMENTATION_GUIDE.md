[NOTE] Authoritative tracking of scope and progress is in plan.md and STATUS.md. This guide provides code patterns only.

# Quick Reference: Adding New Revit Units to go-units

## Template for New Quantity File

```go
package units

var (
	MyQuantity = Quantity("my quantity")
	
	// SI base units
	BaseUnit = newUnit("base unit", "base", MyQuantity, SI)
	KiloBaseUnit = Kilo(BaseUnit)
	MilliBaseUnit = Milli(BaseUnit)
	
	// Imperial/US units
	AlternateUnit = newUnit("alternate unit", "alt", MyQuantity, BI)
	AnotherUnit = newUnit("another unit", "ano", MyQuantity, BI)
)

func init() {
	// Define conversion relationships
	NewRatioConversion(BaseUnit, AlternateUnit, factor)
	
	// Add aliases for Revit compatibility
	BaseUnit.AddAliases("alias name 1", "alias name 2")
	BaseUnit.AddSymbols("alt_symbol")
}
```

## Revit ConversionFactor Usage

From RevitUnits.json, each unit has a ConversionFactor. This is what to use for NewRatioConversion:

**Example from Revit JSON**:
```json
{
  "DisplayName": "Watts per meter",
  "UnitSymbol": "W/m",
  "ConversionFactor": 0.3048
}
```

This means: 1 (base unit) = 0.3048 * (this unit)
So: `NewRatioConversion(WattPerMeter, BaseUnit, 0.3048)`

## High-Priority Implementation Order

### 1. Lighting Units (lighting_units.go)
```go
var (
	Lumen = newUnit("lumen", "lm", LuminousFlux, SI)
	Candela = newUnit("candela", "cd", LuminousIntensity, SI)
	Lux = newUnit("lux", "lx", Illuminance, SI)
	Footcandle = newUnit("footcandle", "ftc", Illuminance, BI)
	LumenPerWatt = newUnit("lumen per watt", "lm/W", Efficacy, SI)
	Footlambert = newUnit("footlambert", "ftL", Luminance, BI)
	CandelaPerSquareMeter = newUnit("candela per square meter", "cd/m²", Luminance, SI)
	CandelaPerSquareFoot = newUnit("candela per square foot", "cd/ft²", Luminance, BI)
)

func init() {
	NewRatioConversion(Lux, Footcandle, 10.763910416709722)
	NewRatioConversion(LumenPerWatt, LumenPerWatt, 1.0)  // Base
	NewRatioConversion(CandelaPerSquareMeter, CandelaPerSquareFoot, 10.763910416709722)
	NewRatioConversion(Footlambert, CandelaPerSquareMeter, 3.1415926535897896)
}
```

### 2. Velocity Units (velocity_units.go)
```go
var (
	MeterPerSecond = newUnit("meter per second", "m/s", Velocity, SI)
	FootPerSecond = newUnit("foot per second", "ft/s", Velocity, BI)
	FootPerMinute = newUnit("foot per minute", "ft/min", Velocity, BI)
	MilePerHour = newUnit("mile per hour", "mph", Velocity, BI)
	KilometerPerHour = newUnit("kilometer per hour", "km/h", Velocity, SI)
	CentimeterPerMinute = newUnit("centimeter per minute", "cm/min", Velocity, SI)
)
```

### 3. Acceleration Units (acceleration_units.go)
```go
var (
	MeterPerSecondSquared = newUnit("meter per second squared", "m/s²", Acceleration, SI)
	FootPerSecondSquared = newUnit("foot per second squared", "ft/s²", Acceleration, BI)
	InchPerSecondSquared = newUnit("inch per second squared", "in/s²", Acceleration, BI)
	KilometerPerSecondSquared = newUnit("kilometer per second squared", "km/s²", Acceleration, SI)
	MilePerSecondSquared = newUnit("mile per second squared", "mi/s²", Acceleration, BI)
)
```

### 4. Angular Speed Units (angular_speed_units.go)
```go
var (
	RevolutionPerMinute = newUnit("revolution per minute", "RPM", AngularSpeed)
	RevolutionPerSecond = newUnit("revolution per second", "RPS", AngularSpeed)
)

func init() {
	NewRatioConversion(RevolutionPerMinute, RevolutionPerSecond, 60)
	// Connect to existing frequency units if compatible
}
```

### 5. Load Units (cooling_load_units.go / heating_load_units.go)
```go
var (
	CoolingLoad = Quantity("cooling load")
	
	Kilowatt = newUnit("kilowatt", "kW", CoolingLoad, SI)
	Watt = newUnit("watt", "W", CoolingLoad, SI)
	BritishThermalUnitPerHour = newUnit("British thermal unit per hour", "Btu/h", CoolingLoad, BI)
	BritishThermalUnitPerSecond = newUnit("British thermal unit per second", "Btu/s", CoolingLoad, BI)
	TonOfRefrigeration = newUnit("ton of refrigeration", "ton", CoolingLoad, BI)
)

func init() {
	NewRatioConversion(Kilowatt, Watt, 1000)
	NewRatioConversion(BritishThermalUnitPerHour, Watt, 0.2930710701722222)
	NewRatioConversion(TonOfRefrigeration, Watt, 3516.853)
}
```

### 6. Moment of Inertia (moment_of_inertia_units.go)
```go
var (
	MomentOfInertia = Quantity("moment of inertia")
	
	MeterToFourth = newUnit("meter to the fourth power", "m⁴", MomentOfInertia, SI)
	FootToFourth = newUnit("foot to the fourth power", "ft⁴", MomentOfInertia, BI)
	InchToFourth = newUnit("inch to the fourth power", "in⁴", MomentOfInertia, BI)
	CentimeterToFourth = newUnit("centimeter to the fourth power", "cm⁴", MomentOfInertia, SI)
	MillimeterToFourth = newUnit("millimeter to the fourth power", "mm⁴", MomentOfInertia, SI)
)

func init() {
	// 1 m⁴ = 863097.48412416 cm⁴
	NewRatioConversion(MeterToFourth, CentimeterToFourth, 863097.48412416)
	// 1 ft⁴ = 20736 in⁴
	NewRatioConversion(FootToFourth, InchToFourth, 20736)
	// Connect SI to BI: 1 m⁴ = 8.630974841241602 ft⁴
	NewRatioConversion(MeterToFourth, FootToFourth, 8.630974841241602)
}
```

## Validation Against RevitUnits.json

For each new quantity, validate by:

```bash
# Check JSON for all units in this quantity
jq '.Quantities[] | select(.DisplayName == "Acceleration") | .Units[]' RevitUnits.json

# Verify ConversionFactors match your NewRatioConversion calls
# Note: ConversionFactor = 1 / (your ratio to base unit)
```

## Common Conversion Factors to Remember

- 1 meter = 3.280839895 feet
- 1 m/s = 3.28084 ft/s
- 1 m² = 10.763910 ft²
- 1 m⁴ = 8.63097 ft⁴
- 1 Pa = 0.000145 psi
- 1 W = 0.000293071 Btu/h
- 1 N·m = 0.737562 lb·ft

## Testing New Quantity

```go
// In corresponding _test.go file
func TestNewQuantity(t *testing.T) {
	v := NewValue(100, SomeUnit)
	result, _ := v.ConvertTo(AnotherUnit)
	expected := expectedValue
	if math.Abs(result - expected) > tolerance {
		t.Errorf("Conversion failed")
	}
}
```

## Checklist for New Quantity

- [ ] Create new file: `quantity_units.go`
- [ ] Define Quantity variable
- [ ] Add all base units with correct symbols
- [ ] Add SI prefixes where applicable
- [ ] Define conversions using ConversionFactor from Revit JSON
- [ ] Add aliases for Revit names
- [ ] Create `quantity_units_test.go` with tests
- [ ] Run `go test ./...` - should pass
- [ ] Verify no duplicate symbols
- [ ] Check conversions match RevitUnits.json ± 0.1%
- [ ] Update plan.md with completion

