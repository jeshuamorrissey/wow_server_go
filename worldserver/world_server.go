package worldserver

import (
	"io"
	"net"
	"strconv"

	c "github.com/jeshuamorrissey/wow_server_go/worldserver/data/dbc/constants"
	"github.com/jeshuamorrissey/wow_server_go/worldserver/data/object"
	"github.com/jeshuamorrissey/wow_server_go/worldserver/data/world"
	"github.com/jeshuamorrissey/wow_server_go/worldserver/system"

	"github.com/jeshuamorrissey/wow_server_go/worldserver/packet"
	"github.com/sirupsen/logrus"
)

func makeSession(config *world.WorldConfig, reader io.Reader, writer io.Writer, log *logrus.Entry, updater *system.Updater, combatManager *system.CombatManager) *system.Session {
	return system.NewSession(
		reader,
		writer,
		opCodeToPacket,
		config,
		log,
		updater,
		combatManager,
	)
}

func setupSession(sess *system.Session) {
	pkt := packet.ServerAuthChallenge{Seed: 0}
	sess.Send(&pkt)
}

func makeObjectUpdatePacket(outOfRangeUpdate system.OutOfRangeUpdate, objectUpdates []system.ObjectUpdate) system.ServerPacket {
	return &packet.ServerUpdateObject{
		OutOfRangeUpdates: outOfRangeUpdate,
		ObjectUpdates:     objectUpdates,
	}
}

func makeAttackerStateUpdatePacker(attacker object.GUID, target object.GUID, attackInfo object.AttackInfo) system.ServerPacket {
	return &packet.ServerAttackerStateUpdate{
		HitInfo:      c.HitInfoNormalSwing,
		Attacker:     attacker,
		Target:       target,
		Damage:       int32(attackInfo.Damage),
		TargetState:  c.AttackTargetStateHit,
		MeleeSpellID: 0,
	}
}

// RunWorldServer takes as input a database and runs an world server referencing
// it.
func RunWorldServer(realmName string, port int, config *world.WorldConfig) {
	log := logrus.WithFields(logrus.Fields{"server": "WORLD", "port": port})
	log.Logger.SetLevel(logrus.TraceLevel)

	// Start updater.
	updater := system.NewUpdater(log, config.ObjectManager, makeObjectUpdatePacket, makeAttackerStateUpdatePacker)
	go updater.Run()

	// Start the combat manager.
	combatManager := system.NewCombatManager(log, config.ObjectManager, updater)

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
		sessLog := logrus.WithFields(logrus.Fields{"server": "WORLD", "account": "???"})
		sess := makeSession(config, conn, conn, sessLog, updater, combatManager)
		setupSession(sess)
		go sess.Run()
	}
}
