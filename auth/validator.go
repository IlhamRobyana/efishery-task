package auth

import (
	"errors"

	"github.com/ilhamrobyana/efishery-task/entity"
)

func checkRegister(request entity.RegisterRequest) error {
	if len(request.Phone) == 0 {
		return errors.New("Phone must not be empty")
	}

	if len(request.Name) == 0 {
		return errors.New("Name must not be empty")
	}

	if request.Role != "Admin" && request.Role != "User" {
		return errors.New("Invalid role")
	}

	return nil
}
