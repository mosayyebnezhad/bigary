package utils

import (
	"errors"
	"os"
)

func Direction() (string, error) {

	location, err := os.Getwd()

	if err != nil {

		return "", errors.New("wrong direction detection")

	}
	return location, nil
}
