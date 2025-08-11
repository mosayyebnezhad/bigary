package command

import (
	"bigary/src/utils/arguments"
	"fmt"
)

func Run(args []string, direction string) error {
	fmt.Println("your command is", args[1])

	var arg map[string]string
	if a, Err := arguments.Arguments(args); Err != nil {

		return Err
	} else {
		arg = a
	}
	fmt.Println(arg)
	return nil

}
