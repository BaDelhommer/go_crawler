package main

import (
	"fmt"
	"net/url"
)

func crawlPage(rawBaseURL, rawCurrentURL string, pages map[string]int, depth int) {
	depth--
	if depth <= 0 {
		return
	}

	currentURLObj, err := url.Parse(rawCurrentURL)
	if err != nil {
		fmt.Printf("Error couldn't parse url '%s': %v\n", rawCurrentURL, err)
		return
	}

	baseURLObj, err := url.Parse(rawBaseURL)
	if err != nil {
		fmt.Printf("Error couldn't parse url '%s': %v\n", rawBaseURL, err)
		return
	}

	if currentURLObj.Hostname() != baseURLObj.Hostname() {
		return
	}

	normalizedURL, err := normalizeURL(rawCurrentURL)
	if err != nil {
		fmt.Printf("Error normalizing url: %v", err)
	}

	if _, visited := pages[normalizedURL]; visited {
		pages[normalizedURL]++
		return
	}

	pages[normalizedURL] = 1

	fmt.Println("crawling ", rawCurrentURL)

	html, err := getHtml(rawCurrentURL)
	if err != nil {
		fmt.Printf("Error getting html: %v", err)
		return
	}

	nextURLs, err := getURLsfromHTML(html, rawBaseURL)
	if err != nil {
		fmt.Printf("Error getting urls from html: %v", err)
		return
	}

	for _, nextURL := range nextURLs {
		crawlPage(rawBaseURL, nextURL, pages, depth)
	}
}
