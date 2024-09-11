package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) > 4 {
		fmt.Printf("Too many arguments %v", len(os.Args))
		os.Exit(1)
	}

	if len(os.Args) < 4 {
		fmt.Printf("Too few arguments %v", len(os.Args))
		os.Exit(1)
	}

	baseURL := os.Args[1]
	maxConcurrency, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Printf("Argument 2 must be int, got %v", err)
	}
	maxPages, err := strconv.Atoi(os.Args[3])
	if err != nil {
		fmt.Printf("Argument 3 must be int, got %v", err)
	}

	cfg, err := configure(baseURL, maxConcurrency, maxPages)
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

	printReport(cfg.pages, baseURL)
}
