package dbc

// String represents a single string within the DBC.
type String uint32

// LocalizedString is a mapping from locale --> string.
type LocalizedString struct {
	EnUS, KoKR, FrFR, DeDE, EnCN, EnTW, EsES, EsMX uint32
	Flags                                          uint32
}
