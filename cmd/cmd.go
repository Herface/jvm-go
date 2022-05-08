package cmd

import (
	"flag"
	"fmt"
	"os"
)

type Cmd struct {
	Help           bool
	Version        bool
	Classpath      string
	Class          string
	Xbootclasspath string
	Args           []string
}

func ParseCmd() *Cmd {
	cmd := &Cmd{}
	flag.Usage = PrintUsage
	flag.BoolVar(&cmd.Help, "help", false, "print help message")
	flag.BoolVar(&cmd.Version, "version", false, "print version info")
	flag.StringVar(&cmd.Classpath, "classpath", "", "specify classpath")
	flag.StringVar(&cmd.Classpath, "cp", "", "specify classpath")
	flag.StringVar(&cmd.Xbootclasspath, "Xbootclasspath", "", "boot classpath")
	flag.Parse()
	args := flag.Args()
	if len(args) > 0 {
		cmd.Class = args[0]
		cmd.Args = args[1:]
	}
	return cmd
}

func PrintUsage() {
	fmt.Printf("Usage: %v [-options] class [args]", os.Args[0])
}
