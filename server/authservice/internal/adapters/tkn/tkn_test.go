package tkn

import (
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"github.com/wralith/aestimatio/server/authservice/internal/core/port"
)

var adaptor port.Token

func TestMain(m *testing.M) {
	adaptor = New([]byte("secret"))
	m.Run()
}

func TestGenerateJWT(t *testing.T) {
	id, email := uuid.New(), "test@mail.com"
	token, err := adaptor.GenerateJWT(id.String(), email)
	require.NoError(t, err)

	isValid := adaptor.VerifyJWT(token)
	require.True(t, isValid)
}

func TestVerifyJWT(t *testing.T) {
	key := []byte("wrong")

	// Wrong key
	tkn := jwt.New(jwt.SigningMethodHS256)
	tknString, _ := tkn.SignedString(key)
	isValid := adaptor.VerifyJWT(tknString)
	require.False(t, isValid)

	// Happy
	claims := jwt.MapClaims{}
	key = []byte("secret")
	tkn = jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tknString, _ = tkn.SignedString(key)
	isValid = adaptor.VerifyJWT(tknString)
	require.True(t, isValid)

	// Expired
	claims["exp"] = time.Now().Add(-60 * time.Minute)
	tkn = jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tknString, _ = tkn.SignedString(key)
	isValid = adaptor.VerifyJWT(tknString)
	require.False(t, isValid)
}
