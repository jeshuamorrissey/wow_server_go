package world

import (
	"encoding/json"
	"io/ioutil"
	"math/big"
	"os"
	"strings"

	"github.com/jeshuamorrissey/wow_server_go/authserver/srp"
	c "github.com/jeshuamorrissey/wow_server_go/worldserver/data/dbc/constants"
	"github.com/jeshuamorrissey/wow_server_go/worldserver/data/object"
	"github.com/sirupsen/logrus"
)

type CharacterSettings struct {
	HideHelm  bool `json:"hide_helm"`
	HideCloak bool `json:"hide_cloak"`
}

type Character struct {
	Name        string            `json:"name"`
	GUID        object.GUID       `json:"guid"`
	HasLoggedIn bool              `json:"has_logged_in"`
	Settings    CharacterSettings `json:"settings"`
}

// Flags returns an set of flags based on the character's state.
func (char *Character) Flags() uint32 {
	var flags uint32
	if char.Settings.HideHelm {
		flags |= uint32(c.CharacterFlagHideHelm)
	}

	if char.Settings.HideCloak {
		flags |= uint32(c.CharacterFlagHideCloak)
	}

	// if char.Settings.IsGhost {
	// 	flags |= uint32(c.CharacterFlagGhost)
	// }

	return flags
}

type Account struct {
	Name      string     `json:"name"`
	Character *Character `json:"character"`

	// The following fields are required for "authentication".
	SaltStr       string `json:"srp_salt"`
	salt          *big.Int
	VerifierStr   string `json:"srp_verifier"`
	verifier      *big.Int
	SessionKeyStr string `json:"srp_session_key"`
	sessionKey    *big.Int
}

// Verifier gets a big.Int version of the account verifier.
func (a *Account) Verifier() *big.Int {
	if a.verifier == nil {
		a.verifier, _ = new(big.Int).SetString(a.VerifierStr, 16)
	}

	return a.verifier
}

// Salt gets a big.Int version of the account salt.
func (a *Account) Salt() *big.Int {
	if a.salt == nil {
		a.salt, _ = new(big.Int).SetString(a.SaltStr, 16)
	}

	return a.salt
}

// SessionKey gets a big.Int version of the account session key.
func (a *Account) SessionKey() *big.Int {
	if a.sessionKey == nil {
		a.sessionKey, _ = new(big.Int).SetString(a.SessionKeyStr, 16)
	}

	return a.sessionKey
}

type WorldConfig struct {
	Name          string          `json:"name"`
	Accounts      []*Account      `json:"accounts"`
	ObjectManager *object.Manager `json:"objects"`
}

func NewWorldConfig(accountName string) *WorldConfig {
	config := WorldConfig{Name: accountName, ObjectManager: object.NewManager(logrus.WithField("system", "object_manager"))}

	salt := srp.GenerateSalt()
	verifier := srp.GenerateVerifier(strings.ToUpper(accountName), strings.ToUpper(accountName), salt)
	config.Accounts = append(config.Accounts, &Account{
		Name:        strings.ToUpper(accountName),
		SaltStr:     salt.Text(16),
		VerifierStr: verifier.Text(16),
	})

	return &config
}

func NewWorldConfigFrom(filepath string) *WorldConfig {
	config := WorldConfig{
		ObjectManager: object.NewManager(logrus.WithField("system", "object_manager")),
	}

	config.LoadFrom(filepath)
	return &config
}

func (wc *WorldConfig) LoadFrom(filepath string) error {
	file, err := os.OpenFile(filepath, os.O_RDONLY, 0555)
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

	return nil
}

func (wc *WorldConfig) SaveTo(filepath string) error {
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
