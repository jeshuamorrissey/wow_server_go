package config

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"strings"

	"github.com/jeshuamorrissey/wow_server_go/server/auth/srp"
	"github.com/jeshuamorrissey/wow_server_go/server/world/data/dynamic"
)

type Config struct {
	Name          string                 `json:"name"`
	Accounts      []*Account             `json:"accounts"`
	ObjectManager *dynamic.ObjectManager `json:"objects"`
}

// NewConfig creates a new config object from scratch with the given account name.
func NewConfig(accountName string) *Config {
	config := Config{Name: accountName, ObjectManager: dynamic.GetObjectManager()}

	salt := srp.GenerateSalt()
	verifier := srp.GenerateVerifier(strings.ToUpper(accountName), strings.ToUpper(accountName), salt)
	config.Accounts = append(config.Accounts, &Account{
		Name:        strings.ToUpper(accountName),
		SaltStr:     salt.Text(16),
		VerifierStr: verifier.Text(16),
	})

	return &config
}

// NewConfigFromJSON creates a config object and populates with data from a JSON file.
func NewConfigFromJSON(jsonFilepath string) *Config {
	config := &Config{
		Accounts:      make([]*Account, 0),
		ObjectManager: dynamic.GetObjectManager(),
	}

	config.LoadFromJSON(jsonFilepath)
	return config
}

// LoadFromJSON overwrites the current config file from a JSON source.
func (wc *Config) LoadFromJSON(jsonFilepath string) error {
	file, err := os.OpenFile(jsonFilepath, os.O_RDONLY, 0555)
	if err != nil {
		return err
	}

	data, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}

	err = json.Unmarshal(data, wc)
	if err != nil {
		return err
	}

	// Initialize all units and players.
	for _, player := range wc.ObjectManager.Players {
		player.Initialize()
	}

	for _, unit := range wc.ObjectManager.Units {
		unit.Initialize()
	}

	return nil
}

// SaveToJSON saves the current config state to a JSON file.
func (wc *Config) SaveToJSON(filepath string) error {
	data, err := json.Marshal(wc)
	if err != nil {
		return err
	}

	file, err := os.OpenFile(filepath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0555)
	if err != nil {
		return err
	}

	defer file.Close()

	_, err = file.Write(data)
	if err != nil {
		return err
	}

	return nil
}
