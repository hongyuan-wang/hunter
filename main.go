package main

import "fmt"
import "hunter/classpath"

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
	cp := classpath.Parse(cmd.classpath)

	data, err := cp.ReadClass(cmd.class)
	if err != nil {
		fmt.Printf("read class %s error, %s ", cmd.class, err)
	} else {

		fmt.Printf("classs data %s\n", data)
	}
	fmt.Printf("start jvm hunter to run %s with %s ...\n", cmd.class, cmd.classpath)
}
