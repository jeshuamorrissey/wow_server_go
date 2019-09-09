package utils

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"strings"
)

const (
	// MagicExpected is the expected value of Magic in the header.
	// As a string of 4 bytes, this is "WDBC".
	MagicExpected = 1128416343
)

// Parse takes as input a filepath and a storage struct and loads the given file into
// the storage.
func Parse(data []byte) (*Header, StringBlock, []byte, error) {
	buffer := bytes.NewBuffer(data)

	// Load the header.
	header := new(Header)
	binary.Read(buffer, binary.LittleEndian, header)

	if header.Magic != MagicExpected {
		return nil, nil, nil, fmt.Errorf("Unexpected magic: wanted '%v' got '%v'", MagicExpected, header.Magic)
	}

	// Load the data.
	recordData := make([]byte, header.NumRecords*header.RecordSize)
	buffer.Read(recordData)

	// Load the string block.
	stringBlockData := make([]byte, header.StringBlockSize)
	buffer.Read(stringBlockData)

	if buffer.Len() > 0 {
		return nil, nil, nil, fmt.Errorf("Malformed header: %v bytes left after data", buffer.Len())
	}

	// Load the string block.
	stringBlock := make(StringBlock)
	stringBlockBuffer := bytes.NewBuffer(stringBlockData)
	currOffset := uint32(0)
	for stringBlockBuffer.Len() > 0 {
		stringBlock[currOffset], _ = stringBlockBuffer.ReadString('\x00')
		stringBlock[currOffset] = strings.Trim(stringBlock[currOffset], "\x00")
		currOffset += uint32(len(stringBlock[currOffset]) + 1)
	}

	// Load the records, but first make sure they are the same size.
	return header, stringBlock, recordData, nil
}
