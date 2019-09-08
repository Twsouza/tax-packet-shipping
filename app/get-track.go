package app

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/twsouza/tax-packet-shipping/common/response"
	"github.com/twsouza/tax-packet-shipping/product"
	"go.mongodb.org/mongo-driver/mongo"
)

// GetTrack endpoint to calc and show a tax from a product
func GetTrack(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	if id == "" {
		response.Unsucessful(http.StatusBadRequest, "Track ID is required", nil, w)
		return
	}

	// convert payload to store struct
	p := &product.Payload{}
	if err := p.GetTrack(id); err != nil {
		log.Println("Err on get track", err)
		if err.Error() == mongo.ErrNoDocuments.Error() {
			response.Unsucessful(http.StatusBadRequest, "Track not found", nil, w)
		} else if err.Error() == "Invalid ID" {
			response.Unsucessful(http.StatusBadRequest, "Invalid ID", nil, w)
		} else {
			response.InternalServerError(w)
		}
		return
	}
	response.Sucessful(http.StatusOK, p, w)
}
