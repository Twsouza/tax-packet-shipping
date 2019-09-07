package response

import (
	"encoding/json"
	"log"
	"net/http"
)

// InvalidJSON is a response for bad or incorrect json
func InvalidJSON(w http.ResponseWriter) {
	r := &defaultResponse{
		Status:  http.StatusUnprocessableEntity,
		Message: "Invalid JSON",
	}

	json, err := json.Marshal(r)
	if err != nil {
		log.Println("Error on json marshal", err)
		InternalServerError(w)
		return
	}

	// set to json
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	// close the response
	w.WriteHeader(r.Status)
	w.Write(json)
}
