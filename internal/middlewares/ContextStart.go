package middlewares

import (
	"context"
	"net/http"
	"time"

	gonanoid "github.com/matoous/go-nanoid/v2"
	contextkeys "github.com/sapiderman/seed-go/internal/contextKeys"
	"github.com/spf13/viper"
)

// ContextStart begin all requests with a x-request-id
func ContextStart(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		ctx := r.Context()
		if ctx == nil {
			t := time.Duration(viper.GetInt("server.request.timeout"))
			ctxNew, cancelFn := context.WithTimeout(context.Background(), t*time.Second)
			defer cancelFn()
			ctx = ctxNew // this is a bit ugly due to := scope
		}

		reqID := r.Header.Get("x-request-id")
		if len(reqID) == 0 {
			reqID, _ = gonanoid.New()
		}
		newContext := context.WithValue(ctx, contextkeys.XRequestID, reqID)
		next.ServeHTTP(w, r.WithContext(newContext))
	})
}
