package dbc

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

// FactionDBC represents the data from faction.dbc.
type FactionDBC struct {
	ID                  uint32
	ReputationIndex     uint32
	ReputationRaceMask  [4]uint32
	ReputationClassMask [4]uint32
	ReputationBase      [4]uint32
	ParentFactionID     uint32
	NameRefs            LocalizedName
}

type Faction struct {
}

// func (f *FactionDBC) ToFaction() *Faction {

// }

// LoadFactions will load the given DBC file into a list of Faction
// structs.
func LoadFactions(filepath string) (map[int]*FactionDBC, error) {
	header, stringBlock, recordData, err := Parse(filepath)
	if err != nil {
		return nil, err
	}

	fmt.Printf("header = %v\n", header)

	// Load the records.
	records := make(map[int]*FactionDBC)
	recordsBuffer := bytes.NewBuffer(recordData)
	for i := 0; i < int(header.NumRecords); i++ {
		record := new(FactionDBC)

		err = binary.Read(recordsBuffer, binary.LittleEndian, record)
		if err != nil {
			return nil, err
		}

		if name, ok := stringBlock[int(record.NameRefs.EnUS)]; ok {
			fmt.Printf("name = %v\n", name)
		}

		records[int(record.ID)] = record
	}

	if recordsBuffer.Len() > 0 {
		return nil, fmt.Errorf("Should have consumed all data, have %v bytes left", recordsBuffer.Len())
	}

	return records, nil
}
