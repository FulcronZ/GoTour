package main

import (
	"fmt"
	"time"
)

type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page
	Fetch(url string) (body string, urls []string, err error)
}

var busyCrawling chan bool // count alive goroutines
var crawledURLs map[string]bool

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum depth.
func Crawl(url string, depth int, fetcher Fetcher) {
	// This implementation doesn't do either:
	if depth <= 0 {
		return
	}

	var isRoot bool
	if len(crawledURLs) == 0 {
		isRoot = true
	}

	body, urls, err := fetcher.Fetch(url)
	switch err {
	case nil:
		crawledURLs[url] = true
		fmt.Printf("found: %s %q\n", url, body)
		for _, u := range urls {
			// Check if crawled before
			if _, crawled := crawledURLs[u]; !crawled {
				busyCrawling <- true // gen a token
				go Crawl(u, depth-1, fetcher)
			}
		}
	default:
		crawledURLs[url] = false
		fmt.Println(err)
	}

	switch isRoot {
	case true:
		// TODO: Find a better way to wake up upon empty channel
		for len(busyCrawling) > 0 {
			time.Sleep(100 * time.Millisecond)
		}
		close(busyCrawling)
		return
	case false:
		<-busyCrawling // return a token
		return
	}
}

func main() {
	// TODO: Find a way round to overcome this limitation
	// A channel is used to check if all gorountines finished
	// This limits the concurrency to the size of channel depth
	busyCrawling = make(chan bool, 2)
	crawledURLs = make(map[string]bool)
	Crawl("http://golang.org/", 4, fetcher)
}

// fakeFetcher is Fetcher that returns canned results.
type fakeFetcher map[string]*fakeResult

type fakeResult struct {
	body string
	urls []string
}

func (f fakeFetcher) Fetch(url string) (string, []string, error) {
	if res, ok := f[url]; ok {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found: %s", url)
}

// fetcher is a populated fakeFetcher
var fetcher = fakeFetcher{
	"http://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"http://golang.org/cmd/",
			"http://golang.org/pkg/",
		},
	},
	"http://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"http://golang.org/",
			"http://golang.org/cmd/",
			"http://golang.org/pkg/fmt/",
			"http://golang.org/pkg/os/",
		},
	},
	"http://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"http://golang.org/",
			"http://golang.org/pkg/",
		},
	},
	"http://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"http://golang.org/",
			"http://golang.org/pkg/",
		},
	},
}
