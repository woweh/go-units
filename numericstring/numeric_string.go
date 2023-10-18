// Package numericstring provides constants with numeric strings for conversion tests
package numericstring

// There are multiple strings for each number, allowing you to use the string
// that is most appropriate and descriptive in the context of the test.
//
// Abbreviations:
// - Txp = Ten to the power of / Ten Exponent
// - Tmn = Ten to the power of negative / Ten Exponent Negative

const (
	// One = 10 ^ 0, the base value
	One = "1"
	// Deca = 10 ^ 1, inverse of Deci
	Deca = "10"
	// Hecto = 10 ^ 2, inverse of Centi
	Hecto = "100"
	// Kilo = 10 ^ 3, inverse of Milli
	Kilo = "1000"
	// Mega = 10 ^ 6, inverse of Micro
	Mega = "1000000"
	// Giga = 10 ^ 9, inverse of Nano
	Giga = "1000000000"
	// Tera = 10 ^ 12, inverse of Pico
	Tera = "1000000000000"
	// Peta = 10 ^ 15, inverse of Femto
	Peta = "1000000000000000"
	// Exa = 10 ^ 18, inverse of Atto
	Exa = "1000000000000000000"
	// Zetta = 10 ^ 21, inverse of Zepto
	Zetta = "1000000000000000000000"
	// Yotta = 10 ^ 24, inverse of Yocto
	Yotta = "1000000000000000000000000"
	// Ronna = 10 ^ 27, inverse of Ronto
	Ronna = "1000000000000000000000000000"
	// Quetta = 10 ^ 30, inverse of Quecto
	Quetta = "1000000000000000000000000000000"
	// Deci = 10 ^ -1, inverse of Deca
	Deci = "0.1"
	// Centi = 10 ^ -2, inverse of Hecto
	Centi = "0.01"
	// Milli = 10 ^ -3, inverse of Kilo
	Milli = "0.001"
	// Micro = 10 ^ -6, inverse of Mega
	Micro = "0.000001"
	// Nano = 10 ^ -9, inverse of Giga
	Nano = "0.000000001"
	// Pico = 10 ^ -12, inverse of Tera
	Pico = "0.000000000001"
	// Femto = 10 ^ -15, inverse of Peta
	Femto = "0.000000000000001"
	// Atto = 10 ^ -18, inverse of Exa
	Atto = "0.000000000000000001"
	// Zepto = 10 ^ -21, inverse of Zetta
	Zepto = "0.000000000000000000001"
	// Yocto = 10 ^ -24, inverse of Yotta
	Yocto = "0.000000000000000000000001"
	// Ronto = 10 ^ -27, inverse of Ronna
	Ronto = "0.000000000000000000000000001"
	// Quecto = 10 ^ -30, inverse of Quetta
	Quecto = "0.000000000000000000000000000001"
	// Txp1 = Ten = Deca = 10 ^ 1, inverse of Deci
	Txp1 = "10"
	// Txp2 = Hundred = Hecto = 10 ^ 2, inverse of Centi
	Txp2 = "100"
	// Txp3 = Thousand = Kilo = 10 ^ 3, inverse of Milli
	Txp3 = "1000"
	// Txp6 = Million = Mega = 10 ^ 6, inverse of Micro
	Txp6 = "1000000"
	// Txp9 = Billion = Giga = 10 ^ 9, inverse of Nano
	Txp9 = "1000000000"
	// Txp12 = Trillion = Tera = 10 ^ 12, inverse of Pico
	Txp12 = "1000000000000"
	// Txp15 = Quadrillion = Peta = 10 ^ 15, inverse of Femto
	Txp15 = "1000000000000000"
	// Txp18 = Quintillion = Exa = 10 ^ 18, inverse of Atto
	Txp18 = "1000000000000000000"
	// Txp21 = Sextillion = Zetta = 10 ^ 21, inverse of Zepto
	Txp21 = "1000000000000000000000"
	// Txp24 = Septillion = Yotta = 10 ^ 24, inverse of Yocto
	Txp24 = "1000000000000000000000000"
	// Tmn1 = Tenth = Deci = 10 ^ -1, inverse of Deca
	Tmn1 = "0.1"
	// Tmn2 = Hundredth = Centi = 10 ^ -2, inverse of Hecto
	Tmn2 = "0.01"
	// Tmn3 = Thousandth = Milli = 10 ^ -3, inverse of Kilo
	Tmn3 = "0.001"
	// Tmn6 = Millionth = Micro = 10 ^ -6, inverse of Mega
	Tmn6 = "0.000001"
	// Tmn9 = Billionth = Nano = 10 ^ -9, inverse of Giga
	Tmn9 = "0.000000001"
	// Tmn12 = Trillionth = Pico = 10 ^ -12, inverse of Tera
	Tmn12 = "0.000000000001"
	// Tmn15 = Quadrillionth = Femto = 10 ^ -15, inverse of Peta
	Tmn15 = "0.000000000000001"
	// Tmn18 = Quintillionth = Atto = 10 ^ -18, inverse of Exa
	Tmn18 = "0.000000000000000001"
	// Tmn21 = Sextillionth = Zepto = 10 ^ -21, inverse of Zetta
	Tmn21 = "0.000000000000000000001"
	// Tmn24 = Septillionth = Yocto = 10 ^ -24, inverse of Yotta
	Tmn24 = "0.000000000000000000000001"
	// Tmn27 = Octillionth = Ronto = 10 ^ -27, inverse of Ronna
	Tmn27 = "0.000000000000000000000000001"
	// Ten = 10 ^ 1 = Deca, inverse of Deci
	Ten = "10"
	// Hundred = 10 ^ 2 = Hecto, inverse of Centi
	Hundred = "100"
	// Thousand = 10 ^ 3 = Kilo, inverse of Milli
	Thousand = "1000"
	// TenThousand = 10 ^ 4
	TenThousand = "10000"
	// HundredThousand = 10 ^ 5
	HundredThousand = "100000"
	// Million = 10 ^ 6 = Mega, inverse of Micro
	Million = "1000000"
	// Billion = 10 ^ 9 = Giga, inverse of Nano
	Billion = "1000000000"
	// Trillion = 10 ^ 12 = Tera, inverse of Pico
	Trillion = "1000000000000"
	// Quadrillion = 10 ^ 15 = Peta, inverse of Femto
	Quadrillion = "1000000000000000"
	// Quintillion = 10 ^ 18 = Exa, inverse of Atto
	Quintillion = "1000000000000000000"
	// Sextillion = 10 ^ 21 = Zetta, inverse of Zepto
	Sextillion = "1000000000000000000000"
	// Septillion = 10 ^ 24 = Yotta, inverse of Yocto
	Septillion = "1000000000000000000000000"
	// Octillion = 10 ^ 27 = Ronna, inverse of Ronto
	Octillion = "1000000000000000000000000000"
	// Tenth = 10 ^ -1 = Deci, inverse of Deca
	Tenth = "0.1"
	// Hundredth = 10 ^ -2 = Centi, inverse of Hecto
	Hundredth = "0.01"
	// Thousandth = 10 ^ -3 = Milli, inverse of Kilo
	Thousandth = "0.001"
	// Millionth = 10 ^ -6 = Micro, inverse of Mega
	Millionth = "0.000001"
	// Billionth = 10 ^ -9 = Nano, inverse of Giga
	Billionth = "0.000000001"
	// Trillionth = 10 ^ -12 = Pico, inverse of Tera
	Trillionth = "0.000000000001"
	// Quadrillionth = 10 ^ -15 = Femto, inverse of Peta
	Quadrillionth = "0.000000000000001"
	// Quintillionth = 10 ^ -18 = Atto, inverse of Exa
	Quintillionth = "0.000000000000000001"
	// Sextillionth = 10 ^ -21 = Zepto, inverse of Zetta
	Sextillionth = "0.000000000000000000001"
	// Septillionth = 10 ^ -24 = Yocto, inverse of Yotta
	Septillionth = "0.000000000000000000000001"
	// Octillionth = 10 ^ -27 = Ronto, inverse of Ronna
	Octillionth = "0.000000000000000000000000001"
)
