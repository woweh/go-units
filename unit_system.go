package units

// UnitSystem is a system of units
type UnitSystem string

// Systems of units
func (s UnitSystem) String() string { return string(s) }

const (
	// SiSystem provides the internal name for International System of Units
	SiSystem UnitSystem = "metric"
	// BiSystem provides the internal name the British Imperial system of units
	BiSystem UnitSystem = "imperial"
	// UsSystem provides the internal name the United States customary system of units
	UsSystem UnitSystem = "us"
	// IecSystem provides the internal name the International Electrotechnical Commission system of units
	IecSystem UnitSystem = "iec"
	// CgsSystem provides the internal name the centimeter-gram-second system of units
	CgsSystem UnitSystem = "cgs"
	// MKpSSystem provides the internal name the MKpS system of units (from French mètre–kilogramme-poids–seconde)
	MKpSSystem UnitSystem = "MKpS"
	// NoSystem provides the internal name for no system of units
	NoSystem UnitSystem = ""
)

// Shorthands for pre-defined unit systems:
var (
	// SI is the International System of Units
	SI = System(SiSystem)
	// BI is the British Imperial system of units
	BI = System(BiSystem)
	// US is the United States customary system of units
	US = System(UsSystem)
	// IEC is the International Electrotechnical Commission system of units
	IEC = System(IecSystem)
	// CGS is the centimeter-gram-second system of units
	CGS = System(CgsSystem)
	// MKpS is the MKpS system of units (from French mètre–kilogramme-poids–seconde)
	MKpS = System(MKpSSystem)
)
