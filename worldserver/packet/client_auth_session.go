package packet

import (
	"encoding/binary"
	"io"
	"strings"

	c "github.com/jeshuamorrissey/wow_server_go/worldserver/data/dbc/constants"
	"github.com/jeshuamorrissey/wow_server_go/worldserver/system"
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
	response.Error = c.AuthOK

	for _, account := range state.Config.Accounts {
		if strings.ToUpper(account.Name) == strings.ToUpper(string(pkt.AccountName)) {
			state.Account = account
		}
	}

	if state.Account == nil {
		response.Error = c.AuthUnknownAccount
	}

	// TODO(jeshua): validate the information sent by the client.
	// If there is no session key, account is invalid.
	if state.Account != nil && state.Account.SessionKey() == nil {
		response.Error = c.AuthBadServerProof
	}

	if response.Error == c.AuthOK {
		state.Log = state.Log.WithField("account", state.Account.Name)
		state.Log.Infof("Account %v authenticated!", state.Account.Name)
	}

	return []system.ServerPacket{response}, nil
}

// OpCode returns the opcode for this packet.
func (pkt *ClientAuthSession) OpCode() system.OpCode {
	return system.OpCodeClientAuthSession
}
