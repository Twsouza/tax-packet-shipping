package product

import "go.mongodb.org/mongo-driver/bson/primitive"

// Payload JSON of a product
type Payload struct {
	ID     primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Name   string             `json:"name"`
	Height float64            `json:"height"`
	Length float64            `json:"length"`
	Width  float64            `json:"width"`
	Weight int                `json:"weight"`
	Price  float64            `json:"price"`
	Tax    float64            `json:"tax,omitempty"`
}
