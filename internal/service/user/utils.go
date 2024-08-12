package user

import (
	"fmt"
	prError "github.com/ukrainskykirill/auth/internal/error"
	"regexp"
)

func validatePassword(pass, confPass string) error {
	if pass != confPass {
		return fmt.Errorf("%w: passwords doesnt match", prError.ErrPassword)
	}
	return nil
}

func validateEmail(email string) error {
	match, err := regexp.MatchString("([a-zA-Z0-9._-]+@[a-zA-Z0-9._-]+.[a-zA-Z0-9_-]+)", email)
	if err != nil {
		return err
	}
	if !match {
		return prError.ErrInvalidEmail
	}
	return nil
}
