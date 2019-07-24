package database

import (
	"database/sql"
	"io/ioutil"
	"log"
	"os"
	"path"

	"github.com/jmoiron/sqlx"

	// Import the SQL driver.
	_ "github.com/mattn/go-sqlite3"
)

var (
	db    *sql.DB
	dbDir string
)

func init() {
	dbDir, err := ioutil.TempDir("", "wowserver")
	if err != nil {
		panic(err)
	}

	dbFilepath := path.Join(dbDir, "wowserver.db")
	db, err = sql.Open("sqlite3", dbFilepath)
	if err != nil {
		panic(err)
	}

	log.Printf("Created database file %s\n", dbFilepath)

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	// Make the tables in the database.
	_, err = sqlx.LoadFile(db, "database/sql/tables.sql")
	if err != nil {
		panic(err)
	}
}

// Cleanup will destroy the generated files and close any open database
// connections. This doesn't _have_ to be called, but it should be if possible.
func Cleanup() {
	db.Close()
	os.RemoveAll(dbDir)
}
