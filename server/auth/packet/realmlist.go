package packet

import (
	"bytes"
	"encoding/binary"
	"io"

	"github.com/jeshuamorrissey/wow_server_go/server/auth/data/static"
	"github.com/jeshuamorrissey/wow_server_go/server/auth/session"
)

// ClientRealmlist packet contains no fields.
type ClientRealmlist struct{}

func (pkt *ClientRealmlist) FromBytes(state *session.State, buffer io.Reader) error {
	var unk uint32
	return binary.Read(buffer, binary.LittleEndian, &unk)
}

// OpCode gets the opcode of the packet.
func (*ClientRealmlist) OpCode() static.OpCode {
	return static.OpCodeRealmlist
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
func (pkt *ServerRealmlist) ToBytes(state *session.State) ([]byte, error) {
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

	return buffer.Bytes(), nil
}

// OpCode returns ServerRealmlistOpCode.
func (*ServerRealmlist) OpCode() static.OpCode {
	return static.OpCodeRealmlist
}

// Handle will check the database for the account and send an appropriate response.
func (pkt *ClientRealmlist) Handle(state *session.State) ([]session.ServerPacket, error) {
	response := new(ServerRealmlist)

	response.Realms = append(response.Realms, Realm{
		Icon:          0,
		Flags:         0,
		Name:          state.Config.Name,
		Address:       "localhost:5001",
		Population:    0,
		NumCharacters: 0,
		Timezone:      0,
	})

	return []session.ServerPacket{response}, nil
}
