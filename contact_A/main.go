package main

import (
	"fmt"
	"strconv"
)

func main() {

	// running server A
	go server()

	for {
		fmt.Print(`
MENU

0. exit
1. send
2. recieve

choice : `)

		var n string
		fmt.Scan(&n)
		i, err := strconv.Atoi(n)
		if err != nil {
			println("Invalid Input!")
		}

		switch i {
		case 0:
			println("Exiting....")
			return
		case 1:
			fmt.Println("Feature running :")
			client()
		default:
			println("Invalid Input!")
		}

	}

}
