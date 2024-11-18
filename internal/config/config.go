package config

import (
	"errors"
	"fmt"

	"github.com/joho/godotenv"
)

var (
	errVariableNotFound = errors.New("environment variable not found")
	errVariableParse    = errors.New("parsing variable")
)

func LoadConfig() error {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	return nil
}
