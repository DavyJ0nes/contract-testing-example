package storage

import (
	"errors"

	"github.com/davyj0nes/contract-testing-example/api"
)

var errNotFound = errors.New("not found")

// UserStorage defines a simple user store
type UserStorage map[string]*api.User

// NewUserStore initliases and returns a UserStorage
func NewUserStore() UserStorage {
	s := make(UserStorage)
	s["steve"] = &api.User{
		Name:  "steve",
		Email: "steve@steve.com",
		Age:   21,
	}

	return s
}

// GetByName pulls user from storage if it exists
func (s UserStorage) GetByName(name string) (*api.User, error) {
	if user, ok := s[name]; ok {
		return user, nil
	}

	return nil, errNotFound
}
