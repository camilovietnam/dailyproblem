package respond

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// HTTPResponder manages the http response types and logic
type HTTPResponder struct {
	w http.ResponseWriter
}

func NewHTTPResponder(w http.ResponseWriter) HTTPResponder {
	return HTTPResponder{
		w: w,
	}
}

func (r *HTTPResponder) WriteJSON(status int, i interface{}) {
	r.w.Header().Set("Content-type", "application/json")
	r.w.WriteHeader(status)
	err := json.NewEncoder(r.w).Encode(i)
	if err != nil {
		fmt.Printf("[❌] Error: could not encode JSON: %v", err)
		r.w.WriteHeader(http.StatusInternalServerError)
	}
}
