package units

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
)

const Data UnitQuantity = "data"
const Bits UnitQuantity = "bits"

var (
	_data = Quantity(Data)
	_bits = Quantity(Bits)

	Byte      = mustCreateNewUnit("byte", "B", _data, SI)
	KiloByte  = Kilo(Byte, Symbols("KB"))
	MegaByte  = Mega(Byte)
	GigaByte  = Giga(Byte)
	TeraByte  = Tera(Byte)
	PetaByte  = Peta(Byte)
	ExaByte   = Exa(Byte)
	ZettaByte = Zetta(Byte)
	YottaByte = Yotta(Byte)

	Kibibyte = mustCreateNewUnit("kibibyte", "KiB", _data, IEC)
	Mebibyte = mustCreateNewUnit("mebibyte", "MiB", _data, IEC)
	Gibibyte = mustCreateNewUnit("gibibyte", "GiB", _data, IEC)
	Tebibyte = mustCreateNewUnit("tebibyte", "TiB", _data, IEC)
	Pebibyte = mustCreateNewUnit("pebibyte", "PiB", _data, IEC)
	Exbibyte = mustCreateNewUnit("exbibyte", "EiB", _data, IEC)
	Zebibyte = mustCreateNewUnit("zebibyte", "ZiB", _data, IEC)
	Yobibyte = mustCreateNewUnit("yobibyte", "YiB", _data, IEC)

	Bit     = mustCreateNewUnit("bit", "b", _bits, SI)
	KiloBit = Kilo(Bit)
	MegaBit = Mega(Bit)
	GigaBit = Giga(Bit)
	TeraBit = Tera(Bit)
	PetaBit = Peta(Bit)
	ExaBit  = Exa(Bit)

	Nibble = mustCreateNewUnit("nibble", "", _data)
)

func init() {
	NewRatioConversion(Nibble, Bit, 4.0)
	NewRatioConversion(Byte, Bit, 8.0)

	NewRatioConversion(Kibibyte, Byte, kb)
	NewRatioConversion(Mebibyte, Byte, mb)
	NewRatioConversion(Gibibyte, Byte, gb)
	NewRatioConversion(Tebibyte, Byte, tb)
	NewRatioConversion(Pebibyte, Byte, pb)
	NewRatioConversion(Exbibyte, Byte, eb)
	NewRatioConversion(Zebibyte, Byte, zb)
	NewRatioConversion(Yobibyte, Byte, yb)
}
