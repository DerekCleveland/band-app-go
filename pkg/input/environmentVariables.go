package input

import (
	"errors"
	"os"
)

type EnvVariables struct {
	UsernameStorage string
	PasswordStorage string
	HostStorage     string
}

func GetEnvVariables() (*EnvVariables, error) {
	var x EnvVariables

	x.UsernameStorage = os.Getenv("USERNAME_STORAGE")
	if x.UsernameStorage == "" {
		return nil, errors.New("expected environment variable USERNAME_STORAGE not set")
	}

	x.PasswordStorage = os.Getenv("PASSWORD_STORAGE")
	if x.PasswordStorage == "" {
		return nil, errors.New("expected environment variable PASSWORD_STORAGE not set")
	}

	x.HostStorage = os.Getenv("HOST_STORAGE")
	if x.HostStorage == "" {
		return nil, errors.New("expected environment variable HOST_STORAGE not set")
	}

	return &x, nil
}