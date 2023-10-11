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
	Bi   = UnitOptionQuantity("bits")
	Data = UnitOptionQuantity("bytes")

	Byte      = newUnit("byte", "B", Data)
	KiloByte  = newUnit("kilobyte", "KB", Data)
	MegaByte  = newUnit("megabyte", "MB", Data)
	GigaByte  = newUnit("gigabyte", "GB", Data)
	TeraByte  = newUnit("terabyte", "TB", Data)
	PetaByte  = newUnit("petabyte", "PB", Data)
	ExaByte   = newUnit("exabyte", "", Data)
	ZettaByte = newUnit("zettabyte", "", Data)
	YottaByte = newUnit("yottabyte", "", Data)

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

	NewRatioConversion(KiloByte, Byte, 1000)
	NewRatioConversion(MegaByte, Byte, 1000000)
	NewRatioConversion(GigaByte, Byte, 1000000000)
	NewRatioConversion(TeraByte, Byte, 1000000000000)
	NewRatioConversion(PetaByte, Byte, 1000000000000000)
	NewRatioConversion(ExaByte, Byte, 1000000000000000000)
	NewRatioConversion(ZettaByte, Byte, 1000000000000000000000)
	NewRatioConversion(YottaByte, Byte, 1000000000000000000000000)

	NewRatioConversion(Kibibyte, Byte, kb)
	NewRatioConversion(Mebibyte, Byte, mb)
	NewRatioConversion(Gibibyte, Byte, gb)
	NewRatioConversion(Tebibyte, Byte, tb)
	NewRatioConversion(Pebibyte, Byte, pb)
	NewRatioConversion(Exbibyte, Byte, eb)
	NewRatioConversion(Zebibyte, Byte, zb)
	NewRatioConversion(Yobibyte, Byte, yb)
}
