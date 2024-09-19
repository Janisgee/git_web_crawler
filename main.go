package main

import (
	"fmt"
	"os"
)

func main() {

	if len(os.Args) < 2 {
		fmt.Printf("no website provided\n")
		return
	}

	if len(os.Args) > 2 {
		fmt.Printf("too many arguments provided\n")
		return
	}

	rawBaseURL := os.Args[1]

	maxConcurrency := 20

	cfg, err := configure(rawBaseURL, maxConcurrency)
	if err != nil {
		fmt.Printf("error in configuring struct:%v", err)
		return
	}

	cfg.wg.Add(1)
	go cfg.crawlPage(rawBaseURL)
	cfg.wg.Wait()

	for normalizeURL, count := range cfg.pages {
		fmt.Printf("%d   - %v\n", count, normalizeURL)
	}

}
