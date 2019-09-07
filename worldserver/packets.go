package worldserver

import (
	"github.com/jeshuamorrissey/wow_server_go/worldserver/packet"
	"github.com/jeshuamorrissey/wow_server_go/worldserver/system"
)

var (
	opCodeToPacket = map[system.OpCode]func() system.ClientPacket{
		system.OpCodeClientAuthSession:       func() system.ClientPacket { return new(packet.ClientAuthSession) },
		system.OpCodeClientCharCreate:        func() system.ClientPacket { return new(packet.ClientCharCreate) },
		system.OpCodeClientCharDelete:        func() system.ClientPacket { return new(packet.ClientCharDelete) },
		system.OpCodeClientCharEnum:          func() system.ClientPacket { return new(packet.ClientCharEnum) },
		system.OpCodeClientItemQuerySingle:   func() system.ClientPacket { return new(packet.ClientItemQuerySingle) },
		system.OpCodeClientNameQuery:         func() system.ClientPacket { return new(packet.ClientNameQuery) },
		system.OpCodeClientPing:              func() system.ClientPacket { return new(packet.ClientPing) },
		system.OpCodeClientPlayerLogin:       func() system.ClientPacket { return new(packet.ClientPlayerLogin) },
		system.OpCodeClientQueryTime:         func() system.ClientPacket { return new(packet.ClientQueryTime) },
		system.OpCodeClientSetActiveMover:    func() system.ClientPacket { return new(packet.ClientSetActiveMover) },
		system.OpCodeClientTutorialFlag:      func() system.ClientPacket { return new(packet.ClientTutorialFlag) },
		system.OpCodeClientUpdateAccountData: func() system.ClientPacket { return new(packet.ClientUpdateAccountData) },
		system.OpCodeClientLogoutRequest:     func() system.ClientPacket { return new(packet.ClientLogoutRequest) },
		system.OpCodeClientStandstatechange:  func() system.ClientPacket { return new(packet.ClientStandStateChange) },

		// Movement packets have the same receiver.
		system.OpCodeClientMoveHeartbeat:        func() system.ClientPacket { return packet.NewClientMovePacket(system.OpCodeClientMoveHeartbeat) },
		system.OpCodeClientMoveStartBackward:    func() system.ClientPacket { return packet.NewClientMovePacket(system.OpCodeClientMoveStartBackward) },
		system.OpCodeClientMoveStartForward:     func() system.ClientPacket { return packet.NewClientMovePacket(system.OpCodeClientMoveStartForward) },
		system.OpCodeClientMoveStartStrafeLeft:  func() system.ClientPacket { return packet.NewClientMovePacket(system.OpCodeClientMoveStartStrafeLeft) },
		system.OpCodeClientMoveStartStrafeRight: func() system.ClientPacket { return packet.NewClientMovePacket(system.OpCodeClientMoveStartStrafeRight) },
		system.OpCodeClientMoveStartTurnLeft:    func() system.ClientPacket { return packet.NewClientMovePacket(system.OpCodeClientMoveStartTurnLeft) },
		system.OpCodeClientMoveStartTurnRight:   func() system.ClientPacket { return packet.NewClientMovePacket(system.OpCodeClientMoveStartTurnRight) },
		system.OpCodeClientMoveStop:             func() system.ClientPacket { return packet.NewClientMovePacket(system.OpCodeClientMoveStop) },
		system.OpCodeClientMoveStopStrafe:       func() system.ClientPacket { return packet.NewClientMovePacket(system.OpCodeClientMoveStopStrafe) },
		system.OpCodeClientMoveStopTurn:         func() system.ClientPacket { return packet.NewClientMovePacket(system.OpCodeClientMoveStopTurn) },
	}
)
