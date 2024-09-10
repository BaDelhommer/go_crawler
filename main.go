package main

import (
	"fmt"
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

	maxDepth := 3

	pages := make(map[string]int)

	crawlPage(baseURL, baseURL, pages, maxDepth)

	for normalizeedURL, count := range pages {
		fmt.Printf("%d - %s\n", count, normalizeedURL)
	}
}
