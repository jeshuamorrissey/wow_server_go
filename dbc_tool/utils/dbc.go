package utils

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"sort"
)

// Record represents a generic record type (which must
// be convertable to a byte array).
type Record interface {
	Clone() Record

	// Information methods.
	NumFields() int
	Strings() []string
	TableName() string

	// Binary serialization/de-serialization methods.
	FromBytes(StringBlock, io.Reader) error
	ToBytes(StringBlock) []byte
}

// Header is the common header of the binary DBC files.
type Header struct {
	Magic           uint32
	NumRecords      uint32
	NumFields       uint32
	RecordSize      uint32
	StringBlockSize uint32
}

// StringBlock is a mapping of offsets --> string values.
type StringBlock map[uint32]string

// FindIndex returns the key matching the given value.
func (sb StringBlock) FindIndex(s string) uint32 {
	for k, v := range sb {
		if v == s {
			return k
		}
	}

	return 0
}

// ToBytes converts the string block to its bytes representation.
func (sb StringBlock) ToBytes() []byte {
	offsets := make([]int, 0, len(sb))
	for k := range sb {
		offsets = append(offsets, int(k))
	}

	sort.Ints(offsets)

	buffer := bytes.NewBufferString("")
	for _, offset := range offsets {
		buffer.WriteString(sb[uint32(offset)])
		if offset > 0 {
			buffer.WriteByte('\x00')
		}
	}

	return buffer.Bytes()
}

// DBC represents a complete blizz-like DBC file.
type DBC struct {
	Records []Record
}

// LoadBinaryDBC loads the given filename using the given record type.
func LoadBinaryDBC(data []byte, sampleRecord Record) (*DBC, error) {
	dbc := new(DBC)

	_, strings, recordData, err := Parse(data)
	if err != nil {
		return nil, err
	}

	// Load the records from the bytes.
	dbc.Records = make([]Record, 0)
	recordDataBuffer := bytes.NewBuffer(recordData)
	for recordDataBuffer.Len() > 0 {
		record := sampleRecord.Clone()
		err = record.FromBytes(strings, recordDataBuffer)
		if err != nil {
			return nil, err
		}

		dbc.Records = append(dbc.Records, record)
	}

	if recordDataBuffer.Len() > 0 {
		return nil, fmt.Errorf("%v bytes of data at end of records", recordDataBuffer.Len())
	}

	return dbc, nil
}

// ToBinary converts the given DBC to binary.
func (dbc *DBC) ToBinary() []byte {
	buffer := bytes.NewBufferString("")

	header, stringBlock := dbc.metadata()
	binary.Write(buffer, binary.LittleEndian, header)
	for _, record := range dbc.Records {
		buffer.Write(record.ToBytes(stringBlock))
	}

	buffer.Write(stringBlock.ToBytes())

	return buffer.Bytes()
}

// Header makes a new header and returns it.
func (dbc *DBC) metadata() (*Header, StringBlock) {
	header := new(Header)

	// Make the string block.
	strings := make([]string, 0)
	stringsSet := make(map[string]bool)
	for _, record := range dbc.Records {
		for _, s := range record.Strings() {
			if _, ok := stringsSet[s]; !ok {
				strings = append(strings, s)
				stringsSet[s] = true
			}
		}
	}

	sort.Strings(strings)

	stringBlock := make(StringBlock)
	stringBlock[0] = "\x00"
	nextOffset := 1
	for _, s := range strings {
		stringBlock[uint32(nextOffset)] = s
		nextOffset += len(s) + 1
	}

	header.Magic = MagicExpected
	header.NumRecords = uint32(len(dbc.Records))
	header.NumFields = uint32(dbc.Records[0].NumFields())
	header.RecordSize = uint32(len(dbc.Records[0].ToBytes(stringBlock)))
	header.StringBlockSize = uint32(nextOffset)

	return header, stringBlock
}
