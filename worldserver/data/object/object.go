package object

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"math/big"

	"github.com/jeshuamorrissey/wow_server_go/common"
	c "github.com/jeshuamorrissey/wow_server_go/worldserver/data/dbc/constants"
)

// UpdateFieldsMap is a mapping from an UpdateField constant to some value.
type UpdateFieldsMap map[c.UpdateField]interface{}

// ToBytes converts an update fields map to bytes.
func (m UpdateFieldsMap) ToBytes(nFields int) []byte {
	mask := big.NewInt(0)
	fields := bytes.NewBufferString("")

	for i := 0; i < nFields; i++ {
		field := c.UpdateField(i)
		valueGeneric, ok := m[field]
		if !ok {
			continue
		}

		switch value := valueGeneric.(type) {
		case float32:
			binary.Write(fields, binary.LittleEndian, value)
			mask.SetBit(mask, i, 1)
		case uint32:
			binary.Write(fields, binary.LittleEndian, value)
			mask.SetBit(mask, i, 1)
		default:
			panic(fmt.Sprintf("Unknown field type %T in update fields (%v)", value, field))
		}
	}

	nBlocks := (nFields + 32 - 1) / 32
	nBytes := (nBlocks * 32) / 8

	fieldBytes := make([]byte, 0)
	fieldBytes = append(fieldBytes, uint8(nBlocks))
	fieldBytes = append(fieldBytes, common.PadBigIntBytes(common.ReverseBytes(mask.Bytes()), nBytes)...)
	fieldBytes = append(fieldBytes, fields.Bytes()...)
	return fieldBytes
}

// Object represents a generic object within the game. All objects
// should implement this interface.
type Object interface {
	// Manager should return the manager associated with this object.
	Manager() *Manager

	// SetManager should update the manager to the given value.
	SetManager(*Manager)

	// GUID should return the full GUID of the object.
	GUID() GUID

	// SetGUID should set the GUID of the object.
	SetGUID(GUID)

	// Location should return the location within the game world. If the
	// object has no actual location, it should use the location of
	// it's container.
	Location() *Location

	// MovementUpdate should return the full movement update for the
	// object.
	MovementUpdate() []byte

	// UpdateFields should return the update fields for the object.
	UpdateFields() UpdateFieldsMap
}
