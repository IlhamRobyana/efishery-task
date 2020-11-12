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
	createdUser, err := authCore.register(*user)

	if err != nil {
		httpStatus := http.StatusInternalServerError
		return c.JSON(httpStatus, map[string]interface{}{"message": err.Error})
	}

	return c.JSON(http.StatusCreated, createdUser)
}

func Login(c echo.Context) error {
	r := new(entity.LoginRequest)
	if err := c.Bind(r); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{"message": "Invalid request data"})
	}

	if err := checkLogin(*r); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{"message": err.Error()})
	}

	authCore := getCore()
	token, err := authCore.login(r.Phone, r.Password)
	if err != nil {
		httpStatus := http.StatusInternalServerError
		return c.JSON(httpStatus, map[string]interface{}{"message": err.Error})
	}

	response := entity.LoginResponse{Token: token}
	return c.JSON(http.StatusOK, response)
}

func getCore() (c *core) {
	c = coreInstance

	if c == nil {
		c = new(core)
		userStorage, _ := storage.GetUserStorage(storage.Postgre)
		tokenStorage, _ := storage.GetTokenStorage(storage.Redis)

		c.userStorage = userStorage
		c.tokenStorage = tokenStorage
		coreInstance = c
	}

	return
}
