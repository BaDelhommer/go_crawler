package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args
	switch {
	case len(args) < 2:
		fmt.Println("no website provided")
		os.Exit(1)
	case len(args) > 2:
		fmt.Println("too many arguments provided")
		os.Exit(1)
	default:
		fmt.Println("starting crawl of: ", args[1])
	}
}
