package middlewares

import (
	"net/http"
	"strings"

	contextkeys "github.com/sapiderman/seed-go/internal/contextKeys"
	log "github.com/sirupsen/logrus"
)

// MyLogger does some stuff
func MyLogger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if strings.HasPrefix(r.URL.Path, "/docs") {
			log.Info("skipping /docs logging.")
		} else {
			// Do stuff

			ctx := r.Context()
			requestID := ctx.Value(contextkeys.XRequestID)

			log.WithFields(log.Fields{
				"method":     r.Method,
				"path":       r.URL.Path,
				"header":     r.Header,
				"request-id": requestID.(string),
			}).Debug("Logger")
			//.Info("Logger")
		}
		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(w, r)
	})
}
