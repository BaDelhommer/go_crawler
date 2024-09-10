package main

import (
	"fmt"
	"net/url"
)

func (cfg *config) crawlPage(rawCurrentURL string) {

	cfg.concurrencyControl <- struct{}{}
	defer func() {
		<-cfg.concurrencyControl
		cfg.wg.Done()
	}()

	currentURLObj, err := url.Parse(rawCurrentURL)
	if err != nil {
		fmt.Printf("Error couldn't parse url '%s': %v\n", rawCurrentURL, err)
		return
	}

	if currentURLObj.Hostname() != cfg.baseURL.Hostname() {
		return
	}

	normalizedURL, err := normalizeURL(rawCurrentURL)
	if err != nil {
		fmt.Printf("error normalizedURL: %v\n", err)
	}

	isFirst := cfg.addPageVisit(normalizedURL)

	if !isFirst {
		return
	}

	fmt.Printf("crawling %s\n", rawCurrentURL)

	htmlBody, err := getHtml(rawCurrentURL)
	if err != nil {
		fmt.Printf("error getHTML: %v\n", err)
		return
	}

	nextURLs, err := getURLsfromHTML(htmlBody, cfg.baseURL)
	if err != nil {
		fmt.Printf("error getURLsFromHTML: %v\n", err)
	}

	for _, nextURL := range nextURLs {
		cfg.wg.Add(1)
		go cfg.crawlPage(nextURL)
	}
}
