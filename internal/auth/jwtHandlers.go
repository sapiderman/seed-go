package auth

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofrs/uuid"
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

// GenerateTokenPair generates access and refresh tokens
func GenerateTokenPair(subject string) (string, string, error) {

	//Creating Access Token
	atKey := config.Get("jwt.accessKey")
	atC := jwt.MapClaims{}

	atExpires := time.Now().Add(time.Minute * 15).Unix()
	atjti, _ := uuid.NewV4()

	atC["authorized"] = true
	atC["sub"] = subject
	atC["exp"] = atExpires
	atC["iss"] = "github.com/sapiderman/seed-go"
	atC["nbf"] = time.Now()
	atC["iat"] = time.Now()
	atC["jti"] = atjti.String()

	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atC)
	accessToken, err := at.SignedString([]byte(atKey))
	if err != nil {
		return "", "", err
	}

	//Creating Refresh Token
	rtKey := config.Get("jwt.refreshKey")
	rtExpires := time.Now().Add(time.Hour * 24 * 7).Unix()

	rtC := jwt.MapClaims{}
	rtC["jti"] = atjti
	rtC["sub"] = subject
	rtC["exp"] = rtExpires
	rt := jwt.NewWithClaims(jwt.SigningMethodHS256, rtC)

	refreshToken, err := rt.SignedString([]byte(rtKey))
	if err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, nil
}

// ExtractToken gets
func ExtractToken(r *http.Request) string {
	bearToken := r.Header.Get("Authorization")
	//normally Authorization the_token_xxx
	strArr := strings.Split(bearToken, " ")
	if len(strArr) == 2 {
		return strArr[1]
	}
	return ""
}

// VerifyAccessToken
func VerifyAccessToken(r *http.Request) (*jwt.Token, error) {
	tokenString := ExtractToken(r)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		//Make sure that the token method conform to "SigningMethodHMAC"
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(config.Get("jwt.accessKey")), nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}

func TokenValid(r *http.Request) error {
	token, err := VerifyAccessToken(r)
	if err != nil {
		return err
	}
	if _, ok := token.Claims.(jwt.Claims); !ok && !token.Valid {
		return err
	}
	return nil
}
