package domain

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestUser_Equal(t *testing.T) {
	id := uuid.New()
	id2 := uuid.New()
	u1 := User{ID: id}
	u2 := User{ID: id}
	u3 := User{ID: id2}

	assert.True(t, u1.Equal(&u2))
	assert.True(t, u2.Equal(&u1))
	assert.False(t, u3.Equal(&u1))
}

func TestNewUser(t *testing.T) {
	type args struct {
		username string
		password string
		email    string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Happy",
			args: args{
				username: "Test",
				password: "1234567",
				email:    "test@mail.com",
			},
			wantErr: false,
		},
		{
			name: "Invalid Mail",
			args: args{
				username: "Test",
				password: "1234567",
				email:    "invalid",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewUser(tt.args.username, tt.args.password, tt.args.email)
			if tt.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				require.Equal(t, tt.args.username, got.Username)
				require.Equal(t, tt.args.password, got.Password)
				require.Equal(t, tt.args.email, got.Email)
			}
		})
	}
}
