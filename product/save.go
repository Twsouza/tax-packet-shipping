package product

import (
	"log"

	"github.com/twsouza/tax-packet-shipping/dao"
)

// Save into database
func (p *Payload) Save() (interface{}, error) {
	insertID, err := dao.Insert(collection, p)
	if err != nil {
		log.Println("Error saving product payload", err)
		return nil, err
	}

	return insertID, nil
}
