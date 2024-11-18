package response

import (
	"encoding/json"
	"net/http"
)

func ResponseJSON(w http.ResponseWriter, data interface{}, statusCode int) {

	w.Header().Set("Contet-Type", "application/json")
	w.WriteHeader(statusCode)

	json.NewEncoder(w).Encode(data)
}
