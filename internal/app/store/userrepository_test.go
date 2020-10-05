package store_test

import (
	"testing"

	"github.com/Karambirov/http-rest-api/internal/app/model"
	"github.com/stretchr/testify/assert"

	"github.com/Karambirov/http-rest-api/internal/app/store"
)

func TestUserRepository_Create(t *testing.T) {
	s, tearDown := store.TestStore(t, databaseURL)
	defer tearDown("users")

	u, err := s.User().Create(&model.User{
		Email: "user@example.org",
	})

	assert.NoError(t, err)
	assert.NotNil(t, u)
}
