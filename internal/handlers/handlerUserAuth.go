package handlers

import (
	"net/http"
)

// UserAuth handles user authentication
func UserAuth(w http.ResponseWriter, s *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

}
