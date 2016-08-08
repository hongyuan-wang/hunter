package main

import "fmt"

func main() {

	cmd := parseCmd()

	if cmd.help {
		fmt.Println("hunter [-options] class [args...]")
	} else if cmd.version {
		fmt.Println("hunter 0.0.1")
	} else if cmd.class != "" {
		startJvm(cmd)
	} else {
		fmt.Println("hunter -help")
	}

}

func startJvm(cmd *Cmd) {
	fmt.Printf("start jvm hunter to run %s with %s ...\n", cmd.class, cmd.classpath)
}
