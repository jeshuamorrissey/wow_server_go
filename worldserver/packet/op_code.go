package packet

import (
	"fmt"

	"github.com/jeshuamorrissey/wow_server_go/common/session"
)

// OpCodes used by the AuthServer.
// TODO(jeshua): Implement all OpCodes.
const ()

// OpCodeName returns a string name for a given OpCode.
func OpCodeName(opCode session.OpCode) string {
	return fmt.Sprintf("UNKNOWN (%v)", opCode)
}
