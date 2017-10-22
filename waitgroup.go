// 经典正例
package main

import (
	"fmt"
	"net/http"
	"sync"
)

var wg sync.WaitGroup

func getURL(url string) {
	a, _ := http.Get(url)
	fmt.Println(a.Status)
	wg.Done()
}

func main() {
	var urls = []string{
		"http://www.orangeapk.com",
		"http://zablog.me/",
		"http://www.pku.edu.cn/",
	}
	for _, url := range urls {
		// Increment the WaitGroup counter.
		wg.Add(1)
		// Launch a goroutine to fetch the URL.
		go getURL(url)
	}
	// Wait for all HTTP fetches to complete.
	wg.Wait()
}
