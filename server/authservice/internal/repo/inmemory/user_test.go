package inmemory

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"github.com/wralith/aestimatio/server/authservice/internal/repo"
)

var db repo.Repo

func TestMain(m *testing.M) {
	db = NewUserRepo()
	m.Run()
}

func TestRepo_Create(t *testing.T) {
	username, password, email := "testcreate", "1234567", "testcreate@mail.com"
	got, err := db.Create(username, password, email)
	require.NoError(t, err)
	require.Equal(t, got.Username, username)
	// Invalid
	_, err = db.Create(username, password, "")
	require.Error(t, err)
}

func TestRepo_Get(t *testing.T) {
	_, err := db.Get(uuid.New())
	require.Equal(t, err, repo.ErrUserNotFound)

	username, password, email := "testget", "1234567", "testget@mail.com"
	created, _ := db.Create(username, password, email)

	got, err := db.Get(created.ID)
	require.NoError(t, err)
	require.True(t, created.Equal(got))
}

func TestRepo_GetByEmail(t *testing.T) {
	username, password, email := "testgetbyemail", "1234567", "testgetbyemail@mail.com"
	created, _ := db.Create(username, password, email)

	got, err := db.GetByEmail(email)
	require.NoError(t, err)
	require.True(t, created.Equal(got))

	_, err = db.GetByEmail("invalidemail@invaliddomain.com")
	require.Equal(t, err, repo.ErrUserNotFound)
}
