package worldserver

import (
	"github.com/jeshuamorrissey/wow_server_go/worldserver/packet"
	"github.com/jeshuamorrissey/wow_server_go/worldserver/system"
)

var (
	opCodeToPacket = map[system.OpCode]func() system.ClientPacket{
		system.OpCodeClientAuthSession:     func() system.ClientPacket { return new(packet.ClientAuthSession) },
		system.OpCodeClientCharCreate:      func() system.ClientPacket { return new(packet.ClientCharCreate) },
		system.OpCodeClientCharDelete:      func() system.ClientPacket { return new(packet.ClientCharDelete) },
		system.OpCodeClientCharEnum:        func() system.ClientPacket { return new(packet.ClientCharEnum) },
		system.OpCodeClientItemQuerySingle: func() system.ClientPacket { return new(packet.ClientItemQuerySingle) },
		system.OpCodeClientNameQuery:       func() system.ClientPacket { return new(packet.ClientNameQuery) },
		system.OpCodeClientPing:            func() system.ClientPacket { return new(packet.ClientPing) },
		system.OpCodeClientPlayerLogin:     func() system.ClientPacket { return new(packet.ClientPlayerLogin) },
		system.OpCodeClientTutorialFlag:    func() system.ClientPacket { return new(packet.ClientTutorialFlag) },
	}
)
