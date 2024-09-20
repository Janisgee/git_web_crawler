package main

import (
	"fmt"
	"net/url"
	"sort"
	"sync"
)

type config struct {
	pages              map[string]int
	maxPages           int
	baseURL            *url.URL
	mu                 *sync.Mutex
	concurrencyControl chan struct{}
	wg                 *sync.WaitGroup
}

func (cfg *config) addPageVisit(normalized string) (isFirst bool) {
	cfg.mu.Lock()
	defer cfg.mu.Unlock()

	_, visited := cfg.pages[normalized]
	if visited {
		cfg.pages[normalized]++
		return false
	}
	cfg.pages[normalized] = 1
	return true
}

func (cfg *config) pagelen() int {
	cfg.mu.Lock()
	defer cfg.mu.Unlock()
	return len(cfg.pages)
}

func printReport(pages map[string]int, baseURL string) {
	//sort pages map
	fmt.Println("=============================")
	fmt.Printf("   REPORT for %s   \n", baseURL)
	fmt.Println("=============================")

	keys := make([]string, 0, len(pages))
	for k := range pages {
		keys = append(keys, k)
	}

	sort.SliceStable(keys, func(i, j int) bool {
		if pages[keys[i]] == pages[keys[j]] {
			return keys[i] < keys[j]
		}
		return pages[keys[i]] > pages[keys[j]]
	})

	for _, k := range keys {
		fmt.Printf("Found %d internal links to %s\n", pages[k], k)
	}
}

func configure(rawbaseURL string, maxConcurrency int, maxPages int) (*config, error) {
	// Initialize pages
	// Parse and set baseURL
	baseURL, err := url.Parse(rawbaseURL)
	if err != nil {
		return nil, fmt.Errorf("couldn't parse base URL: %v", err)
	}

	return &config{
		pages:              make(map[string]int),
		maxPages:           maxPages,
		baseURL:            baseURL,
		mu:                 &sync.Mutex{},
		concurrencyControl: make(chan struct{}, maxConcurrency),
		wg:                 &sync.WaitGroup{},
	}, nil

}
