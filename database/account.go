package database

import (
	"math/big"

	"github.com/jinzhu/gorm"
)

// Account represents a user's account, which contains many characters.
type Account struct {
	gorm.Model

	Name          string `gorm:"unique"`
	VerifierStr   string
	SaltStr       string
	SessionKeyStr *string

	verifier   *big.Int `gorm:"-"`
	salt       *big.Int `gorm:"-"`
	sessionKey *big.Int `gorm:"-"`
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

// SessionKey gets a big.Int version of the account SessionKey.
func (a *Account) SessionKey() *big.Int {
	if a.SessionKeyStr == nil {
		return nil
	}

	if a.sessionKey == nil {
		a.sessionKey, _ = new(big.Int).SetString(*a.SessionKeyStr, 16)
	}

	return a.sessionKey
}
