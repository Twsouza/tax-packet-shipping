package response

import (
	"encoding/json"
	"log"
	"net/http"
)

// Sucessful convert a string to a json and return to user with BODY
func Sucessful(status int, body interface{}, w http.ResponseWriter) {
	json, err := json.Marshal(body)
	// in case of someting get wrong
	if err != nil {
		log.Println("Erro json marshal on response sucessful", err)
		InternalServerError(w)
		return
	}

	// set to json
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	// close response
	w.WriteHeader(status)
	w.Write(json)
}
