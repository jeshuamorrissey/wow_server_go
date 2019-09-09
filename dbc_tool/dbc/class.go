package dbc

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
	"io"
	"io/ioutil"
	"path/filepath"
	"strings"

	"github.com/jeshuamorrissey/wow_server_go/dbc_tool/utils"
	c "github.com/jeshuamorrissey/wow_server_go/worldserver/data/dbc/constants"
)

// Class represents an in-game class.
type Class struct {
	ID          int
	Name        string
	PrimaryStat c.Stat
	PowerType   c.Power
}

type classDBC struct {
	ID              uint32
	IsPlayerClass   uint32
	DamageBonusStat uint32
	PowerType       uint32
	Unk1            uint32
	Name            LocalizedString
	Filename        String
	SpellClassSet   uint32
	Flags           uint32
}

// FromBytes loads this struct from a blizz-like binary record.
func (class *Class) FromBytes(stringBlock utils.StringBlock, buffer io.Reader) error {
	var classDBC classDBC
	err := binary.Read(buffer, binary.LittleEndian, &classDBC)
	if err != nil {
		return err
	}

	class.ID = int(classDBC.ID)
	class.Name = stringBlock[classDBC.Name.EnUS]
	class.PrimaryStat = c.Stat(classDBC.DamageBonusStat)
	class.PowerType = c.Power(classDBC.PowerType)

	return nil
}

// ToBytes converts this struct to a blizz-like binary record.
func (class *Class) ToBytes(stringBlock utils.StringBlock) []byte {
	buffer := bytes.NewBufferString("")

	name := LocalizedString{
		EnUS:  stringBlock.FindIndex(class.Name),
		Flags: 7274527,
	}

	binary.Write(buffer, binary.LittleEndian, uint32(class.ID))
	binary.Write(buffer, binary.LittleEndian, uint32(1)) // IsPlayerClass
	binary.Write(buffer, binary.LittleEndian, uint32(class.PrimaryStat))
	binary.Write(buffer, binary.LittleEndian, uint32(class.PowerType))
	binary.Write(buffer, binary.LittleEndian, String(0)) // PetNameToken
	binary.Write(buffer, binary.LittleEndian, name)
	binary.Write(buffer, binary.LittleEndian, String(stringBlock.FindIndex(strings.ToUpper(class.Name))))
	binary.Write(buffer, binary.LittleEndian, uint32(0)) // SpellClassSet
	binary.Write(buffer, binary.LittleEndian, uint32(0)) // Flags

	return buffer.Bytes()
}

// Strings returns a list of strings embedded in the struct.
func (class *Class) Strings() []string {
	return []string{class.Name, strings.ToUpper(class.Name)}
}

// TableName returns the blizz-like table this struct represents.
func (class *Class) TableName() string {
	return "ChrClasses"
}

// NumFields returns the number of fields for this record.
func (class *Class) NumFields() int {
	return 17
}

// Clone returns a copy of the current object.
func (*Class) Clone() utils.Record {
	return new(Class)
}

// LoadClasses the class information from the JSON data.
func LoadClasses(rootDir string) (*utils.DBC, error) {
	jsonFileContents, err := ioutil.ReadFile(filepath.Join(rootDir, "class.json"))
	if err != nil {
		return nil, err
	}

	// Load the specific records.
	var records []*Class
	err = json.Unmarshal(jsonFileContents, &records)
	if err != nil {
		return nil, err
	}

	// Convert them into generic records.
	genericRecords := make([]utils.Record, 0, len(records))
	for _, record := range records {
		genericRecords = append(genericRecords, record)
	}

	// Return the new DBC.
	return &utils.DBC{
		Records: genericRecords,
	}, nil
}
