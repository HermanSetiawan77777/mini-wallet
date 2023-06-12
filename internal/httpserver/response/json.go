package response

import (
	"encoding/json"
	"net/http"
)

func withJSON(w http.ResponseWriter, httpStatus int, data any) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(httpStatus)
	json.NewEncoder(w).Encode(data)
}
