package storage

import (
	"errors"

	"github.com/ilhamrobyana/efishery-task/auth-api/entity"
	pg "github.com/ilhamrobyana/efishery-task/auth-api/pg_storage"
)

type UserStorage interface {
	Create(user entity.User) (entity.User, error)
	GetByPhone(phone string) (entity.User, error)
}

func GetUserStorage(n int) (UserStorage, error) {
	switch n {
	case Postgre:
		return new(pg.User), nil
	default:
		return nil, errors.New("not implemented yet")
	}
}
