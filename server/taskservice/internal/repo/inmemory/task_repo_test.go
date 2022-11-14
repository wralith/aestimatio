package inmemory

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"github.com/wralith/aestimatio/server/taskservice/internal/core/domain"
)

var repo *InMemoryTaskRepo

func TestMain(m *testing.M) {
	repo = NewInMemoryTaskRepo().(*InMemoryTaskRepo)
	m.Run()
}

func TestNewInMemoryTaskRepo_Get(t *testing.T) {
	userId := uuid.New()
	title, desc := "Test", "Test Desc"
	created, _ := repo.Create(userId, title, desc)

	v, err := repo.Get(created.ID)
	require.NoError(t, err)
	require.Equal(t, title, v.Title)
	require.Equal(t, desc, v.Description)

	_, err = repo.Get(uuid.New())
	require.Error(t, err)
}

func TestNewInMemoryTaskRepo_Create(t *testing.T) {
	id := uuid.New()
	title, desc := "Test", "Test Desc"
	v, err := repo.Create(id, title, desc)

	require.NoError(t, err)
	require.Equal(t, title, v.Title)
	require.Equal(t, desc, v.Description)

	_, err = repo.Create(id, "", "") // Invalid
	require.Error(t, err)
}

func TestNewInMemoryTaskRepo_Update(t *testing.T) {
	userId := uuid.New()
	title, desc := "Test", "Test Desc"
	title2, desc2 := "Test2", "Test Desc2"
	v, _ := repo.Create(userId, title, desc)
	v.UpdateTitle(title2)
	v.UpdateDescription(desc2)

	v2, err := repo.Update(v)

	require.NoError(t, err)
	require.Equal(t, title2, v2.Title)
	require.Equal(t, desc2, v2.Description)

	v3, _ := domain.CreateTask(userId, "Test", "Test Desc")
	_, err = repo.Update(v3)
	require.Error(t, err)
}

func TestNewInMemoryTaskRepo_Delete(t *testing.T) {
	id := uuid.New()
	title, desc := "Test", "Test Desc"
	v, _ := repo.Create(id, title, desc)

	require.NotNil(t, repo.db[v.ID])
	err := repo.Delete(v.ID)
	require.NoError(t, err)
	require.Nil(t, repo.db[v.ID])

	err = repo.Delete(uuid.New())
	require.Error(t, err)
}

func TestInMemoryTaskRepo_List(t *testing.T) {
	userId := uuid.New()
	title, desc := "Test", "Test Desc"

	for i := 0; i < 10; i++ {
		repo.Create(userId, title, desc)
	}

	v := repo.List(userId, 5, 0)
	require.Len(t, v, 5)

	v = repo.List(userId, 999, 0) // Max Limit = Len
	require.Len(t, v, 10)

	v = repo.List(userId, 5, 1)
	require.Len(t, v, 5)
}
