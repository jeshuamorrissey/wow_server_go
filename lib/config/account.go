package config

import "math/big"

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
