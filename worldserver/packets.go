package worldserver

import (
	"github.com/jeshuamorrissey/wow_server_go/common/session"
	"github.com/jeshuamorrissey/wow_server_go/worldserver/packet"
	"github.com/jeshuamorrissey/wow_server_go/worldserver/system"
)

var (
	opCodeToPacket = map[session.OpCode]func() system.ClientPacket{
		packet.OpCodeClientAuthSession: func() system.ClientPacket { return new(packet.ClientAuthSession) },
		packet.OpCodeClientCharCreate:  func() system.ClientPacket { return new(packet.ClientCharCreate) },
		packet.OpCodeClientCharDelete:  func() system.ClientPacket { return new(packet.ClientCharDelete) },
		packet.OpCodeClientCharEnum:    func() system.ClientPacket { return new(packet.ClientCharEnum) },
		packet.OpCodeClientPing:        func() system.ClientPacket { return new(packet.ClientPing) },
		packet.OpCodeClientPlayerLogin: func() system.ClientPacket { return new(packet.ClientPlayerLogin) },
	}
)
