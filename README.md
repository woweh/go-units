[![godocs.io](https://godocs.io/github.com/bcicen/go-units?status.svg)](https://godocs.io/github.com/bcicen/go-units)

# go-units

Go library for manipulating and converting between various units of measurement

## Usage

In the most basic usage, `go-units` may be used to convert a value from one unit to another:

```go
package main

import (
	"fmt"
	u "github.com/bcicen/go-units"
)

func main() {
	// convert a simple float from celsius to fahrenheit
	fmt.Println(u.MustConvertFloat(25.5, u.Celsius, u.Fahrenheit)) // outputs "77.9 fahrenheit"

	// convert a units.Value instance
	val := u.NewValue(25.5, u.Celsius)

	fmt.Println(val)                           // "25.5 celsius"
	fmt.Println(val.MustConvert(u.Fahrenheit)) // "77.9 fahrenheit"
	fmt.Println(val.MustConvert(u.Kelvin))     // "298.65 kelvin"
}
```

### Formatting

Aside from unit conversions, `go-units` may also be used for generating human-readable unit labels, plural names, and symbols:

```go
val := u.NewValue(2.0, u.Nibble)
fmt.Println(val)                     // "2 nibbles"
fmt.Println(val.MustConvert(u.Byte)) // "1 byte"

// value formatting options may also be specified:
opts := u.FmtOptions{
  Label:     true, // append unit name/symbol
  Short:     true, // use unit symbol
  Precision: 3,
}

val = u.NewValue(15.456932, u.KiloMeter)
fmt.Println(val)           // "15.456932 kilometers"
fmt.Println(val.Fmt(opts)) // "15.457 km"
fmt.Println(val.Float())   // "15.456932"
```

### Lookup

The package-level `Find()` method may be used to search for a unit by name, symbol, or alternate spelling:
```go
// symbol
unit, err := u.Find("m")
// name
unit, err := u.Find("meter")
// alternate spelling
unit, err := u.Find("metre")
```

### Custom Units

`go-units` comes with many unit names and symbols builtin; however, new units and conversions can be easily added:

```go
// register custom unit names
Ding := u.NewUnit("ding", "di")
Dong := u.NewUnit("dong", "do")

// there are 100 dongs in a ding
u.NewRatioConversion(Ding, Dong, 100.0)

val := u.NewValue(25.0, Ding)

fmt.Printf("%s = %s\n", val, val.MustConvert(Dong)) // "25 dings = 2500 dongs"

// conversions are automatically registered when using magnitude prefix helper methods
KiloDong := u.Kilo(Dong)
fmt.Println(u.MustConvertFloat(1000.0, Dong, KiloDong)) // "1 kilodong"
```
You can also add aliases to existing units:
```go
SquareInch.AddAliases(
    "square inches", "square in", "square in.", "square ins", "square ins.", "in2", "in^2", "in**2",
    "sq in", "sq in.", "sq ins", "sq ins.", "sqin", "sqin.", "sqins", "□″", "sq inches", "sq inch", "inches/-2",
    "inch/-2", "in/-2", "inches2", "inch2", "\"2", "″2",
)
```

**NOTE:**   
Using unicode symbols (e.g., `㎡`, `㎢`) as aliases is not supported!


### References / Further Reading
Furey, Edward "Conversion Calculators" at https://www.calculatorsoup.com/calculators/conversions/ from CalculatorSoup, https://www.calculatorsoup.com - Online Calculators

The National Institute of Standards and Technology (NIST) - The NIST Guide for the use of the International System of Units -
Appendix B, subsections B.8 Factors for Units Listed Alphabetically and B.9 Factors for units listed by kind of quantity or field of science.
- https://physics.nist.gov/cuu/pdf/sp811.pdf
- https://www.nist.gov/pml/special-publication-811/nist-guide-si-appendix-b-conversion-factors/nist-guide-si-appendix-b8
- https://www.nist.gov/pml/special-publication-811/nist-guide-si-appendix-b-conversion-factors/nist-guide-si-appendix-b9

Wikipedia contributors. "Conversion of units" Wikipedia, The Free Encyclopedia. Wikipedia, The Free Encyclopedia.
- https://en.wikipedia.org/wiki/Conversion_of_units
