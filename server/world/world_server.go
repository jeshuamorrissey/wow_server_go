package world

import (
	"net"
	"strconv"

	"github.com/jeshuamorrissey/wow_server_go/lib/config"
	"github.com/jeshuamorrissey/wow_server_go/server/world/data/dynamic/interfaces"
	"github.com/jeshuamorrissey/wow_server_go/server/world/data/static"
	"github.com/jeshuamorrissey/wow_server_go/server/world/system"

	"github.com/sirupsen/logrus"

	"github.com/jeshuamorrissey/wow_server_go/server/world/packet"
)

var opCodeToPacket = map[static.OpCode]func() interfaces.ClientPacket{
	static.OpCodeClientAuthSession:       func() interfaces.ClientPacket { return new(packet.ClientAuthSession) },
	static.OpCodeClientCharCreate:        func() interfaces.ClientPacket { return new(packet.ClientCharCreate) },
	static.OpCodeClientCharDelete:        func() interfaces.ClientPacket { return new(packet.ClientCharDelete) },
	static.OpCodeClientCharEnum:          func() interfaces.ClientPacket { return new(packet.ClientCharEnum) },
	static.OpCodeClientCreatureQuery:     func() interfaces.ClientPacket { return new(packet.ClientCreatureQuery) },
	static.OpCodeClientItemQuerySingle:   func() interfaces.ClientPacket { return new(packet.ClientItemQuerySingle) },
	static.OpCodeClientLogoutRequest:     func() interfaces.ClientPacket { return new(packet.ClientLogoutRequest) },
	static.OpCodeClientNameQuery:         func() interfaces.ClientPacket { return new(packet.ClientNameQuery) },
	static.OpCodeClientPing:              func() interfaces.ClientPacket { return new(packet.ClientPing) },
	static.OpCodeClientPlayerLogin:       func() interfaces.ClientPacket { return new(packet.ClientPlayerLogin) },
	static.OpCodeClientQueryTime:         func() interfaces.ClientPacket { return new(packet.ClientQueryTime) },
	static.OpCodeClientSetActiveMover:    func() interfaces.ClientPacket { return new(packet.ClientSetActiveMover) },
	static.OpCodeClientStandstatechange:  func() interfaces.ClientPacket { return new(packet.ClientStandStateChange) },
	static.OpCodeClientTutorialFlag:      func() interfaces.ClientPacket { return new(packet.ClientTutorialFlag) },
	static.OpCodeClientUpdateAccountData: func() interfaces.ClientPacket { return new(packet.ClientUpdateAccountData) },
	static.OpCodeClientAttackswing:       func() interfaces.ClientPacket { return new(packet.ClientAttackSwing) },
	static.OpCodeClientAttackstop:        func() interfaces.ClientPacket { return new(packet.ClientAttackStop) },

	// Movement packets have the same receiver.
	static.OpCodeClientMoveHeartbeat: func() interfaces.ClientPacket { return packet.NewClientMovePacket(static.OpCodeClientMoveHeartbeat) },
	static.OpCodeClientMoveSetFacing: func() interfaces.ClientPacket { return packet.NewClientMovePacket(static.OpCodeClientMoveSetFacing) },
	static.OpCodeClientMoveStartBackward: func() interfaces.ClientPacket {
		return packet.NewClientMovePacket(static.OpCodeClientMoveStartBackward)
	},
	static.OpCodeClientMoveStartForward: func() interfaces.ClientPacket { return packet.NewClientMovePacket(static.OpCodeClientMoveStartForward) },
	static.OpCodeClientMoveStartStrafeLeft: func() interfaces.ClientPacket {
		return packet.NewClientMovePacket(static.OpCodeClientMoveStartStrafeLeft)
	},
	static.OpCodeClientMoveStartStrafeRight: func() interfaces.ClientPacket {
		return packet.NewClientMovePacket(static.OpCodeClientMoveStartStrafeRight)
	},
	static.OpCodeClientMoveStartTurnLeft: func() interfaces.ClientPacket {
		return packet.NewClientMovePacket(static.OpCodeClientMoveStartTurnLeft)
	},
	static.OpCodeClientMoveStartTurnRight: func() interfaces.ClientPacket {
		return packet.NewClientMovePacket(static.OpCodeClientMoveStartTurnRight)
	},
	static.OpCodeClientMoveStop:       func() interfaces.ClientPacket { return packet.NewClientMovePacket(static.OpCodeClientMoveStop) },
	static.OpCodeClientMoveStopStrafe: func() interfaces.ClientPacket { return packet.NewClientMovePacket(static.OpCodeClientMoveStopStrafe) },
	static.OpCodeClientMoveStopTurn:   func() interfaces.ClientPacket { return packet.NewClientMovePacket(static.OpCodeClientMoveStopTurn) },
}

func setupSession(sess *system.Session) {
	pkt := packet.ServerAuthChallenge{Seed: 0}
	sess.Send(&pkt)
}

// RunWorldServer takes as input a database and runs an world server referencing
// it.
func RunWorldServer(realmName string, port int, config *config.Config) {
	log := logrus.WithFields(logrus.Fields{"server": "WORLD", "port": port})
	log.Logger.SetLevel(logrus.TraceLevel)

	// Start updater.
	updater := system.NewUpdater(log, config.ObjectManager)
	go updater.Run()

	// Start session handler.
	listener, err := net.Listen("tcp", "127.0.0.1:"+strconv.Itoa(port))
	if err != nil {
		log.Fatalf("Error while opening port: %v\n", err)
	}

	log.Infof("Listening for WORLD connections on :%v...", listener.Addr().String())

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatalf("Error while receiving client connection: %v\n", err)
		}

		log.Printf("Receiving WORLD connection from %v\n", conn.RemoteAddr())
		sess := system.NewSession(
			conn,
			conn,
			opCodeToPacket,
			opCodeToHandler,
			config,
			logrus.WithFields(logrus.Fields{"server": "WORLD", "account": "???"}),
			updater,
		)
		setupSession(sess)
		go sess.Run()
	}
}
