package utils

import (
	"os"
)

func Version() string {
	a, e := os.ReadFile("version")

	if e != nil {
		return ""
	}

	return string(a)
}
