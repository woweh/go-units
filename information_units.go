package units

const Information UnitQuantity = "information"

var (
	_information = Quantity(Information)

	Byte       = mustCreateNewUnit("byte", "B", _information, SI)
	KiloByte   = Kilo(Byte, Symbols("KB"))
	MegaByte   = Mega(Byte)
	GigaByte   = Giga(Byte)
	TeraByte   = Tera(Byte)
	PetaByte   = Peta(Byte)
	ExaByte    = Exa(Byte)
	ZettaByte  = Zetta(Byte)
	YottaByte  = Yotta(Byte)
	RonnaByte  = Ronna(Byte)
	QuettaByte = Quetta(Byte)

	Kibibyte = mustCreateNewUnit("kibibyte", "KiB", _information, IEC)
	Mebibyte = mustCreateNewUnit("mebibyte", "MiB", _information, IEC)
	Gibibyte = mustCreateNewUnit("gibibyte", "GiB", _information, IEC)
	Tebibyte = mustCreateNewUnit("tebibyte", "TiB", _information, IEC)
	Pebibyte = mustCreateNewUnit("pebibyte", "PiB", _information, IEC)
	Exbibyte = mustCreateNewUnit("exbibyte", "EiB", _information, IEC)
	Zebibyte = mustCreateNewUnit("zebibyte", "ZiB", _information, IEC)
	Yobibyte = mustCreateNewUnit("yobibyte", "YiB", _information, IEC)
	Robibyte = mustCreateNewUnit("robi byte", "RiB", _information, IEC)
	Qubibyte = mustCreateNewUnit("qubi byte", "QiB", _information, IEC)

	Bit       = mustCreateNewUnit("bit", "b", _information, SI)
	KiloBit   = Kilo(Bit)
	MegaBit   = Mega(Bit)
	GigaBit   = Giga(Bit)
	TeraBit   = Tera(Bit)
	PetaBit   = Peta(Bit)
	ExaBit    = Exa(Bit)
	ZettaBit  = Zetta(Bit)
	YottaBit  = Yotta(Bit)
	RonnaBit  = Ronna(Bit)
	QuettaBit = Quetta(Bit)

	Nibble = mustCreateNewUnit("nibble", "", _information)
)

func initInformationUnits() {
	NewRatioConversion(Nibble, Bit, 4.0)
	NewRatioConversion(Byte, Bit, 8.0)

	const (
		// byte ratio constants
		_          = iota
		kb float64 = 1 << (10 * iota)
		mb
		gb
		tb
		pb
		eb
		zb
		yb
		rb
		qb
	)
	NewRatioConversion(Kibibyte, Byte, kb)
	NewRatioConversion(Mebibyte, Byte, mb)
	NewRatioConversion(Gibibyte, Byte, gb)
	NewRatioConversion(Tebibyte, Byte, tb)
	NewRatioConversion(Pebibyte, Byte, pb)
	NewRatioConversion(Exbibyte, Byte, eb)
	NewRatioConversion(Zebibyte, Byte, zb)
	NewRatioConversion(Yobibyte, Byte, yb)
	NewRatioConversion(Robibyte, Byte, rb)
	NewRatioConversion(Qubibyte, Byte, qb)
}
