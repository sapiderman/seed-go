package middlewares

import (
	"net/http"
	"strings"

	contextkeys "github.com/sapiderman/seed-go/internal/contextKeys"
	log "github.com/sirupsen/logrus"
)

// sanitizeLogString removes newlines and carriage returns from a string to prevent log forging.
func sanitizeLogString(s string) string {
	s = strings.ReplaceAll(s, "\n", "")
	s = strings.ReplaceAll(s, "\r", "")
	return s
}

// sanitizeHeaders returns a sanitized copy of the headers map.
func sanitizeHeaders(headers http.Header) http.Header {
	safeHeaders := make(http.Header, len(headers))
	for k, v := range headers {
		safeVals := make([]string, len(v))
		for i, val := range v {
			safeVals[i] = sanitizeLogString(val)
		}
		safeHeaders[k] = safeVals
	}
	return safeHeaders
}

// MyLogger does some stuff
func MyLogger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if strings.HasPrefix(r.URL.Path, "/docs") {
			log.Info("skipping /docs logging.")
		} else {
			// Do stuff

			ctx := r.Context()
			requestID := ctx.Value(contextkeys.XRequestID)

			safePath := sanitizeLogString(r.URL.Path)
			safeHeader := sanitizeHeaders(r.Header)

			log.WithFields(log.Fields{
				"method":     r.Method,
				"path":       safePath,
				"header":     safeHeader,
				"request-id": requestID.(string),
			}).Debug("Logger")
			//.Info("Logger")
		}
		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(w, r)
	})
}
