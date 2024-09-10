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

	const maxConcurrency = 3

	cfg, err := configure(baseURL, maxConcurrency)
	if err != nil {
		fmt.Printf("error configure: %v", err)
		return
	}

	fmt.Printf("Beginning crawl of %s", baseURL)

	cfg.wg.Add(1)
	go cfg.crawlPage(baseURL)
	cfg.wg.Wait()

	for normalizedURL, count := range cfg.pages {
		fmt.Printf("%d - %s\n", count, normalizedURL)
	}
}
