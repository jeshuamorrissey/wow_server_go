package packet

import (
	"encoding/binary"
	"io"
)

// ClientPing is sent from the client periodically.
type ClientPing struct {
	Ping    uint32
	Latency uint32
}

// FromBytes reads packet data from the given buffer.
func (pkt *ClientPing) FromBytes(state *State, buffer io.Reader) error {
	binary.Read(buffer, binary.LittleEndian, pkt.Ping)
	binary.Read(buffer, binary.LittleEndian, pkt.Latency)
	return nil
}

// Handle will ensure that the given account exists.
func (pkt *ClientPing) Handle(state *State) ([]ServerPacket, error) {
	response := new(ServerPong)
	response.Pong = pkt.Ping

	return []ServerPacket{response}, nil
}

// OpCode gets the opcode of the packet.
func (*ClientPing) OpCode() OpCode {
	return OpCodeClientPing
}
