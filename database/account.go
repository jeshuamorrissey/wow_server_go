package database

import (
	"context"
	"strings"

	"gitlab.com/jeshuamorrissey/mmo_server/auth_server/srp"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	// CollectionAccount is the name of the Account collection.
	CollectionAccount = "account"
)

// Account represents a user's account, which contains many characters.
type Account struct {
	Name     string
	Verifier string
	Salt     string
}

// NewAccount takes as input a account name + password and generates an Account
// object.
func NewAccount(name, password string) *Account {
	salt := srp.GenerateSalt()
	verifier := srp.GenerateVerifier(
		strings.ToUpper(name),
		strings.ToUpper(password),
		salt)

	account := new(Account)
	account.Name = strings.ToUpper(name)
	account.Verifier = verifier.Text(16)
	account.Salt = salt.Text(16)

	return account
}

// GetAccount will fetch the account with a certain name from the database.
func GetAccount(db *mongo.Database, name string) (*Account, error) {
	ctx := context.TODO()
	account := new(Account)
	err := db.Collection(CollectionAccount).FindOne(
		ctx, bson.M{"name": name}).Decode(account)
	if err != nil {
		return nil, err
	}

	return account, nil
}
