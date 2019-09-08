package response

import (
	"encoding/json"
	"log"
	"net/http"
)

// Unsucessful return a response with message and parameter, log err (if not nil)
// Status must be a valid httpStatus
func Unsucessful(status int, message interface{}, parameter interface{}, w http.ResponseWriter) {
	r := &defaultResponse{
		Status:    status,
		Message:   message,
		Parameter: parameter,
	}

	json, err := json.Marshal(r)
	// in case of someting get wrong
	if err != nil {
		log.Println("Error on json marshal unsucesfull reponse", err)
		InternalServerError(w)
		return
	}

	// set to json
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	// close response
	w.WriteHeader(r.Status)
	w.Write(json)
}
