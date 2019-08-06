package main

import (
	"sync"

	"github.com/jeshuamorrissey/wow_server_go/authserver"
	"github.com/jeshuamorrissey/wow_server_go/authserver/srp"
	"github.com/jeshuamorrissey/wow_server_go/common/database"
	"github.com/jeshuamorrissey/wow_server_go/worldserver"
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"

	// Import the SQL driver.
	_ "github.com/mattn/go-sqlite3"
)

// GenerateTestData will generate all data required to have a reasonable test of the
// game system.
func GenerateTestData(db *gorm.DB) error {
	// Generate some accounts.
	salt := srp.GenerateSalt()
	verifier := srp.GenerateVerifier("JESHUA", "JESHUA", salt)
	account := database.Account{Name: "JESHUA", VerifierStr: verifier.Text(16), SaltStr: salt.Text(16)}
	err := db.Create(&account).Error
	if err != nil {
		return err
	}

	// Make some realms they can connect to.
	realm := database.Realm{Name: "Sydney", Host: "localhost:5001"}
	err = db.Create(&realm).Error
	if err != nil {
		return err
	}

	// Make a character.
	err = db.Create(&database.Character{
		Name: "Jeshua",
		Object: *db.Create(&database.GameObjectPlayer{
			GameObjectUnit: database.GameObjectUnit{
				Race:   database.RaceHuman,
				Class:  database.ClassWarrior,
				Gender: database.GenderMale,

				X: 0.0,
				Y: 0.0,
				Z: 0.0,
				O: 0.0,
			},

			Level: 1,

			SkinColor: 1,
			Face:      1,
			HairStyle: 1,
			HairColor: 1,
			Feature:   1,

			ZoneID: 1,
			MapID:  1,
		}).Value.(*database.GameObjectPlayer),
		Account: account,
		Realm:   realm,
	}).Error
	if err != nil {
		return err
	}

	return nil
}

func main() {
	logrus.SetFormatter(&logrus.TextFormatter{ForceColors: true})
	logrus.SetLevel(logrus.TraceLevel)

	db, err := gorm.Open("sqlite3", ":memory:")
	if err != nil {
		panic(err)
	}

	logrus.Infof("Created in-memory database")

	database.Setup(db)

	defer db.Close()

	// Create database testdata.
	err = GenerateTestData(db)
	if err != nil {
		logrus.Fatalf("Failed to generate test data: %v\n", err)
	}

	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		authserver.RunAuthServer(5000, db)
	}()

	go func() {
		defer wg.Done()
		worldserver.RunWorldServer(5001, db)
	}()

	wg.Wait()
}
