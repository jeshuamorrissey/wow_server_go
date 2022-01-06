package static

// OpCode is the opcode type used by the auth server.
type OpCode int

// OpCodes used by the AuthServer.
// TODO(jeshua): Implement all OpCodes.
const (
	OpCodeLoginChallenge OpCode = 0x00
	OpCodeLoginProof     OpCode = 0x01
	OpCodeRealmlist      OpCode = 0x10
)

// Int returns an int reprentation of the opcode.
func (oc OpCode) Int() int {
	return int(oc)
}
