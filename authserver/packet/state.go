package packet

import (
	"math/big"

	"gitlab.com/jeshuamorrissey/mmo_server/database"
)

// State is a set of object required by all handlers. Things should only be added to
// this structure as they are needed.
type State struct {
	PublicEphemeral  big.Int
	PrivateEphemeral big.Int

	Account *database.Account
}
