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
	// Check if pages map bigger than maxPages
	if cfg.pagelen() >= cfg.maxPages {
		return
	}

	//Make sure the rawCurrentURL is on the same domain as the rawBaseURL.
	parsedCurrentUrl, err := url.Parse(rawCurrentURL)
	if err != nil {
		fmt.Printf("error parsing current URL:%v", err)
		return
	}
	if cfg.baseURL.Hostname() != parsedCurrentUrl.Hostname() {
		return
	}

	//Get a normalized version of the rawCurrentURL
	normalizedCurrentURL, err := normalizeURL(rawCurrentURL)
	if err != nil {
		fmt.Printf("error in normalizing current URL:%v\n", err)
		return
	}

	//check if pages has normalizedCurrentURL
	isFirst := cfg.addPageVisit(normalizedCurrentURL)
	if !isFirst {
		return
	}

	fmt.Printf("crawing current URL:%s\n", rawCurrentURL)

	//Get the HTML from the current URL, and add a print statement
	htmlBody, err := getHTML(rawCurrentURL)
	if err != nil {
		fmt.Printf("error fetching HTML:%v\n", err)
		return
	}

	//Get all the URLs from the response body HTML
	urls, err := getURLsFromHTML(htmlBody, cfg.baseURL)
	if err != nil {
		fmt.Printf("error getting URLs from HTML:%v\n", err)
		return
	}

	for _, url := range urls {
		cfg.wg.Add(1)
		go cfg.crawlPage(url)
	}

}
