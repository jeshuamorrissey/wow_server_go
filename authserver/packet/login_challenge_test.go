package packet_test

import (
	"testing"

	"gotest.tools/assert"

	"github.com/jinzhu/gorm"
	"github.com/jeshuamorrissey/wow_server_go/authserver/packet"
	"github.com/jeshuamorrissey/wow_server_go/authserver/srp"
	"github.com/jeshuamorrissey/wow_server_go/common/database"

	// Import the SQL driver.
	_ "github.com/mattn/go-sqlite3"
)

func makeFakeDB() *gorm.DB {
	db, err := gorm.Open("sqlite3", ":memory:")
	if err != nil {
		panic(err)
	}

	database.Setup(db)

	return db
}

func makeClientLoginChallengePacket(
	gameName string,
	version [3]uint8,
	build int,
	platform, os, locale string,
	timezoneOffset int,
	ipAddress int,
	accountName string) *packet.ClientLoginChallenge {
	pkt := &packet.ClientLoginChallenge{
		Version:        version,
		Build:          uint16(build),
		TimezoneOffset: uint32(timezoneOffset),
		IPAddress:      uint32(ipAddress),
		AccountName:    []byte(accountName),
	}

	copy(pkt.GameName[:], gameName)
	copy(pkt.Platform[:], platform)
	copy(pkt.OS[:], os)
	copy(pkt.Locale[:], locale)

	return pkt
}

func TestHandle(t *testing.T) {
	db := makeFakeDB()

	// Make some fake data.
	db.Create(&database.Account{
		Name:        "TEST",
		VerifierStr: "10",
		SaltStr:     "20",
	})

	pkt := makeClientLoginChallengePacket(
		packet.SupportedGameName,
		packet.SupportedGameVersion,
		packet.SupportedGameBuild,
		"x86",
		"Win",
		"enUS",
		10,
		0,
		"TEST",
	)

	state := packet.NewState(db)

	// Check the response.
	responses, err := pkt.Handle(state)
	assert.NilError(t, err)
	assert.Equal(t, state.Account.Name, "TEST")
	assert.Equal(t, len(responses), 1)

	publicEphemeralExp := srp.GeneratePublicEphemeral(
		state.Account.Verifier(),
		&state.PrivateEphemeral)
	response := responses[0].(*packet.ServerLoginChallenge)
	assert.Equal(t, response.Error, packet.LoginOK)
	assert.Equal(t, response.B.Text(16), publicEphemeralExp.Text(16))
	assert.Equal(t, response.Salt.Text(16), "20")
}

func TestHandleWithError(t *testing.T) {
	db := makeFakeDB()

	// Make some fake data.
	db.Create(&database.Account{
		Name:        "TEST",
		VerifierStr: "10",
		SaltStr:     "20",
	})

	var tests = []struct {
		gameName      string
		gameVersion   [3]uint8
		gameBuild     int
		accountName   string
		expectedError packet.LoginErrorCode
	}{
		{
			packet.SupportedGameName,
			packet.SupportedGameVersion,
			packet.SupportedGameBuild,
			"TEST",
			packet.LoginOK,
		},
		{
			packet.SupportedGameName,
			packet.SupportedGameVersion,
			packet.SupportedGameBuild,
			"UNKNOWN",
			packet.LoginUnknownAccount,
		},
		{
			"GAME",
			packet.SupportedGameVersion,
			packet.SupportedGameBuild,
			"TEST",
			packet.LoginFailed,
		},
		{
			packet.SupportedGameName,
			[3]uint8{5, 5, 5},
			packet.SupportedGameBuild,
			"TEST",
			packet.LoginBadVersion,
		},
		{
			packet.SupportedGameName,
			packet.SupportedGameVersion,
			1000,
			"TEST",
			packet.LoginBadVersion,
		},
	}

	for _, test := range tests {
		pkt := makeClientLoginChallengePacket(
			test.gameName,
			test.gameVersion,
			test.gameBuild,
			"x86",
			"Win",
			"enUS",
			10,
			0,
			test.accountName,
		)

		state := packet.NewState(db)

		// Check the response.
		responses, err := pkt.Handle(state)
		assert.NilError(t, err)
		assert.Equal(t, len(responses), 1)

		response := responses[0].(*packet.ServerLoginChallenge)
		assert.Equal(t, response.Error, test.expectedError)
	}
}
