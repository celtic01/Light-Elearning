package tests

import (
	"testing"

	model "github.com/celtic01/Light-Elearning/api/models"
	"github.com/stretchr/testify/require"
)

func TestCreateUser(t *testing.T) {
	err := refreshUserTable()
	if err != nil {
		t.Errorf("Error refreshing user table %v\n", err)
	}
	user := model.User{
		Username: "testuser",
		Email:    "what@email.com",
		Password: "password",
	}
	savedUser, err := user.SaveUser(DB)
	if err != nil {
		t.Errorf("Error creating user %v\n", err)
		return
	}
	require.Equal(t, user.ID, uint(0x1))
	require.Equal(t, user.Username, savedUser.Username)
	require.Equal(t, user.Email, savedUser.Email)
}

func TestGetUser(t *testing.T) {
	err := refreshUserTable()
	if err != nil {
		t.Errorf("Error refreshing user table %v\n", err)
	}
	user, err := seedOneUser()
	if err != nil {
		t.Errorf("Error seeding user %v\n", err)
	}
	dbUser, err := user.GetUser(DB, user.ID)
	if err != nil {
		t.Errorf("Error getting user %v\n", err)
	}
	require.Equal(t, dbUser.ID, user.ID)
	require.Equal(t, dbUser.Username, user.Username)
	require.Equal(t, dbUser.Email, user.Email)
	require.Equal(t, dbUser.Password, user.Password)
}
