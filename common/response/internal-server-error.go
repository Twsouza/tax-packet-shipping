package response

import (
	"encoding/json"
	"net/http"
)

// InternalServerError returns error 500 and log error
func InternalServerError(w http.ResponseWriter) {
	// some generic response
	var r defaultResponse
	r.Status = http.StatusInternalServerError
	r.Message = "Something went wrong"
	json, _ := json.Marshal(r)

	// set to json
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	// close the response
	w.WriteHeader(r.Status)
	w.Write(json)
}
