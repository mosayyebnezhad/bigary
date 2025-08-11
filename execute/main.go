package main

import (
	"bigary/src/utils"
	"fmt"
)

func main() {
	var location string
	if l, e := utils.Direction(); e != nil {
		fmt.Println("‼️ ", e)
		return
	} else {
		location = l
	}

	var args map[string]string
	if a, Err := utils.Arguments(); Err != nil {
		fmt.Println("‼️ ", Err)
		return
	} else {
		args = a
	}

	fmt.Println(args["name"], location)

}
