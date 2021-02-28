package handlers

import "net/http"

// NotFound customizes not found message
func NotFound(w http.ResponseWriter, r *http.Request) {

	// w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotFound)

	// http.ServeFile(w, r, "./static/404.html")
}
