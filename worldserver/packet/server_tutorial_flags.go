package packet

import (
	"bytes"
	"math/big"

	"github.com/jeshuamorrissey/wow_server_go/common"

	"github.com/jeshuamorrissey/wow_server_go/worldserver/data/static"
	"github.com/jeshuamorrissey/wow_server_go/worldserver/system"
)

// ServerTutorialFlags is sent back in response to ClientPing.
type ServerTutorialFlags struct{}

// ToBytes writes out the packet to an array of bytes.
func (pkt *ServerTutorialFlags) ToBytes(state *system.State) ([]byte, error) {
	buffer := bytes.NewBufferString("")

	// Convert the binary array to a bitmask.
	mask := big.NewInt(0)
	for i, isDone := range state.Character.Tutorials {
		if isDone {
			mask.SetBit(mask, i, 1)
		}
	}

	buffer.Write(common.PadBigIntBytes(common.ReverseBytes(mask.Bytes()), 8))

	return buffer.Bytes(), nil
}

// OpCode gets the opcode of the packet.
func (*ServerTutorialFlags) OpCode() static.OpCode {
	return static.OpCodeServerTutorialFlags
}
