package packet

import (
	"bytes"
	"encoding/binary"
	"io"

	"github.com/jeshuamorrissey/wow_server_go/common/database"
	"github.com/jeshuamorrissey/wow_server_go/common/session"
)

// ClientRealmlist packet contains no fields.
type ClientRealmlist struct{}

func (pkt *ClientRealmlist) Read(buffer io.Reader) error {
	var unk uint32
	return binary.Read(buffer, binary.LittleEndian, &unk)
}

// Realm is information required to send as part of the realmlist.
type Realm struct {
	Icon          uint32
	Flags         uint8
	Name          string
	Address       string
	Population    float32
	NumCharacters uint8
	Timezone      uint8
}

// ServerRealmlist is made up of a list of realms.
type ServerRealmlist struct {
	Realms []Realm
}

// Bytes converts the ServerRealmlist packet to an array of bytes.
func (pkt *ServerRealmlist) Bytes(stateBase session.State) []byte {
	realmsBuffer := bytes.NewBufferString("")

	binary.Write(realmsBuffer, binary.LittleEndian, uint32(0)) // unk
	realmsBuffer.WriteByte(uint8(len(pkt.Realms)))

	for _, realm := range pkt.Realms {
		binary.Write(realmsBuffer, binary.LittleEndian, realm.Icon)
		binary.Write(realmsBuffer, binary.LittleEndian, realm.Flags)
		realmsBuffer.WriteString(realm.Name + "\x00")
		realmsBuffer.WriteString(realm.Address + "\x00")
		binary.Write(realmsBuffer, binary.LittleEndian, realm.Population)
		binary.Write(realmsBuffer, binary.LittleEndian, realm.NumCharacters)
		binary.Write(realmsBuffer, binary.LittleEndian, realm.Timezone)
		binary.Write(realmsBuffer, binary.LittleEndian, uint8(0)) // unk
	}

	binary.Write(realmsBuffer, binary.LittleEndian, uint16(2)) // unk

	// Make the real buffer, which has the length at the start.
	buffer := bytes.NewBufferString("")
	binary.Write(buffer, binary.LittleEndian, uint16(realmsBuffer.Len()))
	buffer.Write(realmsBuffer.Bytes())

	return buffer.Bytes()
}

// OpCode returns ServerRealmlistOpCode.
func (*ServerRealmlist) OpCode() session.OpCode {
	return OpCodeRealmlist
}

// Handle will check the database for the account and send an appropriate response.
func (pkt *ClientRealmlist) Handle(stateBase session.State) ([]session.ServerPacket, error) {
	// state := stateBase.(*State)
	response := new(ServerRealmlist)

	// Get information from the session.
	var realms []database.Realm
	err := stateBase.DB().Find(&realms).Error
	if err != nil {
		return nil, err
	}

	for _, realm := range realms {
		response.Realms = append(response.Realms, Realm{
			Icon:          0,
			Flags:         0,
			Name:          realm.Name,
			Address:       realm.Host,
			Population:    0,
			NumCharacters: 0,
			Timezone:      0,
		})
	}

	return []session.ServerPacket{response}, nil
}
