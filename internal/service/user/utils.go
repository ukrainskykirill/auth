package user

import (
	"fmt"
	"regexp"
)

func validatePassword(pass, confPass string) error {
	if pass != confPass {
		return fmt.Errorf("password does not match")
	}
	return nil
}

func validateEmail(email string) error {
	match, err := regexp.MatchString("([a-zA-Z0-9._-]+@[a-zA-Z0-9._-]+.[a-zA-Z0-9_-]+)", email)
	if err != nil {
		return err
	}
	if !match {
		return fmt.Errorf("email does not match")
	}
	return nil
}
