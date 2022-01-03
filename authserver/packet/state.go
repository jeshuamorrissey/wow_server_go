package packet

import (
	"math/big"

	"github.com/jeshuamorrissey/wow_server_go/common/session"
	"github.com/jeshuamorrissey/wow_server_go/worldserver/data/world"
	"github.com/sirupsen/logrus"
)

// State represents all information required by the AuthServer.
type State struct {
	config  *world.WorldConfig
	log     *logrus.Entry
	session *session.Session

	PublicEphemeral  big.Int
	PrivateEphemeral big.Int

	Account *world.Account
}

// NewState creates a new state based on the given DB connection.
func NewState(config *world.WorldConfig, log *logrus.Entry) *State {
	return &State{config: config, log: log}
}

// DB returns a reference to the Database object stored in this state.
func (s *State) Config() *world.WorldConfig {
	return s.config
}

// Log returns a reference to the Database object stored in this state.
func (s *State) Log() *logrus.Entry {
	return s.log
}

// SetSession updates the local session to point to a session.
func (s *State) SetSession(sess *session.Session) {
	s.session = sess
}

// Session gets a reference to the associated session.
func (s *State) Session() *session.Session {
	return s.session
}

// AddLogField adds a new field to the logger for this state.
func (s *State) AddLogField(key string, value interface{}) {
	*s.log = *s.log.WithField(key, value)
}
