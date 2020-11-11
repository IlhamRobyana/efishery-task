package auth

import (
	"github.com/ilhamrobyana/efishery-task/entity"
	"github.com/ilhamrobyana/efishery-task/storage"
)

type core struct {
	userStorage storage.UserStorage
}

func (c *core) register(user entity.User) (createdUser entity.User, e error) {
	createdUser, e = c.userStorage.Create(user)
	return
}
