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

	return nil
}
