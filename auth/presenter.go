package auth

import (
	"net/http"

	"github.com/ilhamrobyana/efishery-task/entity"
	"github.com/ilhamrobyana/efishery-task/storage"
	"github.com/labstack/echo"
)

var coreInstance *core

func Register(c echo.Context) error {
	r := new(entity.RegisterRequest)
	if err := c.Bind(r); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{"message": "Invalid request data"})
	}

	if err := checkRegister(*r); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{"message": err.Error()})
	}

	user := new(entity.User)
	user.Phone = r.Phone
	user.Name = r.Name
	user.Role = r.Role

	authCore := getCore()
	createdUser, e := authCore.register(*user)

	if e != nil {
		httpStatus := http.StatusInternalServerError
		return c.JSON(httpStatus, map[string]interface{}{"message": e.Error})
	}

	return c.JSON(http.StatusCreated, createdUser)
}

func getCore() (c *core) {
	c = coreInstance

	if c == nil {
		c = new(core)
		userStorage, _ := storage.GetUserStorage(storage.Postgre)

		c.userStorage = userStorage
		coreInstance = c
	}

	return
}
