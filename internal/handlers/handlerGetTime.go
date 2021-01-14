package handlers

import (
	"net/http"

	log "github.com/sirupsen/logrus"
)

// HandlerHello handles /hello calls
func getTime(w http.ResponseWriter, s *http.Request) {

	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)

	resp, err := http.Get("http://example.com/")
	if err != nil {
		log.Error("error found", err)
	}

	_, err = w.Write(resp)
	if err != nil {
		log.Error("resp error: ", err)
	}

}
