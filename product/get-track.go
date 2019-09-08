package product

import (
	"errors"
	"log"

	"github.com/twsouza/tax-packet-shipping/dao"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// GetTrack returns a track from product
func (p *Payload) GetTrack(ID string) error {
	// convert to object ID
	objectID, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		log.Println("Err on s.Get->primitive.ObjectIDFromHex", err)
		return errors.New("Invalid ID")
	}

	filter := bson.D{{"_id", objectID}}

	return dao.FindOne(collection, filter, p)
}
