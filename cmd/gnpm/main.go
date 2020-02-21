package main

import (
	"fmt"
	"os"
)

func usage() {
	fmt.Printf("Usage: npm <command>\n\nwhere <command> is one of:\n    install\n")
}

func run(args []string) int {
	if len(args) == 0 {
		usage()
		return 0
	}

	switch args[0] {
	case "install":
		if err := install(args[1:]); err != nil {
			fmt.Println("Error: " + err.Error())
			return 1
		}
	default:
		usage()
	}

	return 0
}

func main() {
	os.Exit(run(os.Args[1:]))
}
