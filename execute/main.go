package main

import (
	"bigary/src/command"
	"bigary/src/utils"
	"fmt"
	"os"
)

func main() {
	var location string
	if l, e := utils.Direction(); e != nil {
		fmt.Println("‼️ ", e)
		return
	} else {
		location = l
	}

	argument := os.Args

	switch len(argument) {
	case 1:
		utils.Help()
		return

	case 2:
		err := command.OneCommand(argument)
		if err != nil {
			fmt.Println(err)
			return
		}
		return
	default:
		if err := command.Run(argument, location); err != nil {
			fmt.Println(err)
			return
		}

	}
}
