package main

import "flag"

//hunter [-options] class [args...]
type Cmd struct {
	help      bool
	version   bool
	classpath string
	class     string
	args      []string
}

func parseCmd() *Cmd {
	cmd := &Cmd{}

	flag.BoolVar(&cmd.help, "help", false, "show help message")
	flag.BoolVar(&cmd.version, "version", false, "jvm version")
	flag.StringVar(&cmd.classpath, "classpath", "", "java classpath")

	flag.Parse()

	args := flag.Args()
	if len(args) > 0 {
		cmd.class = args[0]
		cmd.args = args[1:]
	}
	return cmd
}
