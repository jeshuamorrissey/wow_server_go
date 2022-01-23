package handlers

import (
	"strings"

	"github.com/jeshuamorrissey/wow_server_go/server/world/data/dynamic/interfaces"
	"github.com/jeshuamorrissey/wow_server_go/server/world/data/static"
	"github.com/jeshuamorrissey/wow_server_go/server/world/packet"
	"github.com/jeshuamorrissey/wow_server_go/server/world/system"
)

// Handle will ensure that the given account exists.
func HandleClientAuthSession(pkt *packet.ClientAuthSession, state *system.State) ([]interfaces.ServerPacket, error) {
	response := new(packet.ServerAuthResponse)
	response.Error = static.AuthOK

	for _, account := range state.Config.Accounts {
		if strings.ToUpper(account.Name) == strings.ToUpper(string(pkt.AccountName)) {
			state.Account = account
		}
	}

	if state.Account == nil {
		response.Error = static.AuthUnknownAccount
	}

	// TODO(jeshua): validate the information sent by the client.
	// If there is no session key, account is invalid.
	if state.Account != nil && state.Account.SessionKey() == nil {
		response.Error = static.AuthBadServerProof
	}

	if response.Error == static.AuthOK {
		state.Log = state.Log.WithField("account", state.Account.Name)
		state.Log.Infof("Account %v authenticated!", state.Account.Name)
	}

	return []interfaces.ServerPacket{response}, nil
}
