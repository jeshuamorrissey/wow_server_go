package packet

import (
	"fmt"

	"github.com/jeshuamorrissey/wow_server_go/common/session"
)

// OpCodes used by the AuthServer.
// TODO(jeshua): Implement all OpCodes.
const (
	ClientLoginChallengeOpCode session.OpCode = 0x00
	ClientLoginProofOpCode     session.OpCode = 0x01
	ClientRealmlistOpCode      session.OpCode = 0x10

	ServerLoginChallengeOpCode session.OpCode = 0x00
	ServerLoginProofOpCode     session.OpCode = 0x01
	ServerRealmlistOpCode      session.OpCode = 0x10
)

// OpCodeName returns a string name for a given OpCode.
func OpCodeName(opCode session.OpCode) string {
	if opCode == ClientLoginChallengeOpCode {
		return "LOGIN_CHALLENGE"
	} else if opCode == ClientLoginProofOpCode {
		return "LOGIN_PROOF"
	} else if opCode == ClientRealmlistOpCode {
		return "REALMLIST"
	}

	return fmt.Sprintf("UNKNOWN (%v)", opCode)
}
