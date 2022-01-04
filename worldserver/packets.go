package worldserver

import (
	"github.com/jeshuamorrissey/wow_server_go/worldserver/data/static"
	"github.com/jeshuamorrissey/wow_server_go/worldserver/packet"
	"github.com/jeshuamorrissey/wow_server_go/worldserver/system"
)

var (
	opCodeToPacket = map[static.OpCode]func() system.ClientPacket{
		static.OpCodeClientAuthSession:       func() system.ClientPacket { return new(packet.ClientAuthSession) },
		static.OpCodeClientCharCreate:        func() system.ClientPacket { return new(packet.ClientCharCreate) },
		static.OpCodeClientCharDelete:        func() system.ClientPacket { return new(packet.ClientCharDelete) },
		static.OpCodeClientCharEnum:          func() system.ClientPacket { return new(packet.ClientCharEnum) },
		static.OpCodeClientCreatureQuery:     func() system.ClientPacket { return new(packet.ClientCreatureQuery) },
		static.OpCodeClientItemQuerySingle:   func() system.ClientPacket { return new(packet.ClientItemQuerySingle) },
		static.OpCodeClientLogoutRequest:     func() system.ClientPacket { return new(packet.ClientLogoutRequest) },
		static.OpCodeClientNameQuery:         func() system.ClientPacket { return new(packet.ClientNameQuery) },
		static.OpCodeClientPing:              func() system.ClientPacket { return new(packet.ClientPing) },
		static.OpCodeClientPlayerLogin:       func() system.ClientPacket { return new(packet.ClientPlayerLogin) },
		static.OpCodeClientQueryTime:         func() system.ClientPacket { return new(packet.ClientQueryTime) },
		static.OpCodeClientSetActiveMover:    func() system.ClientPacket { return new(packet.ClientSetActiveMover) },
		static.OpCodeClientStandstatechange:  func() system.ClientPacket { return new(packet.ClientStandStateChange) },
		static.OpCodeClientTutorialFlag:      func() system.ClientPacket { return new(packet.ClientTutorialFlag) },
		static.OpCodeClientUpdateAccountData: func() system.ClientPacket { return new(packet.ClientUpdateAccountData) },
		static.OpCodeClientAttackswing:       func() system.ClientPacket { return new(packet.ClientAttackSwing) },

		// Movement packets have the same receiver.
		static.OpCodeClientMoveHeartbeat:        func() system.ClientPacket { return packet.NewClientMovePacket(static.OpCodeClientMoveHeartbeat) },
		static.OpCodeClientMoveSetFacing:        func() system.ClientPacket { return packet.NewClientMovePacket(static.OpCodeClientMoveSetFacing) },
		static.OpCodeClientMoveStartBackward:    func() system.ClientPacket { return packet.NewClientMovePacket(static.OpCodeClientMoveStartBackward) },
		static.OpCodeClientMoveStartForward:     func() system.ClientPacket { return packet.NewClientMovePacket(static.OpCodeClientMoveStartForward) },
		static.OpCodeClientMoveStartStrafeLeft:  func() system.ClientPacket { return packet.NewClientMovePacket(static.OpCodeClientMoveStartStrafeLeft) },
		static.OpCodeClientMoveStartStrafeRight: func() system.ClientPacket { return packet.NewClientMovePacket(static.OpCodeClientMoveStartStrafeRight) },
		static.OpCodeClientMoveStartTurnLeft:    func() system.ClientPacket { return packet.NewClientMovePacket(static.OpCodeClientMoveStartTurnLeft) },
		static.OpCodeClientMoveStartTurnRight:   func() system.ClientPacket { return packet.NewClientMovePacket(static.OpCodeClientMoveStartTurnRight) },
		static.OpCodeClientMoveStop:             func() system.ClientPacket { return packet.NewClientMovePacket(static.OpCodeClientMoveStop) },
		static.OpCodeClientMoveStopStrafe:       func() system.ClientPacket { return packet.NewClientMovePacket(static.OpCodeClientMoveStopStrafe) },
		static.OpCodeClientMoveStopTurn:         func() system.ClientPacket { return packet.NewClientMovePacket(static.OpCodeClientMoveStopTurn) },
	}
)
