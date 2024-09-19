package main

import (
	"fmt"
	"net/url"
	"strings"

	"golang.org/x/net/html"
)

func getURLsFromHTML(htmlBody string, baseURL *url.URL) ([]string, error) {

	htmlReader := strings.NewReader(htmlBody) // Create io.Reader
	doc, err := html.Parse(htmlReader)        // Create a tree of html.Node
	if err != nil {
		return nil, fmt.Errorf("couldn't parse HTML: %v", err)

	}

	var traverse func(*html.Node)
	var urlSlice []string

	traverse = func(node *html.Node) {
		if node.Type == html.ElementNode && node.Data == "a" {
			for _, anchor := range node.Attr {
				if anchor.Key == "href" {
					href, err := url.Parse(anchor.Val)
					if err != nil {
						fmt.Printf("couldn't parse href value %s: %v\n", anchor.Val, err)
						continue
					}
					resolvedUrl := baseURL.ResolveReference(href)
					urlSlice = append(urlSlice, resolvedUrl.String())
				}

			}
		}
		child := node.FirstChild // Start with the first child node

		// Continue looping as long as the child is not nil
		for child != nil {
			traverse(child) // Call traverseNodes on the current child
			// Move to the next sibling
			child = child.NextSibling
		}
	}

	traverse(doc)

	return urlSlice, nil
}
