package auth

import (
	"net/http"
	"strings"
	"time"

	// "github.com/go-jose/go-jose/v3"

	"github.com/sapiderman/seed-go/internal/config"
	log "github.com/sirupsen/logrus"
)

var (
	jwtLog = log.WithField("go", "jwtModule")
)

// JwtMiddleware validates token in transit
func JwtMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// Bypass /health, /docs
		if strings.HasPrefix(r.URL.Path, "/health") {
			next.ServeHTTP(w, r)
		}

		next.ServeHTTP(w, r)
	})
}

// GenerateTokens generates the access and refresh tokens
func GenerateTokens() (string, string, error) {

	key := config.Get("jwt.key")
	sig, err := jwt.NewSigner(jwt.SigningKey{Algorithm: jwt.HS256, Key: key}, (&jwt.SignerOptions{}).WithType("JWT"))
	if err != nil {
		panic(err)
	}

	cl := jwt.Claims{
		Subject:   "subject",
		Issuer:    "github.com/sapiderman/seed-go",
		NotBefore: jwt.NewNumericDate(time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)),
		Audience:  jwt.Audience{"leela", "fry"},
		IssuedAt:  time.Now().Unix(),
		ID:        "1234567890-ABCDEFGHIJKLMNOP",
		Expired:   time.Now().Add(time.Hour * 24).Unix(),
		Type:      "ACCESS",
	}
	accessToken, err := jwt.Signed(sig).Claims(cl).CompactSerialize()
	if err != nil {
		panic(err)
	}

	rfClaim := jwt.Config{
		Expired: time.Now().Add(time.Hour * 24 * 7).Unix(),
		Issuer:  "github.com/sapiderman/seed-go",
		Type:    "REFRESH",
		Subject: "1",
	}

	refreshToken, err := jwt.Signed(sig).Claims(rfClaim)
	if err != nil {
		panic(err)
	}

	return accessToken, refreshToken, nil
}
