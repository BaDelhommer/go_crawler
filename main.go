package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	args := os.Args
	var baseURL string
	switch {
	case len(args) < 2:
		fmt.Println("no website provided")
		os.Exit(1)
	case len(args) > 2:
		fmt.Println("too many arguments provided")
		os.Exit(1)
	default:
		baseURL = args[1]
		fmt.Println("starting crawl of: ", baseURL)
	}

	doc, err := getHtml(baseURL)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(doc)
}
