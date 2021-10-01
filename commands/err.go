package commands

import (
	"errors"
)

func UserNameAleadyExists() error {
	return errors.New("username already exists")
}
