package storage

import (
	"errors"

	redis "github.com/ilhamrobyana/efishery-task/auth-api/redis_storage"
)

type TokenStorage interface {
	IsAvailable(token string) bool
	Create(token string) error
	Update(token string, oldToken string) error
	Delete(token string) error
}

func GetTokenStorage(n int) (TokenStorage, error) {
	switch n {
	case Redis:
		return new(redis.TokenStorage), nil
	default:
		return nil, errors.New("Not implemented yet")
	}
}
