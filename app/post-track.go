package app

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/twsouza/tax-packet-shipping/common/response"
	"github.com/twsouza/tax-packet-shipping/product"
)

// TrackProd endpoint to calc and show a tax from a product
func TrackProd(w http.ResponseWriter, r *http.Request) {
	// convert payload to store struct
	p := &product.Payload{}
	if err := json.NewDecoder(r.Body).Decode(p); err != nil {
		log.Println("Incorrect json payload", err)
		response.InvalidJSON(w)
		return
	}

	p.Tax = p.CalcTax()
	insertID, err := p.Save()
	if err != nil {
		log.Println("Err on save product in database", err)
		response.Unsucessful(http.StatusBadRequest, "Unable to track product, try again later", p, w)
		return
	}

	id := make(map[string]interface{})
	id["id"] = insertID
	response.Sucessful(http.StatusCreated, id, w)
}
