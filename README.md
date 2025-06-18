[![Go Reference](https://pkg.go.dev/badge/github.com/Necoro/go-units.svg)](https://pkg.go.dev/github.com/Necoro/go-units)

# go-units

Go library for manipulating and converting between various units of measurement.

This is a fork of [woweh's go-units](https://github.com/woweh/go-units), which in turn is a major fork of [bcicen's go-units](https://github.com/bcicen/go-units).

## Principles
- Name and Aliases are case-***in***sensitive  
- Symbols are case-*sensitive* (e.g. MegaMeter Mm <> MilliMeter mm)
- Names, symbols and aliases must be unique! => Looking up a unit must return 1 exact match!

## Usage

In the most basic usage, `go-units` may be used to convert a value from one unit to another:

```go
package main

import (
	"fmt"
	u "github.com/Necoro/go-units"
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

You can also add aliases or symbols to existing units:
```go
// add alternative spellings and translations
CubicMeter.AddAliases("cubic metre", "Quadratmeter", "Vierkante meter", "Metr kwadratowy", "Mètre carré")
// add additional symbols
CubicMeter.AddSymbols("m3", "m^3", "m**3", "cum", "cbm", "CBM", "M3")
```

### Notes regarding unit and conversion definitions 
When you specify a ***function conversion*** (`NewConversionFromFn`), you must also specify the *inverse* conversion.

Specifying a ***ratio conversion*** (`NewRatioConversion`) also registers the inverse conversions.  
For ratio conversions you don't need to specify the inverse conversions.

The metric factory methods, like `u.Kilo`, will create the new unit with metric prefixes for name and symbols, and will register ratio conversions between the base unit and the derived metric unit.  
There is no need to separately specify the ratio conversion.
