package helper

import (
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/ilhamrobyana/efishery-task/entity"
)

func GenerateToken(user entity.User) (token string, e error) {
	t := jwt.New(jwt.SigningMethodHS256)
	claims := t.Claims.(jwt.MapClaims)
	claims["name"] = user.Name
	claims["phone"] = user.Phone
	claims["role"] = user.Role
	claims["timestamp"] = time.Now()

	secret := os.Getenv("TOKEN_SECRET")
	token, e = t.SignedString([]byte(secret))
	return
}
