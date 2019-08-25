package worldserver

import (
	"github.com/jeshuamorrissey/wow_server_go/worldserver/packet"
)

var (
	opCodeToPacket = map[packet.OpCode]func() packet.ClientPacket{
		packet.OpCodeClientAuthSession: func() packet.ClientPacket { return new(packet.ClientAuthSession) },
		packet.OpCodeClientCharCreate:  func() packet.ClientPacket { return new(packet.ClientCharCreate) },
		packet.OpCodeClientCharDelete:  func() packet.ClientPacket { return new(packet.ClientCharDelete) },
		packet.OpCodeClientCharEnum:    func() packet.ClientPacket { return new(packet.ClientCharEnum) },
		packet.OpCodeClientPing:        func() packet.ClientPacket { return new(packet.ClientPing) },
		packet.OpCodeClientPlayerLogin: func() packet.ClientPacket { return new(packet.ClientPlayerLogin) },
	}
)
