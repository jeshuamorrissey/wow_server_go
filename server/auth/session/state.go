package session

import (
	"math/big"

	"github.com/sirupsen/logrus"

	"github.com/jeshuamorrissey/wow_server_go/lib/config"
	"github.com/jeshuamorrissey/wow_server_go/server/auth/data/static"
)

// State represents all information required by the AuthServer.
type State struct {
	log     *logrus.Entry
	session *Session

	// Mapping from OpCode --> Packet generation function.
	opCodeToPacket map[static.OpCode]func() ClientPacket

	// The global configuration.
	Config *config.Config

	// The following are SRP specific numbers which are generated as part
	// of the authentication flow.
	PublicEphemeral  *big.Int
	PrivateEphemeral *big.Int

	// The user's account is fetched after the first packet is processed.
	Account *config.Account
}

// NewState creates a new state based on the given DB connection.
func NewState(config *config.Config, log *logrus.Entry, opCodeToPacket map[static.OpCode]func() ClientPacket) *State {
	return &State{
		Config:           config,
		log:              log,
		opCodeToPacket:   opCodeToPacket,
		PublicEphemeral:  new(big.Int),
		PrivateEphemeral: new(big.Int),
	}
}

// AddLogField will add a new field to the logger for this session.
func (s *State) AddLogField(key string, value interface{}) {
	s.log = s.log.WithField(key, value)
}
