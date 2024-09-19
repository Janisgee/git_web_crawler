package main

import (
	"fmt"
	"net/url"
)

func crawlPage(rawBaseURL, rawCurrentURL string, pages map[string]int) {
	//Make sure the rawCurrentURL is on the same domain as the rawBaseURL.
	parsedBaseURL, err := url.Parse(rawBaseURL)
	if err != nil {
		fmt.Printf("error parsing base URL:%v\n", err)
		return
	}
	parsedCurrentUrl, err := url.Parse(rawCurrentURL)
	if err != nil {
		fmt.Printf("error parsing current URL:%v", err)
		return
	}
	if parsedBaseURL.Hostname() != parsedCurrentUrl.Hostname() {
		return
	}

	//Get a normalized version of the rawCurrentURL
	normalizedCurrentURL, err := normalizeURL(rawCurrentURL)
	if err != nil {
		fmt.Printf("error in normalizing current URL:%v\n", err)
		return
	}

	//check if pages has normalizedCurrentURL
	_, ok := pages[normalizedCurrentURL]
	if ok {
		pages[normalizedCurrentURL]++
		return
	}
	pages[normalizedCurrentURL] = 1

	fmt.Printf("crawing current URL:%s\n", rawCurrentURL)

	//Get the HTML from the current URL, and add a print statement
	htmlBody, err := getHTML(rawCurrentURL)
	if err != nil {
		fmt.Printf("error fetching HTML:%v\n", err)
		return
	}

	//Get all the URLs from the response body HTML
	urls, err := getURLsFromHTML(htmlBody, rawBaseURL)
	if err != nil {
		fmt.Printf("error getting URLs from HTML:%v\n", err)
		return
	}

	for _, urlKey := range urls {
		crawlPage(rawBaseURL, urlKey, pages)
	}

}
