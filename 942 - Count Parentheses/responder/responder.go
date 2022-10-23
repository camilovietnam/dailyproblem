package responder

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func WriteJson(w http.ResponseWriter, status int, i any) {
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(status)
	err := json.NewEncoder(w).Encode(i)
	if err != nil {
		fmt.Printf("[❌] Error: could not encode JSON: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}
