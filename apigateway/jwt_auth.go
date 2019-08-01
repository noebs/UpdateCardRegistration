package gateway

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

var SECRETKEY, _ = generateSecretKey(50)

func GenerateJWT(serviceID string) (string, error) {
	// Create a new token object, specifying signing method and the claims
	// you would like it to contain.
	expiresAt := time.Now().Add(time.Minute * 5).UTC().Unix()

	claims := TokenClaims{
		serviceID,
		jwt.StandardClaims{
			ExpiresAt: expiresAt,
			Issuer:    "noebs",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign and get the complete encoded token as a string using the secret
	if tokenString, err := token.SignedString(SECRETKEY); err == nil {
		fmt.Println(tokenString)
		return tokenString, nil
	} else {
		return "", err
	}
}

func VerifyJWT(tokenString string, claims jwt.Claims) (jwt.Claims, error) {

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {

		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		} else {
			return SECRETKEY, nil
		}
	})
	if err != nil {
		return nil, err
	}
	return token.Claims, err
}

type TokenClaims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}
