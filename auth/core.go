package auth

import (
	"errors"

	"github.com/dgrijalva/jwt-go"
	"github.com/ilhamrobyana/efishery-task/entity"
	"github.com/ilhamrobyana/efishery-task/helper"
	"github.com/ilhamrobyana/efishery-task/storage"
)

type core struct {
	userStorage  storage.UserStorage
	tokenStorage storage.TokenStorage
}

func (c *core) register(user entity.User) (createdUser entity.User, e error) {
	createdUser, e = c.userStorage.Create(user)
	return
}

func (c *core) login(phone, password string) (token string, e error) {
	user, e := c.userStorage.GetByPhone(phone)
	if e != nil {
		return
	}

	if password != user.Password {
		e = errors.New("Phone or Password is wrong")
	}

	token, e = helper.GenerateToken(user)
	if e != nil {
		return
	}

	e = c.tokenStorage.Create(token)
	return
}

func (c *core) validateToken(token string) (claims jwt.MapClaims, e error) {
	if !c.tokenStorage.IsAvailable(token) {
		return nil, errors.New("Token is invalid")
	}
	claims, e = helper.ParseToken(token)
	return
}
