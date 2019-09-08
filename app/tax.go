package app

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/twsouza/tax-packet-shipping/common/response"
	"github.com/twsouza/tax-packet-shipping/product"
)

// CalcProdTax endpoint to calc and show a tax from a product
func CalcProdTax(w http.ResponseWriter, r *http.Request) {
	// convert payload to store struct
	p := &product.Payload{}
	if err := json.NewDecoder(r.Body).Decode(p); err != nil {
		log.Println("Incorrect json payload", err)
		response.InvalidJSON(w)
		return
	}

	tax := make(map[string]float64)
	tax["tax"] = p.CalcTax()

	response.Sucessful(http.StatusOK, tax, w)
}
