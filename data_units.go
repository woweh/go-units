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

var (
	Bi   = Quantity("bits")
	Data = Quantity("bytes")

	Byte      = newUnit("byte", "B", Data, BaseUnit)
	KiloByte  = Kilo(Byte, Symbols("KB"))
	MegaByte  = Mega(Byte)
	GigaByte  = Giga(Byte)
	TeraByte  = Tera(Byte)
	PetaByte  = Peta(Byte)
	ExaByte   = Exa(Byte)
	ZettaByte = Zetta(Byte)
	YottaByte = Yotta(Byte)

	Kibibyte = newUnit("kibibyte", "KiB", Data, IEC)
	Mebibyte = newUnit("mebibyte", "MiB", Data, IEC)
	Gibibyte = newUnit("gibibyte", "GiB", Data, IEC)
	Tebibyte = newUnit("tebibyte", "TiB", Data, IEC)
	Pebibyte = newUnit("pebibyte", "PiB", Data, IEC)
	Exbibyte = newUnit("exbibyte", "EiB", Data, IEC)
	Zebibyte = newUnit("zebibyte", "ZiB", Data, IEC)
	Yobibyte = newUnit("yobibyte", "YiB", Data, IEC)

	Bit     = newUnit("bit", "b", Bi)
	KiloBit = Kilo(Bit)
	MegaBit = Mega(Bit)
	GigaBit = Giga(Bit)
	TeraBit = Tera(Bit)
	PetaBit = Peta(Bit)
	ExaBit  = Exa(Bit)

	Nibble = newUnit("nibble", "", Data)
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
