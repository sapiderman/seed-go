package logger

import (
	"net/http"

	log "github.com/sirupsen/logrus"
)

// MyLogger does some stuff
func MyLogger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// Do stuff
		log.WithFields(log.Fields{
			"method": r.Method,
			"path":   r.URL.Path,
			"header": r.Header,
		}).Info("Logger")

		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(w, r)
	})
}
