package packet

import "github.com/jeshuamorrissey/wow_server_go/worldserver/objects"

// ServerUpdateObject is the UPDATE_OBJECT packet.
type ServerUpdateObject struct {
	object objects.GameObject
}

// Bytes converts the packet into an array of bytes.
func (pkt *ServerUpdateObject) Bytes() []byte {
	return []byte{}
}

// OpCode returns the OpCode of the packet.
func (pkt *ServerUpdateObject) OpCode() OpCode {
	return OpCodeServerUpdateObject
}
