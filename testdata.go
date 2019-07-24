package main

import "gitlab.com/jeshuamorrissey/mmo_server/database"

// GenerateTestData will generate all data required to have a reasonable test of the
// game system.
func GenerateTestData() error {
	// Generate some accounts.
	account := database.NewAccount("jeshua", "jeshua")
	err := account.Save()
	if err != nil {
		return err
	}

	// Make some realms they can connect to.
	err = database.NewRealm("Sydney", "localhost:5001").Save()
	if err != nil {
		return err
	}

	err = database.NewRealm("Brisbane", "localhost:5002").Save()
	if err != nil {
		return err
	}

	return nil
}
