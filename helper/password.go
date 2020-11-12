package helper

import (
	"fmt"

	"github.com/sethvargo/go-password/password"
)

func GeneratePassword(length int) (str string) {
	pass, err := password.Generate(4, 0, 0, false, true)
	if err != nil {
		panic(fmt.Sprintf("%v", err))
	}
	return pass
}
