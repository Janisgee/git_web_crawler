package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	if len(os.Args) < 4 {
		fmt.Printf("please provide enough argument.\n")
		fmt.Printf("usage: crawler <baseURL> <maxConcurrency> <maxPages>\n")
		return
	}

	if len(os.Args) > 4 {
		fmt.Printf("too many arguments provided.")
		return
	}

	//Configurable via command-line args.
	//Convert maxConcurrency & MaxPage from string to int
	rawBaseURL := os.Args[1]
	maxConcurrency, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Printf("error converting MaxCurrency from string to int:%v", err)
		return
	}
	maxPages, err := strconv.Atoi(os.Args[3])
	if err != nil {
		fmt.Printf("error converting maxPages from string to int:%v", err)
		return
	}

	cfg, err := configure(rawBaseURL, maxConcurrency, maxPages)
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

	//Print report
	printReport(cfg.pages, rawBaseURL)
}
