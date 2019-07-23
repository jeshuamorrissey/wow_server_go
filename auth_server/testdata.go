package main

import (
	"context"
	"time"

	"gitlab.com/jeshuamorrissey/mmo_server/database"
	"go.mongodb.org/mongo-driver/mongo"
)

func clearCollections(ctx context.Context, db *mongo.Database) error {
	err := db.Collection("account").Drop(ctx)
	if err != nil {
		return err
	}

	return nil
}

// GenerateTestData will generate all data required to have a reasonable test of the
// game system. This includes creating accounts, characters, items, monsters,
// quests, ...
func GenerateTestData(db *mongo.Database) error {
	ctx, ctxCancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer ctxCancel()

	err := clearCollections(ctx, db)
	if err != nil {
		return err
	}

	// Generate account data.
	accounts := db.Collection(database.CollectionAccount)
	_, err = accounts.InsertOne(ctx, database.NewAccount("jeshua", "jeshua"))
	if err != nil {
		return nil
	}

	return nil
}
