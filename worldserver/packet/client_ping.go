package packet

import (
	"encoding/binary"
	"io"

	"github.com/jeshuamorrissey/wow_server_go/common/session"
)

// ClientPing is sent from the client periodically.
type ClientPing struct {
	Ping    uint32
	Latency uint32
}

func (pkt *ClientPing) Read(buffer io.Reader) error {
	binary.Read(buffer, binary.LittleEndian, pkt.Ping)
	binary.Read(buffer, binary.LittleEndian, pkt.Latency)
	return nil
}

// Handle will ensure that the given account exists.
func (pkt *ClientPing) Handle(stateBase session.State) ([]session.ServerPacket, error) {
	response := new(ServerPong)
	response.Pong = pkt.Ping

	return []session.ServerPacket{response}, nil
}
