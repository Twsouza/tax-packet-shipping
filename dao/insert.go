package dao

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
)

// Insert a doc into a collection
func Insert(collection string, doc interface{}) (interface{}, error) {
	// connect it to db
	client, err := connect()
	db := client.Database(config.DbParams.Name)
	defer client.Disconnect(context.Background())

	if err != nil {
		return nil, err
	}

	// convert it to BSON, to be saved in the db
	bDoc, err := bson.Marshal(doc)
	if err != nil {
		return nil, err
	}

	// insert into the database
	res, err := db.Collection(collection).InsertOne(context.Background(), bDoc)
	return res.InsertedID, err
}
