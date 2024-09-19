package main

import (
	"fmt"
	"os"
)

func main() {

	if len(os.Args) < 2 {
		fmt.Printf("no website provided\n")
		os.Exit(1)
	}

	if len(os.Args) > 2 {
		fmt.Printf("too many arguments provided\n")
		os.Exit(1)
	}

	rawBaseURL := os.Args[1]

	pages := make(map[string]int)

	crawlPage(rawBaseURL, rawBaseURL, pages)

	for urls, counts := range pages {
		fmt.Printf("%d   Key: %s\n", counts, urls)
	}

}
