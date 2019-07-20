package packet

import (
	"encoding/binary"

	"gitlab.com/jeshuamorrissey/wow-server-go/packet"
)

// ClientLoginChallenge encodes information about a new connection to the
// login server.
type ClientLoginChallenge struct {
	errorCode      uint8
	length         uint16
	gameName       [4]byte
	version        [3]uint8
	build          uint16
	platform       [4]byte
	os             [4]byte
	locale         [4]byte
	timezoneOffset uint32
	ipAddress      uint32
	accountNameLen uint8
	accountName    []byte
}

func (pkt *ClientLoginChallenge) Process(process packet.ProcessFunc) {
	process(binary.LittleEndian)
}
