package main

import (
	"fmt"
	"net/url"
	"strings"

	"golang.org/x/net/html"
)

func getURLsfromHTML(htmlBody, rawBaseURL string) ([]string, error) {
	urls := []string{}
	doc, err := html.Parse(strings.NewReader(htmlBody))
	if err != nil {
		return []string{}, err
	}

	var f func(*html.Node) ([]string, error)
	f = func(n *html.Node) ([]string, error) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, a := range n.Attr {
				if a.Key == "href" {
					parsedURL, err := url.Parse(a.Val)
					if err != nil {
						fmt.Printf("error parsing urls: %v", err)
					}

					if !parsedURL.IsAbs() {
						joinedURL, err := url.JoinPath(rawBaseURL, parsedURL.String())
						if err != nil {
							fmt.Println("error joining urls: ", err)
						}
						urls = append(urls, joinedURL)
					} else {
						urls = append(urls, parsedURL.String()+"/")
					}
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
		return urls, nil
	}
	return f(doc)
}
