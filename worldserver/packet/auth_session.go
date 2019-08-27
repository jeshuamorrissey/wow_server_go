package packet

import (
	"encoding/binary"
	"io"

	"github.com/jeshuamorrissey/wow_server_go/common/database"
	"github.com/jeshuamorrissey/wow_server_go/worldserver/system"
	"github.com/jinzhu/gorm"
)

// ClientAuthSession is the initial message sent from the server
// to the client to start authorization.
type ClientAuthSession struct {
	BuildNumber      uint32
	AccountName      []byte
	ClientSeed       uint32
	ClientProof      [20]byte
	AddonSize        uint32
	AddonsCompressed []byte
}

// FromBytes reads a ClientAuthSession pcket from the byter buffer.
func (pkt *ClientAuthSession) FromBytes(state *system.State, buffer io.Reader) error {
	var unk uint32

	binary.Read(buffer, binary.LittleEndian, &pkt.BuildNumber)
	binary.Read(buffer, binary.LittleEndian, &unk)

	// Null-terminated account name
	var b byte
	binary.Read(buffer, binary.LittleEndian, &b)
	for b != '\x00' {
		pkt.AccountName = append(pkt.AccountName, b)
		binary.Read(buffer, binary.LittleEndian, &b)
	}

	binary.Read(buffer, binary.LittleEndian, &pkt.ClientSeed)
	binary.Read(buffer, binary.LittleEndian, &pkt.ClientProof)
	binary.Read(buffer, binary.LittleEndian, &pkt.AddonSize)

	pkt.AddonsCompressed = make([]byte, pkt.AddonSize)
	buffer.Read(pkt.AddonsCompressed)
	return nil
}

// Handle will ensure that the given account exists.
func (pkt *ClientAuthSession) Handle(state *system.State) ([]system.ServerPacket, error) {
	response := new(ServerAuthResponse)
	response.Error = AuthOK

	state.Account = new(database.Account)
	err := state.DB.Where(&database.Account{Name: string(pkt.AccountName)}).First(state.Account).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			response.Error = AuthUnknownAccount
		} else {
			return nil, err
		}
	}

	// TODO(jeshua): validate the information sent by the client.
	// If there is no session key, account is invalid.
	if state.Account.SessionKey() == nil {
		response.Error = AuthBadServerProof
	}

	if response.Error == AuthOK {
		state.Log = state.Log.WithField("account", state.Account.Name)
		state.Log.Infof("Account %v authenticated!", state.Account.Name)
	}

	return []system.ServerPacket{response}, nil
}

// OpCode returns the opcode for this packet.
func (pkt *ClientAuthSession) OpCode() system.OpCode {
	return system.OpCodeClientAuthSession
}
