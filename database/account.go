package database

import (
	"math/big"
	"strings"

	"gitlab.com/jeshuamorrissey/mmo_server/authserver/srp"
)

const (
	accountGet    = "SELECT id, name, verifier, salt FROM Account WHERE name = ?"
	accountInsert = "INSERT INTO Account(name, verifier, salt) VALUES(?, ?, ?)"
	accountUpdate = "UPDATE Account(name, verifier, salt) VALUES(?, ?, ?) WHERE id = ?"
)

// Account represents a user's account, which contains many characters.
type Account struct {
	ID       int64
	Name     string
	Verifier big.Int
	Salt     big.Int
}

// NewAccount takes as input a account name + password and generates an Account
// object.
func NewAccount(name, password string) *Account {
	account := new(Account)

	account.ID = -1
	account.Name = strings.ToUpper(name)
	account.Salt.Set(srp.GenerateSalt())
	account.Verifier.Set(srp.GenerateVerifier(
		account.Name,
		strings.ToUpper(password),
		&account.Salt))

	return account
}

// GetAccount retrieves an account from the database given it's name.
func GetAccount(name string) (*Account, error) {
	stmt, err := db.Prepare(accountGet)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	account := new(Account)

	var verifier, salt string
	err = stmt.QueryRow(name).Scan(&account.ID, &account.Name, &verifier, &salt)
	if err != nil {
		return nil, err
	}

	account.Verifier.SetString(verifier, 16)
	account.Salt.SetString(salt, 16)

	return account, nil
}

// Save will write the account to the database. This will also insert it if it
// doesn't already exist.
func (a *Account) Save() error {
	if a.ID < 0 {
		stmt, err := db.Prepare(accountInsert)
		if err != nil {
			return err
		}

		result, err := stmt.Exec(a.Name, a.Verifier.Text(16), a.Salt.Text(16))
		if err != nil {
			return err
		}

		a.ID, err = result.LastInsertId()
		if err != nil {
			return err
		}
	} else {
		stmt, err := db.Prepare(accountUpdate)
		if err != nil {
			return err
		}

		_, err = stmt.Exec(a.Name, a.Verifier.Text(16), a.Salt.Text(16), a.ID)
		if err != nil {
			return err
		}
	}

	return nil
}
