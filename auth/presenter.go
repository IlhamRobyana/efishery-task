package auth

import (
	"net/http"
	"strings"

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

func ValidateToken(c echo.Context) error {
	r := c.Request().Header.Get("Authorization")
	if len(r) == 0 {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{"message": "Bearer token is empty"})
	}
	token := strings.Split(r, " ")[1]

	authCore := getCore()
	claims, err := authCore.validateToken(token)
	if err != nil {
		httpStatus := http.StatusInternalServerError
		return c.JSON(httpStatus, map[string]interface{}{"message": err.Error})
	}
	return c.JSON(http.StatusOK, claims)
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
