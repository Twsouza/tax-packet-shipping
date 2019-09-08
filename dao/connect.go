package dao

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// connect to the database
func connect() (*mongo.Client, error) {
	var err error
	client, err := mongo.NewClient(options.Client().ApplyURI(config.DbParams.Host))
	if err != nil {
		log.Println("Err db.connect->mongo.NewClient", err)
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), config.DbParams.TimeOut*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		log.Println("Err db.connect->context.WithTimeout", err)
		return nil, err
	}

	return client, nil
}
