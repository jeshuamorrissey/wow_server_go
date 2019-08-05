package packet

import (
	"encoding/binary"
	"io"

	"github.com/jeshuamorrissey/wow_server_go/common/database"
	"github.com/jeshuamorrissey/wow_server_go/common/session"
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

func (pkt *ClientAuthSession) Read(buffer io.Reader) error {
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
func (pkt *ClientAuthSession) Handle(stateBase session.State) ([]session.ServerPacket, error) {
	state := stateBase.(*State)
	response := new(ServerAuthResponse)
	response.Error = AuthOK

	err := stateBase.DB().Where(&database.Account{Name: string(pkt.AccountName)}).First(&state.Account).Error
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
		stateBase.AddLogField("account", state.Account.Name)
		stateBase.Log().Infof("Account %v authenticated!", state.Account.Name)
	}

	return []session.ServerPacket{response}, nil
}
