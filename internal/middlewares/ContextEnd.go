package middlewares

import (
	"net/http"

	contextkeys "github.com/sapiderman/seed-go/internal/contextKeys"
	log "github.com/sirupsen/logrus"
)

// ContextEnd closes the log with last requestid
func ContextEnd(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		ctx := r.Context()
		requestID := ctx.Value(contextkeys.XRequestID).(string)

		next.ServeHTTP(w, r)
		log.Debug("Done with http.request-id: ", requestID)
	})
}
