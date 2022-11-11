package domain

import (
	"github.com/google/uuid"
	"github.com/wralith/aestimatio/server/authservice/pkg/vld"
)

type User struct {
	ID       uuid.UUID
	Username string `validate:"required,min=3"`
	Password string `validate:"required,min=6"`
	Email    string `validate:"required,email"`
}

func (u *User) Equal(target *User) bool {
	return u.ID == target.ID
}

func NewUser(username, password, email string) (*User, error) {
	id := uuid.New()
	user := &User{
		ID:       id,
		Username: username,
		Password: password,
		Email:    email,
	}

	validate := vld.GetValidator()
	err := validate.Struct(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}
