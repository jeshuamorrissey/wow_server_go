package packet_test

import (
	"testing"

	"gotest.tools/assert"

	"github.com/jinzhu/gorm"
	"gitlab.com/jeshuamorrissey/mmo_server/authserver/packet"
	"gitlab.com/jeshuamorrissey/mmo_server/authserver/srp"
	"gitlab.com/jeshuamorrissey/mmo_server/database"

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

func TestHandle(t *testing.T) {
	db := makeFakeDB()

	// Make some fake data.
	db.Create(&database.Account{
		Name:        "TEST",
		VerifierStr: "10",
		SaltStr:     "20",
	})

	// Create some fake state.
	pkt := packet.ClientLoginChallenge{
		AccountName: []byte("TEST"),
	}

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
	assert.Equal(t, response.Error, uint8(0))
	assert.Assert(t, response.B.String() != "0")
	assert.Equal(t, response.B.Text(16), publicEphemeralExp.Text(16))
	assert.Equal(t, response.Salt.Text(16), "20")
}
