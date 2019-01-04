package router

import (
	"net/http"
)

// New creates a new router and attaches routes
func New() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/data", basicHandler)

	return mux
}

func basicHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	w.Write([]byte(`{"message": "hey"}`))
}
