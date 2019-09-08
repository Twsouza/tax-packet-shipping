package dao

import (
	"context"
	"log"
)

// FindOne find one doc within a collection, given a filter
func FindOne(collection string, filter interface{}, doc interface{}) error {
	client, err := connect()
	db := client.Database(config.DatabaseName)
	defer client.Disconnect(context.Background())

	log.Println("FindOne db connect err", err)
	if err != nil {
		return err
	}

	err = db.Collection(collection).FindOne(context.Background(), filter).Decode(doc)
	log.Println("FindOne db.Collection err:", err)

	return err
}
