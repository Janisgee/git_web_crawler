package main

import (
	"fmt"
	"net/url"
	"sync"
)

type config struct {
	pages              map[string]int
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

func configure(rawbaseURL string, maxConcurrency int) (*config, error) {
	// Initialize pages
	// Parse and set baseURL
	baseURL, err := url.Parse(rawbaseURL)
	if err != nil {
		return nil, fmt.Errorf("couldn't parse base URL: %v", err)
	}

	return &config{
		pages:              make(map[string]int),
		baseURL:            baseURL,
		mu:                 &sync.Mutex{},
		concurrencyControl: make(chan struct{}, maxConcurrency),
		wg:                 &sync.WaitGroup{},
	}, nil

}
