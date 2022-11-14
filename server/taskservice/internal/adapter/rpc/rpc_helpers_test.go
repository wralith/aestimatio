package rpc

import (
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func Test_getSubFromJWT(t *testing.T) {
	id := uuid.New()
	claims := jwt.MapClaims{}
	claims["sub"] = id
	claims["exp"] = time.Now().Add(60 * time.Minute).Unix()
	claims["iss"] = "aestimatio"
	claims["iat"] = time.Now().Unix()

	tkn := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	str, err := tkn.SignedString([]byte("test-secret"))
	require.NoError(t, err)
	got, err := getSubFromJWT(str)
	require.NoError(t, err)
	require.Equal(t, id, got)
}
