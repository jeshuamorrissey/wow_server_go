package main

import (
	"sync"

	"github.com/jeshuamorrissey/wow_server_go/authserver"
	"github.com/jeshuamorrissey/wow_server_go/authserver/srp"
	"github.com/jeshuamorrissey/wow_server_go/common/data"
	c "github.com/jeshuamorrissey/wow_server_go/common/data/constants"
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
	realmSydney := database.Realm{Name: "Sydney", Host: "localhost:5001"}
	err = db.Create(&realmSydney).Error
	if err != nil {
		return err
	}

	// Make a character.
	equipment := []*database.EquippedItem{}
	for slot, item := range data.GetStartingEquipment(c.ClassWarrior, c.RaceHuman) {
		equipment = append(equipment, &database.EquippedItem{
			Slot: slot,
			Item: &database.GameObjectItem{
				GameObjectBase: database.GameObjectBase{
					Entry: item.Entry,
				},
			},
		})
	}

	inventory := []*database.BaggedItem{}
	for i, item := range data.GetStartingItems(c.ClassWarrior, c.RaceHuman) {
		inventory = append(inventory, &database.BaggedItem{
			Slot: i,
			Item: &database.GameObjectItem{
				GameObjectBase: database.GameObjectBase{
					Entry: item.Entry,
				},
			},
		})
	}

	charJeshua := database.Character{
		Name: "Jeshua",
		Object: database.GameObjectPlayer{
			GameObjectUnit: database.GameObjectUnit{
				Race:   c.RaceHuman,
				Class:  c.ClassWarrior,
				Gender: c.GenderMale,

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

			Equipment: equipment,
			Inventory: inventory,
			Bags:      []*database.GameObjectContainer{},
		},
		AccountID: account.ID,
		RealmID:   realmSydney.ID,
	}

	db.Create(&charJeshua)

	return nil
}

func main() {
	logrus.SetFormatter(&logrus.TextFormatter{ForceColors: true})
	logrus.SetLevel(logrus.TraceLevel)

	// Load constant data.
	logrus.Info("Loading items.json.gz...")
	err := data.LoadItems("D:\\Users\\Jeshua\\go\\src\\github.com\\jeshuamorrissey\\wow_server_go\\common\\data\\items.json.gz")
	if err != nil {
		panic(err)
	}

	logrus.Infof("Done! Loaded %v items.", len(data.Items))

	logrus.Infof("%v", data.Items[25])

	// Setup test database.
	db, err := gorm.Open("sqlite3", ":memory:")
	if err != nil {
		panic(err)
	}

	db = db.Set("gorm:auto_preload", true)

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
		worldserver.RunWorldServer("Sydney", 5001, db)
	}()

	wg.Wait()
}
