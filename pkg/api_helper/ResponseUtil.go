package api_helper

import (
	"encoding/json"
	"net/http"
)

func WriteResponse(w http.ResponseWriter, status int, dataObject interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	usersJSON, _ := json.Marshal(dataObject)
	w.Write(usersJSON)
}
