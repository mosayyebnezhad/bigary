package command

import (
	"bigary/src/utils"
	"bigary/src/utils/arguments"
	"fmt"
	"strings"
)

func OneCommand(argumen []string) error {

	command, err := arguments.Command(argumen)

	if err != nil {
		return err
	}

	switch {
	case strings.Contains(command, "health"):
		utils.ShovelMessage(command)
	case strings.Contains(command, "version"):
		fmt.Println(utils.Version())

	default:
		isHelp := strings.Contains(command, "help")
		if !isHelp {
			fmt.Println("command not found !")
		}

		utils.Help()
	}

	return nil

}
