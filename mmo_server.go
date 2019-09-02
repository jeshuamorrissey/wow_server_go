package main

import (
	"sync"

	"github.com/jeshuamorrissey/wow_server_go/worldserver/data/dbc"
	"github.com/jeshuamorrissey/wow_server_go/worldserver/data/object"

	"github.com/jeshuamorrissey/wow_server_go/authserver"
	"github.com/jeshuamorrissey/wow_server_go/authserver/srp"
	"github.com/jeshuamorrissey/wow_server_go/common/database"
	"github.com/jeshuamorrissey/wow_server_go/worldserver"
	c "github.com/jeshuamorrissey/wow_server_go/worldserver/data/dbc/constants"
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"

	// Import the SQL driver.
	_ "github.com/mattn/go-sqlite3"
)

// GenerateTestData will generate all data required to have a reasonable test of the
// game system.
func GenerateTestData(om *object.Manager, db *gorm.DB) error {
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
	charObj, err := database.NewCharacter(
		om,
		&account, &realmSydney,
		"Jeshua",
		c.RaceHuman, c.ClassWarrior, c.GenderMale,
		1, 1, 1, 1, 1)
	if err != nil {
		return err
	}

	db.Create(charObj)

	return nil
}

func main() {
	logrus.SetFormatter(&logrus.TextFormatter{ForceColors: true})
	logrus.SetLevel(logrus.DebugLevel)

	// Load constant data.
	// logrus.Info("Loading units.json.gz...")
	// err := dbc.LoadUnits("D:\\Users\\Jeshua\\go\\src\\github.com\\jeshuamorrissey\\wow_server_go\\worldserver\\data\\dbc\\units.json.gz")
	// if err != nil {
	// 	panic(err)
	// }

	logrus.Info("Loading starting_stats.json...")
	err := dbc.LoadStartingStats("D:\\Users\\Jeshua\\go\\src\\github.com\\jeshuamorrissey\\wow_server_go\\worldserver\\data\\dbc\\starting_stats.json")
	if err != nil {
		panic(err)
	}

	logrus.Info("Loading starting_locations.json...")
	err = dbc.LoadStartingLocations("D:\\Users\\Jeshua\\go\\src\\github.com\\jeshuamorrissey\\wow_server_go\\worldserver\\data\\dbc\\starting_locations.json")
	if err != nil {
		panic(err)
	}

	// Setup object manager.
	om := object.NewManager(logrus.WithField("system", "object_manager"))

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
	err = GenerateTestData(om, db)
	if err != nil {
		logrus.Fatalf("Failed to generate test data: %v\n", err)
	}

	// go om.Run()

	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		authserver.RunAuthServer(5000, db)
	}()

	go func() {
		defer wg.Done()
		worldserver.RunWorldServer("Sydney", 5001, om, db)
	}()

	wg.Wait()
}
