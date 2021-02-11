package handlers

import (
	"net/http"

	log "github.com/sirupsen/logrus"
)

// UserAuth handles user authentication
func UserAuth(w http.ResponseWriter, s *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	log.Warn("not impelemted yet..but OK")

	// TODO: implement body

}
