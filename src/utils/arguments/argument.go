package arguments

import (
	"errors"
	"strings"
)

func Arguments(argument []string) (map[string]string, error) {

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
					return nil, errors.New("argument syntax is not valid, 2th input most be One of them : key`=`value | --key value | -key value ")
				}
				arguments[splt[0]] = splt[1]
				// fmt.Println(value)
			}
		}

	}

	return arguments, nil
}

func Command(argument []string) (string, error) {
	arg := argument[1]

	if len(arg) < 1 {
		return "", errors.New("command is most be at least 2 word")
	}

	if strings.HasPrefix(arg, "--") {
		return strings.TrimPrefix(arg, "--"), nil
	}

	if strings.HasPrefix(arg, "-") {
		return strings.TrimPrefix(arg, "-"), nil
	}

	return arg, nil
}
