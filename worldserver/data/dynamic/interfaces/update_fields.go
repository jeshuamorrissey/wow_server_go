package interfaces

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"math/big"

	"github.com/jeshuamorrissey/wow_server_go/common"
	"github.com/jeshuamorrissey/wow_server_go/worldserver/data/static"
)

// UpdateFieldsMap is a mapping from an UpdateField constant to some value.
// Values stored can be either floating point numbers of unsigned integers (both 32-bit).
type UpdateFieldsMap map[static.UpdateField]interface{}

// ToBytes converts an update fields map to bytes.
func (m UpdateFieldsMap) ToBytes(nFields int) []byte {
	mask := big.NewInt(0)
	fields := bytes.NewBufferString("")

	for i := 0; i < nFields; i++ {
		field := static.UpdateField(i)
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
