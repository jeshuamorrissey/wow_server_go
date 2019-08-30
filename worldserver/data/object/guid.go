package object

import (
	"encoding/binary"

	c "github.com/jeshuamorrissey/wow_server_go/worldserver/data/dbc/constants"
)

// GUID is a globally-unique identifier used to identify objects within
// the database.
type GUID uint64

// MakeGUID constructs a new GUID from a low/high pair.
func MakeGUID(low uint32, high c.HighGUID) GUID {
	return GUID(uint64(high)<<32 | uint64(low))
}

// High returns the high-part of the GUID. This correspondings to a HighGUID
// value.
func (guid GUID) High() c.HighGUID {
	return c.HighGUID(guid >> 32)
}

// Low returns the low-part of the GUID. This is the part that can be modified
// by the server.
func (guid GUID) Low() uint32 {
	return uint32(guid)
}

// Pack returns a minimal version of the GUID as an array of bytes.
func (guid GUID) Pack() []byte {
	guidBytes := make([]byte, 8)
	binary.LittleEndian.PutUint64(guidBytes, uint64(guid))

	mask := uint8(0)
	packedGUID := make([]byte, 0)
	for i, b := range guidBytes {
		if b != 0 {
			mask |= (1 << uint(i))
			packedGUID = append(packedGUID, b)
		}
	}

	return append([]byte{mask}, packedGUID...)
}
