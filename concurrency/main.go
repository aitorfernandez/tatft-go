package main

import (
	"fmt"
)

type WebsiteChecker func(string) bool

type result struct {
	string
	bool
}

func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)
	resultChannel := make(chan result)

	for _, url := range urls {
		go func(u string) {
			resultChannel <- result{u, wc(u)}
		}(url)
	}

	for i := 0; i < len(urls); i++ {
		r := <-resultChannel
		results[r.string] = r.bool
	}
	return results

	// race condition bug
	// go test main.go main_test.go -race
	// results := make(map[string]bool, len(urls))
	// for _, url := range urls {
	// 	go func(u string) {
	// 		results[u] = wc(u)
	// 	}(url)
	// }
	// return results
}

func main() {
	fmt.Println("tatft")
}
