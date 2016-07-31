package main

import "fmt"

func main() {

	cmd := parseCmd()

	if cmd.help {
		fmt.Println("hunter [-options] class [args...]")
	} else if cmd.version {
		fmt.Println("hunter 0.0.1")
	} else {
		fmt.Println("hunter -help")
	}

}
