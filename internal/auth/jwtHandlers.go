package auth

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gofrs/uuid"
	"github.com/golang-jwt/jwt"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
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
	atKey := viper.GetString("jwt.accessKey")
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
	rtKey := viper.GetString("jwt.refreshKey")
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

// VerifyAccessToken ..
func VerifyAccessToken(r *http.Request) (*jwt.Token, error) {
	tokenString := ExtractToken(r)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		//Make sure that the token method conform to "SigningMethodHMAC"
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(viper.GetString("jwt.accessKey")), nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}

// TokenValid ...
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

// RefreshTokens ...
func RefreshTokens(refreshToken string) (string, string, error) {
	logf := jwtLog.WithField("fn", "RefreshTokens")

	//verify the token
	rtKey := viper.GetString("jwt.refreshKey")
	token, err := jwt.Parse(refreshToken, func(token *jwt.Token) (interface{}, error) {
		//Make sure that the token method conform to "SigningMethodHMAC"
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			logf.Error("unexpeted signing key")
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(rtKey), nil
	})
	//if there is an error, the token must have expired
	if err != nil {
		logf.Error(err)
		return "", "", err
	}
	//is token valid?
	if _, ok := token.Claims.(jwt.Claims); !ok && !token.Valid {
		logf.Error(err)
		return "", "", err
	}
	//Since token is valid, get the uuid:
	claims, ok := token.Claims.(jwt.MapClaims) //the token claims should conform to MapClaims
	if ok && token.Valid {
		subject, ok := claims["sub"].(string) //convert the interface to string
		if !ok {
			logf.Error(err)
			return "", "", err
		}

		// userId, err := strconv.ParseUint(fmt.Sprintf("%.f", claims["user_id"]), 10, 64)
		// if err != nil {
		// 	c.JSON(http.StatusUnprocessableEntity, "Error occurred")
		// 	return
		// }
		//Delete the previous Refresh Token
		// deleted, delErr := DeleteAuth(refreshUuid)
		// if delErr != nil || deleted == 0 { //if any goes wrong
		// 	logf.Error(delErr)
		// 	return "", "", delErr
		// }
		//Create new pairs of refresh and access tokens
		at, _, createErr := GenerateTokenPair(subject)
		if err != nil {
			logf.Error(createErr)
			return "", "", createErr
		}
		//save the tokens metadata to redis
		// saveErr := CreateAuth(userId, ts)
		// if saveErr != nil {
		// 	c.JSON(http.StatusForbidden, saveErr.Error())
		// 	return
		// }

		return at, refreshToken, nil
	}
	logf.Error("refresh expired")
	return "", "", nil
}
