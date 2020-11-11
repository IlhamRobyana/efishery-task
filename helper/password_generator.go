package helper

import (
	"fmt"

	"github.com/sethvargo/go-password/password"
)

func GeneratePassword(length int) (str string) {
	pass, err := password.Generate(4, 2, 2, false, false)
	if err != nil {
		panic(fmt.Sprintf("%v", err))
	}
	return pass
}
