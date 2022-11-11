package tkn

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/wralith/aestimatio/server/authservice/internal/core/port"
)

var (
	ErrSignatureInvalid = "invalid signaute"
)

type TokenAdaptor struct {
	key []byte
}

func New(key []byte) port.Token {
	return &TokenAdaptor{key: key}
}

func (a *TokenAdaptor) GenerateJWT(id string, email string) (string, error) {
	claims := jwt.MapClaims{}
	claims["exp"] = time.Now().Add(60 * time.Minute).Unix()
	claims["iat"] = time.Now().Unix()
	claims["sub"] = id
	claims["iss"] = "aestimatio"
	claims["email"] = email

	tkn := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tknString, err := tkn.SignedString(a.key)
	if err != nil {
		return "", err
	}

	return tknString, nil
}

func (a *TokenAdaptor) VerifyJWT(tkn string) bool {
	token, err := jwt.Parse(tkn, func(t *jwt.Token) (interface{}, error) {
		return a.key, nil
	})

	if err != nil || !token.Valid {
		return false
	}

	return true
}
