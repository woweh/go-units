package units

var (
	// Information is the unit quantity for information.
	Information = NewUnitQuantity("information")
	Byte        = Information.MustCreateUnit("byte", "B", SI)
	KiloByte    = Kilo(Byte, Symbols("KB"))
	MegaByte    = Mega(Byte)
	GigaByte    = Giga(Byte)
	TeraByte    = Tera(Byte)
	PetaByte    = Peta(Byte)
	ExaByte     = Exa(Byte)
	ZettaByte   = Zetta(Byte)
	YottaByte   = Yotta(Byte)
	RonnaByte   = Ronna(Byte)
	QuettaByte  = Quetta(Byte)

	Kibibyte = Information.MustCreateUnit("kibibyte", "KiB", IEC)
	Mebibyte = Information.MustCreateUnit("mebibyte", "MiB", IEC)
	Gibibyte = Information.MustCreateUnit("gibibyte", "GiB", IEC)
	Tebibyte = Information.MustCreateUnit("tebibyte", "TiB", IEC)
	Pebibyte = Information.MustCreateUnit("pebibyte", "PiB", IEC)
	Exbibyte = Information.MustCreateUnit("exbibyte", "EiB", IEC)
	Zebibyte = Information.MustCreateUnit("zebibyte", "ZiB", IEC)
	Yobibyte = Information.MustCreateUnit("yobibyte", "YiB", IEC)
	Robibyte = Information.MustCreateUnit("robi byte", "RiB", IEC)
	Qubibyte = Information.MustCreateUnit("qubi byte", "QiB", IEC)

	Bit       = Information.MustCreateUnit("bit", "b", SI)
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

	Nibble = Information.MustCreateUnit("nibble", "")
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
