package utils

import (
	"errors"
	"os"
	"strings"
)

func Arguments() (map[string]string, error) {
	argument := os.Args

	arg := argument[2:]

	var arguments = make(map[string]string)

	for index, value := range arg {
		if strings.HasPrefix(value, "-") || strings.HasPrefix(value, "--") {

			if value[0] == '-' && value[1] == '-' {
				value = strings.TrimPrefix(value, "--")
			}

			if value[0] == '-' && value[1] != '-' {
				value = strings.TrimPrefix(value, "-")
			}
			if len(arg)-1 < index+1 {
				return nil, errors.New("last argument can not be key")
			}
			arguments[value] = arg[index+1]
			arg[index+1] = ""
		} else {
			if value != "" {
				splt := strings.Split(value, "=")
				if len(splt) != 2 {
					return nil, errors.New("does not contain `=`")
				}
				arguments[splt[0]] = splt[1]
				// fmt.Println(value)
			}
		}

	}

	return arguments, nil
}
