package utils

import (
	"fmt"
	"time"
)

func ShovelMessage(do string) {
	i := 0
	for {
		i++

		var bill string

		if i%2 == 0 {
			bill = "🪏 "
		} else {
			bill = " 🪏"
		}

		fmt.Printf("%s %s \r", do, bill)

		time.Sleep(time.Millisecond * 300)

		if i == 15 {
			break
		}
	}
}
