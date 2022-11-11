package service

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/wralith/aestimatio/server/authservice/internal/core/port"
	"github.com/wralith/aestimatio/server/authservice/internal/repo/inmemory"
)

var svc port.Service

func TestMain(m *testing.M) {
	repo := inmemory.NewUserRepo()
	svc = New(repo)
	m.Run()
}

func Test_service_Login(t *testing.T) {
	email, password := "test@test.com", "123456"
	_, err := svc.Login(email, password)
	require.Error(t, err)

	_, err = svc.Register("test", password, email)
	require.NoError(t, err)
	user, err := svc.Login(email, password)
	require.NoError(t, err)
	require.Equal(t, user.Email, email)

	password = "wrong-password"
	_, err = svc.Login(email, password)
	require.Equal(t, err, ErrInvalidCredentials)
}

func Test_service_Register(t *testing.T) {
	email, password := "test@test.com", "123456"
	user, err := svc.Register("test", password, email)
	require.NoError(t, err)
	require.Equal(t, user.Email, email)

	email = "invalid"
	_, err = svc.Register("test", password, email)
	require.Error(t, err)

}
